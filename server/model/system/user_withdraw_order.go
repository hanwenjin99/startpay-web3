package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type UserWithDrawOrder struct {
	global.GVA_MODEL
	MerchantId   int64  `json:"merchantId"  gorm:"column:merchantId" `
	MerchantName string `json:"merchantName" gorm:"default:;comment:merchantName"`
	BankId       string `json:"bankId" gorm:"default:;comment:bankId"`
	BankTitle    string `json:"bankTitle"  gorm:"default:;comment:bankTitle"`
	Currency     string `json:"currency" gorm:"default:;comment:currency"`
	Chain        string `json:"chain"  gorm:"default:;comment:chain"`

	Amount        float64 `json:"amount"  gorm:"index;comment:amount" `
	Fee           float64 `json:"fee" gorm:"default:;comment:fee"`
	RemittanceFee float64 `json:"remittanceFee" gorm:"default:;comment:remittanceFee"`
	TotalAmount   float64 `json:"totalAmount" gorm:"default:;comment:totalAmount"`

	Status             int    `json:"status"  gorm:"index;comment:status" `
	StatusName         string `json:"statusName" gorm:"default:;comment:statusName"`
	TxInfo             string `json:"txInfo" gorm:"default:;comment:txInfo"`
	TxCertificationUrl string `json:"txCertificationUrl"  gorm:"default:;comment:txCertificationUrl"`
	TxReference        string `json:"txReference" gorm:"default:;comment:txReference"`
	InputNote          string `json:"inputNote"  gorm:"index;comment:inputNote" `
	Supplier           string `json:"supplier" gorm:"default:;comment:supplier"`
	AdminMemo          string `json:"adminMemo" gorm:"default:;comment:adminMemo"`
	Memo               string `json:"memo" gorm:"default:;comment:memo"`
	AssembleAddress    string `json:"assemble_address" gorm:"default:;comment:assemble_address"`
}

func (UserWithDrawOrder) TableName() string {
	return "user_withdraw_order"
}
