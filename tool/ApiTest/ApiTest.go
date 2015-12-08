//api test
package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"html/template"
	"io"
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
//渲染视图
func RenderView(w io.Writer,viewName string, data map[string]interface{}){
	//绑定默认数据
	if _,ok :=data["appInfo"];!ok{
		data["appInfo"]=AppInfo
	}

	viewCnt := ReadView(viewName)
	tmpl,err := template.New(viewName).Parse(viewCnt)
	if err != nil { panic(err) }
	err = tmpl.Execute(w, data)
	if err != nil { panic(err) }
}
//首页
func index(w http.ResponseWriter, r *http.Request){
	RenderView(w,"index",map[string]interface{}{})
}