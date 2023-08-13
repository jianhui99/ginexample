package main

import (
	"ginexample/config"
	"ginexample/controller"
	mysql "ginexample/db/mysql"
	redis "ginexample/db/redis"
)

func main() {
	config.Init()
	mysql.Init()
	redis.Init()
	controller.Init()
}
