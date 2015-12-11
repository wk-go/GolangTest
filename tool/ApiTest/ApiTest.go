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
	_"reflect"
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
	"/api_conf":api_conf,
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
		cnt_map map[string]interface{}
		err error
		api_identify string
	)
	if val, ok := data["api_identify"]; len(val) > 0 && ok {
		api_identify = val[0]
	}
	if api_identify == "" {
		log.Println("post: api_identify is empty")
		return
	}

	cnt_map, err = c.ReadApiConf(api_identify)
	if err == nil {
		for key, val := range CvtMapIntf(data){
			cnt_map[key]=val
		}
	}else {
		cnt_map=CvtMapIntf(data)
		log.Println(err)
	}
	if len(cnt_map) <= 0 {
		log.Println("post: nothing to save")
		return
	}
	c.WriteApiConf(api_identify,cnt_map)
}
//读取配置文件内容
func (c *ApiConfig)ReadApiConf(api_identify string) (map[string]interface{},error) {
	var cnt_map_itf map[string]interface{}
	api_conf, _ := Config["api"].(map[string]string)
	filename := fmt.Sprintf("%v/%v.json", api_conf["path"], api_identify)
	content, err := ioutil.ReadFile(filename)
	//log.Printf("The string content of file %v:%v\n",api_identify,string(content))
	json.Unmarshal(content, &cnt_map_itf)
	//log.Printf("The content of file %v:%v\n",api_identify,cnt_map_itf)
	//cnt_map := CvtMapStr(cnt_map_itf)
	log.Printf("Reading the api config: %v\n",api_identify)
	return cnt_map_itf,err
}
//将配置信息写入配置文件
func (c *ApiConfig)WriteApiConf(api_identify string,cnt_map map[string]interface{}) (error) {
	api_conf, _ := Config["api"].(map[string]string)
	filename := fmt.Sprintf("%v/%v.json", api_conf["path"], api_identify)
	content,_ :=json.Marshal(cnt_map)
	log.Printf("Update the api config: %v\n",api_identify)
	return ioutil.WriteFile(filename, content, 755)
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
	//log.Println("Data map:",formData)
	for key, val := range formData {
		//log.Printf("Convert map to string:[%v]=%v",key,val)
		if val, ok:= val.(string); ok {
			data[key] = []string{val}
			continue
		}
		if val, ok:= val.([]string); ok {
			data[key] = val
		}
	}
	//log.Println("Convert map to string map:",data)
	return data
}
//首页
func index(w http.ResponseWriter, req *http.Request) {
	RenderView(w, "index", map[string]interface{}{})
}

//添加新配置
func api_conf(w http.ResponseWriter, req *http.Request) {
	edit := false
	req.ParseForm()
	log.Printf("api_conf handle start\n")
	if act, ok := req.Form["act"];ok{
		conf := &ApiConfig{}

		if len(req.PostForm) > 0 {
			if _,err :=conf.ReadApiConf(req.FormValue("api_identify"));act[0] == "add" && err==nil {
				http.Redirect(w,req,fmt.Sprintf("/api_conf?act=edit&api=%v",req.FormValue("api_identify")),302);
			}
			conf.SaveConfig(req.PostForm)
			http.Redirect(w,req,fmt.Sprintf("/api_conf?act=edit&api=%v",req.FormValue("api_identify")),302);
		}
		if act[0] == "edit"{
			if api_identify, ok := req.Form["api"]; ok{
				if len(req.PostForm) <= 0{
					conf_data,err :=conf.ReadApiConf(api_identify[0])
					if err ==nil {
						for _, key := range []string{"api_host", "api_description", "api_identify", "api_name"} {
							str, _ := conf_data[key].(string)
							req.PostForm[key] = []string{str}
						}
					}
				}
			}
			edit = true
		}
	} else {
		http.Redirect(w, req, "/api_conf?act=add", 301)
	}
	log.Println(req.PostForm)
	log.Printf("api_conf render\n")
	RenderView(w, "add", map[string]interface{}{"req":req,"edit":edit})
	log.Printf("api_conf handle end\n")
}