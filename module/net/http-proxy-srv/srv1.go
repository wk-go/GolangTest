package main

import (
    "fmt"
    "net/http"
    "log"
    "flag"
    "encoding/base64"
    "io/ioutil"
    "bufio"
    "strings"
    "net"
)


func handle(w http.ResponseWriter, r *http.Request){
    fmt.Println(":::r:::",r)
    rawReq, _ := base64.StdEncoding.DecodeString(r.PostForm.Get("data"))
    fmt.Println(":::r.PostForm:::",r.PostForm)
    fmt.Println(":::rawReq:::", rawReq)
    respBody,_ := ioutil.ReadAll(r.Body)
    fmt.Println(":::body:::", string(respBody))
    bodyByte,_ :=base64.StdEncoding.DecodeString(string(respBody))
    fmt.Println(":::bodyByte:::", string(bodyByte))
    realReq, _ := http.ReadRequest(bufio.NewReader(strings.NewReader(string(bodyByte))))
    fmt.Println(":::realReq.Host:::", realReq.Host)
    //fmt.Fprint(w,"HTTP/1.1 200 OK\r\n\r\nhello world!")
    w.Write([]byte("hello world!"))
}

//路由表
var UrlRoute = map[string]http.HandlerFunc{
    "/":handle,
}
var localAddr string
func main() {
    flag.StringVar(&localAddr, "l", "0.0.0.0:8080", "本地监听IP:端口")

    //绑定路由
    for key, value := range UrlRoute {
        http.Handle(key, value)
    }
    log.Printf("Server is starting:%s\n", localAddr)
    log.Fatal(http.ListenAndServe(localAddr, nil))
}


func encodeData(data []byte) {
    for i, _ := range data {
        data[i] ^= 100;
    }
}
