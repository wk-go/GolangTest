package main

import (
	//"github.com/gin-gonic/gin"
	"reflect"
	"fmt"
)
type Router struct {

}
func (r *Router) Add(method,path string, handler interface{}){
	handlerType := reflect.TypeOf(handler)
	fmt.Println("handler type",handlerType)
}

func (r *Router) UrlTo(endpoint string, values ...interface{}) string{
	return ""
}