//websocket server2 聊天室
package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
)

type ws_client struct{
	client map[int]*websocket.Conn
}
func (self *ws_client) Append(item *websocket.Conn) {
	self.client[len(self.client)] = item
}
func (self *ws_client) Push(item *websocket.Conn) {
	self.client[len(self.client)] = item
}

func (self *ws_client) Pop() *websocket.Conn {
	tree, ok := self.client[len(self.client) - 1]
	if ok {
		delete(self.client, len(self.client) - 1)
		return tree
	}
	return nil
}
func (self *ws_client)SendMsg(msg string) (error){
	if !self.Empty(){
		for _,ws := range self.client{
			if err := websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				fmt.Println(ws)
				return err
			}
		}
	}
	return nil
}

func (self *ws_client) Empty() bool {
	return len(self.client) == 0
}

var WSC *ws_client

func Echo(ws *websocket.Conn) {
	var err error
	WSC.Push(ws)
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)
		go WSC.SendMsg(msg)
		/*if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}*/
	}
}

func main() {
	WSC = &ws_client{make(map[int]*websocket.Conn)}

	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}