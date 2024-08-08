package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type UserTransferOrder struct {
	global.GVA_MODEL
	MerchantId    int64   `json:"merchantId"  gorm:"column:merchantId" `
	MerchantName  string  `json:"merchantName" gorm:"column:merchantName"`
	FromAddress   string  `json:"from_addres" gorm:"column:from_addres"`
	ProjetcId     string  `json:"projetcId" gorm:"column:projetcId"`
	Currency      string  `json:"currency" gorm:"column:currency"`
	Chain         string  `json:"chain"  gorm:"column:chain"`
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
	ToAddress          string `json:"to_address" gorm:"column:to_address"`
}

func (UserTransferOrder) TableName() string {
	return "user_transfer_order"
}
