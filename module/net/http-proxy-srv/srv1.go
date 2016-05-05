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
    //"io"
    "time"
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
    var host,port,err = net.SplitHostPort(realReq.Host)
    if err != nil{
        host=realReq.Host
        port="80"
    }
    fmt.Println(":::addrInfo:::", host, port)
    remoteSrv, err:= net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", host, port))
    if err != nil {
        fmt.Println("resolve ", realReq.Host, " failed. err:", err)
        return
    }
    remote, err := net.DialTCP("tcp", nil, remoteSrv)
    if err != nil {
        return
    }
    finishChannel := make(chan bool, 1)
    go remote.Write(bodyByte)
    /*var realResp = make([]byte, 204800)
    respLen, err :=io.ReadAtLeast(remote,realResp,16);
    fmt.Println(":::realResp[:respLen]:::", realResp[:respLen])
    w.Write(realResp[:respLen])*/
    pipeThenClose(remote,w, finishChannel)
    <-finishChannel
    remote.Close()
}

func pipeThenClose(src *net.TCPConn, dst http.ResponseWriter , finish chan bool) {
    defer func(){
        src.CloseRead()
        finish <- true
    }()

    buf := make([]byte, 5120)
    for {
        src.SetReadDeadline(time.Now().Add(60 * time.Second))
        n, err := src.Read(buf);
        if n > 0 {
            data := buf[0:n]
            //encodeData(data)
            if _, err := dst.Write(data); err != nil {
                break
            }
        }
        if err != nil {
            break
        }
    }
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
