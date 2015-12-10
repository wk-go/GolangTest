//api test
package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"html/template"
	"io"
	"encoding/json"
)
//app配置
var AppInfo = map[string]string{
	"app_name": "HTTP API Test",
	"version": "0.0.1",
	"descriptin":"用于快速测试Http接口",
	"assets":"./web/assets",
	"views":"./web/views",
}
var Config = map[string]interface{}{
	"tpml":map[string]string{
		"header":"./web/views/header.html",
		"footer":"./web/views/footer.html",
	},
	"api":map[string]string{
		"path":"./api_config",
	},
}
//路由表
var UrlRoute = map[string]http.HandlerFunc{
	"/":index,
	"/add":add,
}
func main() {
	//设置静态文件
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(AppInfo["assets"]))))
	//绑定路由
	for key, value := range UrlRoute {
		http.Handle(key, value)
	}
	log.Println("Server is starting")
	log.Fatal(http.ListenAndServe("localhost:10000", nil))
}
//读取视图文件
func ReadView(viewName string) (string) {
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(fmt.Sprintf("%v/%v.html", AppInfo["views"], viewName)); err != nil {
		log.Println(err)
		return ""
	}
	return string(content)
}
//渲染视图
func RenderView(w io.Writer, viewName string, data map[string]interface{}) {
	//绑定默认数据
	if _, ok := data["appInfo"]; !ok {
		data["appInfo"] = AppInfo
	}
	if _, ok := data["config"]; !ok {
		data["config"] = Config
	}

	//渲染默认页头
	if _, ok := data["header"]; !ok {
		viewCnt := ReadView("header")
		tmpl, err := template.New(viewName).Parse(viewCnt)
		if err != nil { panic(err) }
		err = tmpl.Execute(w, data)
		if err != nil { panic(err) }
	}

	//渲染业务视图
	viewCnt := ReadView(viewName)
	tmpl, err := template.New(viewName).Parse(viewCnt)
	if err != nil { panic(err) }
	err = tmpl.Execute(w, data)
	if err != nil { panic(err) }

	//渲染默认页脚
	if _, ok := data["footer"]; !ok {
		viewCnt := ReadView("footer")
		tmpl, err := template.New(viewName).Parse(viewCnt)
		if err != nil { panic(err) }
		err = tmpl.Execute(w, data)
		if err != nil { panic(err) }
	}
}

//处理配置信息
type ApiConfig struct {

}
//处理从表单收到的配置数据
func (c *ApiConfig)HandleConfig(data map[string]interface{}) (map[string]interface{}) {
	return nil
}
//保存配置信息到文件
func (c *ApiConfig)SaveConfig(data map[string][]string) {
	var (
		content []byte
		cnt_map_itf map[string]interface{}
		cnt_map map[string][]string
		err error
		filename string
	)
	if val, ok := data["api_identify"]; len(val) > 0 && ok {
		api_conf, _ := Config["api"].(map[string]string)
		filename = fmt.Sprintf("%v/%v.json", api_conf["path"], val[0])

		content, err = ioutil.ReadFile(filename)
		json.Unmarshal(content, cnt_map_itf)
		cnt_map = CvtMapStr(cnt_map_itf)
		if err == nil {
			for key, val := range data{
				cnt_map[key]=val
			}
		}else {
			cnt_map=data
			log.Println(err)
		}

		content, err = json.Marshal(CvtMapIntf(cnt_map))
	}
	if filename == "" {
		return
	}

	if len(content) > 0 {
		ioutil.WriteFile(filename, content, 755)
	}
}
//转换:将map[string][]string转换为map[string]interface{}
func CvtMapIntf(formData map[string][]string) (map[string]interface{}) {
	var data = make(map[string]interface{})
	for key, val := range formData {
		if len(val) == 1 {
			data[key] = val[0]
			continue
		}
		data[key] = val
	}
	return data
}
//转换:将map[string]interface{}转换为map[string][]string
func CvtMapStr(formData map[string]interface{}) (map[string][]string) {
	var data = make(map[string][]string)
	for key, val := range formData {
		if val, ok:= val.(string); ok {
			data[key] = []string{val}
			continue
		}
		if val, ok:= val.([]string); ok {
			data[key] = val
		}
	}
	return data
}
//首页
func index(w http.ResponseWriter, req *http.Request) {
	RenderView(w, "index", map[string]interface{}{})
}

//添加新配置
func add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	conf := &ApiConfig{}
	if len(req.Form) > 0 {
		conf.SaveConfig(req.Form)
	}
	RenderView(w, "add", map[string]interface{}{"req":req})
}