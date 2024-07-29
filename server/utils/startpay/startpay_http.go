package web3api

import (
	"bytes"
	"encoding/json"
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

type ProjectList struct {
	Code int `json:"code"`
	Data struct {
		Content []struct {
			AppKey          string `json:"appKey"`
			AssembleAddress string `json:"assembleAddress"`
			AssembleChain   string `json:"assembleChain"`
			CallbackDomain  string `json:"callbackDomain"`
			CallbackUrl     string `json:"callbackUrl"`
			CreateTime      int64  `json:"createTime"`
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

type Web3CreateProjectResponse struct {
	Code    int               `json:"code"`
	Data    CreateProjectData `json:"data"`
	Message string            `json:"message"`
}

type CreateProjectData struct {
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

type Web3GetSecretResponse struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type GetAccountInfoRespons struct {
	Code    int              `json:"code"`
	Data    []GetAccountInfo `json:"data"`
	Message string           `json:"message"`
}

type GetAccountInfo struct {
	Address   string `json:"address"`
	Balance   string `json:"balance"`
	Chain     string `json:"chain"`
	Currency  string `json:"currency"`
	UsdtPrice string `json:"usdtPrice"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	ProjectId string `json:"projectId"`
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

type Web3CreatetransferReturn struct {
	Code int `json:"code"`
	Data struct {
		Amount      string `json:"amount"`
		Chain       string `json:"chain"`
		CreateTime  int    `json:"createTime"`
		Currency    string `json:"currency"`
		FromAddress string `json:"fromAddress"`
		Gas         string `json:"gas"`
		GasPrice    string `json:"gasPrice"`
		Id          string `json:"id"`
		Status      string `json:"status"`
		ToAddress   string `json:"toAddress"`
		TxTime      int    `json:"txTime"`
		Txid        string `json:"txid"`
	} `json:"data"`
	Message string `json:"message"`
}

/*type Web3ListtransferReturn struct {
	Content []struct {
		Amount      string `json:"amount"`
		Chain       string `json:"chain"`
		CreateTime  uint   `json:"createTime"`
		Currency    string `json:"currency"`
		FromAddress string `json:"fromAddress"`
		Gas         string `json:"gas"`
		GasPrice    string `json:"gasPrice"`
		Id          string `json:"id"`
		Status      string `json:"status"`
		ToAddress   string `json:"toAddress"`
		TxTime      uint   `json:"txTime"`
		Txid        string `json:"txid"`
	} `json:"content"`
	Last       bool `json:"last"`
	Page       int  `json:"page"`
	PageSize   int  `json:"page_size"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
}*/

type Web3ListtransferReturn struct {
	Data struct {
		Content []struct {
			Id          string `json:"id"`
			FromAddress string `json:"fromAddress"`
			ToAddress   string `json:"toAddress"`
			Chain       string `json:"chain"`
			Currency    string `json:"currency"`
			Amount      string `json:"amount"`
			Gas         string `json:"gas"`
			GasPrice    string `json:"gasPrice"`
			Status      string `json:"status"`
			Txid        string `json:"txid"`
			TxTime      uint   `json:"txTime"`
			CreateTime  uint   `json:"createTime"`
		} `json:"content"`
		TotalPages int  `json:"total_pages"`
		Last       bool `json:"last"`
		Page       int  `json:"page"`
		PageSize   int  `json:"page_size"`
		Total      int  `json:"total"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Web3DepositListReturn struct {
	Data struct {
		Total    int `json:"total"`
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
		Data     []struct {
			Id                string `json:"id"`
			MerchantAddressId string `json:"merchantAddressId"`
			Chain             string `json:"chain"`
			Token             string `json:"token"`
			Address           string `json:"address"`
			DepositAmount     string `json:"depositAmount"`
			Amount            string `json:"amount"`
			AmountUsd         string `json:"amountUsd"`
			Status            string `json:"status"`
			CreateTime        uint   `json:"createTime"`
			FinishedTime      int    `json:"finishedTime"`
			TxHash            string `json:"txHash"`
			FromAddress       string `json:"fromAddress"`
			SettleTxid        string `json:"settleTxid"`
			TxList            string `json:"txList"`
			OrderModeStatus   string `json:"orderModeStatus"`
		} `json:"data"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
