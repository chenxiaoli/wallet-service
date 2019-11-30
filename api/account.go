package api

import (
	"github.com/chenxiaoli/wallet-service/handler"
	"github.com/chenxiaoli/wallet-service/middleware"

	"github.com/gin-gonic/gin"
)

func GetAccountAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.Use(middleware.JWTAuth())

	v1.GET("/account", handler.ListAccount)

}
