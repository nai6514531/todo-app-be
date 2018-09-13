package common

import (
	"fmt"

	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

func SetUpSession() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	database := viper.GetInt("redis.database")
	secret := viper.GetString("session.secret")

	store, err := redis.NewStore(database, "tcp", addr, password, []byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	SessionStore = store
}

var (
	SessionStore redis.Store
)
