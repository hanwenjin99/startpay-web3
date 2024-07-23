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
