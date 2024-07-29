package web3api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type StartpayWeb3Api struct {
	ApiKey    string
	ApiSecret string
	Host      string
	client    *HttpClient
}

func NewStartpayWeb3Api() *StartpayWeb3Api {

	return &StartpayWeb3Api{
		client:    NewHttpClient(),
		ApiKey:    global.GVA_CONFIG.StartpayWeb3.ApiKey,
		ApiSecret: global.GVA_CONFIG.StartpayWeb3.ApiSecret,
		Host:      global.GVA_CONFIG.StartpayWeb3.Host,
	}
}

func (s *StartpayWeb3Api) CeateProject(assembleChain string, projectName string, settleCurrency string) (*Web3CreateProjectResponse, error) {

	ApiKey := global.GVA_CONFIG.StartpayWeb3.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	postURL := "https://" + Host + "/project/create"
	srcStr := "POST" + Host + "/project/create?assembleChain=" + assembleChain + "&name=" + projectName + "&settleCurrency=" + settleCurrency + strtm
	signStr, err := s.SignMessage(srcStr)

	postHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	postBody := map[string]interface{}{
		"assembleChain":  assembleChain,
		"name":           projectName,
		"settleCurrency": settleCurrency,
	}

	postResponse, err := client.Post(postURL, postHeaders, postBody)
	if err != nil {
		fmt.Println("POST 请求错误:", err)
	} else {
		fmt.Println("POST 请求响应:", string(postResponse))
	}

	ProjectInfo := Web3CreateProjectResponse{}
	json.Unmarshal(postResponse, &ProjectInfo)

	return &ProjectInfo, nil
}

func (s *StartpayWeb3Api) GetProjectList(Page int, PageSize int, status string, projectID string) (*ProjectList, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := global.GVA_CONFIG.StartpayWeb3.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	pagenums := strconv.Itoa(Page)
	pagesizes := strconv.Itoa(PageSize)
	srcStr := "GET" + Host + "/project/list?page=" + pagenums + "&pageSize=" + pagesizes + "&projectId=" + projectID + "&status=" + status + strtm
	signStr, err := s.SignMessage(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}
	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	getURL := "https://" + Host + "/project/list?page=" + pagenums + "&pageSize=" + pagesizes + "&projectId=" + projectID + "&status=" + status

	global.GVA_LOG.Error("test web3",
		zap.Any("signStr", signStr),
		zap.Any("srcStr", srcStr),
		zap.Any("getURL", getURL),
	)

	getResponse, err := client.Get(getURL, getHeaders)
	if err != nil {
		fmt.Println("GEt 请求错误:", err)
	} else {
		fmt.Println("GEt 请求响应:", string(getResponse))
	}

	pplist := ProjectList{}
	json.Unmarshal(getResponse, &pplist)
	return &pplist, nil
}

func (s *StartpayWeb3Api) GetChainListInfo() (*Web3ChainListRespons, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := global.GVA_CONFIG.StartpayWeb3.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	srcStr := "GET" + Host + "currency/list" + strtm
	signStr, err := s.SignMessage(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}

	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	getURL := "https://" + Host + "/currency/list"

	global.GVA_LOG.Error("test web3",
		zap.Any("signStr", signStr),
		zap.Any("srcStr", srcStr),
		zap.Any("getURL", getURL),
	)

	getResponse, err := client.Get(getURL, getHeaders)
	if err != nil {
		fmt.Println("GEt 请求错误:", err)
	} else {
		fmt.Println("GEt 请求响应:", string(getResponse))
	}

	chainlist := Web3ChainListRespons{}
	json.Unmarshal(getResponse, &chainlist)
	return &chainlist, nil
}

func (s *StartpayWeb3Api) Web3TransferCreate(requst systemReq.CreateTransferRequest) (*Web3CreatetransferReturn, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := s.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	global.GVA_LOG.Info("Web3TransferCreate web3",
		zap.Any("requst", requst),
	)

	srcStr := "GET" + Host + "/transaction/transfer?amount=" + requst.Amount + "&asset=" + requst.Asset + "&chain=" + requst.Chain + "&toAddress=" + requst.ToAddress + strtm
	signStr, err := s.SignMessage2(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}
	postURL := "https://" + Host + "/transaction/transfer"

	postHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	postBody := map[string]interface{}{
		"amount":    requst.Amount,
		"asset":     requst.Asset,
		"chain":     requst.Chain,
		"toAddress": requst.ToAddress,
	}

	global.GVA_LOG.Info("Web3TransferCreate web3",
		zap.Any("signStr", signStr),
		zap.Any("srcStr", srcStr),
		zap.Any("POStURL", postURL),
		zap.Any("ApiKey", ApiKey),
	)

	postResponse, err := client.Post(postURL, postHeaders, postBody)
	if err != nil {
		global.GVA_LOG.Error("Web3TransferCreate web3",
			zap.Any("POST 请求错误", err.Error()),
		)
	} else {
		global.GVA_LOG.Info("Web3TransferCreate web3",
			zap.Any("POST 请求响应:", postResponse),
		)
	}

	transferRecord := Web3CreatetransferReturn{}
	json.Unmarshal(postResponse, &transferRecord)
	global.GVA_LOG.Info("Web3TransferCreate web3",
		zap.Any("transferRecord:", transferRecord),
	)
	return &transferRecord, nil
}

func (s *StartpayWeb3Api) GetDepositOrder(projectid string, requst systemReq.GetWeb3Requst) (*Web3DepositListReturn, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := s.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	pagenums := strconv.Itoa(requst.Page)
	pagesizes := strconv.Itoa(requst.PageSize)

	srcStr := "GET" + Host + "/wallet/deposit/history?address=" + projectid + "&chain=" + requst.Chain + "&page=" + pagenums + "&pageSize=" + pagesizes +
		"&token=" + requst.Currency + strtm
	signStr, err := s.SignMessage2(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}

	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	getURL := "https://" + Host + "/wallet/deposit/history?address=" + projectid + "&chain=" + requst.Chain + "&page=" + pagenums + "&pageSize=" + pagesizes +
		"&token=" + requst.Currency

	global.GVA_LOG.Error("test web3",
		zap.Any("signStr", signStr),
		zap.Any("srcStr", srcStr),
		zap.Any("getURL", getURL),
		zap.Any("ApiKey", ApiKey),
	)

	getResponse, err := client.Get(getURL, getHeaders)
	if err != nil {
		fmt.Println("GEt 请求错误:", err)
	} else {
		fmt.Println("GEt 请求响应:", string(getResponse))
	}

	depositlist := Web3DepositListReturn{}
	json.Unmarshal(getResponse, &depositlist)
	return &depositlist, nil
}

func (s *StartpayWeb3Api) Web3TransferList(projectid string, requst systemReq.GetWeb3Requst) (*Web3ListtransferReturn, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := s.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	pagenums := strconv.Itoa(requst.Page)
	pagesizes := strconv.Itoa(requst.PageSize)

	srcStr := "GET" + Host + "/wallet/transfer/history?chain=" + requst.Chain + "&page=" + pagenums + "&pageSize=" + pagesizes +
		"&projectId=" + projectid + "&token=" + requst.Currency + strtm
	signStr, err := s.SignMessage2(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}

	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	getURL := "https://" + Host + "/wallet/transfer/history?chain=" + requst.Chain + "&page=" + pagenums + "&pageSize=" + pagesizes +
		"&projectId=" + projectid + "&token=" + requst.Currency

	global.GVA_LOG.Error("test web3",
		zap.Any("signStr", signStr),
		zap.Any("srcStr", srcStr),
		zap.Any("getURL", getURL),
		zap.Any("ApiKey", ApiKey),
	)

	getResponse, err := client.Get(getURL, getHeaders)
	if err != nil {
		fmt.Println("GEt 请求错误:", err)
	} else {
		fmt.Println("GEt 请求响应:", string(getResponse))
	}

	transferlist := Web3ListtransferReturn{}
	json.Unmarshal(getResponse, &transferlist)
	return &transferlist, nil
}

func (s *StartpayWeb3Api) GetAccountInfo() (*GetAccountInfoRespons, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := s.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	srcStr := "GET" + Host + "/wallet/project/account?" + strtm
	signStr, err := s.SignMessage2(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}

	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	getURL := "https://" + Host + "/wallet/project/account"

	global.GVA_LOG.Error("test web3",
		zap.Any("signStr", signStr),
		zap.Any("srcStr", srcStr),
		zap.Any("getURL", getURL),
		zap.Any("ApiKey", ApiKey),
	)

	getResponse, err := client.Get(getURL, getHeaders)
	if err != nil {
		fmt.Println("GEt 请求错误:", err)
	} else {
		fmt.Println("GEt 请求响应:", string(getResponse))
	}

	Accoubtlist := GetAccountInfoRespons{}
	json.Unmarshal(getResponse, &Accoubtlist)
	return &Accoubtlist, nil
}

func (s *StartpayWeb3Api) GetProjectSecret(projectId string) (*Web3GetSecretResponse, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := global.GVA_CONFIG.StartpayWeb3.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	projectID := projectId
	srcStr := "GET" + Host + "/project/secret?projectId=" + projectID + strtm
	signStr, err := s.SignMessage(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}
	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}

	getURL := "https://" + Host + "/project/secret?projectId=" + projectID
	getResponse, err := client.Get(getURL, getHeaders)
	if err != nil {
		fmt.Println("POST 请求错误:", err)
	} else {
		fmt.Println("POST 请求响应:", string(getResponse))
	}

	SecretInfo := Web3GetSecretResponse{}
	json.Unmarshal(getResponse, &SecretInfo)

	return &SecretInfo, nil
}

func (s *StartpayWeb3Api) SignMessage(message string) (string, error) {
	apiSecret := global.GVA_CONFIG.StartpayWeb3.ApiSecret
	mac := hmac.New(sha1.New, []byte(apiSecret))
	mac.Write([]byte(message))
	macHex := mac.Sum(nil)
	// 转为 Base64 编码
	signStr := base64.StdEncoding.EncodeToString(macHex)
	return signStr, nil
}

func (s *StartpayWeb3Api) SignMessage2(message string) (string, error) {
	apiSecret := s.ApiSecret
	mac := hmac.New(sha1.New, []byte(apiSecret))
	mac.Write([]byte(message))
	macHex := mac.Sum(nil)
	// 转为 Base64 编码
	signStr := base64.StdEncoding.EncodeToString(macHex)
	return signStr, nil
}
