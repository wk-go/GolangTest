package main

import (
	"flag"
	"fmt"
)

// golang flag 库基本用法

// params: -listen=:8081 -username susan -age=22 -score 98

func main() {
	var conf Configure
	flag.StringVar(&conf.Listen, "listen", ":8080", "")
	username := flag.String("username", "", "username")
	flag.UintVar(&conf.Age, "age", 20, "age")
	flag.Float64Var(&conf.Score, "score", 100, "score")
	flag.Parse()
	conf.Username = *username
	fmt.Printf("%#v\n", conf)
}

type Configure struct {
	Listen   string
	Username string
	Age      uint
	Score    float64
}
