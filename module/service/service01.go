package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

//windows下需要管理员权限
var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	log.Println("hello")
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "xxxGoServiceTestxxx",
		DisplayName: "xxx Go Service Test xxx",
		Description: "xxx This is a test Go service. xxx",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			if err = s.Install(); err != nil {
				log.Println(err)
			} else {
				log.Println("服务安装成功")
			}
			return
		}

		if os.Args[1] == "remove" {
			if err = s.Uninstall(); err != nil {
				log.Println(err)
			} else {
				log.Println("服务卸载成功")
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
