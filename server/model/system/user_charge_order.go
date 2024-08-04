package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type UserChargeOrder struct {
	global.GVA_MODEL
	MerchantId   int64  `json:"merchantId"  gorm:"column:merchantId" `
	MerchantName string `json:"merchantName" gorm:"column:merchantName"`
	BankId       string `json:"bankId" gorm:"column:bankId"`
	ProjetcId    string `json:"projetcId" gorm:"column:projetcId"`
	BankTitle    string `json:"bankTitle"  gorm:"column:bankTitle"`
	Currency     string `json:"currency" gorm:"column:currency"`
	Chain        string `json:"chain"  gorm:"column:chain"`

	Amount        float64 `json:"amount"  gorm:"column:amount" `
	Fee           float64 `json:"fee" gorm:"column:fee"`
	RemittanceFee float64 `json:"remittanceFee" gorm:"column:remittanceFee"`
	TotalAmount   float64 `json:"totalAmount" gorm:"column:totalAmount"`

	Status             int    `json:"status"  gorm:"column:status" `
	StatusName         string `json:"statusName" gorm:"column:statusName"`
	TxInfo             string `json:"txInfo" gorm:"column:txInfo"`
	TxCertificationUrl string `json:"txCertificationUrl"  gorm:"column:txCertificationUrl"`
	TxReference        string `json:"txReference" gorm:"column:txReference"`
	InputNote          string `json:"inputNote"  gorm:"column:inputNote" `
	Supplier           string `json:"supplier" gorm:"column:supplier"`
	AdminMemo          string `json:"adminMemo" gorm:"column:adminMemo"`
	Memo               string `json:"memo" gorm:"column:memo"`
	AssembleAddress    string `json:"assemble_address" gorm:"column:assemble_address"`
}

func (UserChargeOrder) TableName() string {
	return "user_charge_order"
}
