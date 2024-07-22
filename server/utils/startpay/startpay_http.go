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
	Data struct {
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
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
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
