package main

import "github.com/gin-gonic/gin"

type Controller struct {
	Title string
}

//group front
type FrontController struct {
	Controller
}

func (self *FrontController) Index(c *gin.Context){
	c.String(200,"FrontIndex!")
}


//group admin
type AdminController struct {
	Controller
}

func (self *AdminController) Index(c *gin.Context){
	c.String(200,"AdminIndex")
}
