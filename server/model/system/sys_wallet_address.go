package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Userwallet struct {
	global.GVA_MODEL
	MerchantId int64  `json:"merchantId"  gorm:"column:merchantId" `
	Name       string `json:"name" gorm:"default:;comment:name"`
	Address    string `json:"address" gorm:"default:;comment:address"`
	Chain      string `json:"chain"  gorm:"default:;comment:chain"`
	IsInternal string `json:"is_internal" gorm:"default:;comment:is_internal"`
}

func (Userwallet) TableName() string {
	return "user_wallet_address"
}
