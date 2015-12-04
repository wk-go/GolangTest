package main

import(
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	resp, err := http.Get("http://www.baidu.com")
	if err !=nil{
		log.Println(err)
	}
	fmt.Println(resp.Header)
	//var resp_body *[]byte
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}