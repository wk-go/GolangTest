package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	Title string
	DB *gorm.DB
}

//group front
type FrontController struct {
	Controller
}

func (self *FrontController) Index(c *gin.Context){
	c.HTML(200,"front/index",gin.H{"title":"Index","content":"This is index page."})
}
func (self *FrontController) View(c *gin.Context){
	c.HTML(200, "front/view", gin.H{"title": "View", "content": "This is view page."})
}


//group admin
type AdminController struct {
	Controller
	Session sessions.Session
}
func (self *AdminController) MiddleWarePrepare(c *gin.Context){
	self.Session = sessions.Default(c)

	c.Next()
}
func (self *AdminController) isLogin(c *gin.Context) bool{
	self.Session = sessions.Default(c)
	if id,ok := self.Session.Get("id").(uint); ok && id > 0{
		return true
	}
	return false
}

func (self *AdminController) Index(c *gin.Context){
	c.HTML(200,"admin/index",gin.H{"title": "Gin Test"})
}
func (self *AdminController) Statistics(c *gin.Context){
	c.String(200,"AdminStatistics")
}

func (self *AdminController) SessionTest(c *gin.Context){
	session := self.Session
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

func (self * AdminController) Login(c *gin.Context){
	if self.isLogin(c){
		c.Redirect(302,"/admin")
		return
	}
	if  c.Request.Method == "POST"{
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "" && password != ""{
			user := &User{Username:username}
			self.DB.First(&user)
			if user.ID != uint(0) && user.Password == password{
				session := self.Session
				session.Set("id", user.ID)
				session.Set("username", user.Username)
				session.Save()
				c.JSON(200, gin.H{"code": 1, "message": "登录成功"})
				return
			}
			c.JSON(200,gin.H{"code": -1,"message": "登录失败"})
			return
		}
	}
	c.HTML(200,"admin/login",gin.H{"title": "Login"})
}
func (self *AdminController) Logout(c *gin.Context){
	session := self.Session
	session.Clear()
	session.Save()
	c.JSON(200,gin.H{"message": "Logout Success"})
}