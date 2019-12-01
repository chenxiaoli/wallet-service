package admin

import (
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/chenxiaoli/wallet-service/models"
	"github.com/gin-gonic/gin"
)

// list currency
func ListCurrency(c *gin.Context) {
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

// add currency
func AddCurrency(c *gin.Context) {
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

// add currency
func UpdateCurrency(c *gin.Context) {
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

// delete currency
func DeleteCurrency(c *gin.Context) {
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
