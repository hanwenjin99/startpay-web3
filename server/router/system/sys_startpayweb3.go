package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StartpayRouter struct{}

func (s *StartpayRouter) InitStartpayRouter(Router *gin.RouterGroup) {
	web3Router := Router.Group("web3").Use(middleware.OperationRecord())
	web3RouterWithoutRecord := Router.Group("web3")
	startpayWeb3Api := v1.ApiGroupApp.SystemApiGroup.StartpayWeb3Api
	{
		web3RouterWithoutRecord.POST("create_project", startpayWeb3Api.CreateProject)
		web3RouterWithoutRecord.GET("project_list", startpayWeb3Api.GetProjectList)
		web3RouterWithoutRecord.GET("wallet_list", startpayWeb3Api.GetWalletList)
		web3RouterWithoutRecord.GET("account_info", startpayWeb3Api.GetAccountInfo)
		web3RouterWithoutRecord.GET("chain_list", startpayWeb3Api.GetChainListInfo)
		web3RouterWithoutRecord.GET("deposit_address", startpayWeb3Api.GetDepositAddress)
		web3RouterWithoutRecord.GET("deposit_order", startpayWeb3Api.GetAdepositOrder)

	}
	{
		web3Router.POST("get_projectlist", startpayWeb3Api.GetProjectList)
	}
}
