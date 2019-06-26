// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"golang_test/framework/beego/WebIM/models"
)

type Subscription struct {
	Archive []models.Event      // All the events from the archive.
	New     <-chan models.Event // New events coming in.
}

func newEvent(ep models.EventType, roomId, user, msg string) models.Event {
	return models.Event{ep, roomId, user, int(time.Now().Unix()), msg}
}

func Join(user, roomId string, ws *websocket.Conn) {
	beego.Info("User [", user, "] Join Room:", roomId)
	subscribe <- Subscriber{Name: user, RoomId: roomId, Conn: ws}
}

func Leave(user string) {
	beego.Info("User [", user, "] leave Room")
	unsubscribe <- user
}

type Subscriber struct {
	Name   string
	RoomId string
	Conn   *websocket.Conn // Only for WebSocket users; otherwise nil.
}
type Room struct {
	Id          string     //room id
	Name        string     //room name
	Limit       int        //room User number limit
	Count       int        //current user number
	Subscribers *list.List //user list
}

func (this *Room) Leave(username string) {
	for sub := this.Subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(*Subscriber).Name == username {
			this.Subscribers.Remove(sub)
			beego.Info("User [", sub.Value.(*Subscriber).Name, "] leave Room:", sub.Value.(*Subscriber).RoomId)
		}
	}
}

func newRoom(roomId string) {
	if IsRoomExist(rooms, roomId) {
		return
	}
	room := Room{Id: roomId}
	rooms.PushBack(room)
}
func GetRoom(roomId string) *Room {
	for sub := rooms.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(*Room).Id == roomId {
			return sub.Value.(*Room)
		}
	}
	return nil
}
func GetRoomElement(roomId string) *list.Element {
	for sub := rooms.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(*Room).Id == roomId {
			return sub
		}
	}
	return nil
}

func IsRoomExist(rooms *list.List, roomId string) bool {
	for sub := rooms.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(*Room).Id == roomId {
			return true
		}
	}
	return false
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// Channel for exit users.
	unsubscribe = make(chan string, 10)
	// Send events here to publish them.
	publish = make(chan models.Event, 10)
	// Long polling waiting list.
	waitingList = list.New()
	subscribers = list.New()
	rooms       = list.New()
)

// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(&sub) // Add user to the end of list.
				// Publish a JOIN event.
				room := GetRoom(sub.RoomId)
				room.Subscribers.PushBack(&sub)
				publish <- newEvent(models.EVENT_JOIN, sub.RoomId, sub.Name, "")
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event)
			models.NewArchive(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(*Subscriber).Name == unsub {
					subscribers.Remove(sub)

					user := sub.Value.(*Subscriber)

					room := GetRoom(user.RoomId)
					//beego.Info(":::room.Subscribers:::",room.Subscribers)
					room.Leave(unsub)
					//beego.Info(":::room.Subscribers:::",room.Subscribers)

					// Clone connection.
					ws := user.Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, sub.Value.(*Subscriber).RoomId, unsub, "") // Publish a LEAVE event.
					break
				}
			}
		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(*Subscriber).Name == user {
			return true
		}
	}
	return false
}
