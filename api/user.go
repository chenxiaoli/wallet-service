package api

import (
	"github.com/chenxiaoli/wallet-service/handler"
	"github.com/gin-gonic/gin"
)

// @Summary 登录
//@Accept  application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string}  string
// @Router /auth/login [post]
func LoginAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.POST("/auth/login", handler.Login)
}
