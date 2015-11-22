//websocket server2 聊天室
package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
	"container/list"
	"reflect"
	"encoding/json"
	"time"
)

type Msg_data struct{
	status 			int32 			//状态码
	_t 				string			//类型
	msg 			string			//消息内容
	to_user 		string			//消息接收者
	from_user		string			//消息发送者
}

type Client struct{
	client_id 		string
	nickname 		string
	ws 				*websocket.Conn
	msg 			*list.List
}

type Client_list struct{
	list 			*list.List		//客户端列表
	msg_list		*list.List		//群发消息列表
	msg_ch			chan int		//群发消息开关
}

var c_list *Client_list
//添加客户端
func (cl *Client_list)Append(data interface{}) (*list.Element){
	return cl.list.PushBack(data)
}

//处理客户端发送的消息
func (cl *Client_list)msg_handle(msg string, ws *websocket.Conn){
	var msg_tmp interface{}
	if err := json.Unmarshal([]byte(msg), &msg_tmp); err != nil{
		fmt.Println(err)
	}
	if msg_data, ok := msg_tmp.(map[string]interface{}); ok{
		switch msg_data["_t"] {
		case "normal":
				if msg_data["to_user"] == ""{
					cl.msg_list.PushBack(msg_data)
					cl.msg_ch <- 1
				}
		case "set":

		}
	}
}


//群发消息
func (cl *Client_list)SendMsg(msg string) (error){
	if cl.list.Len() > 0{
		for element := cl.list.Front(); element != nil;  {
			if client, ok := element.Value.(*Client); ok{
				if err := websocket.Message.Send(client.ws, msg); err != nil {
					fmt.Println("Can't send")
					element = element.Next();
					cl.list.Remove(element.Prev())
					continue
				}
			}else{
				fmt.Println(element.Value)
				fmt.Println(reflect.TypeOf(element.Value))
			}
			element = element.Next()
		}
	}
	return nil
}
//处理群发消息
func (cl *Client_list)SendNormal() (error){
	if cl.msg_list.Len() > 0 {
		e := cl.msg_list.Front();
		if msg_data, ok := e.Value.(map[string]interface{}); ok{
			if msg, err := json.Marshal(msg_data); err == nil{
				cl.SendMsg(string(msg))
			}
		}
		cl.msg_list.Remove(e)
	}
	return nil
}

//群发go程
func (cl *Client_list) start_mass(){
	for{
		select {
		case <-cl.msg_ch:
			cl.SendNormal()
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

//接收客户端连接请求
func Accept(ws *websocket.Conn) {
	var err error
	c_list.Append(&Client{ws:ws,msg:list.New()})
	fmt.Println("client connected")
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received from client: " + reply)
		c_list.msg_handle(reply, ws)
	}
}

func main() {
	c_list = &Client_list{
		list:		list.New(),
		msg_list:	list.New(),
		msg_ch:		make(chan int,100),
	}

	go c_list.start_mass()

	http.Handle("/", websocket.Handler(Accept))
	fmt.Println("start serving")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}