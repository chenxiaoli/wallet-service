package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	// CollectionCustomer holds the name of the customer collection
	CollectionCustomer = "customer"
	CollectionCoin     = "coin"
)

type Currency struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Symbol  string        `json:"symbol" bson:"symbol"`
	Name    string        `json:"name" bson:"name"`
	Comment string        `json:"comment" bson:"comment"`
}

type Account struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id"`
	Currency       Currency      `json:"currency" bson:"currency"`             //货币类别
	Balance        float64       `json:"balance" bson:"balance"`               //余额
	BlockedBalance float64       `json:"blockedBalance" bson:"blockedBalance"` //冻结余额
	Status         string        `json:"status" bson:"status"`                 //账户状态, 正常(avail),冻结(blocked)
}

type AccountDetail struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CurrentRevenue float64       `json:"currentRevenue" bson:"currentRevenue"` //本期收入
	CurrentOutlay  float64       `json:"currentOutlay" bson:"currentOutlay"`   //本期支出
	OpeningBalance float64       `json:"openingBalance" bson:"openingBalance"` //上期余额
	CurrentBalance float64       `json:"currentBalance" bson:"currentBalance"` //本期余额
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`           //发生日期
	TxId           string        `json:"txId" bson:"txId"`                     //记账事务ID
}

type BlockedDetail struct {
	//冻结记录
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CurrentRevenue float64       `json:"currentRevenue" bson:"currentRevenue"` //本期收入
	CurrentOutlay  float64       `json:"currentOutlay" bson:"currentOutlay"`   //本期支出
	OpeningBalance float64       `json:"openingBalance" bson:"openingBalance"` //上期余额
	CurrentBalance float64       `json:"currentBalance" bson:"currentBalance"` //本期余额
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`           //发生日期
	Comment        string        `json:"comment" bson:"comment"`               //备注
	TxId           string        `json:"txId" bson:"txId"`                     //记账事务ID
}
