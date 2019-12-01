package models

import "gopkg.in/mgo.v2/bson"

type Chain struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string        `json:"name" bson:"name"`
	Comment string        `json:"comment" bson:"Comment"`
}

type Coin struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Chain         Chain         `json:"chain,omitempty" bson:"chain,omitempty"`       //区块链
	CoinType      string        `json:"coinType,omitempty" bson:"coinType,omitempty"` //货币类别， 代币（token） or 货币（coin）
	Symbol        string        `json:"symbol" bson:"symbol"`                         //代号
	Name          string        `json:"name" bson:"name"`                             //名称
	Comment       string        `json:"comment" bson:"comment"`                       //备注
	Currency      Currency      `json:"currency" bson:"currency"`                     //货币类别
	BlockCount    int           `json:"blockCount" bson:"blockCount"`                 //到账区块确认数
	MinDeposit    float64       `json:"minDeposit" bson:"minDeposit"`                 //最小充值
	WithdrawalFee float64       `json:"withdrawalFee" bson:"withdrawalFee"`           //提现手续费
	FeeType       string        `json:"feeType" bson:"feeType"`                       //提现手续费类型: fixed(固定),percent(百分比)
}

type DepositAddress struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Account Account       `json:"account" bson:"account"` //资金账号
	Coin    Coin          `json:"coin" bson:"coin"`       //币种
	Address string        `json:"address" bson:"address"`
}

type DepositRecord struct {
	//充值记录
	//区块确认数=LatestBlockNumber-BlockNumber
	BlockNumber       int64  `json:"blockNumber" bson:"blockNumber"`             //事务所在区块链高度
	TransactionId     string `json:"transactionId" bson:"transactionId"`         //区块链事务ID
	Address           string `json:"address" bson:"address"`                     //充值地址
	Value             string `json:"value" bson:"value"`                         //充值金额
	LatestBlockNumber int64  `json:"latestBlockNumber" bson:"latestBlockNumber"` //当前区块高度
	Status            string `json:"status" bson:"status"`                       //充值确认状态
}

type WithdrawalAddress struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Account       Account       `json:"account" bson:"account"`
	Coin          Coin          `json:"coin" bson:"coin"`
	Address       string        `json:"address" bson:"address"`
	Status        string        `json:"status" bson:"status"`               //提现状态
	BlockNumber   int64         `json:"blockNumber" bson:"blockNumber"`     //汇出事务所在区块链高度
	TransactionId string        `json:"transactionId" bson:"transactionId"` //汇出区块链事务ID
}
