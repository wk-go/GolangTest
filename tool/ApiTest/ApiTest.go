//api test
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"html/template"
	"io"
	"encoding/json"
	"strings"
	"strconv"
	"path/filepath"
	"mime/multipart"
	"bytes"
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
	"/api_show":api_show,
	"/test_api":test_api,
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
		confMap map[string]interface{}
		groupMap map[string]interface{}
		itemMap map[string]interface{}
		items map[string]interface{}
		itemId string
		err error
		apiID string
		groupID string
	)

	if val, ok := data["api_id"]; len(val) > 0 && ok {
		apiID = val[0]
	}
	if val, ok := data["group_id"]; len(val) > 0 && ok {
		groupID = val[0]
		//delete(data,"group_id")
	}
	if apiID == "" {
		log.Println("post: api_id is empty")
		return
	}

	confMap, err = c.ReadApiConf(apiID)



	var group map[string]interface{}
	if _, ok := confMap["group"]; !ok {
		group = make(map[string]interface{})
	}else {
		group, _ = confMap["group"].(map[string]interface{})
	}

	if err == nil {
		if _, ok := group[groupID]; ok {
			groupMap = group[groupID].(map[string]interface{})
		}else {
			groupMap = make(map[string]interface{})
		}
		t_data := CvtMapIntf(data)
		itemMap = make(map[string]interface{})
		itemMap["name"] = t_data["item_name"]
		itemMap["url"] = t_data["item_url"]
		if str, ok := itemMap["url"].(string); ok {
			itemId = strings.Replace(str, "/", "_", -1);
		}
		itemMap["dataType"] = t_data["item_dataType"]
		methodType := map[string][]interface{}{
			"get":make([]interface{}, 0, 40),
			"post":make([]interface{}, 0, 40),
		}
		//solve post and get params
		for method,_ := range methodType {
			for i := 0; true; i++ {
				var tmp = make(map[string]interface{})
				if _, ok := t_data[fmt.Sprintf("%v[%v][field]",method, i)]; ok {
					for _,sub_field := range []string{"label","field","type","value","des"} {
						tmp[sub_field] = t_data[fmt.Sprintf("%v[%v][%v]", method, i,sub_field)]
					}
					methodType[method] = append(methodType[method], tmp)
				}else {
					break
				}
			}
			itemMap[fmt.Sprintf("%vField",method)] = methodType[method];
		}

		if _, ok := groupMap["items"]; !ok {
			items = make(map[string]interface{})
		}else {
			items, _ = groupMap["items"].(map[string]interface{})
		}
		fmt.Printf("items:%v;item len:%v\n", items, len(items))
		items[itemId] = itemMap
	}
	if len(confMap) <= 0 {
		log.Println("post: nothing to save")
		return
	}
	groupMap["items"] = items
	group[groupID] = groupMap
	confMap["group"] = group

	c.WriteApiConf(apiID, confMap)
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
	log.Printf("api_conf handle start\n")
	req.ParseForm()
	fmt.Println(":::form:::",req.Form)
	act :=req.FormValue("act")
	fmt.Println(":::act:::",act)
	if act != "" {
		conf := &ApiConfig{}

		if len(req.PostForm) > 0 {
			if _, err := conf.ReadApiConf(req.FormValue("api_id")); act == "add" && err == nil {
				http.Redirect(w, req, fmt.Sprintf("/api_conf?act=edit&api=%v", req.FormValue("api_id")), 302);
			}
			conf.SaveConfig(req.PostForm)
			http.Redirect(w, req, fmt.Sprintf("/api_conf?act=edit&api=%v", req.FormValue("api_id")), 302);
		}
		if act == "edit" && req.FormValue("api") != "" && len(req.PostForm) <= 0 {
			conf_data, err := conf.ReadApiConf(req.FormValue("api"))
			if err == nil {
				for _, key := range []string{"api_host", "api_description", "api_id", "api_name"} {
					str, _ := conf_data[key].(string)
					req.PostForm[key] = []string{str}
				}
			}
			edit = true
		}

		log.Printf("api_conf render\n")
		RenderView(w, "api_conf_add", map[string]interface{}{"req":req, "edit":edit})
	} else {//没有参数显示列表
		log.Println("api_conf show list")
		api_setting, _ := Config["api"].(map[string]string)
		fmt.Println(":::api_setting:::", api_setting)
		apiConf := new(ApiConfig)
		apiMap := make(map[string]interface{})
		if files, err := filepath.Glob(fmt.Sprintf("%v/*.json", api_setting["path"])); err == nil {
			fmt.Println(":::files:::", files)
			for _, filename := range files {
				fmt.Println(":::file:::", filename)
				apiId := filename[strings.LastIndex(filename, string(filepath.Separator)) + 1:strings.LastIndex(filename, ".json")]
				fmt.Println(":::api_id:::", apiId)
				apiInfo := make(map[string]interface{})
				if conf, err := apiConf.ReadApiConf(apiId); err == nil {
					for _, key := range []string{"api_id", "api_name", "api_host", "api_description"} {
						apiInfo[key] = conf[key]
					}
				}else {
					log.Println("api_show:", err)
				}
				apiMap[apiId] = apiInfo
			}
		}else {
			log.Println(err)
		}

		log.Printf("api_conf render\n")
		RenderView(w, "api_conf_list", map[string]interface{}{"req":req, "apiMap":apiMap})
	}
	log.Println(req.PostForm)
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
		log.Printf("api_group render\n")
		RenderView(w, "api_group", map[string]interface{}{"req":req, "edit":edit})
	} else {//分组列表
		log.Printf("api_group group list\n")
	}
	//log.Println("req.PostForm:", req.PostForm)
	log.Printf("api_group handle end\n")
}
//api_item start ------------------------------------------------------------------------
type TplWriter struct {
	content []byte
}
func (t *TplWriter)Write(b []byte) (n int, err error) {
	if t.content == nil {
		t.content = make([]byte, len(b))
	}
	for i := 0; i < len(b); i++ {
		t.content = append(t.content, b[i])
	}
	return len(b), nil
}
func (t *TplWriter)ReadContent() (cnt string, n int) {
	if len(t.content) <= 0 {
		return "", 0
	}
	cnt = string(t.content)
	t.content = make([]byte, 0)
	return cnt, len(cnt)
}
func _item_field(data map[string][]string) (post template.HTML, get template.HTML, postLen int, getLen int) {
	valTpl := `<div class="form-group field-config">
                <label class="col-sm-2 control-label">{{.methodName}}参数</label>
                <div class="col-sm-10">
                    <input name="{{.method}}[{{.idx}}][label]" value="{{.label}}" placeholder="标签" class="form-control col-sm-2">
                    <input name="{{.method}}[{{.idx}}][field]" value="{{.field}}" placeholder="字段" class="form-control col-sm-2">
                    <select name="{{.method}}[{{.idx}}][type]" class="form-control col-sm-2">
                        <option value="text"{{ if eq .type "text"}} selected{{end}}>text</option>
                        <option value="password"{{ if eq .type "password"}} selected{{end}}>password</option>
                        <option value="textarea"{{if eq .type "textarea"}} selected{{end}}>textarea</option>
                        <option value="select"{{if eq .type "select"}} selected{{end}}>select</option>
                        <!--<option value="checkbox"{{if eq .type "checkbox"}} selected{{end}}>checkbox</option>
                        <option value="radio"{{if eq .type "radio"}} selected{{end}}>radio</option>-->
                        <option value="file"{{if eq .type "file"}} selected{{end}}>file</option>
                    </select>
                    <input name="{{.method}}[{{.idx}}][value]" value="{{.value}}" placeholder="目标值" class="form-control col-sm-2">
                    <input name="{{.method}}[{{.idx}}][des]" value="{{.des}}" placeholder="描述" class="form-control col-sm-2">
                </div>
            </div>`
	field := map[string]string{
		"post":"",
		"get":"",
	}
	fieldLen := map[string]int{
		"post":0,
		"get":0,
	}
	tplWriter := new(TplWriter)

	for method, _ := range field {
		var tmp = make(map[string]string)
		tmp["method"] = method
		tmp["methodName"] = strings.ToUpper(method)
		for i := 0; true; i++ {
			fieldLen[method] = i
			tmp["idx"] = strconv.Itoa(i)
			if _, ok := data[fmt.Sprintf("%v[%v][field]", method, i)]; ok {
				for _,sub_field := range []string{"label","field","type","value","des"} {
					tmp[sub_field] = data[fmt.Sprintf("%v[%v][%v]", method, i,sub_field)][0]
				}
				fmt.Println(":::tmp:::",tmp)
				tmpl, err := template.New("val").Parse(valTpl)
				if err != nil { panic(err) }
				err = tmpl.Execute(tplWriter, tmp)
				if err != nil { panic(err) }
				if tplStr, l := tplWriter.ReadContent(); l > 0 {
					field[method] = strings.Join([]string{field[method], tplStr}, "")
				}
			}else {
				break
			}
		}
	}
	return template.HTML(field["post"]), template.HTML(field["get"]), fieldLen["post"], fieldLen["get"]
}

//添加配置项
func api_item(w http.ResponseWriter, req *http.Request) {
	edit := false
	req.ParseForm()
	log.Printf("api_item handle start\n")
	log.Println("req.From:", req.Form)
	if act, ok := req.Form["act"]; ok {
		conf := &ApiConfig{}

		if len(req.PostForm) > 0 {
			conf.SaveItem(req.PostForm)
			if act[0] == "add" {
				item_id := req.FormValue("item")
				if item_id == "" {
					item_id = strings.Replace(req.PostFormValue("item_url"), "/", "_", -1)
				}
				http.Redirect(w, req, fmt.Sprintf("/api_item?act=edit&api=%v&group=%v&item=%v", req.FormValue("api_id"), req.FormValue("group_id"), item_id), 302);
			}
		}

		if _, err := conf.ReadApiConf(req.FormValue("api")); act[0] == "add" && err == nil {
			conf_data, err := conf.ReadApiConf(req.FormValue("api"))
			if err == nil {
				for _, key := range []string{"api_id", "group_id", "item_id", "item_name", "item_url", "item_dataType"} {
					if key == "group_id" {
						req.PostForm[key] = []string{req.FormValue("group")}
						continue
					}
					if _, ok := req.Form[key]; !ok {
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

		if act[0] == "edit" && req.FormValue("api") != "" && req.FormValue("group") != "" && len(req.PostForm) <= 0 {
			if req.FormValue("item") != "" {
				req.PostForm["api_id"] = []string{req.FormValue("api")}
				conf_data, err := conf.ReadApiConf(req.FormValue("api"))
				if err == nil {
					if group, ok := conf_data["group"]; ok {
						if group, ok := group.(map[string]interface{}); ok {
							req.PostForm["group_id"] = []string{req.FormValue("group")}
							if gTmp, ok := group[req.FormValue("group")].(map[string]interface{}); ok {
								itemsMap, _ := gTmp["items"].(map[string]interface{})
								if itemMap, ok := itemsMap[req.FormValue("item")].(map[string]interface{}); ok {
									var methodMap = map[string]string{
										"getField":"get",
										"postField":"post",
									}
									for key, val := range itemMap {
										if _, ok := methodMap[key]; ok {
											field, _ := val.([]interface{})
											for k, _ := range field {
												field_map, _ := field[k].(map[string]interface{})
												for k_field, v := range field_map {
													if str, ok := v.(string); ok {
														req.PostForm[fmt.Sprintf("%v[%v][%v]", methodMap[key], k, k_field)] = []string{str}
													}else {
														req.PostForm[fmt.Sprintf("%v[%v][%v]", methodMap[key], k, k_field)] = []string{""}
													}
												}
											}
											continue
										}
										if str, ok := val.(string); ok {
											req.PostForm[fmt.Sprintf("item_%v", key)] = []string{str}
										}
									}
								}else {
									req.PostForm["group_name"] = []string{""}
								}
							}
						}
					}
					edit = true
				}
			}else {
				http.Redirect(w, req, fmt.Sprintf("/api_item?act=add&api=%v&group=%v", req.FormValue("api"), req.FormValue("group")), 302);
			}
		}
	} else {
		//http.Redirect(w, req, "/", 302)
	}
	postItem, getItem, postLen, getLen := _item_field(req.PostForm)
	//fmt.Printf(postItem, getItem)
	log.Println("req.PostForm:", req.PostForm)
	log.Printf("api_item render\n")
	RenderView(w, "api_item", map[string]interface{}{
		"req":req,
		"edit":edit,
		"postItem":postItem,
		"getItem":getItem,
		"postLen":postLen,
		"getLen":getLen,
	})
	log.Printf("api_item handle end\n")

}
//api_item end ------------------------------------------------------------------------
//api_show start ------------------------------------------------------------------------
//显示接口信息
func api_show(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Printf("api_item handle start\n")
	fmt.Println(":::req.From:::", req.Form)
	if act := req.FormValue("act"); act == "" {
		data := make(map[string]interface{})
		api_setting, _ := Config["api"].(map[string]string)
		fmt.Println(":::api_setting:::", api_setting)
		apiConf := new(ApiConfig)
		apiMap := make(map[string]interface{})
		if files, err := filepath.Glob(fmt.Sprintf("%v/*.json", api_setting["path"])); err == nil {
			fmt.Println(":::files:::", files)
			for _, filename := range files {
				fmt.Println(":::file:::", filename)
				apiId := filename[strings.LastIndex(filename, string(filepath.Separator)) + 1:strings.LastIndex(filename, ".json")]
				fmt.Println(":::api_id:::", apiId)
				apiInfo := make(map[string]interface{})
				if conf, err := apiConf.ReadApiConf(apiId); err == nil {
					for _, key := range []string{"api_id", "api_name", "api_host", "api_description"} {
						apiInfo[key] = conf[key]
					}
				}else {
					log.Println("api_show:", err)
				}

				apiMap[apiId] = apiInfo
			}
		}else {
			log.Println(err)
		}
		data["apiList"] = apiMap
		if apiListJson, err := json.Marshal(apiMap); err == nil {
			data["apiListJson"] = template.JS(string(apiListJson))
		}
		RenderView(w, "api_show", data)
	}else {
		if act == "get_group" {
			apiConf := new(ApiConfig)
			if apiId := req.Form.Get("api_id"); apiId != "" {
				if conf, err := apiConf.ReadApiConf(apiId); err == nil {
					if conf, err := json.Marshal(conf["group"]); err == nil {
						fmt.Fprint(w, string(conf))
					}
				}else {
					log.Println("api_show:", err)
				}
			}else {

			}
		}
	}
}
//api_show end ------------------------------------------------------------------------
//solve test form
func test_api(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Printf("test_api handle start\n")
	fmt.Println(":::form1:::", req.Form)


	fmt.Println(":::form1:::", req.PostFormValue)
	targetHOST := req.PostFormValue("host")
	targetURL := req.PostFormValue("url")
	fmt.Println(":::host:::", targetHOST)
	fmt.Println(":::url:::", targetURL)
	if targetURL[0:1] != "/" {
		targetURL = strings.Join([]string{"/", targetURL}, "")
	}

	//solve post and get parameters
	postParam, getParam := url.Values{}, url.Values{}
	for key, _ := range req.Form {
		keySplit := strings.Split(key, ":")
		fmt.Println(":::keySplit param:::", keySplit)
		if len(keySplit) >= 2 && keySplit[1] != "" {
			if keySplit[0] == "post" {
				postParam.Add(keySplit[1], req.PostFormValue(key))
			}else {
				getParam.Add(keySplit[1], req.PostFormValue(key))
			}
		}
	}
	finalURL := fmt.Sprintf("http://%v%v", targetHOST, targetURL)
	if len(getParam) > 0 {
		finalURL = fmt.Sprintf("http://%v?%v", finalURL, getParam.Encode())
	}

	var resp *http.Response
	var respErr error
	if len(postParam) > 0 {
		fmt.Println(":::url:::", finalURL)
		/*file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()*/

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		/*part, err := writer.CreateFormFile(paramName, path)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)*/

		for key, _ := range postParam {
			_ = writer.WriteField(key, postParam.Get(key))
		}
		err := writer.Close()
		if err != nil {
			return
		}
		request, err := http.NewRequest("POST", finalURL, body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		if err == nil {
			resp, respErr = http.DefaultClient.Do(request)
		}
	}else {
		resp, respErr = http.Get(finalURL)

	}

	if respErr != nil {
		log.Printf("test_api error:%v http.Get(%v)\n", respErr, finalURL)
	}else {
		fmt.Println(":::resp:::", resp)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(":::body:::", string(body))
			apiResp := make(map[string]interface{})

			apiResp["Status"] = resp.Status
			apiResp["StatusCode"] = resp.StatusCode
			apiResp["Proto"] = resp.Proto
			apiResp["Header"] = resp.Header
			apiResp["Body"] = string(body)
			apiResp["ContentLength"] = resp.ContentLength
			//fmt.Fprint(w,string(body))
			if respData, err := json.Marshal(apiResp); err == nil {
				fmt.Fprint(w, string(respData))
			}
		}
	}

	if data, err := json.Marshal(req.Form); err == nil {
		fmt.Println(string(data))
	}
	log.Printf("test_api handle end\n")
}
