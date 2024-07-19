package web3api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type ProjectList struct {
	Code int `json:"code"`
	Data struct {
		Content []struct {
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
		} `json:"content"`
		Last       bool `json:"last"`
		Page       int  `json:"page"`
		PageSize   int  `json:"page_size"`
		Total      int  `json:"total"`
		TotalPages int  `json:"total_pages"`
	} `json:"data"`
	Message string `json:"message"`
}

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

func (s *StartpayWeb3Api) CeateProject(assembleChain string, projectName string, settleCurrency string) (string, error) {

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
	return string(postResponse), nil
}

func (s *StartpayWeb3Api) GetProjectList(pagenum string, pagesize string, status string) (*ProjectList, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := global.GVA_CONFIG.StartpayWeb3.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	srcStr := "GET" + Host + "/project/list?page=" + pagenum + "&pageSize=" + pagesize + "&status=" + status + strtm
	signStr, err := SignMessage(srcStr)

	if err != nil {
		fmt.Println("SignMessage err")
	}

	getHeaders := map[string]string{
		"FP-API-KEY":   ApiKey,
		"FP-SIGN":      signStr,
		"FP-TIMESTAMP": strtm,
	}
	getURL := "https://" + Host + "/project/list?page=1&pageSize=20&status=ACTIVE"
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

func (s *StartpayWeb3Api) GetProjectSecret(projectId string) (string, error) {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	ApiKey := global.GVA_CONFIG.StartpayWeb3.ApiKey
	Host := global.GVA_CONFIG.StartpayWeb3.Host
	client := NewHttpClient()

	projectID := projectId
	srcStr := "GET" + Host + "/project/secret?projectId=" + projectID + strtm
	signStr, err := SignMessage(srcStr)

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

	return string(getResponse), nil
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
