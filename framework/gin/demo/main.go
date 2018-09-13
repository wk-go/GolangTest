package main

import "github.com/gin-gonic/gin"

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
	{
		adminGroup.GET("",adminCtrl.Index)
		adminGroup.GET("/statistics",adminCtrl.Statistics)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
