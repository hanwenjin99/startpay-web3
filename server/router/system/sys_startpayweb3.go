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
		web3RouterWithoutRecord.GET("merchant/bankAccount/list", startpayWeb3Api.GetbankAccountList)
		web3RouterWithoutRecord.POST("merchant/bankAccount/create", startpayWeb3Api.BankAccountCreate)
		web3RouterWithoutRecord.POST("merchant/bankAccount/delete", startpayWeb3Api.BankAccountDelete)
		web3RouterWithoutRecord.GET("merchant/contact/list", startpayWeb3Api.UserContactList)
		web3RouterWithoutRecord.POST("merchant/contact/create", startpayWeb3Api.UserContactCreate)
		web3RouterWithoutRecord.POST("merchant/contact/delete", startpayWeb3Api.UserContactDelete)

		web3RouterWithoutRecord.GET("merchant/withdraw/list", startpayWeb3Api.WithdrawOrderList)
		web3RouterWithoutRecord.GET("admin/withdraw/list", startpayWeb3Api.WithdrawOrderList)
		web3RouterWithoutRecord.POST("merchant/withdraw/create", startpayWeb3Api.WithdrawOrderCreate)
		web3RouterWithoutRecord.POST("merchant/withdraw/update", startpayWeb3Api.WithdrawOrderUpdate)
		web3RouterWithoutRecord.POST("admin/withdraw/update", startpayWeb3Api.WithdrawOrderUpdate)

	}
	{
		web3Router.POST("get_projectlist", startpayWeb3Api.GetProjectList)
	}
}
