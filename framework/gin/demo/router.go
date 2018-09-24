package main

import (
	//"github.com/gin-gonic/gin"
	"reflect"
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
)
type RouterInfo struct{
	HttpMethod []string
	RelativePath string
	Endpoint string
}
type Router struct {
	Engine *gin.Engine
	Group *gin.RouterGroup
	routerList map[string]*RouterInfo
}
func NewRouter(engine *gin.Engine) *Router{
	return &Router{
		Engine: engine,
		routerList: make(map[string]*RouterInfo),
	}
}
func (r *Router) Add(httpMethod,relativePath string, controller interface{}, method string){
	controllerType := reflect.TypeOf(controller)
	controllerMethod, flag := controllerType.MethodByName(method)
	endpoint := strings.Replace(controllerType.String()+"."+method, "*", "", -1)
	if !flag {
		panic("router error")
	}
	handler := func(c *gin.Context){
		controllerMethod.Func.Call([]reflect.Value{reflect.ValueOf(controller),reflect.ValueOf(c)})
	}
	if r.Group != nil{
		r.Group.Handle(httpMethod,relativePath,handler)
		relativePath = r.Group.BasePath() + relativePath
	}else{
		r.Engine.Handle(httpMethod,relativePath,handler)
	}
	routerInfo, ok :=  r.routerList[endpoint]
	//fmt.Println(routerInfo, ok)
	if !ok {
		routerInfo = &RouterInfo{HttpMethod:[]string{}, Endpoint:endpoint,RelativePath:relativePath}
	}
	flag = true
	for _,val := range routerInfo.HttpMethod{
		if val == httpMethod{
			flag = false
		}
	}
	if flag {
		routerInfo.HttpMethod = append(routerInfo.HttpMethod, httpMethod)
	}
	if r.routerList == nil{
		r.routerList = map[string]*RouterInfo{
			endpoint: routerInfo,
		}
	}else{
		r.routerList[endpoint] = routerInfo
	}
}

func (r *Router) UrlTo(endpoint string, values ...interface{}) string{
	routerInfo, ok := r.routerList[endpoint]
	if !ok {
		return ""
	}
	if len(values)%2 != 0 {
		//logs.Warn("urlfor params must key-value pair")
		return ""
	}
	params := make(map[string]string)
	if len(values) > 0 {
		key := ""
		for k, v := range values {
			if k%2 == 0 {
				key = fmt.Sprint(v)
			} else {
				params[key] = fmt.Sprint(v)
			}
		}
	}
	urlResult := Strtr(routerInfo.RelativePath, params)
	queryStr := ""
	separator :=""
	for key, val := range params{
		if strings.Index(key, ":") != -1{
			continue
		}
		queryStr += separator + key + "=" + val
		separator = "&"
	}
	if len(queryStr) > 0{
		return urlResult + "?" + queryStr
	}
	return urlResult
}