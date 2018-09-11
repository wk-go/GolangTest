package main

import (
	"gopkg.in/robfig/cron.v2"
	"fmt"
	"log"
)

//定时任务
func main() {
	stop := make(chan bool, 1)
	timeZone := "Asia/Shanghai"
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
	crontab.AddFunc(fmt.Sprintf(specTpl, "* 07 13 * * *"), func() {
		log.Println("######crontab 11:21 run#########")
	})

	crontab.Start()
	<-stop
}
