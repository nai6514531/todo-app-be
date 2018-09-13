package api

import (
	"todo-app-be/src/server/common"
	"todo-app-be/src/server/controller/api"
	"todo-app-be/src/server/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ConfigSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Options(sessions.Options{MaxAge: 10})
		c.Set("session", session)
		c.Next()
	}
}
func Api(r *gin.Engine) {

	var (
		todo = &controller.TodoController{}
		user = &controller.UserController{}
	)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(sessions.Sessions("todoSessKey", common.SessionStore))
	r.Use(middleware.ConfigSession())
	r.POST("/api/login", user.Login)

	api := r.Group("/api", middleware.Auth())
	api.GET("/incr", func(c *gin.Context) {
		session := c.MustGet("session").(sessions.Session)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	api.GET("/ping", todo.GetList)

}
