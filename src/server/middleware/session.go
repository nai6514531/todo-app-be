package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ConfigSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		maxAge := viper.GetInt("session.maxAge")
		session.Options(sessions.Options{MaxAge: maxAge})
		c.Set("session", session)
		c.Next()
	}
}
