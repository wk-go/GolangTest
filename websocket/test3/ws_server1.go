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
)

type Msg_data struct{
	status 			int32 			//状态码
	_t 				string			//类型
	msg 			string			//消息内容
	to_user 		string			//消息接收者
	from_user		string			//消息发送者
}

type Client struct{
	ws 				*websocket.Conn
	client_id 		string
	nickname 		string
	msg 			*list.List
}

type Client_list struct{
	list 			*list.List
}

var c_list *Client_list
//
func (cl *Client_list)Append(data interface{}) (*list.Element){
	return cl.list.PushBack(data)
}
//发送消息
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

//接收客户端连接请求
func Accept(ws *websocket.Conn) {
	var err error
	c_list.Append(&Client{ws:ws,msg:list.New()})
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received from client: " + reply)
		msg_handle(reply, ws)
	}
}

//处理客户端发送的消息
func msg_handle(msg string, ws *websocket.Conn){
	var msg_tmp interface{}
	if err := json.Unmarshal([]byte(msg), &msg_tmp); err != nil{
		fmt.Println(err)
	}
	if msg_data, ok := msg_tmp.(map[string]interface{}); ok{
		if(msg_data["_t"] == "normal" && msg_data["to_user"] == ""){
			go c_list.SendMsg(msg)
		}
	}
}

func main() {
	c_list = &Client_list{list:list.New()}

	http.Handle("/", websocket.Handler(Accept))
	fmt.Println("start serving")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}