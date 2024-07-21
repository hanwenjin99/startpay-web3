package request

// Register User register structure
type CreateProject struct {
	AssembleChain   string `json:"assembleChain" example:"assembleChain"`
	Name            string `json:"name" example:"name"`
	SettleCurrency  string `json:"settleCurrency" example:"settleCurrency"`
	AssembleAddress string `json:"assemble_address"  example:"assemble_address"`
}
