package common

type ClearDB struct {
	TableName    string
	CompareField string
	Interval     string
}

var WEB3CHAINLIST = map[string]string{
	"ETH":     "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
	"BSC":     "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/png",
	"TRON":    "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/TRX.png",
	"POLYGON": "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
	"BTC":     "https://token-talk.oss-cn-shenzhen.aliyuncs.com/icon/wallet-btc.png?x-oss-process=image/resize,w_150",
}

var WEB3GASLIST = map[string]string{
	"ETH":     "ETH",
	"BSC":     "BNB",
	"TRON":    "TRX",
	"POLYGON": "MATIC",
	"BTC":     "BTC",
}

var WEB3TOKENLISTAll = map[string]string{
	"USDT":  "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
	"USDC":  "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdc.png",
	"BNB":   "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png",
	"MATIC": "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/MATIC.png",
	"BTC":   "https://token-talk.oss-cn-shenzhen.aliyuncs.com/icon/wallet-btc.png?x-oss-process=image/resize,w_150",
	"ETH":   "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png",
}
