package main

import "github.com/gin-gonic/gin"
type Router struct {

}
func (r *Router) Add(method,path string, endpoint gin.HandlerFunc){
}

func (r *Router) UrlTo(endpoint string, values ...interface{}) string{
	return ""
}