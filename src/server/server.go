package server

import (
	"todo-app-be/src/server/route/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// "github.com/thinkerou/favicon"
)

func SetUpServer() {
	indexHtml := "./src/static/index.html"
	r := gin.New()
	// r.Use(favicon.New("./src/static/favicon.ico"))
	r.Use(static.Serve("/", static.LocalFile("./src/static/", false)))
	r.LoadHTMLGlob(indexHtml)
	r.NoRoute(func(c *gin.Context) {
		c.File(indexHtml)
	})
	api.Api(r)
	r.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))

}
