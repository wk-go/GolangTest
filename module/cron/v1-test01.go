package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
	"fmt"
)

func main() {
	cronV1()
}

func cronV1() {
	stop := make(chan bool, 1)
	TimeLocation,_ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(TimeLocation)

	crontab := cron.NewWithLocation(TimeLocation)
	//crontab := cron.New()

	//every 20 seconds
	crontab.AddFunc("*/20 * * * * *", func() {
		log.Println("crontab 20s run")
	})
	// every 50 seconds
	crontab.AddFunc("*/50 * * * * *", func() {
		log.Println("crontab 50s run")
	})
	// 定时
	min := now.Minute()+1
	hour := now.Hour()
	if min == 60{
		min = 0
		hour = hour+1
		if hour == 24{
			hour = 0
		}
	}
	spec := fmt.Sprintf("0 %d %d * * *",min, hour)
	log.Printf("It will run at: %s",spec)
	crontab.AddFunc(spec, func() {
		log.Printf("######crontab %s run#########\n", spec)
	})
	crontab.Start()
	<-stop
}
