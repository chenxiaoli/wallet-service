package admin

import (
	"github.com/chenxiaoli/wallet-service/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
)

// list chain
func ListChain(c *gin.Context) {
	return
}

//
func AddChain(c *gin.Context) {
	return
}

func UpdateChain(c *gin.Context) {

}

func DeleteChain(c *gin.Context) {
	return
}

// add coin
func AddCoin(c *gin.Context) {
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

// update coin
func UpdateCoin(c *gin.Context) {
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

func GetDepositAddress(c *gin.Context) {
	return
}

func GetWithdrawalAddress(c *gin.Context) {
	return
}

func GetDepositList(c *gin.Context) {
	return
}

func GetWithdrawalList(c *gin.Context) {
	return
}

//提现审核
func ReviewWithdrawal(c *gin.Context) {

}
