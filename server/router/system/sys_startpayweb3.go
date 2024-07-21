package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StartpayRouter struct{}

func (s *StartpayRouter) InitStartpayRouter(Router *gin.RouterGroup) {
	web3Router := Router.Group("/backend/bound_address").Use(middleware.OperationRecord())
	web3RouterWithoutRecord := Router.Group("/backend/bound_address")
	startpayWeb3Api := v1.ApiGroupApp.SystemApiGroup.StartpayWeb3Api
	{
		web3Router.POST("create_wallet", startpayWeb3Api.CreateWallet)
		web3Router.GET("list", startpayWeb3Api.GetWalletList)
	}
	{
		web3RouterWithoutRecord.POST("get_projectlist", startpayWeb3Api.GetUserList)
		web3RouterWithoutRecord.GET("get_projectlist", startpayWeb3Api.GetUserInfo)
	}
}
