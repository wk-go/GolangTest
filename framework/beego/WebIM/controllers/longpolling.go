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
	"git.oschina.net/walkskyer/GolangTest/framework/beego/WebIM/models"
)

// LongPollingController handles long polling requests.
type LongPollingController struct {
	baseController
}

// Join method handles GET requests for LongPollingController.
func (this *LongPollingController) Join() {
	// Safe check.
	uname := this.GetString("uname")
	roomId := this.GetString("roomid")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	// Join chat room.
	Join(uname, roomId, nil)

	this.TplName = "longpolling.html"
	this.Data["IsLongPolling"] = true
	this.Data["UserName"] = uname

	this.Data["roomid"] = roomId
	this.Data["room"] = GetRoom(roomId)
}

// Post method handles receive messages requests for LongPollingController.
func (this *LongPollingController) Post() {
	this.TplName = "longpolling.html"

	uname := this.GetString("uname")
	roomId := this.GetString("roomid")
	content := this.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(models.EVENT_MESSAGE,roomId, uname, content)
}

// Fetch method handles fetch archives requests for LongPollingController.
func (this *LongPollingController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	roomId := this.GetString("roomid")
	if err != nil {
		return
	}

	events := models.GetEvents(int(lastReceived),roomId)
	if len(events) > 0 {
		this.Data["json"] = events
		this.ServeJSON()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	this.Data["json"] = models.GetEvents(int(lastReceived),roomId)
	this.ServeJSON()
}