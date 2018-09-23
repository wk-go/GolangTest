package main

import (
	//"github.com/gin-gonic/gin"
	"reflect"
	"github.com/gin-gonic/gin"
)
type Router struct {
	Engine *gin.Engine
	Group *gin.RouterGroup
}
func (r *Router) Add(httpMethod,relativePath string, controller interface{}, method string){
	controllerType := reflect.TypeOf(controller)
	controllerMethod, flag := controllerType.MethodByName(method)
	if !flag {
		panic("router error")
	}
	handler := func(c *gin.Context){
		controllerMethod.Func.Call([]reflect.Value{reflect.ValueOf(controller),reflect.ValueOf(c)})
	}
	if r.Group != nil{
		r.Group.Handle(httpMethod,relativePath,handler)
		return
	}
	r.Engine.Handle(httpMethod,relativePath,handler)
}

func (r *Router) UrlTo(endpoint string, values ...interface{}) string{
	return ""
}