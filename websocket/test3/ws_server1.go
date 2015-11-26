//websocket server2 聊天室
package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
	"container/list"
	"reflect"
	"encoding/json"
	"time"
	"fmt"
)

type Msg_data struct {
	status    int32  //状态码
	_t        string //类型
	msg       string //消息内容
	to_user   string //消息接收者
	from_user string //消息发送者
}

type Client struct {
	client_id string
	nickname  string
	ws        *websocket.Conn
	msg       *list.List
}

//发送消息
func (c *Client)SendMsg(msg string) (err error) {
	if err := websocket.Message.Send(c.ws, msg); err != nil {
		return err
	}
	return nil
}
//追加个人消息
func (c *Client)AppendMsg(msg string) (*list.Element) {
	log.Println(fmt.Sprintf("client append message  :%v",msg))
	return c.msg.PushBack(msg)
}
//发送消息
func (c *Client)SendMsgSingle() (err error) {
	for element := c.msg.Front(); element != nil; {
		if msg, ok := element.Value.(string); ok {
			log.Println(fmt.Sprintf("client single message can't send:%v",msg))
			if err := c.SendMsg(msg); err != nil {
				log.Println(fmt.Sprintf("client single message can't send:%v-%v",err,msg))
				return err
			}
		}
		c.msg.Remove(element)
		element = c.msg.Front()
	}
	return nil
}
//设置昵称
func (c *Client)SetNickname(nickname string) {
	c.nickname = nickname
	log.Println("client set nickname:" + c.nickname)
}


type ClientList struct {
	list       *list.List //客户端列表
	msg_list   *list.List //群发消息列表
	msg_ch     chan int   //群发消息开关
	msg_single chan int   //单发消息
}
//添加客户端
func (cl *ClientList)Append(data interface{}) (*list.Element) {
	return cl.list.PushBack(data)
}
//删除客户端
func (cl *ClientList)RemoveClient(c *Client) (*list.Element) {
	if cl.list.Len() > 0 {
		for element := cl.list.Front(); element != nil; {
			if client, ok := element.Value.(*Client); ok && client == c {
				log.Println(fmt.Sprintf("client:(nickname:%v) has been deleted", c.nickname))
				cl.list.Remove(element)
				return element
			}
			element = element.Next()
		}
	}
	return nil
}

//处理客户端发送的消息
func (cl *ClientList)msg_handle(msg string, client *Client) {
	var msg_tmp interface{}
	if err := json.Unmarshal([]byte(msg), &msg_tmp); err != nil {
		log.Println(err)
	}
	if msg_data, ok := msg_tmp.(map[string]interface{}); ok {
		switch msg_data["_t"] {
		case "normal":
			log.Println("received normal msg:send")
			if msg_data["to_user"] == "" {
				cl.msg_list.PushBack(msg_data)
				cl.msg_ch <- 1
			}
			log.Println("received normal msg:sended")
		case "set":
			if msg_data["to_user"] == "" && msg_data["nickname"] != "" {//设置昵称
				log.Println("set nickname")
				var flag = true
				for element := cl.list.Front(); element != nil; element = element.Next() {
					if client, ok := element.Value.(*Client); ok && client.nickname == msg_data["nickname"] {
						flag = false
					}
				}
				log.Println("set nickname:checked")
				nickname_msg := map[string]string{
					"status":"1",
					"_t":"set",
					"msg":"昵称修改成功",
					"nickname" : msg_data["nickname"].(string),
					"to_user":msg_data["nickname"].(string),
					"from_user":msg_data["nickname"].(string),
				}
				if flag {
					log.Println("set nickname:start")
					if nickname, ok := msg_data["nickname"].(string); ok {
						log.Println("set nickname:running")
						client.SetNickname(nickname)
						log.Println("set nickname:success")
					}
				}else{
					nickname_msg["status"] = "0"
					nickname_msg["msg"] = "您设置的昵称可能已经被别人使用，请重新设置"
					log.Println("set nickname:duplicate")
				}

				if msg, err := json.Marshal(nickname_msg); err == nil {
					client.AppendMsg(string(msg))
					cl.msg_single <- 1
					log.Println("set nickname:send result")
				}
			}
		}
	}
}


//群发消息
func (cl *ClientList)SendMsg(msg string) (error) {
	if cl.list.Len() > 0 {
		for element := cl.list.Front(); element != nil; {
			if client, ok := element.Value.(*Client); ok {
				if err := client.SendMsg(msg); err != nil {
					log.Println("Can't send")
					if e1 := element.Next(); e1 != nil {
						element = e1;
						if e2 := element.Prev(); e2 != nil {
							cl.list.Remove(e2)
						}
					}else{
						cl.list.Remove(element)
						element = nil
					}
					continue
				}
			}else {
				log.Println(element.Value)
				log.Println(reflect.TypeOf(element.Value))
			}
			element = element.Next()
		}
	}
	return nil
}
//处理群发消息
func (cl *ClientList)SendNormal() (error) {
	if cl.msg_list.Len() > 0 {
		e := cl.msg_list.Front();
		if msg_data, ok := e.Value.(map[string]interface{}); ok {
			if msg, err := json.Marshal(msg_data); err == nil {
				cl.SendMsg(string(msg))
			}
		}
		cl.msg_list.Remove(e)
	}
	return nil
}

//处理群发消息
func (cl *ClientList)SendSingle() (error) {
	if cl.list.Len() > 0 {
		for element := cl.list.Front(); element != nil; element = element.Next() {
			if client, ok := element.Value.(*Client); ok {
				go client.SendMsgSingle()
			}
		}
	}
	return nil
}

//群发go程
func (cl *ClientList) Start_send() {
	for {
		select {
		case <-cl.msg_ch:
			cl.SendNormal()
		case <-cl.msg_single:
			cl.SendSingle()
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

var cList *ClientList

//接收客户端连接请求
func Accept(ws *websocket.Conn) {
	var err error
	currClient := &Client{ws:ws, msg:list.New()}
	cList.Append(currClient)
	log.Println("client connected")
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Println("Can't receive")
			cList.RemoveClient(currClient)
			break
		}
		log.Println("Received from client: " + reply)
		cList.msg_handle(reply, currClient)
	}
}

func main() {
	cList = &ClientList{
		list:        list.New(),
		msg_list:    list.New(),
		msg_ch:        make(chan int, 100),
		msg_single:        make(chan int, 2000),
	}

	go cList.Start_send()

	http.Handle("/", websocket.Handler(Accept))
	log.Println("start serving")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}