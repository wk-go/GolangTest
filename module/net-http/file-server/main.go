package main

import (
	"flag"
	"net/http"
)

// 一个简易的http文件服务器

func main() {
	var listen string
	var path string
	flag.StringVar(&listen, "listen", ":8080", "listen address")
	flag.StringVar(&path, "path", "./", "path")
	flag.Parse()
	println("正在监听地址: " + listen)
	http.Handle("/", http.FileServer(http.Dir(path)))
	if err := http.ListenAndServe(listen, nil); err != nil {
		panic(err)
	}
}
