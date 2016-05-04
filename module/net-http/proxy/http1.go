package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "fmt"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(":::r:::", r)
    req, _ := http.NewRequest(r.Method, r.RequestURI, r.Body)
    fmt.Println(":::req1:::",req)
    for k, v := range r.Header {
        for _, vv := range v {
            req.Header.Add(k, vv)
        }
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

    fmt.Println(":::req2:::",req)
    client := &http.Client{}
    resp, err := client.Do(req)

    if err!=nil {
        fmt.Println("error:",err)
    }

    defer resp.Body.Close()


    for k, v := range resp.Header {
        for _, vv := range v {
            w.Header().Add(k, vv)
        }
    }

    for _, value := range resp.Request.Cookies() {
        w.Header().Add(value.Name,value.Value)
    }

    w.WriteHeader(resp.StatusCode)

    result, err := ioutil.ReadAll(resp.Body)
    if (err != nil) {
        fmt.Println("error:",err)
    }
    fmt.Println(":::result:::",string(result)[:100])
    _,err = w.Write(result)
    if (err != nil) {
        fmt.Println("error:",err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Start serving on port 8888")
    err := http.ListenAndServe(":8888", nil)
    if (err!=nil) {
        log.Println("error:",err)
    }
}