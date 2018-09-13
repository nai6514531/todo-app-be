package common

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

func SetUpConfig() {
	var (
		conf = flag.String("conf", "./runtime.config", "config file path")
	)

	flag.Parse()

	viper.SetConfigName(*conf)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:", err.Error())
	}
}
