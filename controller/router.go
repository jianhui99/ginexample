package controller

import (
	"ginexample/controller/middleware"

	"github.com/gin-gonic/gin"
)

func route(engine *gin.Engine) {
	engine.Use(middleware.CorsMiddleware)
	rg := engine.Group("/api/v1")
	{
		rg.GET("/examples", GinHandleResultError(HandleGetExample))
		rg.GET("/examples/:id", GinHandleResultError(HandleGetExampleDetail))
	}

}
