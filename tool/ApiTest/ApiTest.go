//api test
package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)
//app配置
var AppInfo = map[string]string{
	"app_name": "HTTP API Test",
	"version": "0.0.1",
	"descriptin":"用于快速测试Http接口",
	"assets":"./web/assets",
	"views":"./web/views",
}
//路由表
var UrlRoute = map[string]http.HandlerFunc{
	"/":index,
}
func main(){
	//设置静态文件
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(AppInfo["assets"]))))
	//绑定路由
	for key,value := range UrlRoute{
		http.Handle(key,value)
	}
	log.Println("Server is starting")
	log.Fatal(http.ListenAndServe("localhost:10000", nil))
}
//读取视图文件
func ReadView(viewName string)(string)  {
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(fmt.Sprintf("%v/%v.html",AppInfo["views"], viewName)); err!=nil{
		log.Println(err)
		return ""
	}
	return string(content)
}
//首页
func index(w http.ResponseWriter, r *http.Request){
	content :=ReadView("index")
	fmt.Fprint(w, content)
}