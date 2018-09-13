package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/mustache"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func _SetupDB(name string, readOnly bool) *gorm.DB {
	isDevelopment := viper.GetBool("isDevelopment")
	instance := "wr"
	if readOnly {
		instance = "r"
	}
	prefix := mustache.Render("mysql.database.{{name}}.{{instance}}", map[string]interface{}{
		"name":     name,
		"instance": instance,
	})
	dialect := viper.GetString(prefix + ".dialect")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetString(prefix + ".port")
	maxIdle := viper.GetInt(prefix + ".max-idle")
	maxOpen := viper.GetInt(prefix + ".max-open")
	url := mustache.Render("{{user}}:{{password}}@tcp({{host}}:{{port}})/{{database}}?charset=utf8&parseTime=True&loc=Local", map[string]interface{}{
		"user":     user,
		"password": password,
		"database": database,
		"host":     host,
		"port":     port,
	})
	db, err := gorm.Open(dialect, url)
	if err != nil {
		fmt.Println("err", err)
		// Logger.Warningln("failed to connect database:", database, err.Error())
		// panic("failed to connect database")
	}
	db.LogMode(isDevelopment)
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	return db
}

func SetupDB() {

	TodoDB_WR = _SetupDB("todo", false)

	//MNDB_WR = _SetupDB("mnzn", false)
	//MNDB_R = _SetupDB("mnzn", true)
}

var (
	TodoDB_WR *gorm.DB
)
