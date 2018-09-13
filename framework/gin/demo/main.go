package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"path/filepath"
	"github.com/gin-contrib/multitemplate"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	frontCtrl := &FrontController{}
	frontGroup := r.Group("/")
	{
		frontGroup.GET("", frontCtrl.Index)
		frontGroup.GET("view", frontCtrl.View)
	}

	adminCtrl := &AdminController{}
	adminGroup := r.Group("/admin")
	store := cookie.NewStore([]byte("secret"))
	adminGroup.Use(sessions.Sessions("mysession", store))

	r.HTMLRender = loadTemplates("views/admin")
	{
		adminGroup.Static("/assets", "views/admin/assets")
		adminGroup.GET("",adminCtrl.Index)
		adminGroup.GET("/login",adminCtrl.Login)
		adminGroup.POST("/login",adminCtrl.Login)
		adminGroup.GET("/logout",adminCtrl.Logout)
		adminGroup.GET("/statistics",adminCtrl.Statistics)
		adminGroup.GET("/session-test",adminCtrl.SessionTest)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
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
		r.AddFromFiles(filepath.Base(templatesDir) + "/" + filepath.Base(include), files...)
	}
	return r
}
