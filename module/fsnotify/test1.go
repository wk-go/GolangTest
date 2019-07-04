package main

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	//watcher, err := fsnotify.NewWatcher()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	var tick = time.Tick(50 * time.Millisecond)
	var evts = make([]fsnotify.Event, 0)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("[Modified file]:", event.Name)
					evts = append(evts, event)
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("[Created file]:", event.Name)
					evts = append(evts, event)
				}
			case <-tick:
				//通过增加时间间隔来过滤编辑器多次操作事件(算是目前比较有效的解决办法)
				if len(evts) == 0 {
					continue
				}
				log.Println("EVTS:", evts)
				evts = make([]fsnotify.Event, 0)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
