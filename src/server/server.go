package server

import (
	"todo-app-be/src/server/route/api"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetUpServer() {

	r := gin.New()
	api.Api(r)

	r.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))

}
