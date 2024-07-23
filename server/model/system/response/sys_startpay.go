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

type GetAccountInfoRespons struct {
	AutoSettle  bool `json:"autoSettle"`
	AccountInfo []struct {
		Currency               string      `json:"currency"`
		Chain                  string      `json:"chain"`
		WithdrawEnable         bool        `json:"withdrawEnable"`
		Balance                float64     `json:"balance"`
		AmountUsd              float64     `json:"amountUsd"`
		CurrencyIcon           string      `json:"currencyIcon"`
		CurrencyName           interface{} `json:"currencyName"`
		ChainIcon              string      `json:"chainIcon"`
		TransferFeeRate        float64     `json:"transferFeeRate"`
		GasToken               string      `json:"gasToken"`
		GasTokenToUsd          float64     `json:"gasTokenToUsd"`
		GasTokenIcon           string      `json:"gasTokenIcon"`
		GasAmount              float64     `json:"gasAmount"`
		UsdPrice               float64     `json:"usdPrice"`
		WithdrawFeeBoundAmount int         `json:"withdrawFeeBoundAmount"`
		WithdrawFeeRateDown    float64     `json:"withdrawFeeRateDown"`
		WithdrawFeeRateUp      float64     `json:"withdrawFeeRateUp"`
		RemittanceFeeAmount    int         `json:"remittanceFeeAmount"`
		WithdrawFeeRate1       float64     `json:"withdrawFeeRate1"`
		WithdrawFeeRate2       float64     `json:"withdrawFeeRate2"`
		WithdrawFeeRate3       float64     `json:"withdrawFeeRate3"`
		WithdrawFeeRate4       float64     `json:"withdrawFeeRate4"`
		WithdrawFeeRate5       float64     `json:"withdrawFeeRate5"`
		WithdrawFeeRate6       float64     `json:"withdrawFeeRate6"`
		WithdrawFeeRate7       float64     `json:"withdrawFeeRate7"`
	} `json:"accountInfo"`
	AssetValuationAmount   float64 `json:"assetValuationAmount"`
	AssetValuationCurrency string  `json:"assetValuationCurrency"`
}

type Web3ChainListRespons struct {
	Data struct {
		ETH []struct {
			Symbol    string  `json:"symbol"`
			Contract  *string `json:"contract,omitempty"`
			Icon      string  `json:"icon"`
			Decimals  int     `json:"decimals"`
			MinCharge string  `json:"minCharge"`
			Cact      string  `json:"cact,omitempty"`
		} `json:"ETH"`
		BSC []struct {
			Symbol    string  `json:"symbol"`
			Contract  *string `json:"contract"`
			Icon      string  `json:"icon"`
			Decimals  int     `json:"decimals"`
			MinCharge string  `json:"minCharge"`
		} `json:"BSC"`
		TRON []struct {
			Symbol    string `json:"symbol"`
			Contract  string `json:"contract"`
			Icon      string `json:"icon"`
			Decimals  int    `json:"decimals"`
			MinCharge string `json:"minCharge"`
		} `json:"TRON"`
		POLYGON []struct {
			Symbol    string  `json:"symbol"`
			Contract  *string `json:"contract"`
			Icon      string  `json:"icon"`
			Decimals  int     `json:"decimals"`
			MinCharge string  `json:"minCharge"`
		} `json:"POLYGON"`
		BTC []struct {
			Symbol    string      `json:"symbol"`
			Contract  interface{} `json:"contract"`
			Icon      string      `json:"icon"`
			Decimals  int         `json:"decimals"`
			MinCharge string      `json:"minCharge"`
		} `json:"BTC"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserWalletRespons struct {
	Content    []UserDEsposit `json:"content"`
	TotalPages int            `json:"total_pages"`
	Last       bool           `json:"last"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	Total      interface{}    `json:"total"`
}

type UserDEsposit struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Chain      string `json:"chain"`
	ChainIcon  string `json:"chainIcon"`
	IsInternal bool   `json:"isInternal"`
}

type UserBankRespons struct {
	Id              string      `json:"id"`
	MerchantId      string      `json:"merchantId"`
	Region          string      `json:"region"`
	RemittanceType  string      `json:"remittanceType"`
	EnterpriseTitle string      `json:"enterpriseTitle"`
	BankTitle       string      `json:"bankTitle"`
	BankCode        string      `json:"bankCode"`
	FedWire         string      `json:"fedWire"`
	ReceiverNumber  string      `json:"receiverNumber"`
	ReceiverName    string      `json:"receiverName"`
	ReceiverAddress string      `json:"receiverAddress"`
	ReferenceField  interface{} `json:"referenceField"`
	CreateTime      time.Time   `json:"createTime"`
	UpdateTime      time.Time   `json:"updateTime"`
}
