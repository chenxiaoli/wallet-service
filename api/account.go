package api

import (
	"github.com/chenxiaoli/wallet-service/handler"
	"github.com/chenxiaoli/wallet-service/middleware"

	"github.com/gin-gonic/gin"
)

// @Summary 获取当前登录用户的资金账户各币种基本信息(余额)
// @Produce  json
// @Success 200 {objects}  AccountVo
// @Router /account/me [get]
// @Security ApiKeyAuth
func GetAccountAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.Use(middleware.JWTAuth())
	v1.GET("/account/me", handler.ListAccount)
}

// @Summary 获取资金账户指定货币的账号详情 (资金流水)
// @produce json
// @Param currency query string true "currency"
// @Success 200
// @Router /account/me/currency/{currencyId} [get]
// @Security ApiKeyAuth
func GetAccountDetailAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.Use(middleware.JWTAuth())
	v1.GET("/account/me/currency/:currencyId", handler.ListAccount)
}

// @Summary 充值
// @Produce  json
// @Param account body models.Account   true "deposit"
// @Success 200 {string} AccountVo
// @Router /deposit/callback [post]
func DepositAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.GET("/deposit/callback", handler.Deposit)
}

//@Summary 充值地址
//@Description 获取当前登录用户充值地址
//@Produce json
func GetDepositAddressAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.GET("/deposit/address", handler.GetDepositAddress)
}

// @Summary 充值
// @Produce  json
// @Param account body models.Account   true "deposit"
// @Success 200 {string} AccountVo
// @Router /deposit/callback [post]
func TransferAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.GET("/deposit/callback", handler.Deposit)
}
