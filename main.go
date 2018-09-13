package main

import (
	"todo-app-be/src/server"
	"todo-app-be/src/server/common"
)

func main() {
	common.SetUpConfig()

	common.SetupDB()

	common.SetUpRedis()

	common.SetUpSession()
	// common.Redis.Set("key:user:", 123, time.Duration(12*time.Second))
	// common.Redis.Set("key:sssss:", "this is test", time.Duration(24*time.Hour))

	// time.Sleep(13 * time.Second)
	// id, err := common.Redis.Get("key:user:").Int64()
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Println("id", id)
	server.SetUpServer()
	// common.SetUpCommon()

	// common.SetUpLogger()

	// common.SetUpSession()

	// common.SetupDB()

	// common.SetUpRedis()

	// common.SetUpUserRedis()

	// //common.SetUpRateLimiter()

	// common.SetupCORS()

	// //common.SetUpSecure()

	// cron.SetUpCron()

	// server.SetUpServer()
}
