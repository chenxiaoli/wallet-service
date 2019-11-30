package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	// CollectionCustomer holds the name of the customer collection
	CollectionCustomer = "customer"
)

type User struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

type Currency struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Symbol  string        `json:"symbol" bson:"symbol"`
	Name    string        `json:"name" bson:"name"`
	Comment string        `json:"comment" bson:"comment"`
}

type Account struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id"`
	Currency Currency      `json:"currency" bson:"currency"` //货币类别
	Balance  int64         `json:"balance" bson:"balance"`   //余额
}

type AccountDetail struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CurrentRevenue int64         `json:"currentRevenue" bson:"currentRevenue"` //本期收入
	CurrentOutlay  int64         `json:"currentOutlay" bson:"currentOutlay"`   //本期支出
	OpeningBalance int64         `json:"openingBalance" bson:"openingBalance"` //上期余额
	CurrentBalance int64         `json:"currentBalance" bson:"currentBalance"` //本期余额
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`           //发生日期
}

type Chain struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

type Coin struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Chain    Chain         `json:"chain,omitempty" bson:"chain,omitempty"`       //区块链
	CoinType string        `json:"coinType,omitempty" bson:"coinType,omitempty"` //货币类别， 代币（token） or 货币（coin）
	Symbol   string        `json:"symbol" bson:"symbol"`                         //代号
	Name     string        `json:"name" bson:"name"`                             //名称
	Comment  string        `json:"comment" bson:"comment"`                       //备注
	Currency Currency      `json:"currency" bson:"currency"`                     //货币类别
}

type DepositAddress struct {
}

type WithdrawalAddress struct {
}
