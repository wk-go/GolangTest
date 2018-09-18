package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/jinzhu/gorm"
	"strings"
)

type Controller struct {
	Title string
	DB *gorm.DB
}

//group front
type FrontController struct {
	Controller
}

func (ctrl *FrontController) Index(c *gin.Context){
	c.HTML(200,"front/index",gin.H{"title":"Index","content":"This is index page."})
}
func (ctrl *FrontController) View(c *gin.Context){
	c.HTML(200, "front/view", gin.H{"title": "View", "content": "This is view page."})
}


//group admin
type AdminController struct {
	Controller
	Session sessions.Session
}
func (ctrl *AdminController) MiddleWarePrepare(c *gin.Context){
	ctrl.Session = sessions.Default(c)

	if !ctrl.isLogin(c) && c.Request.RequestURI != "/admin/login" && strings.Index(c.Request.RequestURI, "assets") == -1 {
		c.Redirect(302, "/admin/login")
		return
	}
	c.Next()
}
func (ctrl *AdminController) isLogin(c *gin.Context) bool{
	ctrl.Session = sessions.Default(c)
	if id,ok := ctrl.Session.Get("id").(uint); ok && id > 0{
		return true
	}
	return false
}

func (ctrl *AdminController) Index(c *gin.Context){
	c.HTML(200,"admin/index",gin.H{"title": "Gin Test"})
}
func (ctrl *AdminController) Statistics(c *gin.Context){
	c.String(200,"AdminStatistics")
}

func (ctrl *AdminController) SessionTest(c *gin.Context){
	session := ctrl.Session
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

func (ctrl * AdminController) Login(c *gin.Context){
	if ctrl.isLogin(c){
		c.Redirect(302,"/admin")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")

	if  c.Request.Method == "POST"{
		if username != "" && password != ""{
			user := &User{Username:username}
			ctrl.DB.First(&user)
			if user.ID != uint(0) && user.Password == password{
				session := ctrl.Session
				session.Set("id", user.ID)
				session.Set("username", user.Username)
				session.Save()
				c.Redirect(302, "/admin")
				//c.JSON(200, gin.H{"code": 1, "message": "登录成功"})
				return
			}
			//c.JSON(200,gin.H{"code": -1,"message": "登录失败"})
			//return
		}
	}
	c.HTML(200,"admin/login",gin.H{"title": "Login", "username": username, "password": password})
}
func (ctrl *AdminController) Logout(c *gin.Context){
	session := ctrl.Session
	session.Clear()
	session.Save()
	//c.JSON(200,gin.H{"message": "Logout Success"})
	c.Redirect(302, "login")
}
/*************** Article start *********************/
type ArticleController struct{
	AdminController
}
func (ctrl *ArticleController) Index(){

}

func (ctrl *ArticleController) Create(){

}

func (ctrl *ArticleController) Update(){

}

func (ctrl *ArticleController) Delete(){

}
/*************** Article  end  *********************/