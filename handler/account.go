package handler

import (
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/chenxiaoli/wallet-service/models"
	"github.com/gin-gonic/gin"
)

// List all Account
func ListAccount(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var products []models.Account
	err := db.C(models.CollectionCustomer).Find(nil).All(&products)
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
		"data":   products,
	})
}
