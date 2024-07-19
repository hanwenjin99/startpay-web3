package web3api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func HttpRequest(
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any) (*http.Response, error) {
	// 创建URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	// 将数据编码为JSON
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// 创建请求
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// 返回响应，让调用者处理
	return resp, nil
}

// HttpClient 封装了一个 HTTP 客户端
type HttpClient struct {
	client *http.Client
}

// NewHttpClient 创建一个新的 HttpClient 实例
func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Get 发送一个 GET 请求
func (c *HttpClient) Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Post 发送一个 POST 请求
func (c *HttpClient) Post(url string, headers map[string]string, body interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {
	apiKey := "32e78838-a81e-4172-8008-d03f5102a2e7"
	//apiSecret := "0B010C5D9A3BF3BC49D4CA76713ABD4C207E4E71"
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	strtm := fmt.Sprintf("%d", timestamp)

	// alohaschen@foxmail.com
	// AppKey=32e78838-a81e-4172-8008-d03f5102a2e7
	// AppSecret=0B010C5D9A3BF3BC49D4CA76713ABD4C207E4E71

	client := NewHttpClient()

	// GET 请求示例
	/*
		url := "https://api.example.com/get"
		headers := map[string]string{
			"assembleChain":  "POLYGON",
			"name":           "startpay-test",
			"settleCurrency": "POLYGON",
			"FP-API-KEY":     apiKey,
			"FP-SIGN":        signStr,
			"FP-TIMESTAMP":   strtm,
		}
		response, err := client.Get(url, headers)
		if err != nil {
			fmt.Println("GET 请求错误:", err)
		} else {
			fmt.Println("GET 请求响应:", string(response))
		}*/

	// POST 请求示例
	/*
		postURL := "https://api.satogate.io/project/create"

		//srcStr := "GETcvm.tencentcloudapi.com/?Action=DescribeInstances&InstanceIds.0=ins-09dx96dg&Limit=20&Nonce=11886&Offset=0&Region=ap-guangzhou&SecretId=32e78838-a81e-4172-8008-d03f5102a2e7&Timestamp=1465185768&Version=2017-03-121681973331"
		assembleChain := "POLYGON"
		projectName := "startpay-test"
		settleCurrency := "USDT"

		srcStr := "POSTapi.satogate.io/project/create?assembleChain=" + assembleChain + "&name=" + projectName + "&settleCurrency=" + settleCurrency + strtm
		// 计算 HMAC-SHA1
		mac := hmac.New(sha1.New, []byte(apiSecret))
		mac.Write([]byte(srcStr))
		macHex := mac.Sum(nil)

		// 转为 Base64 编码
		signStr := base64.StdEncoding.EncodeToString(macHex)

		postHeaders := map[string]string{
			"FP-API-KEY":   apiKey,
			"FP-SIGN":      signStr,
			"FP-TIMESTAMP": strtm,
		}
		postBody := map[string]interface{}{
			"assembleChain":  assembleChain,
			"name":           projectName,
			"settleCurrency": settleCurrency,
		}
		fmt.Println("postURL", postURL)
		fmt.Println("srcStr", srcStr)
		fmt.Println("signStr", signStr)
		fmt.Println("POST head:", postHeaders)
		fmt.Println("POST postBody:", postBody)
		postResponse, err := client.Post(postURL, postHeaders, postBody)
		if err != nil {
			fmt.Println("POST 请求错误:", err)
		} else {
			fmt.Println("POST 请求响应:", string(postResponse))
		}*/

	srcStr1 := "GETapi.satogate.io/project/list?page=1&pageSize=20&status=ACTIVE" + strtm
	signStr1, err := SignMessage(srcStr1)

	if err != nil {
		fmt.Println("SignMessage err")
	}

	postHeaders := map[string]string{
		"FP-API-KEY":   apiKey,
		"FP-SIGN":      signStr1,
		"FP-TIMESTAMP": strtm,
	}
	/*postBody = map[string]interface{}{
		"page":     1,
		"pageSize": 20,
		"status":   "ACTIVE",
	}*/

	postURL1 := "https://api.satogate.io/project/list?page=1&pageSize=20&status=ACTIVE"

	fmt.Println("getRL", postURL1)
	fmt.Println("srcStr", srcStr1)
	fmt.Println("signStr", signStr1)
	fmt.Println("get head:", postHeaders)

	postResponse, err := client.Get(postURL1, postHeaders)
	if err != nil {
		fmt.Println("GEt 请求错误:", err)
	} else {
		fmt.Println("GEt 请求响应:", string(postResponse))
	}

	pplist := ProjectList{}
	json.Unmarshal(postResponse, &pplist)

	for _, value := range pplist.Data.Content {

		apiKey := apiKey
		projectID := value.Id
		srcStr1 := "GETapi.satogate.io/project/secret?projectId=" + projectID + strtm
		signStr1, err := SignMessage(srcStr1)

		if err != nil {
			fmt.Println("SignMessage err")
		}

		postHeaders = map[string]string{
			"FP-API-KEY":   apiKey,
			"FP-SIGN":      signStr1,
			"FP-TIMESTAMP": strtm,
		}
		/*postBody := map[string]interface{}{
			"projectId": projectID,
		}*/

		postURL := "https://api.satogate.io/project/secret?projectId=" + projectID

		fmt.Println("getRL", postURL)
		fmt.Println("srcStr", srcStr1)
		fmt.Println("signStr", signStr1)
		fmt.Println("get head:", postHeaders)

		postResponse, err = client.Get(postURL, postHeaders)
		if err != nil {
			fmt.Println("POST 请求错误:", err)
		} else {
			fmt.Println("POST 请求响应:", string(postResponse))
		}
	}

}

func SignMessage(message string) (string, error) {
	apiSecret := "0B010C5D9A3BF3BC49D4CA76713ABD4C207E4E71"
	mac := hmac.New(sha1.New, []byte(apiSecret))
	mac.Write([]byte(message))
	macHex := mac.Sum(nil)
	// 转为 Base64 编码
	signStr := base64.StdEncoding.EncodeToString(macHex)
	return signStr, nil
}
