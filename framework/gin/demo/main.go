package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	{
		adminGroup.GET("",adminCtrl.Index)
		adminGroup.GET("/login",adminCtrl.Login)
		adminGroup.POST("/login",adminCtrl.Login)
		adminGroup.GET("/logout",adminCtrl.Logout)
		adminGroup.GET("/statistics",adminCtrl.Statistics)
		adminGroup.GET("/session-test",adminCtrl.SessionTest)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
