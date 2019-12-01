package handler

import (
	"fmt"
	"github.com/chenxiaoli/wallet-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	fmt.Printf("password %s , username: %s", password, username)
	auth := service.AuthApiService{}
	token, err := auth.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"token":  token,
	})
}
