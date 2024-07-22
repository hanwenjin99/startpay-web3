package web3api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
