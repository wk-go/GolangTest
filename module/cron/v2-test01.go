package main

import (
	"gopkg.in/robfig/cron.v2"
	"fmt"
	"log"
	"time"
)

//定时任务
func main() {
	stop := make(chan bool, 1)
	timeZone := "Asia/Shanghai"

	TimeLocation,_ := time.LoadLocation(timeZone)
	now := time.Now().In(TimeLocation)

	specTpl := "TZ=" + timeZone + " %s"

	crontab := cron.New()

	//every 20 seconds
	crontab.AddFunc(fmt.Sprintf(specTpl, "*/20 * * * * *"), func() {
		log.Println("crontab 20s run")
	})
	// every 50 seconds
	crontab.AddFunc(fmt.Sprintf(specTpl, "*/50 * * * * *"), func() {
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
	crontab.AddFunc(fmt.Sprintf(specTpl, spec), func() {
		log.Printf("######crontab %s run#########", spec)
	})

	crontab.Start()
	<-stop
}
