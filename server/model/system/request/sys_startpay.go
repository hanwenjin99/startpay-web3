package request

// Register User register structure
type CreateProject struct {
	AssembleChain   string `json:"assembleChain" example:"assembleChain"`
	Name            string `json:"name" example:"name"`
	SettleCurrency  string `json:"settleCurrency" example:"settleCurrency"`
	AssembleAddress string `json:"assemble_address"  example:"assemble_address"`
}

type GetProjectList struct {
	Page     int `json:"page" example:"page"`
	PageSize int `json:"pageSize" example:"pageSize"`
}

type GetWalletList struct {
	Page     int `json:"page" example:"page"`
	PageSize int `json:"pageSize" example:"pageSize"`
}
