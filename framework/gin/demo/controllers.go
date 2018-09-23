package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/jinzhu/gorm"
	"strings"
	"net/http"
	"strconv"
	"fmt"
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
func (ctrl *AdminController) HTML(c *gin.Context, code int, name string, obj gin.H){
	obj["ctrl"] = ctrl

	session := sessions.Default(c)
	obj["session"] = session

	obj["username"] = session.Get("username")
	c.HTML(code, name, obj)
}

func (ctrl *AdminController) MiddleWareSurroundings(c *gin.Context){
	if !ctrl.isLogin(c) && c.Request.RequestURI != "/admin/login" && strings.Index(c.Request.RequestURI, "assets") == -1 {
		c.Redirect(302, "/admin/login")
		return
	}
	c.Next()
}
func (ctrl *AdminController) isLogin(c *gin.Context) bool{
	session := sessions.Default(c)
	if id,ok := session.Get("id").(uint); ok && id > 0{
		return true
	}
	return false
}

func (ctrl *AdminController) Index(c *gin.Context){
	ctrl.HTML(c,200,"admin/index",gin.H{"title": "Gin Test"})
}
func (ctrl *AdminController) Statistics(c *gin.Context){
	c.String(200,"AdminStatistics")
}

func (ctrl *AdminController) SessionTest(c *gin.Context){
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
			ctrl.DB.First(user)
			if user.ID != uint(0) && user.Password == password{
				session := sessions.Default(c)
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
	ctrl.HTML(c,200,"admin/login",gin.H{"title": "Login", "username": username, "password": password})
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
func (ctrl *ArticleController) Index(c *gin.Context){
	ctrl.Title = "文章列表"
	var (
		perPage  = 20
		count  = 0
		page int
		models []*Article
	)
	pageQuery := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageQuery)
	if err !=nil || page <= 0{
		page = 1
	}
	ctrl.DB.Offset((page-1)*perPage).Limit(perPage).Find(&models).Count(&count)

	ctrl.HTML(c,200,"admin/article-index", gin.H{"title": ctrl.Title, "models": models})
}

func (ctrl *ArticleController) Create(c *gin.Context){
	ctrl.Title = "添加文章"
	model := Article{}
	if c.Request.Method == "POST"{
		if err := c.Bind(&model); err == nil {
			ctrl.DB.Create(&model)
			c.Redirect(302, fmt.Sprintf("/admin/article/update/%d", model.ID))
			return
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	ctrl.HTML(c,200,"admin/article-edit", gin.H{"title": ctrl.Title, "model": model})
}

func (ctrl *ArticleController) Update(c *gin.Context){
	ctrl.Title = "更新"
	idInt,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		panic(err)
	}
	id := uint(idInt)
	model := ctrl.getModel(id)
	ctrl.HTML(c,200,"admin/article-edit", gin.H{"title": ctrl.Title, "model": model})
}

func (ctrl *ArticleController) Delete(c *gin.Context){
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil{
		panic(err)
		c.Redirect(302,"/admin/article/index")
		return
	}
	model := Article{}
	model.ID = uint(id)
	ctrl.DB.Delete(&model)
	c.Redirect(302,"/admin/article/index")
}

func (ctrl *ArticleController) getModel(id uint) *Article{
	model := &Article{}
	if err := ctrl.DB.First(&model, id).Error; err != nil {
		panic(err)
	}
	return model
}
/*************** Article  end  *********************/