package middleware

import (
	"Tiktok/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status-msg,omitempty"`
}

// Auth 鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		// 没有携带token
		if len(token) == 0 {
			// 没有token，阻止后面函数执行
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				// token有误，阻止后面函数执行
				c.Abort()
				c.JSON(http.StatusUnauthorized, Response{
					StatusCode: -1,
					StatusMsg:  "Token Error",
				})
			} else {
				log.Println("Token correct")
			}
			c.Set("userID", claims.ID)
			c.Next()
		}
	}
}
