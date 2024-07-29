package system

import (
	"fmt"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	MyCommon "github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const GLOBAL_Fee = 0.04
const RemittanceFee = 50
const MaxRemittanceFee = 100000

var WEB3TOKENLIST = map[string][]Web3Chain{
	"ETH": {
		{
			Chain:     "ETH",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
			Symbol:    "ETH",
			Contract:  "",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
			Decimals:  18,
			MinCharge: "1",
		},
		{
			Chain:     "ETH",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
			Symbol:    "USDT",
			Contract:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
			Decimals:  6,
			MinCharge: "3000",
		},
		{
			Chain:     "ETH",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
			Symbol:    "USDC",
			Contract:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
			Decimals:  6,
			MinCharge: "3000",
		},
	},

	"BNB": {
		{
			Chain:     "BSC",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
			Symbol:    "BNB",
			Contract:  "",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png",
			Decimals:  18,
			MinCharge: "1",
		},
		{
			Chain:     "BSC",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
			Symbol:    "USDT",
			Contract:  "0x55d398326f99059ff775485246999027b3197955",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
			Decimals:  18,
			MinCharge: "100",
		},
		{
			Chain:     "BSC",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
			Symbol:    "USDC",
			Contract:  "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
			Decimals:  18,
			MinCharge: "100",
		},
	},
	"TRON": {
		{
			Chain:     "TRON",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/TRX.png",
			Symbol:    "USDT",
			Contract:  "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
			Decimals:  6,
			MinCharge: "1000",
		},
		{
			Chain:     "TRON",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/TRX.png",
			Symbol:    "USDC",
			Contract:  "TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
			Decimals:  6,
			MinCharge: "1000",
		},
	},
	"POLYGON": {
		{
			Chain:     "POLYGON",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
			Symbol:    "MATIC",
			Contract:  "",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
			Decimals:  18,
			MinCharge: "170",
		},
		{
			Chain:     "POLYGON",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
			Symbol:    "USDT",
			Contract:  "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
			Decimals:  6,
			MinCharge: "100",
		},
		{
			Chain:     "POLYGON",
			Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
			Symbol:    "USDC",
			Contract:  "0x3c499c542cef5e3811e1192ce70d8cc03d5c3359",
			Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
			Decimals:  6,
			MinCharge: "100",
		},
	},
	"BTC": {
		{
			Chain:     "BTC",
			Chainicon: "https://token-talk.oss-cn-shenzhen.aliyuncs.com/icon/wallet-btc.png?x-oss-process=image/resize,w_150",
			Symbol:    "BTC",
			Contract:  "",
			Icon:      "https://token-talk.oss-cn-shenzhen.aliyuncs.com/icon/wallet-btc.png?x-oss-process=image/resize,w_150",
			Decimals:  8,
			MinCharge: "0.0005",
		},
	},
}

type GasQuoteList struct {
	TransferCountGradeInfo []struct {
		TransferCountGrade int     `json:"transferCountGrade"`
		GasTotalAmount     float64 `json:"gasTotalAmount"`
		GasTotalAmountUsd  float64 `json:"gasTotalAmountUsd"`
		UsdtTotalAmount    float64 `json:"usdtTotalAmount"`
		UsdtTotalAmountUsd float64 `json:"usdtTotalAmountUsd"`
	} `json:"transferCountGradeInfo"`
	GasToken                     string  `json:"gasToken"`
	GasOnceAmount                float64 `json:"gasOnceAmount"`
	GasUSDAmount                 float64 `json:"gasUSDAmount"`
	GasTokenIcon                 string  `json:"gasTokenIcon"`
	ChainIcon                    string  `json:"chainIcon"`
	UsdtIcon                     string  `json:"usdtIcon"`
	MerchantGasBalance           int     `json:"merchantGasBalance"`
	MerchantGasLeftTransferCount int     `json:"merchantGasLeftTransferCount"`
}

var GASQUOTELIST = map[string]interface{}{
	"transferCountGradeInfo": []map[string]interface{}{
		{
			"transferCountGrade": 5,
			"gasTotalAmount":     0.001440,
			"gasTotalAmountUsd":  4.685306400,
			"usdtTotalAmount":    4.6872000,
			"usdtTotalAmountUsd": 4.6872000,
		},
		{
			"transferCountGrade": 20,
			"gasTotalAmount":     0.005760,
			"gasTotalAmountUsd":  18.741225600,
			"usdtTotalAmount":    18.7488000,
			"usdtTotalAmountUsd": 18.7488000,
		},
		{
			"transferCountGrade": 50,
			"gasTotalAmount":     0.014400,
			"gasTotalAmountUsd":  46.853064000,
			"usdtTotalAmount":    46.8720000,
			"usdtTotalAmountUsd": 46.8720000,
		},
		{
			"transferCountGrade": 100,
			"gasTotalAmount":     0.028800,
			"gasTotalAmountUsd":  93.706128000,
			"usdtTotalAmount":    93.7440000,
			"usdtTotalAmountUsd": 93.7440000,
		},
	},
	"gasToken":                     "ETH",
	"gasOnceAmount":                0.000288,
	"gasUSDAmount":                 0.937061280,
	"gasTokenIcon":                 "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
	"chainIcon":                    "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
	"usdtIcon":                     "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
	"merchantGasBalance":           0,
	"merchantGasLeftTransferCount": 0,
}

var WEB3TOKENINFO map[string]Web3Chain = map[string]Web3Chain{
	"ETH-ETH": {
		Chain:     "ETH",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
		Symbol:    "ETH",
		Contract:  "",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
		Decimals:  18,
		MinCharge: "1",
	},
	"ETH-USDT": {
		Chain:     "ETH",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
		Symbol:    "USDT",
		Contract:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
		Decimals:  6,
		MinCharge: "3000",
	},
	"ETH-USDC": {
		Chain:     "ETH",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
		Symbol:    "USDC",
		Contract:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
		Decimals:  6,
		MinCharge: "3000",
	},
	"BSC-BNB": {
		Chain:     "BSC",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
		Symbol:    "BNB",
		Contract:  "",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png",
		Decimals:  18,
		MinCharge: "1",
	},
	"BSC-USDT": {
		Chain:     "BSC",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
		Symbol:    "USDT",
		Contract:  "0x55d398326f99059ff775485246999027b3197955",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
		Decimals:  18,
		MinCharge: "100",
	},
	"BSC-USDC": {
		Chain:     "BSC",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
		Symbol:    "USDC",
		Contract:  "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
		Decimals:  18,
		MinCharge: "100",
	},
	"TRON-USDT": {
		Chain:     "TRON",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/TRX.png",
		Symbol:    "USDT",
		Contract:  "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
		Decimals:  6,
		MinCharge: "1000",
	},
	"TRON-USDC": {
		Chain:     "TRON",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/TRX.png",
		Symbol:    "USDC",
		Contract:  "TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
		Decimals:  6,
		MinCharge: "1000",
	},
	"POLYGON-MATIC": {
		Chain:     "POLYGON",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
		Symbol:    "MATIC",
		Contract:  "",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
		Decimals:  18,
		MinCharge: "170",
	},
	"POLYGON-USDT": {
		Chain:     "POLYGON",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
		Symbol:    "USDT",
		Contract:  "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
		Decimals:  6,
		MinCharge: "100",
	},
	"POLYGON-USDC": {
		Chain:     "POLYGON",
		Chainicon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
		Symbol:    "USDC",
		Contract:  "0x3c499c542cef5e3811e1192ce70d8cc03d5c3359",
		Icon:      "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
		Decimals:  6,
		MinCharge: "100",
	},
	"BTC-BTC": {
		Chain:     "BTC",
		Chainicon: "https://token-talk.oss-cn-shenzhen.aliyuncs.com/icon/wallet-btc.png?x-oss-process=image/resize,w_150",
		Symbol:    "BTC",
		Contract:  "",
		Icon:      "https://token-talk.oss-cn-shenzhen.aliyuncs.com/icon/wallet-btc.png?x-oss-process=image/resize,w_150",
		Decimals:  8,
		MinCharge: "0.0005",
	},
}

type Web3Chain struct {
	Chain     string `json:"chain"`
	Chainicon string `json:"chainicon"`
	Symbol    string `json:"symbol"`
	Contract  string `json:"contract,omitempty"`
	Icon      string `json:"icon"`
	Decimals  int    `json:"decimals"`
	MinCharge string `json:"minCharge"`
	Cact      string `json:"cact,omitempty"`
}
type StartpayWeb3Api struct{}

func (b *StartpayWeb3Api) CreateProject(c *gin.Context) {
	var r systemReq.CreateProject
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Error("test", zap.Any("CreateProject", r))
	err = utils.Verify(r, utils.CreateProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	UserId := utils.GetUserID(c)

	project := &system.SysProject{AssembleChain: r.AssembleChain, ProName: r.Name, SettleCurrency: r.SettleCurrency, AssembleAddress: r.AssembleAddress}
	project.UserId = UserId

	global.GVA_LOG.Error("test", zap.Any("project", project))

	ProjectReturn, err := StartpayWeb3Service.CreateProject(*project)

	global.GVA_LOG.Error("test", zap.Any("ProjectReturn", ProjectReturn))

	if err != nil {
		global.GVA_LOG.Error("创建项目失败!", zap.Error(err))
		response.FailWithDetailed("创建项目失败", "创建项目失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysProjectResponse{Project: ProjectReturn}, "创建项目成功", c)
}

// func (b*StartpayWeb3Api) GetWalletList(c *gin.Context, user system.SysUser) {
func (b *StartpayWeb3Api) GetProjectList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20
	err := c.ShouldBindQuery(&r)
	if err != nil {
		global.GVA_LOG.Error("test", zap.Any("GetProjectList", r),
			zap.Any("err", err.Error()),
		)
		//response.FailWithMessage(err.Error(), c)
		//return
	}

	global.GVA_LOG.Error("test", zap.Any("GetProjectList", r))

	err = utils.Verify(r, utils.GetProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("test", zap.Any("Verify", err.Error()))
		return
	}

	userId := utils.GetUserID(c)

	ProjectReturn, err := StartpayWeb3Service.GetProjectList(userId, r.Page, r.PageSize)
	if err != nil {
		global.GVA_LOG.Error("获取项目失败!", zap.Error(err))
		response.FailWithDetailed("获取项目失败", "获取项目失败", c)
		return
	}

	response.OkWithDetailed(ProjectReturn.Data, "获取项目成功", c)
}

func (b *StartpayWeb3Api) GetWalletList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindQuery(&r)
	if err != nil {
		//response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("test welcome", zap.Any("err", err.Error()))
		//return
	}
	global.GVA_LOG.Error("test", zap.Any("GetWalletList", r))
	err = utils.Verify(r, utils.GetWalletVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)

	global.GVA_LOG.Error("test", zap.Any("userId111", userId))

	ProjectReturn, err := StartpayWeb3Service.GetProjectList(userId, r.Page, r.PageSize)
	if err != nil {
		global.GVA_LOG.Error("获取钱包失败!", zap.Error(err))
		response.FailWithDetailed("获取钱包失败!", "获取钱包失败", c)
		return
	}
	WalletResp := systemRes.GetWalletRespons{}

	WalletResp.Page = ProjectReturn.Data.Page
	WalletResp.PageSize = ProjectReturn.Data.PageSize
	WalletResp.Last = ProjectReturn.Data.Last
	WalletResp.TotalPages = ProjectReturn.Data.TotalPages
	WalletResp.Total = ProjectReturn.Data.Total

	for _, conv := range ProjectReturn.Data.Content {
		wlc := systemRes.WalletInfo{}
		wlc.Id = conv.Id
		wlc.Address = conv.AssembleAddress
		wlc.Chain = conv.AssembleChain
		wlc.MerchantAddressId = userId
		wlc.ProjectId = &conv.Id
		wlc.ProjectName = &conv.Name
		wlc.CreateTime = time.Unix(conv.CreateTime, 0)
		WalletResp.Content = append(WalletResp.Content, wlc)
	}

	response.OkWithDetailed(WalletResp, "获取钱包成功", c)
}

func (b *StartpayWeb3Api) GetWeb3Announcement(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindQuery(&r)
	if err != nil {
		//response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("test welcome", zap.Any("err", err.Error()))
		//return
	}

	war := systemRes.Web3AnnouncementRespons{}

	war.TotalCount = 3

	aa1 := systemRes.Web3AnnouncementInfo{Id: "1", Title: "2023年8月StartPay 品牌正式启动", Content: "欢迎有识之士加入", Time: "2024-01-18 15:18:02"}
	aa2 := systemRes.Web3AnnouncementInfo{Id: "1", Title: "2024年4月StartPay 启动web3 2B业务", Content: "给予跨海企业跨境业务支持", Time: "2024-03-18 15:18:02"}
	war.AnnouncementList = append(war.AnnouncementList, aa1)
	war.AnnouncementList = append(war.AnnouncementList, aa2)
	aa3 := systemRes.Web3AnnouncementInfo{Id: "1", Title: "2024年8月StartPay 期待web3业务正式上线", Content: "欢迎欢迎!!!!!", Time: "2024-07-18 15:18:02"}
	war.AnnouncementList = append(war.AnnouncementList, aa3)

	response.OkWithDetailed(war, "获取account成功", c)
}

func (b *StartpayWeb3Api) GetWeb3Dashboard(c *gin.Context) {
	userId := utils.GetUserID(c)

	AccountReturn, err := StartpayWeb3Service.GetAccountInfo(userId)
	if err != nil {
		global.GVA_LOG.Error("获取account失败!", zap.Error(err))
		response.FailWithDetailed("获取account失败", "获取account失败", c)
		return
	}
	aacountResp := systemRes.Web3Dashboard{}
	totalUsdPrice := float64(0)
	for _, avalue := range AccountReturn {
		Balance, _ := strconv.ParseFloat(avalue.Balance, 64)
		UsdPrice, _ := strconv.ParseFloat(avalue.UsdtPrice, 64)
		totalUsdPrice += UsdPrice * Balance
	}
	aacountResp.AssetValuationAmount = totalUsdPrice
	aacountResp.AssetValuationCurrency = "USTD"
	aacountResp.TotalDepositAmountToday = 0
	aacountResp.TotalDepositAmountYesterday = 0
	aacountResp.TotalDepositAmount7Days = 0
	aacountResp.TotalDepositAmount30Days = 0
	aacountResp.TotalWithdrawAmountToday = 0
	aacountResp.TotalWithdrawAmountYesterday = 0
	aacountResp.TotalWithdrawAmount7Days = 0
	aacountResp.TotalWithdrawAmount30Days = 0
	response.OkWithDetailed(aacountResp, "获取account成功", c)
}
func (b *StartpayWeb3Api) GetAccountInfo(c *gin.Context) {
	userId := utils.GetUserID(c)

	AccountReturn, err := StartpayWeb3Service.GetAccountInfo(userId)
	if err != nil {
		global.GVA_LOG.Error("获取account失败!", zap.Error(err))
		response.FailWithDetailed("获取account失败", "获取account失败", c)
		return
	}

	aacountResp := systemRes.GetAccountInfoRespons{}

	for _, avalue := range AccountReturn {
		accountres := systemRes.Web3AccountInfo{}
		accountres.Balance, _ = strconv.ParseFloat(avalue.Balance, 64)
		accountres.UsdPrice, _ = strconv.ParseFloat(avalue.UsdtPrice, 64)
		accountres.AmountUsd = accountres.UsdPrice * accountres.Balance
		accountres.Address = avalue.Address
		accountres.WalletName = avalue.Name

		keys := avalue.Chain + "-" + avalue.Currency
		SymbolInfo := WEB3TOKENINFO[keys]

		accountres.Chain = avalue.Chain
		accountres.ID = avalue.Id
		accountres.Currency = avalue.Currency
		accountres.ChainIcon = SymbolInfo.Chainicon
		accountres.CurrencyIcon = SymbolInfo.Icon
		accountres.CurrencyName = SymbolInfo.Symbol
		accountres.WithdrawEnable = true
		accountres.RemittanceFeeAmount = RemittanceFee
		accountres.WithdrawFeeBoundAmount = MaxRemittanceFee
		accountres.WithdrawFeeRate1 = GLOBAL_Fee
		accountres.WithdrawFeeRate2 = GLOBAL_Fee
		accountres.WithdrawFeeRate3 = GLOBAL_Fee
		accountres.WithdrawFeeRate4 = GLOBAL_Fee
		accountres.WithdrawFeeRate5 = GLOBAL_Fee
		accountres.WithdrawFeeRate6 = GLOBAL_Fee
		accountres.WithdrawFeeRate7 = GLOBAL_Fee

		aacountResp.AccountInfo = append(aacountResp.AccountInfo, accountres)
	}
	response.OkWithDetailed(aacountResp, "获取account成功", c)
}

func (b *StartpayWeb3Api) GetChainListInfo(c *gin.Context) {
	chainList := make([]systemRes.Web3ChainListRespons, 0)
	for key, value := range MyCommon.WEB3CHAINLIST {
		chainInfo := systemRes.Web3ChainListRespons{Name: key, Icon: value}
		chainList = append(chainList, chainInfo)
	}
	response.OkWithDetailed(chainList, "获取链列表成功", c)
}

func (b *StartpayWeb3Api) GetWeb3Quote(c *gin.Context) {
	response.OkWithDetailed(GASQUOTELIST, "get quote OK", c)
}

func (b *StartpayWeb3Api) GetTokenListInfo(c *gin.Context) {
	var r systemReq.GetTokenInfoReq
	err := c.ShouldBindQuery(&r)
	if err != nil {
		//response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("test welcome", zap.Any("err", err.Error()))
		//return
	}
	TokenList := make([]systemRes.Web3ChainListRespons, 0)
	if r.Chain != "" {
		if data, ok := WEB3TOKENLIST[r.Chain]; ok {
			for _, value := range data {
				tokenInfo := systemRes.Web3ChainListRespons{Name: value.Symbol, Icon: value.Icon}
				TokenList = append(TokenList, tokenInfo)
			}
		}

	} else {
		for key, value := range MyCommon.WEB3TOKENLISTAll {
			tokenInfo := systemRes.Web3ChainListRespons{Name: key, Icon: value}
			TokenList = append(TokenList, tokenInfo)
		}
	}

	response.OkWithDetailed(TokenList, "获取token列表成功", c)
}

func (b *StartpayWeb3Api) GetDepositAddress(c *gin.Context) {

	//userId := utils.GetUserID(c)
}

func (b *StartpayWeb3Api) Web3TransferCreate(c *gin.Context) {

	var r systemReq.CreateTransferRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("Web3TransferCreate ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("Web3TransferCreate web3 db before", zap.Any("Web3TransferCreate", r))
	userId := utils.GetUserID(c)

	data, err := StartpayWeb3Service.Web3TransferCreate(userId, r)
	if err != nil {
		global.GVA_LOG.Error("转账失败!", zap.Error(err))
		response.FailWithMessage("转账失败", c)
		return
	}
	response.OkWithDetailed(data, "转账成功", c)
}

func (b *StartpayWeb3Api) Web3TransferList(c *gin.Context) {

	var r systemReq.GetWeb3Requst
	r.Page = 1
	r.PageSize = 20
	err := c.ShouldBindQuery(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)
	list, err := StartpayWeb3Service.Web3TransferList(userId, r)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(list, "获取成功", c)
}

func (b *StartpayWeb3Api) GetdepositOrder(c *gin.Context) {
	var r systemReq.GetWeb3Requst
	r.Page = 1
	r.PageSize = 20
	err := c.ShouldBindQuery(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetdepositOrder web3 db before", zap.Any("GetdepositOrder", r))
	userId := utils.GetUserID(c)
	list, err := StartpayWeb3Service.GetDepositOrder(userId, r)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}

func (b *StartpayWeb3Api) GetbankAccountList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindQuery(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)
	list, _, err := StartpayWeb3Service.GetbankAccountList(userId, r.Page, r.PageSize)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	blist := make([]systemRes.UserBankRespons, 0)

	for _, uwvalue := range list {
		ubs := systemRes.UserBankRespons{}
		ubs.ReceiverNumber = uwvalue.ReceiverNumber
		ubs.Id = strconv.FormatInt(int64(uwvalue.ID), 10)
		ubs.MerchantId = strconv.FormatInt(int64(uwvalue.MerchantId), 10)
		ubs.BankTitle = uwvalue.BankTitle
		ubs.BankCode = uwvalue.BankCode
		ubs.Region = uwvalue.Region
		ubs.FedWire = uwvalue.FedWire
		ubs.EnterpriseTitle = uwvalue.EnterpriseTitle
		ubs.ReceiverAddress = uwvalue.ReceiverAddress
		ubs.ReceiverName = uwvalue.ReceiverName
		ubs.CreateTime = uwvalue.CreatedAt
		ubs.ReferenceField = uwvalue.ReferenceField
		ubs.RemittanceType = uwvalue.RemittanceType
		ubs.UpdateTime = uwvalue.UpdatedAt
		blist = append(blist, ubs)
	}

	response.OkWithDetailed(blist, "获取成功", c)
}

func (b *StartpayWeb3Api) BankAccountCreate(c *gin.Context) {
	var r systemReq.CreateBank
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)

	userBank := &system.UserBank{BankTitle: r.BankTitle, ReceiverName: r.ReceiverName, ReceiverAddress: r.ReceiverAddress, BankCode: r.BankCode,
		EnterpriseTitle: r.EnterpriseTitle, FedWire: r.FedWire, ReceiverNumber: r.ReceiverNumber, Region: r.Region, RemittanceType: r.RemittanceType}
	userBank.MerchantId = int64(userId)

	err = StartpayWeb3Service.BankAccountCreate(userBank)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed("false", "创建失败", c)
		return
	}
	response.OkWithDetailed("true", "创建成功", c)
}

func (b *StartpayWeb3Api) BankAccountDelete(c *gin.Context) {
	var r systemReq.DeleteDBank
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("BankAccountDelete web3 db before", zap.Any("BankAccountDelete", r))
	userbank := &system.UserBank{}
	u64, err := strconv.ParseUint(r.Id, 10, 64)
	if err != nil {
		global.GVA_LOG.Error("删除bank id错误!", zap.Error(err))
		response.FailWithDetailed("false", "删除bank id错误", c)
		return
	}
	userbank.ID = uint(u64)

	err = StartpayWeb3Service.BankAccountDelete(userbank)
	if err != nil {
		global.GVA_LOG.Error("删除bank失败!", zap.Error(err))
		response.FailWithDetailed("false", "删除bank 失败", c)
		return
	}
	response.OkWithDetailed("true", "删除bank成功", c)
}

func (b *StartpayWeb3Api) UserContactList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindQuery(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)
	list, total, err := StartpayWeb3Service.UserContactList(userId, r.Page, r.PageSize)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	uws := systemRes.UserWalletRespons{}

	for _, uwvalue := range list {
		uds := systemRes.UserDEsposit{}
		uds.Chain = uwvalue.Chain
		uds.ChainIcon = MyCommon.WEB3CHAINLIST[uwvalue.Chain]
		uds.Name = uwvalue.Name
		uds.Id = strconv.FormatInt(int64(uwvalue.ID), 10)
		uds.Address = uwvalue.Address
		uds.IsInternal = false
		uws.Content = append(uws.Content, uds)
	}
	uws.Page = r.Page
	uws.PageSize = r.PageSize
	uws.Total = total
	uws.TotalPages = int(total)/r.PageSize + 1
	uws.Last = false
	response.OkWithDetailed(uws, "获取成功", c)
}
func (b *StartpayWeb3Api) UserContactCreate(c *gin.Context) {
	var r systemReq.CreateDespoitAddress
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("UserContactCreate web3 db before", zap.Any("UserContactCreate", r))
	userId := utils.GetUserID(c)

	userAddress := &system.Userwallet{Address: r.Address, Name: r.Name, Chain: r.Chain}
	userAddress.MerchantId = int64(userId)

	err = StartpayWeb3Service.UserContactCreate(userAddress)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed("false", "创建失败", c)
		return
	}
	response.OkWithDetailed("true", "创建成功", c)
}
func (b *StartpayWeb3Api) UserContactDelete(c *gin.Context) {
	var r systemReq.DeleteDespoitAddress
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("UserContactDelete web3 db before", zap.Any("UserContactDelete", r))

	userAddress := &system.Userwallet{}
	u64, err := strconv.ParseUint(r.Id, 10, 64)
	if err != nil {
		global.GVA_LOG.Error("删除收款地址id错误!", zap.Error(err))
		response.FailWithDetailed("false", "删除收款地址id错误", c)
		return
	}
	userAddress.ID = uint(u64)

	err = StartpayWeb3Service.UserContactDelete(userAddress)
	if err != nil {
		global.GVA_LOG.Error("删除收款地址失败!", zap.Error(err))
		response.FailWithDetailed("false", "删除收款地址k 失败", c)
		return
	}
	response.OkWithDetailed("true", "删除收款地址成功", c)
}

func (b *StartpayWeb3Api) WithdrawOrderList(c *gin.Context) {

	var r systemReq.GetWeb3Requst
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindQuery(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)

	//struserId := fmt.Sprintf("%u", userId)

	list, _, err := StartpayWeb3Service.WithdrawOrderList(userId, &r)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	uws := systemRes.UserWithdrawOrderRespons{}

	for _, uwvalue := range list {

		uds := systemRes.UserWithdrawOrder{}

		iid := int(uwvalue.ID)
		uds.Id = strconv.Itoa(iid)

		bkId, err := StartpayWeb3Service.BankAccountInfo(uwvalue.BankId)

		if err != nil {
			uds.BankInfo = bkId.BankTitle
			uds.BankAccount.BankTitle = bkId.BankTitle
			uds.BankAccount.EnterpriseTitle = bkId.EnterpriseTitle
			uds.BankAccount.BankCode = bkId.BankCode
			uds.BankAccount.ReceiverName = bkId.ReceiverName
			uds.BankAccount.ReceiverNumber = bkId.ReceiverNumber
			uds.BankAccount.Region = bkId.Region
		}

		uds.BankInfo = uwvalue.BankTitle
		uds.Currency = uwvalue.Currency
		uds.Chain = uwvalue.Chain
		uds.Memo = uwvalue.Memo
		uds.AdminMemo = uwvalue.AdminMemo
		uds.InputNote = uwvalue.InputNote
		uds.TxInfo = uwvalue.TxInfo
		uds.StatusName = uwvalue.StatusName
		uds.CreateTime = uwvalue.UpdatedAt
		uds.Status = fmt.Sprintf("%v", uwvalue.Status)
		uds.MerchantId = fmt.Sprintf("%v", uwvalue.MerchantId)
		uds.RemittanceFee = uwvalue.RemittanceFee
		uds.Amount = uwvalue.Amount
		uds.TotalAmount = uwvalue.TotalAmount
		uws.Content = append(uws.Content, uds)
	}
	response.OkWithDetailed(uws, "获取成功", c)

}

func (b *StartpayWeb3Api) WithdrawOrderCreate(c *gin.Context) {

	var r systemReq.CreateWithdrawOrderRequst
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("UserContactCreate web3 db before", zap.Any("UserContactCreate", r))
	userId := utils.GetUserID(c)

	Iamount, _ := strconv.ParseFloat(r.Amount, 10)

	uwo := &system.UserWithDrawOrder{
		Currency:      r.Currency,
		Chain:         r.Chain,
		Amount:        Iamount,
		BankId:        r.BankAccountId,
		InputNote:     r.Note,
		Fee:           GLOBAL_Fee,
		RemittanceFee: RemittanceFee,
		TotalAmount:   Iamount*(1+GLOBAL_Fee) + RemittanceFee,
	}
	uwo.MerchantId = int64(userId)

	err = StartpayWeb3Service.WithdrawOrderCreate(uwo)
	if err != nil {
		global.GVA_LOG.Error("创建取现订单失败!", zap.Error(err))
		response.FailWithDetailed("false", "创建取现订单失败", c)
		return
	}
	response.OkWithDetailed("true", "创建取现订单成功", c)

}

func (b *StartpayWeb3Api) AdminWithdrawOrderUpdate(c *gin.Context) {

	var r systemReq.UpdateWithdrawOrderRequst
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("AdminWithdrawOrderUpdate ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Info("AdminWithdrawOrderUpdate ShouldBindJSON ", zap.Any("request", r))

	err = StartpayWeb3Service.AdminWithdrawOrderUpdate(&r)
	if err != nil {
		global.GVA_LOG.Error("更新取现订单失败!", zap.Error(err))
		response.FailWithDetailed("false", "更新取现订单失败", c)
		return
	}
	response.OkWithDetailed("true", "更新取现订单成功", c)

}

func (b *StartpayWeb3Api) WithdrawOrderUpdate(c *gin.Context) {
	var r systemReq.UpdateWithdrawOrderRequst
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Info("AdminWithdrawOrderUpdate ShouldBindJSON ", zap.Any("request", r))

	err = StartpayWeb3Service.WithdrawOrderUpdate(&r)
	if err != nil {
		global.GVA_LOG.Error("更新取现订单失败!", zap.Error(err))
		response.FailWithDetailed("false", "更新取现订单失败", c)
		return
	}
	response.OkWithDetailed("true", "更新取现订单成功", c)
}

func (b *StartpayWeb3Api) UserBillSummary(c *gin.Context) {
	response.OkWithDetailed("true", "get成功", c)
}

func (b *StartpayWeb3Api) UserBillList(c *gin.Context) {
	response.OkWithDetailed("true", "get成功", c)
}

func (b *StartpayWeb3Api) UserBillExport(c *gin.Context) {
	response.OkWithDetailed("true", "get成功", c)
}
