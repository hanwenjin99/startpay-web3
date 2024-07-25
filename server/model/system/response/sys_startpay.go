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
	AutoSettle             bool              `json:"autoSettle"`
	AccountInfo            []Web3AccountInfo `json:"accountInfo"`
	AssetValuationAmount   float64           `json:"assetValuationAmount"`
	AssetValuationCurrency string            `json:"assetValuationCurrency"`
}

type Web3AccountInfo struct {
	WalletName             string  `json:"Walletname"`
	Currency               string  `json:"currency"`
	Chain                  string  `json:"chain"`
	WithdrawEnable         bool    `json:"withdrawEnable"`
	Balance                float64 `json:"balance"`
	AmountUsd              float64 `json:"amountUsd"`
	Address                string  `json:"address"`
	CurrencyIcon           string  `json:"currencyIcon"`
	CurrencyName           string  `json:"currencyName"`
	ChainIcon              string  `json:"chainIcon"`
	TransferFeeRate        float64 `json:"transferFeeRate"`
	GasToken               string  `json:"gasToken"`
	GasTokenToUsd          float64 `json:"gasTokenToUsd"`
	GasTokenIcon           string  `json:"gasTokenIcon"`
	GasAmount              float64 `json:"gasAmount"`
	UsdPrice               float64 `json:"usdPrice"`
	WithdrawFeeBoundAmount int     `json:"withdrawFeeBoundAmount"`
	WithdrawFeeRateDown    float64 `json:"withdrawFeeRateDown"`
	WithdrawFeeRateUp      float64 `json:"withdrawFeeRateUp"`
	RemittanceFeeAmount    int     `json:"remittanceFeeAmount"`
	WithdrawFeeRate1       float64 `json:"withdrawFeeRate1"`
	WithdrawFeeRate2       float64 `json:"withdrawFeeRate2"`
	WithdrawFeeRate3       float64 `json:"withdrawFeeRate3"`
	WithdrawFeeRate4       float64 `json:"withdrawFeeRate4"`
	WithdrawFeeRate5       float64 `json:"withdrawFeeRate5"`
	WithdrawFeeRate6       float64 `json:"withdrawFeeRate6"`
	WithdrawFeeRate7       float64 `json:"withdrawFeeRate7"`
}

type Web3ChainListRespons struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
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

type UserWithdrawOrderRespons struct {
	Data struct {
		Content    []UserWithdrawOrder `json:"content"`
		TotalPages int                 `json:"total_pages"`
		Last       bool                `json:"last"`
		Page       int                 `json:"page"`
		PageSize   int                 `json:"page_size"`
		Total      float64             `json:"total"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserWithdrawOrder struct {
	Id                 string          `json:"id"`
	MerchantId         string          `json:"merchantId"`
	MerchantName       string          `json:"merchantName"`
	Currency           string          `json:"currency"`
	Chain              string          `json:"chain"`
	ChainIcon          string          `json:"chainIcon"`
	Amount             float64         `json:"amount"`
	Fee                float64         `json:"fee"`
	RemittanceFee      float64         `json:"remittanceFee"`
	TotalAmount        float64         `json:"totalAmount"`
	BankInfo           string          `json:"bankInfo"`
	BankAccount        UserBankRespons `json:"bankAccount"`
	Status             string          `json:"status"`
	StatusName         string          `json:"statusName"`
	TxInfo             string          `json:"txInfo"`
	TxCertificationUrl string          `json:"txCertificationUrl"`
	TxReference        string          `json:"txReference"`
	CreateTime         time.Time       `json:"createTime"`
	InputNote          string          `json:"inputNote"`
	CurrencyIcon       string          `json:"currencyIcon"`
	Supplier           string          `json:"supplier"`
	Memo               string          `json:"memo"`
	AdminMemo          string          `json:"adminMemo"`
}
