package router

import (
	"Tiktok/api"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) Init(router *gin.RouterGroup) {
	r := router.Group("user")
	{
		r.GET("", middleware.Auth(), api.UserInfo)
		r.POST("register", api.Register)
		r.POST("login", api.Login)
	}
}
