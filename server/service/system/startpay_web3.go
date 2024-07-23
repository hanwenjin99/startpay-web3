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
	if !errors.Is(global.GVA_DB.Where("user_id = ? and pro_name=? and settle_currency=? and assemble_chain=?",
		u.UserId, u.ProName, u.SettleCurrency, u.AssembleChain).First(&project).Error, gorm.ErrRecordNotFound) { // 判断项目名是否注册
		return projectInter, errors.New("项目名已注册")
	}

	global.GVA_LOG.Error("test", zap.Any("SysProject", u),
		zap.Any("project", project),
	)

	web3 := web3api.StartpayWeb3Api{}

	web3DataP, err := web3.CeateProject(u.AssembleChain, u.ProName, u.SettleCurrency)

	global.GVA_LOG.Error("test", zap.Any("web3DataP", web3DataP))
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

	global.GVA_LOG.Error("test", zap.Any("web3DataP", web3Datas))

	if err != nil {
		return u, err
	}

	u.AppSecret = web3Datas.Data

	err = global.GVA_DB.Create(&u).Error
	if err != nil {
		global.GVA_LOG.Error("test", zap.Any("create project db err", err.Error()))
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

	global.GVA_LOG.Error("test", zap.Any("userId", userId),
		zap.Any("Page", Page),
		zap.Any("PageSize", PageSize),
		zap.Any("projectlist", projectlist),
	)

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

	return web3.GetProjectList(Page, PageSize, "ACTIVE", stringProjectid)
}

func (s *StartpayWeb3Service) GetAccountInfo() (*web3api.GetAccountInfoRespons, error) {
	web3 := web3api.StartpayWeb3Api{}
	return web3.GetAccountInfo()
}

func (s *StartpayWeb3Service) GetChainListInfo() (*web3api.Web3ChainListRespons, error) {
	web3 := web3api.StartpayWeb3Api{}
	return web3.GetChainListInfo()
}

func (s *StartpayWeb3Service) GetDepositAddress(userId uint, Page int, PageSize int) (*web3api.ProjectList, error) {
	web3 := web3api.StartpayWeb3Api{}

	var projectlist []system.SysProject

	_, err := global.GVA_DB.Where("user_id = ? ", userId).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户项目失败")
	}

	global.GVA_LOG.Error("test", zap.Any("userId", userId),
		zap.Any("Page", Page),
		zap.Any("PageSize", PageSize),
		zap.Any("projectlist", projectlist),
	)

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

	return web3.GetProjectList(Page, PageSize, "ACTIVE", stringProjectid)
}

func (s *StartpayWeb3Service) GetAdepositOrder(userId uint, Page int, PageSize int) (*web3api.ProjectList, error) {
	web3 := web3api.StartpayWeb3Api{}

	var projectlist []system.SysProject

	_, err := global.GVA_DB.Where("user_id = ? ", userId).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户项目失败")
	}

	global.GVA_LOG.Error("test", zap.Any("userId", userId),
		zap.Any("Page", Page),
		zap.Any("PageSize", PageSize),
		zap.Any("projectlist", projectlist),
	)

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

	return web3.GetProjectList(Page, PageSize, "ACTIVE", stringProjectid)
}
