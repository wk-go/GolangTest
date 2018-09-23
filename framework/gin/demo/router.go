package main

import (
	//"github.com/gin-gonic/gin"
	"reflect"
	"fmt"
)
type Router struct {

}
func (r *Router) Add(method,path string, controller, handler interface{}){
	controllerType := reflect.TypeOf(controller)
	fmt.Println("controllerType:", controllerType)
	handlerType := reflect.TypeOf(handler)
	fmt.Println("handlerType.Kind():",handlerType.Kind())
	fmt.Println("handler type:",handlerType)
	handlerValue := reflect.TypeOf(handler)
	fmt.Println("handlerValue",handlerValue)
}

func (r *Router) UrlTo(endpoint string, values ...interface{}) string{
	return ""
}