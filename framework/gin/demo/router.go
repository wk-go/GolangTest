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
func (r *Router) Add(httpMethod,relativePath string, controller, handler interface{}){
	controllerType := reflect.TypeOf(controller)
	controllerMethod, flag := controllerType.MethodByName("TestRouter")
	if !flag {
		panic("router error")
	}
	if r.Group != nil{
		r.Group.Handle(httpMethod,relativePath,func(c *gin.Context){
			controllerMethod.Func.Call([]reflect.Value{reflect.ValueOf(controller),reflect.ValueOf(c)})
		})
		return
	}
	r.Engine.Handle(httpMethod,relativePath,func(c *gin.Context){
		controllerMethod.Func.Call([]reflect.Value{reflect.ValueOf(controller),reflect.ValueOf(c)})
	})
}

func (r *Router) UrlTo(endpoint string, values ...interface{}) string{
	return ""
}