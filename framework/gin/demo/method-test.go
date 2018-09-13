package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

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
func (self *FrontController) View(c *gin.Context){
	c.String(200,"FrontView!")
}


//group admin
type AdminController struct {
	Controller
}

func (self *AdminController) Index(c *gin.Context){
	c.String(200,"AdminIndex")
}
func (self *AdminController) Statistics(c *gin.Context){
	c.String(200,"AdminStatistics")
}

func (self *AdminController) SessionTest(c *gin.Context){
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil{
		count = 0
	} else {
		count = v.(int)
		count ++
	}
	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{"count": count})
}
