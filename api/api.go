package api

import (
	"github.com/gin-gonic/gin"
)

func RunHTTPServer(engine *gin.Engine) {
	engine.Group("auth")
	{
		GetAccountAPI(engine)
		LoginAPI(engine)
	}

}
