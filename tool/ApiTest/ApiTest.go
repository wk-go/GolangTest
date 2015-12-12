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
	"/api_conf":api_conf,
	"/api_group":api_group,
	"/api_item":api_item,
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
		api_id string
	)
	if val, ok := data["api_id"]; len(val) > 0 && ok {
		api_id = val[0]
	}
	if api_id == "" {
		log.Println("post: api_id is empty")
		return
	}

	cnt_map, err = c.ReadApiConf(api_id)
	if err == nil {
		for key, val := range CvtMapIntf(data) {
			cnt_map[key] = val
		}
	}else {
		cnt_map = CvtMapIntf(data)
		log.Println(err)
	}
	if len(cnt_map) <= 0 {
		log.Println("post: nothing to save")
		return
	}
	c.WriteApiConf(api_id, cnt_map)
}
//保存分组信息到配置文件
func (c *ApiConfig)SaveGroup(data map[string][]string) {
	var (
		conf_map map[string]interface{}
		group_map map[string]interface{}
		err error
		api_id string
		group_id string
	)

	if val, ok := data["api_id"]; len(val) > 0 && ok {
		api_id = val[0]
	}
	if val, ok := data["group_id"]; len(val) > 0 && ok {
		group_id = val[0]
		//delete(data,"group_id")
	}
	if api_id == "" {
		log.Println("post: api_id is empty")
		return
	}

	conf_map, err = c.ReadApiConf(api_id)



	var group map[string]interface{}
	if _, ok := conf_map["group"]; !ok {
		group = make(map[string]interface{})
	}else {
		group, _ = conf_map["group"].(map[string]interface{})
	}

	if err == nil {
		if _, ok := group[group_id]; ok {
			group_map = group[group_id].(map[string]interface{})
		}else {
			group_map = make(map[string]interface{})
		}
		t_data := CvtMapIntf(data)
		group_map["name"] = t_data["group_name"]
	}
	if len(conf_map) <= 0 {
		log.Println("post: nothing to save")
		return
	}
	group[group_id] = group_map
	conf_map["group"] = group

	c.WriteApiConf(api_id, conf_map)
}
//保存分组项目信息到配置文件
func (c *ApiConfig)SaveItem(data map[string][]string) {
	var (
		conf_map map[string]interface{}
		group_map map[string]interface{}
		//group_item map[string]interface{}
		err error
		api_id string
		group_id string
	)

	if val, ok := data["api_id"]; len(val) > 0 && ok {
		api_id = val[0]
	}
	if val, ok := data["group_id"]; len(val) > 0 && ok {
		group_id = val[0]
		//delete(data,"group_id")
	}
	if api_id == "" {
		log.Println("post: api_id is empty")
		return
	}

	conf_map, err = c.ReadApiConf(api_id)



	var group map[string]interface{}
	if _, ok := conf_map["group"]; !ok {
		group = make(map[string]interface{})
	}else {
		group, _ = conf_map["group"].(map[string]interface{})
	}

	if err == nil {
		if _, ok := group[group_id]; ok {
			group_map = group[group_id].(map[string]interface{})
		}else {
			group_map = make(map[string]interface{})
		}
		t_data := CvtMapIntf(data)
		group_map["name"] = t_data["group_name"]
	}
	if len(conf_map) <= 0 {
		log.Println("post: nothing to save")
		return
	}
	group[group_id] = group_map
	conf_map["group"] = group

	c.WriteApiConf(api_id, conf_map)
}
//读取配置文件内容
func (c *ApiConfig)ReadApiConf(api_id string) (map[string]interface{}, error) {
	var cnt_map_itf map[string]interface{}
	api_conf, _ := Config["api"].(map[string]string)
	filename := fmt.Sprintf("%v/%v.json", api_conf["path"], api_id)
	content, err := ioutil.ReadFile(filename)
	//log.Printf("The string content of file %v:%v\n",api_id,string(content))
	json.Unmarshal(content, &cnt_map_itf)
	//log.Printf("The content of file %v:%v\n",api_id,cnt_map_itf)
	//cnt_map := CvtMapStr(cnt_map_itf)
	log.Printf("Reading the api config: %v\n", api_id)
	return cnt_map_itf, err
}
//将配置信息写入配置文件
func (c *ApiConfig)WriteApiConf(api_id string, cnt_map map[string]interface{}) (error) {
	api_conf, _ := Config["api"].(map[string]string)
	filename := fmt.Sprintf("%v/%v.json", api_conf["path"], api_id)
	content, _ := json.Marshal(cnt_map)
	log.Printf("Update the api config: %v\n", api_id)
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
		if val, ok := val.(string); ok {
			data[key] = []string{val}
			continue
		}
		if val, ok := val.([]string); ok {
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
	if act, ok := req.Form["act"]; ok {
		conf := &ApiConfig{}

		if len(req.PostForm) > 0 {
			if _, err := conf.ReadApiConf(req.FormValue("api_id")); act[0] == "add" && err == nil {
				http.Redirect(w, req, fmt.Sprintf("/api_conf?act=edit&api=%v", req.FormValue("api_id")), 302);
			}
			conf.SaveConfig(req.PostForm)
			http.Redirect(w, req, fmt.Sprintf("/api_conf?act=edit&api=%v", req.FormValue("api_id")), 302);
		}
		if act[0] == "edit" && req.FormValue("api") != "" && len(req.PostForm) <= 0 {
			conf_data, err := conf.ReadApiConf(req.FormValue("api"))
			if err == nil {
				for _, key := range []string{"api_host", "api_description", "api_id", "api_name"} {
					str, _ := conf_data[key].(string)
					req.PostForm[key] = []string{str}
				}
			}
			edit = true
		}
	} else {
		http.Redirect(w, req, "/api_conf?act=add", 301)
	}
	log.Println(req.PostForm)
	log.Printf("api_conf render\n")
	RenderView(w, "add", map[string]interface{}{"req":req, "edit":edit})
	log.Printf("api_conf handle end\n")
}

//添加配置项分组
func api_group(w http.ResponseWriter, req *http.Request) {
	edit := false
	req.ParseForm()
	log.Printf("api_group handle start\n")
	if act, ok := req.Form["act"]; ok {
		conf := &ApiConfig{}

		if len(req.PostForm) > 0 {
			conf.SaveGroup(req.PostForm)
			http.Redirect(w, req, fmt.Sprintf("/api_group?act=edit&api=%v&group=%v", req.FormValue("api_id"), req.FormValue("group_id")), 302);
		}

		if _, err := conf.ReadApiConf(req.FormValue("api")); act[0] == "add" && err == nil {
			conf_data, err := conf.ReadApiConf(req.FormValue("api"))
			if err == nil {
				for _, key := range []string{"api_id", "group_name", "group_id"} {
					if _, ok := req.PostForm[key]; !ok {
						if str, ok := conf_data[key].(string); ok {
							req.PostForm[key] = []string{str}
						}else {
							req.PostForm[key] = []string{""}
						}
					}
				}
			}
		}else {
			log.Println(err)
		}
		log.Println(req.PostForm)

		if act[0] == "edit" && req.FormValue("api") != "" && len(req.PostForm) <= 0 {
			if req.FormValue("group") != "" {
				req.PostForm["api_id"] = []string{req.FormValue("api")}
				conf_data, err := conf.ReadApiConf(req.FormValue("api"))
				if err == nil {
					if group, ok := conf_data["group"]; ok {
						if group, ok := group.(map[string]interface{}); ok {
							req.PostForm["group_id"] = []string{req.FormValue("group")}
							if gTmp, ok := group[req.FormValue("group")].(map[string]interface{}); ok {
								if str, ok := gTmp["name"].(string); ok {
									req.PostForm["group_name"] = []string{str}
								}
							}else {
								req.PostForm["group_name"] = []string{""}
							}
						}
					}
				}
				edit = true
			}else {
				http.Redirect(w, req, fmt.Sprintf("/api_group?act=add&api=%v", req.FormValue("api")), 302);
			}
		}
	} else {
		http.Redirect(w, req, "/", 302)
	}
	//log.Println("req.PostForm:", req.PostForm)
	log.Printf("api_group render\n")
	RenderView(w, "api_group", map[string]interface{}{"req":req, "edit":edit})
	log.Printf("api_group handle end\n")
}
//添加配置项
func api_item(w http.ResponseWriter, req *http.Request) {
	edit := false
	req.ParseForm()
	log.Printf("api_item handle start\n")
	if act, ok := req.Form["act"]; ok {
		conf := &ApiConfig{}

		if len(req.PostForm) > 0 {
			conf.SaveItem(req.PostForm)
			//http.Redirect(w, req, fmt.Sprintf("/api_item?act=edit&api=%v&group=%v&item=%v", req.FormValue("api_id"), req.FormValue("group_id"),, req.FormValue("item_id")), 302);
		}

		if _, err := conf.ReadApiConf(req.FormValue("api")); act[0] == "add" && err == nil {
			conf_data, err := conf.ReadApiConf(req.FormValue("api"))
			if err == nil {
				for _, key := range []string{"api_id", "group_name", "group_id"} {
					if _, ok := req.PostForm[key]; !ok {
						if str, ok := conf_data[key].(string); ok {
							req.PostForm[key] = []string{str}
						}else {
							req.PostForm[key] = []string{""}
						}
					}
				}
			}
		}else {
			log.Println(err)
		}
		log.Println(req.PostForm)

		if act[0] == "edit" && req.FormValue("api") != "" && req.FormValue("group") != "" && len(req.PostForm) <= 0 {
			if req.FormValue("item") != "" {
				req.PostForm["api_id"] = []string{req.FormValue("api")}
				conf_data, err := conf.ReadApiConf(req.FormValue("api"))
				if err == nil {
					if group, ok := conf_data["group"]; ok {
						if group, ok := group.(map[string]interface{}); ok {
							req.PostForm["group_id"] = []string{req.FormValue("group")}
							if gTmp, ok := group[req.FormValue("group")].(map[string]interface{}); ok {
								if str, ok := gTmp["name"].(string); ok {
									req.PostForm["group_name"] = []string{str}
								}
							}else {
								req.PostForm["group_name"] = []string{""}
							}
						}
					}
				}
				edit = true
			}else {
				//http.Redirect(w, req, fmt.Sprintf("/api_item?act=add&api=%v", req.FormValue("api")), 302);
			}
		}
	} else {
		//http.Redirect(w, req, "/", 302)
	}
	//log.Println("req.PostForm:", req.PostForm)
	log.Printf("api_item render\n")
	RenderView(w, "api_item", map[string]interface{}{"req":req, "edit":edit})
	log.Printf("api_item handle end\n")
}