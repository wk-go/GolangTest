package main

import (
    "fmt"
    "net"
    "runtime"
    //"time"
    "flag"
    "os"
    //"github.com/wkServerService/asocks-go/src/asocks"
    "io"
    "net/http"
    //"bytes"
    //"mime/multipart"
    "encoding/base64"
    //"log"
    "io/ioutil"
    "strings"
)

func handleConnection(conn *net.TCPConn) (err error) {
    closed := false

    defer func() {
        if !closed {
            conn.Close()
        }
    }()

    var reqLen int
    buf := make([]byte, 2048)
    reqContent := make([]byte, 0, 8)


    reqLen, err = io.ReadAtLeast(conn, buf, 2)
    if err == io.EOF {
        fmt.Println(":::reading err:::",err)
    }
    //fmt.Println(":::reqLen:::", reqLen)
    //fmt.Println(":::buf:::", buf)
    reqContent=buf[:reqLen]

    //fmt.Println(":::reqLen:::", reqLen)
    //fmt.Println(":::reqContent:::", reqContent)
    //fmt.Println(":::reqContent:::", string(reqContent))
    //conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello world!"))
    //var cliReq *http.Request
    var resp *http.Response

    finalURL := fmt.Sprintf("http://%v", serverAddr, )
    fmt.Println(":::url:::", finalURL)

    bodyStr := fmt.Sprintf("data=%s", base64.StdEncoding.EncodeToString(reqContent))
    fmt.Println(":::bodyStr:::", bodyStr)
    bodyByte := strings.NewReader(base64.StdEncoding.EncodeToString(reqContent))
    fmt.Println(":::bodyByte:::", bodyByte)
    /*cliReq, err = http.NewRequest("POST", finalURL, bodyByte)
    cliReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    fmt.Println(":::cliReq:::", cliReq)
    if err == nil {
        resp, err = http.DefaultClient.Do(cliReq)
        defer resp.Body.Close()
    }*/
    resp, err = http.Post(finalURL,
        "application/x-www-form-urlencoded",
        bodyByte)
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    fmt.Println(":::resp:::", resp)

    respBody,_ := ioutil.ReadAll(resp.Body)
    fmt.Println(":::body:::", string(respBody))
    conn.Write(respBody)
    return nil
}

func encodeData(data []byte) {
    for i, _ := range data {
        data[i] ^= 100
    }
}

func printUsage() {
    fmt.Printf("Usage:%s -s server_addr:server_port -l local_addr:local_port\n", os.Args[0])
}

var localAddr string
var serverAddr string
var server net.TCPAddr

func main() {
    flag.StringVar(&localAddr, "l", "127.0.0.1:8888", "本地监听IP:端口")
    flag.StringVar(&serverAddr, "s", "127.0.0.1:8080", "服务器IP:端口")
    flag.Parse()

    if serverAddr == "" {
        printUsage()
        return
    }

    i, err := net.ResolveTCPAddr("tcp", serverAddr)
    if err != nil {
        fmt.Println("resolve ", serverAddr, " failed. err:", err)
        return
    }
    server = *i

    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)

    bindAddr, _ := net.ResolveTCPAddr("tcp", localAddr)
    ln, err := net.ListenTCP("tcp", bindAddr)
    if err != nil {
        fmt.Println("listen error:", err)
        return
    }
    defer ln.Close()

    fmt.Println("listening ", ln.Addr())
    fmt.Println("server:", server.String())

    for {
        conn, err := ln.AcceptTCP()
        if err != nil {
            fmt.Println("accept error:", err)
            continue
        }

        go handleConnection(conn)
    }
}
