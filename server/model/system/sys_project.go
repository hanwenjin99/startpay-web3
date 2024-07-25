package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysProject struct {
	global.GVA_MODEL
	UserId          uint   `json:"user_id" gorm:"index;comment:用户user_id"`
	ProUuid         string `json:"pro_uuid"  gorm:"index;comment:项目ID" `
	ProName         string `json:"pro_name" gorm:"default:;comment:项目名称"`
	AppKey          string `json:"app_key" gorm:"default:;comment:appKeyid"`
	AppSecret       string `json:"app_secret"  gorm:"default:;comment:AppSecret"`
	SettleCurrency  string `json:"settle_currency" gorm:"default:;comment:settleCurrency"`
	AssembleChain   string `json:"assemble_chain"  gorm:"index;comment:assembleChain"`
	AssembleAddress string `json:"assemble_address"  gorm:"index;comment:assembleAddress"`
	CallbackDomain  string `json:"callback_domain"  gorm:"default:;comment:callbackDomain"`
	CallbackUrl     string `json:"callback_url" gorm:"default:;comment:callbackUrl"`
	Status          int    `json:"status" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

func (SysProject) TableName() string {
	return "sys_project"
}
