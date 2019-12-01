package main

import (
	"github.com/chenxiaoli/wallet-service/api"
	"github.com/chenxiaoli/wallet-service/db"
	_ "github.com/chenxiaoli/wallet-service/docs"
	"github.com/chenxiaoli/wallet-service/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const Port = "9002"

// @title Wallet Service API
// @version 1.0
// @description 这是钱包服务的restful api.

// @host localhost:9002
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db.Connect()
	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(middleware.Connect)

	// consul健康检查回调函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// 注册服务到consul
	middleware.ConsulRegister()
	// 从consul中发现服务
	//middleware.ConsulFindServer()

	r.Use(middleware.Cors())
	api.RunHTTPServer(r)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + Port)
}
