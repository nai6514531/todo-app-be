package middleware

import (
	"todo-app-be/src/server/enity"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	authMsg = map[string]string{
		"011001": "用户未登录",
	}
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet("session").(sessions.Session)
		userId := session.Get("userId")
		if userId == nil {
			result := &enity.Result{"011001", false, nil, authMsg["011001"]}
			c.JSON(401, result)
			c.Abort()
		}
		// 延长session
		session.Set("userId", userId)
		session.Save()
		c.Next()
	}
}
