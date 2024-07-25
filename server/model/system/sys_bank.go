package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type UserBank struct {
	global.GVA_MODEL
	MerchantId      int64  `json:"merchantId"  gorm:"column:merchantId" `
	Region          string `json:"region" gorm:"default:;comment:region"`
	RemittanceType  string `json:"remittanceType" gorm:"column:remittanceType"`
	EnterpriseTitle string `json:"enterpriseTitle"  gorm:"column:enterpriseTitle"`
	BankTitle       string `json:"bankTitle" gorm:"default:;comment:bankTitle"`
	BankCode        string `json:"bankCode"  gorm:"column:bankCode"`
	FedWire         string `json:"fedWire"  gorm:"column:fedWire"`
	ReceiverNumber  string `json:"receiverNumber"  gorm:"column:receiverNumber"`
	ReceiverName    string `json:"receiverName" gorm:"column:receiverName"`
	ReceiverAddress string `json:"receiverAddress" gorm:"column:receiverAddress"`
	ReferenceField  string `json:"referenceField" gorm:"column:referenceField"`
}

func (UserBank) TableName() string {
	return "user_bank"
}
