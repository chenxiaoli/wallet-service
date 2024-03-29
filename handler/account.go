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
	var accounts []models.Account
	err := db.C(models.CollectionCustomer).Find(nil).All(&accounts)
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
		"data":   accounts,
	})
}
func ListCoin(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var coins []models.Coin
	err := db.C(models.CollectionCoin).Find(nil).All(&coins)
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
		"data":   coins,
	})
}

func ListCurrency(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var coins []models.Coin
	err := db.C(models.CollectionCoin).Find(nil).All(&coins)
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
		"data":   coins,
	})
}

func Withdrawal(c *gin.Context) {

	address, _ := c.GetPostForm("address")
	value, _ := c.GetPostForm("value")

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"value":   value,
	})
}

func GetBalance(c *gin.Context) {

	coin, _ := c.GetPostForm("coin")
	c.JSON(http.StatusOK, gin.H{
		"address": coin,
	})

}
