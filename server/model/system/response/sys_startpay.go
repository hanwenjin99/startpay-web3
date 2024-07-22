package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"time"
)

type SysProjectResponse struct {
	Project system.SysProject `json:"project"`
}

type CreateProjectResponse struct {
	AppKey          string `json:"appKey"`
	AssembleAddress string `json:"assembleAddress"`
	AssembleChain   string `json:"assembleChain"`
	CallbackDomain  string `json:"callbackDomain"`
	CallbackUrl     string `json:"callbackUrl"`
	CreateTime      int    `json:"createTime"`
	Id              string `json:"id"`
	Name            string `json:"name"`
	PaymentPageUrl  string `json:"paymentPageUrl"`
	SettleCurrency  string `json:"settleCurrency"`
	Status          string `json:"status"`
}

type GetProjectRespons struct {
	Content []struct {
		Id              string      `json:"id"`
		UserId          interface{} `json:"userId"`
		MerchantId      string      `json:"merchantId"`
		Name            string      `json:"name"`
		AppKey          string      `json:"appKey"`
		AppSecret       string      `json:"appSecret"`
		SettleCurrency  string      `json:"settleCurrency"`
		AssembleChain   string      `json:"assembleChain"`
		AssembleAddress string      `json:"assembleAddress"`
		CallbackDomain  interface{} `json:"callbackDomain"`
		CallbackUrl     interface{} `json:"callbackUrl"`
		PaymentPageUrl  interface{} `json:"paymentPageUrl"`
		CreateTime      time.Time   `json:"createTime"`
		MerchantName    string      `json:"merchantName"`
		MerchantEmail   interface{} `json:"merchantEmail"`
	} `json:"content"`
	TotalPages int         `json:"total_pages"`
	Last       bool        `json:"last"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      interface{} `json:"total"`
}

type GetWalletRespons struct {
	Content    []WalletInfo `json:"content"`
	TotalPages int          `json:"total_pages"`
	Last       bool         `json:"last"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
	Total      interface{}  `json:"total"`
}

type WalletInfo struct {
	Id                string      `json:"id"`
	Chain             string      `json:"chain"`
	Address           string      `json:"address"`
	MerchantAddressId interface{} `json:"merchantAddressId"`
	ProjectId         *string     `json:"projectId"`
	ProjectName       *string     `json:"projectName"`
	CreateTime        time.Time   `json:"createTime"`
}
