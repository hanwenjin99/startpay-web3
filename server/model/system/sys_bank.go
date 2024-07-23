package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type UserBank struct {
	global.GVA_MODEL
	MerchantId      int64  `json:"merchantId"  gorm:"index;comment:merchantId" `
	Region          string `json:"region" gorm:"default:;comment:region"`
	RemittanceType  string `json:"remittanceType" gorm:"default:;comment:remittanceType"`
	EnterpriseTitle string `json:"enterpriseTitle"  gorm:"default:;comment:enterpriseTitle"`
	BankTitle       string `json:"bankTitle" gorm:"default:;comment:bankTitle"`
	BankCode        string `json:"bankCode"  gorm:"index;comment:bankCode"`
	FedWire         string `json:"fedWire"  gorm:"index;comment:fedWire"`
	ReceiverNumber  string `json:"receiverNumber"  gorm:"default:;comment:receiverNumber"`
	ReceiverName    string `json:"receiverName" gorm:"default:;comment:receiverName"`
	ReceiverAddress string `json:"receiverAddress" gorm:"default:;comment:receiverAddress"`
	ReferenceField  string `json:"referenceField" gorm:"default:;comment:referenceField"`
}

func (UserBank) TableName() string {
	return "user_bank"
}
