package main

import (
	"gopkg.in/robfig/cron.v1"
	"log"
)

func main() {
	cronV1()
}

func cronV1() {
	stop := make(chan bool, 1)
	//crontab := cron.NewWithLocation(TimeLocation)
	crontab := cron.New()

	//every 20 seconds
	crontab.AddFunc("*/20 * * * * *", func() {
		log.Println("crontab 20s run")
	})
	// every 50 seconds
	crontab.AddFunc("*/50 * * * * *", func() {
		log.Println("crontab 50s run")
	})

	// 定时
	crontab.AddFunc("* 13 13 * * *", func() {
		log.Println("######crontab 11:21 run#########")
	})
	crontab.Start()
	<-stop
}
