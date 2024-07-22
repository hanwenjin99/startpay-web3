package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	web3api "github.com/flipped-aurora/gin-vue-admin/server/utils/startpay"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StartpayWeb3Service struct{}

func (s *StartpayWeb3Service) CreateProject(u system.SysProject) (projectInter system.SysProject, err error) {
	var project system.SysProject
	if !errors.Is(global.GVA_DB.Where("user_id = ? and ProName=? and settle_currency=? and AssembleChain=?",
		u.UserId, u.ProName, u.SettleCurrency, u.AssembleChain).First(&project).Error, gorm.ErrRecordNotFound) { // 判断项目名是否注册
		return projectInter, errors.New("项目名已注册")
	}

	web3 := web3api.StartpayWeb3Api{}

	web3DataP, err := web3.CeateProject(u.AssembleChain, u.ProName, u.SettleCurrency)
	if err != nil {
		return u, err
	}
	u.AppKey = web3DataP.Data.AppKey
	u.AssembleAddress = web3DataP.Data.AssembleAddress
	u.ProUuid = web3DataP.Data.Id
	//u.CreatedAt = web3DataP.Data.CreateTime
	u.CallbackDomain = web3DataP.Data.CallbackDomain
	u.CallbackUrl = web3DataP.Data.CallbackUrl
	u.Status = web3DataP.Data.Status

	web3Datas, err := web3.GetProjectSecret(u.ProUuid)
	if err != nil {
		return u, err
	}

	u.AppSecret = web3Datas.Data

	err = global.GVA_DB.Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (s *StartpayWeb3Service) GetProjectList(userId uint, Page int, PageSize int) (*web3api.ProjectList, error) {
	web3 := web3api.StartpayWeb3Api{}

	var projectlist []system.SysProject

	_, err := global.GVA_DB.Where("user_id = ? ", userId).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户项目失败")
	}

	global.GVA_LOG.Error("test", zap.Any("userId", userId))

	stringProjectid := ""
	for index, pvalue := range projectlist {
		if len(projectlist)-1 == index {
			stringProjectid += pvalue.ProUuid
		} else {
			stringProjectid += pvalue.ProUuid + ","
		}
	}

	if stringProjectid == "" {
		return nil, errors.New("没有查询到符合条件的项目")
	}

	return web3.GetProjectList(Page, PageSize, stringProjectid, "ACTIVE")
}
