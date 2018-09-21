package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"path/filepath"
	"github.com/gin-contrib/multitemplate"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func main() {

	//db init
	DB, err := gorm.Open("sqlite3", "cache/my.db")
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&User{}, &Article{}, &Category{})
	admin := &User{Username: "admin", Password: "123"}
	DB.FirstOrCreate(&admin)
	DB.LogMode(true)

	r := gin.Default()
	render := loadTemplates("views/admin")
	render = frontTemplates("views/front", render)
	r.HTMLRender = render

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	frontCtrl := &FrontController{}
	frontCtrl.DB = DB
	frontGroup := r.Group("/")
	{
		frontGroup.GET("", frontCtrl.Index)
		frontGroup.GET("view", frontCtrl.View)
	}

	adminCtrl := &AdminController{}
	adminCtrl.DB = DB
	articleCtrl := &ArticleController{}
	articleCtrl.DB = DB
	adminGroup := r.Group("/admin")
	store := cookie.NewStore([]byte("secret"))
	adminGroup.Use(sessions.Sessions("mysession", store))
	adminGroup.Use(adminCtrl.MiddleWarePrepare)
	router := &Router{}
	{
		router.Add("POST","/login", adminCtrl.Login)

		adminGroup.Static("/assets", "views/admin/assets")
		adminGroup.GET("", adminCtrl.Index)
		adminGroup.GET("/login", adminCtrl.Login)
		adminGroup.POST("/login", adminCtrl.Login)
		adminGroup.GET("/logout", adminCtrl.Logout)
		adminGroup.GET("/statistics", adminCtrl.Statistics)
		adminGroup.GET("/Session-test", adminCtrl.SessionTest)

		//article
		adminGroup.GET("/article/index", articleCtrl.Index)
		adminGroup.GET("/article/create", articleCtrl.Create)
		adminGroup.POST("/article/create", articleCtrl.Create)
		adminGroup.GET("/article/update/:id", articleCtrl.Update)
		adminGroup.POST("/article/update/:id", articleCtrl.Update)
		adminGroup.GET("/article/delete/:id", articleCtrl.Delete)

	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func frontTemplates(templatesDir string, render multitemplate.Renderer) multitemplate.Renderer {
	tpls := []string{"index", "view"}
	for _, tpl := range tpls {
		render.AddFromFiles("front/"+tpl, templatesDir+"/"+tpl+".html")
	}
	return render
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		files := append(layouts, include)
		r.AddFromFiles(filepath.Base(templatesDir)+"/"+strings.Replace(filepath.Base(include), ".html", "", 1), files...)
	}
	return r
}
