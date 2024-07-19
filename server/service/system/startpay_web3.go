package system

import (
	"github.com/gofrs/uuid/v5"
)

type StartpayWeb3Service struct{}

func (s *StartpayWeb3Service) CreateProject(uuid uuid.UUID, projectName string) (err error) {
	/*var project system.SysProject

	global.GVA_MODEL
	UUID            uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	ProUuid         string    `json:"pro_uuid"  gorm:"index;comment:项目UUID" `
	ProName         string    `json:"pro_name" gorm:"default:;comment:项目名称"`
	AppKey          string    `json:"app_key" gorm:"default:;comment:appKeyid"`
	AppSecret       string    `json:"app_secret"  gorm:"default:;comment:AppSecret"`
	SettleCurrency  string    `json:"settle_currency" gorm:"default:;comment:settleCurrency"`
	AssembleChain   string    `json:"assemble_chain"  gorm:"index;comment:assembleChain"`
	AssembleAddress string    `json:"assemble_address"  gorm:"index;comment:assembleAddress"`
	CallbackDomain  string    `json:"callback_domain"  gorm:"default:;comment:callbackDomain"`
	CallbackUrl     string    `json:"callback_url" gorm:"default:;comment:callbackUrl"`
	Status          string    `json:"status" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`

	user := &system.SysProject{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}

	web3 := web3api.StartpayWeb3Api{}

	web3.CeateProject("","",)
	web3.GetProjectSecret()

	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	*/
	return nil
}

func (s *StartpayWeb3Service) GetProjectList() (*string, error) {

	return nil, nil
}
