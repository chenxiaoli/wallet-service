package service

import (
	"github.com/chenxiaoli/wallet-service/models"
)

func GetBalance(user *models.User, currency *models.Currency) float64 {
	return 0
}

func Withdrawal(user *models.User, coin *models.Coin) float64 {
	return 0
}

func Deposit(address string) {

}

func GetDepositAddress(coin string) string {
	return ""
}
