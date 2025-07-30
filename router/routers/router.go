package routers

import (
	"Tiktok/middleware"
	"Tiktok/router"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	root := r.Group("tiktok")
	{
		root.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		rg := router.Routers
		rg.User.Init(root)
	}
	return r
}
