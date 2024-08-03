package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type UserContractInfo struct {
	global.GVA_MODEL
	MerchantId int64 `json:"feebound" gorm:"column:feebound"`
	Maxbound   int64 `json:"max_bound" gorm:"column:max_bound"`

	WithdrawEnable int `json:"withdraw_enable" gorm:"column:withdraw_enable"`
	ChargeEnable   int `json:"charge_enable" gorm:"column:charge_enable"`
	TransferEnable int `json:"transfer_enable" gorm:"column:transfer_enable"`

	currency          string  `json:"currency"  gorm:"column:currency" `
	Chain             string  `json:"chain" gorm:"column:chain"`
	WithdrawFeeamount float64 `json:"withdraw_feeamount" gorm:"column:withdraw_feeamount"`
	DespositFeeamount float64 `json:"desposit_feeamount" gorm:"column:desposit_feeamount"`
	ChargeFeeamount   float64 `json:"charge_feeamount" gorm:"column:charge_feeamount"`
	TransferFeeamount float64 `json:"transfer_feeamount" gorm:"column:transfer_feeamount"`
	WithdrawFeerate1  float64 `json:"withdraw_feerate1" gorm:"column:withdraw_feerate1"`
	WithdrawFeerate2  float64 `json:"withdraw_feerate2" gorm:"column:withdraw_feerate2"`
	DespositFeerate1  float64 `json:"desposit_feerate1" gorm:"column:desposit_feerate1"`
	DespositFeerate2  float64 `json:"desposit_feerate2" gorm:"column:desposit_feerate2"`
	ChargeFeerate1    float64 `json:"charge_feerate1" gorm:"column:charge_feerate1"`
	ChargeFeerate2    float64 `json:"charge_feerate2" gorm:"column:charge_feerate2"`
	TransferFeerate1  float64 `json:"transfer_feerate1" gorm:"column:transfer_feerate1"`
	TransferFeerate2  float64 `json:"transfer_feerate2" gorm:"column:transfer_feerate2"`
}

func (UserContractInfo) TableName() string {
	return "user_contract"
}
