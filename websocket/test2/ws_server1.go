//websocket server2 聊天室
package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
	"container/list"
	"reflect"
)

type Client struct{
	ws *websocket.Conn
	msg *list.List
}

var client_list *list.List

func SendMsg(msg string) (error){
	if client_list.Len() > 0{
		for element := client_list.Front(); element != nil;  {
			if client, ok := element.Value.(*Client); ok{
				if err := websocket.Message.Send(client.ws, msg); err != nil {
					fmt.Println("Can't send")
					element = element.Next();
					client_list.Remove(element.Prev())
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

func Echo(ws *websocket.Conn) {
	var err error
	client_list.PushBack(&Client{ws:ws,msg:list.New()})
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)
		go SendMsg(msg)
		/*if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}*/
	}
}

func main() {
	client_list = list.New()

	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}