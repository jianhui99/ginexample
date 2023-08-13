package controller

import (
	"ginexample/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	engine := gin.New()
	isProd := config.GetEnv("APP_ENV") == "production"
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		engine.Use(gin.Logger())
	}
	engine.Use(gin.Recovery())
	route(engine)
	err := engine.Run(":" + config.GetEnv("APP_PORT"))

	if err != nil {
		panic(err)
	}
}
