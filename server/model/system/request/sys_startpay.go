package request

// Register User register structure
type CreateProject struct {
	AssembleChain   string `json:"assembleChain" example:"assembleChain"`
	Name            string `json:"name" example:"name"`
	SettleCurrency  string `json:"settleCurrency" example:"settleCurrency"`
	AssembleAddress string `json:"assemble_address"  example:"assemble_address"`
}

type GetCommonPageInfo struct {
	Page     int `json:"page" example:"page"`
	PageSize int `json:"pageSize" example:"pageSize"`
}

type GetTokenInfoReq struct {
	Chain string `form:"chain"`
}

type CreateBank struct {
	BankCode        string `json:"bankCode" example:"bankCode"`
	BankTitle       string `json:"bankTitle" example:"bankTitle"`
	EnterpriseTitle string `json:"enterpriseTitle" example:"enterpriseTitle"`
	FedWire         string `json:"fedWire"  example:"fedWire"`
	Id              string `json:"id"  example:"id"`
	ReceiverAddress string `json:"receiverAddress"  example:"receiverAddress"`
	ReceiverName    string `json:"receiverName"  example:"receiverName"`
	ReceiverNumber  string `json:"receiverNumber"  example:"receiverNumber"`
	Region          string `json:"region"  example:"region"`
	RemittanceType  string `json:"remittanceType"  example:"remittanceType"`
	EmailCode       string `json:"emailCode"`
	GoogleCode      string `json:"googleCode"`
}

type CreateDespoitAddress struct {
	Name       string `json:"name" example:"name"`
	Chain      string `json:"chain" example:"chain"`
	Address    string `json:"address"  example:"address"`
	EmailCode  string `json:"emailCode"`
	GoogleCode string `json:"googleCode"`
}

type DeleteDBank struct {
	EmailCode  string `json:"emailCode"`
	GoogleCode string `json:"googleCode"`
	Id         string `json:"id"`
}

type DeleteDespoitAddress struct {
	EmailCode  string `json:"emailCode"`
	GoogleCode string `json:"googleCode"`
	Id         string `json:"id"`
	Language   string `json:"language"`
}

type GetWithdrawOrderRequst struct {
	PageSize      int    `json:"pageSize"`
	Page          int    `json:"page"`
	ContactSearch string `json:"contactSearch"`
	Txid          string `json:"txid"`
	Currency      string `json:"currency"`
}

type CreateWithdrawOrderRequst struct {
	Amount        string `json:"amount"`
	BankAccountId string `json:"bankAccountId"`
	ProjectId     string `json:"projectId"`
	Chain         string `json:"chain"`
	Currency      string `json:"currency"`
	EmailCode     string `json:"emailCode"`
	GoogleCode    string `json:"googleCode"`
	Language      string `json:"language"`
	Note          string `json:"note"`
	PayPassword   string `json:"payPassword"`
}

type UpdateWithdrawOrderRequst struct {
	Id          string `json:"id"`
	EmailCode   string `json:"emailCode"`
	GoogleCode  string `json:"googleCode"`
	Status      string `json:"status"`
	Memo        string `json:"memo"`
	PayPassword string `json:"payPassword"`
}

type CreateTransferRequest struct {
	Amount      string  `json:"amount"`
	Asset       string  `json:"asset"`
	Chain       string  `json:"chain"`
	ProjectId   string  `json:"projectId"`
	EmailCode   string  `json:"emailCode"`
	GasDefault  float64 `json:"gasDefault"`
	GoogleCode  string  `json:"googleCode"`
	PayPassword string  `json:"payPassword"`
	ToAddress   string  `json:"toAddress"`
	ID          string  `json:"id"`
}

type GetWeb3Requst struct {
	PageSize         int    `json:"pageSize"`
	Page             int    `json:"page"`
	Search1          string `json:"Search1"`
	Search2          string `json:"Search2"`
	Currency         string `json:"currency"`
	Chain            string `json:"chain"`
	ID               string `json:"id"`
	StartTime        string `json:"startTime"`
	DepositOrderType string `json:"depositOrderType"`
	EndTime          string `json:"startTime"`
}
