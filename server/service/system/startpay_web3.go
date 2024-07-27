package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	MyCommon "github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	web3api "github.com/flipped-aurora/gin-vue-admin/server/utils/startpay"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type StartpayWeb3Service struct {
}

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

	if web3DataP.Data.Status == "ACTIVE" {
		u.Status = 1
	} else {
		u.Status = 2
	}

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

func (s *StartpayWeb3Service) GetAccountInfo(userId uint) ([]web3api.GetAccountInfo, error) {

	var projectlist []system.SysProject
	_, err := global.GVA_DB.Where("user_id = ? ", userId).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户项目失败")
	}

	web3AccountList := make([]web3api.GetAccountInfo, 0)
	for _, pvalue := range projectlist {
		web3 := web3api.StartpayWeb3Api{ApiKey: pvalue.AppKey, ApiSecret: pvalue.AppSecret}
		webrResp, err := web3.GetAccountInfo()
		if err != nil {
			continue
		}
		for _, data := range webrResp.Data {
			data.Name = pvalue.ProName
			web3AccountList = append(web3AccountList, data)
		}

	}
	return web3AccountList, nil
}
func (s *StartpayWeb3Service) Web3TransferCreate(userId uint, request systemReq.CreateTransferRequest) (*string, error) {
	var projectlist []system.SysProject

	_, err := global.GVA_DB.Where("user_id = ? and id =? ", userId, request.ID).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户交易密钥失败")
	}

	for _, pvalue := range projectlist {
		web3 := web3api.StartpayWeb3Api{ApiKey: pvalue.AppKey, ApiSecret: pvalue.AppSecret}
		webrResp, err := web3.Web3TransferCreate(request)
		if err != nil {
			return nil, err
		}
		return &webrResp.Data.Id, nil
	}

	return nil, errors.New("交易失败")
}

func (s *StartpayWeb3Service) Web3TransferList(userId uint, request systemReq.GetWeb3Requst) (*systemRes.TransferListRespons, error) {
	var projectlist []system.SysProject

	_, err := global.GVA_DB.Where("user_id = ? ", userId).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户交易密钥失败")
	}

	transferRes := systemRes.TransferListRespons{}
	for _, pvalue := range projectlist {
		web3 := web3api.StartpayWeb3Api{ApiKey: pvalue.AppKey, ApiSecret: pvalue.AppSecret}
		webrResp, err := web3.Web3TransferList(pvalue.ProUuid, request)
		if err != nil {
			continue
		}
		for _, data := range webrResp.Content {
			tri := systemRes.TransferRecordInfo{}
			tri.ToAddress = data.ToAddress
			tri.Chain = data.Chain
			tri.Currency = data.Currency
			tri.Id = data.Id
			tri.Amount, _ = strconv.Atoi(data.Amount)
			tri.UsdPrice = 0
			tri.CurrencyIcon = MyCommon.WEB3TOKENLISTAll[data.Currency]
			tri.ChainIcon = MyCommon.WEB3CHAINLIST[data.Chain]
			tri.AmountUsd = 0
			tri.Status = data.Status
			tri.MerchantId = userId
			tri.CreateTime = data.CreateTime
			tri.Note = ""
			tri.UpdateTime = data.CreateTime
			tri.AuditRecordId = ""
			tri.ContractAddress = ""
			tri.ErrorMessage = ""
			tri.Fee = 0
			tri.FromAddress = data.FromAddress
			tri.Gas, _ = strconv.ParseFloat(data.Gas, 64)
			tri.GasPrice, _ = strconv.ParseFloat(data.GasPrice, 64)
			tri.GasToken = MyCommon.WEB3GASLIST[data.Chain]
			tri.Nonce = ""
			tri.Tag = ""
			tri.Txid = data.Txid
			tri.TxTime = data.TxTime
			transferRes.Content = append(transferRes.Content, tri)
		}

	}

	return &transferRes, nil
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

func (s *StartpayWeb3Service) GetDepositOrder(userId uint, request systemReq.GetWeb3Requst) (*systemRes.DepositOrederRespons, error) {
	var projectlist []system.SysProject

	_, err := global.GVA_DB.Where("user_id = ? ", userId).Find(&projectlist).Rows()

	if err != nil {
		return nil, errors.New("查询用户交易密钥失败")
	}

	depositRes := systemRes.DepositOrederRespons{}
	for _, pvalue := range projectlist {
		web3 := web3api.StartpayWeb3Api{ApiKey: pvalue.AppKey, ApiSecret: pvalue.AppSecret}
		webrResp, err := web3.GetDepositOrder(pvalue.AssembleAddress, request)
		if err != nil {
			continue
		}
		for _, data := range webrResp.Data {
			tri := systemRes.DepositOrederInfo{}
			tri.Id = data.Id
			tri.Amount, _ = strconv.ParseFloat(data.Amount, 64)
			tri.AssembleAddress = data.Address
			tri.FromAddress = data.FromAddress
			tri.AmountUsd, _ = strconv.ParseFloat(data.AmountUsd, 64)
			tri.CreateTime = data.CreateTime
			tri.Status = data.Status
			tri.Chain = data.Chain
			tri.MerchantId = userId
			tri.ChainIcon = MyCommon.WEB3CHAINLIST[data.Chain]
			tri.Asset = data.Token
			tri.Address = data.Address
			tri.ProjectName = pvalue.ProName
			tri.ProjectId = pvalue.ProUuid
			tri.MerchantAddressId = data.MerchantAddressId
			tri.AssetIcon = MyCommon.WEB3TOKENLISTAll[data.Token]
			tri.TxHash = data.TxHash
			tri.DepositAmountUsd = 0
			tri.DepositAmount, _ = strconv.ParseFloat(data.DepositAmount, 64)
			depositRes.Content = append(depositRes.Content, tri)
		}

	}

	return &depositRes, nil
}

func (s *StartpayWeb3Service) GetbankAccountList(userId uint, Page int, PageSize int) (list []system.UserBank, total int64, err error) {

	limit := PageSize
	offset := PageSize * (Page - 1)
	db := global.GVA_DB.Model(&system.UserBank{})
	var bankList []system.UserBank
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Where("merchantId = ?", userId).Find(&bankList).Error
	return bankList, total, err
}

func (s *StartpayWeb3Service) UserContactList(userId uint, Page int, PageSize int) (list []system.Userwallet, total int64, err error) {

	limit := PageSize
	offset := PageSize * (Page - 1)
	db := global.GVA_DB.Model(&system.Userwallet{})
	var walletList []system.Userwallet
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Where("merchantId = ?", userId).Find(&walletList).Error
	return walletList, total, err
}

func (s *StartpayWeb3Service) BankAccountCreate(ub *system.UserBank) error {
	var userBank system.UserBank
	if !errors.Is(global.GVA_DB.Where(" merchantId = ? and region= ? and bankCode = ? and bankTitle = ? and receiverNumber = ? ", ub.MerchantId, ub.Region, ub.BankCode, ub.BankTitle, ub.ReceiverNumber).First(&userBank).Error, gorm.ErrRecordNotFound) {
		return errors.New("银行账户已添加")
	}
	err := global.GVA_DB.Create(&ub).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *StartpayWeb3Service) BankAccountDelete(ub *system.UserBank) error {
	if err := global.GVA_DB.Where("id = ? ", ub.ID).Delete(&system.UserBank{}).Error; err != nil {
		return errors.New("删除bank信息失败")
	}
	return nil
}

func (s *StartpayWeb3Service) UserContactCreate(uw *system.Userwallet) error {
	var userWallet system.Userwallet
	if !errors.Is(global.GVA_DB.Where(" merchantId = ? and address = ? and chain = ? and name=? ", uw.MerchantId, uw.Address, uw.Chain, uw.Name).First(&userWallet).Error, gorm.ErrRecordNotFound) {
		return errors.New("收款地址已添加")
	}
	err := global.GVA_DB.Create(&uw).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *StartpayWeb3Service) UserContactDelete(uw *system.Userwallet) error {

	if err := global.GVA_DB.Where("id = ? ", uw.ID).Delete(&system.Userwallet{}).Error; err != nil {
		return errors.New("删除收款地址失败")
	}

	return nil
}

func (s *StartpayWeb3Service) WithdrawOrderList(userId string, reInfo *systemReq.GetWithdrawOrderRequst) (list []system.UserWithDrawOrder, total int64, err error) {
	limit := reInfo.PageSize
	offset := reInfo.PageSize * (reInfo.Page - 1)
	db := global.GVA_DB.Model(&system.UserWithDrawOrder{})
	var uwoList []system.UserWithDrawOrder
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Where("merchantId = ?", userId).Find(&uwoList).Error
	return uwoList, total, err
}

func (s *StartpayWeb3Service) WithdrawOrderCreate(uwo *system.UserWithDrawOrder) error {
	err := global.GVA_DB.Create(&uwo).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *StartpayWeb3Service) AdminWithdrawOrderUpdate(req *systemReq.UpdateWithdrawOrderRequst) error {
	st, _ := strconv.Atoi(req.Status)
	uwo := &system.UserWithDrawOrder{
		Status:    st,
		AdminMemo: req.Memo,
	}
	uwo.UpdatedAt = time.Now()
	if err := global.GVA_DB.Where("id = ? ", req.Id).Updates(uwo).Error; err != nil {
		return errors.New("更新取现订单失败")
	}

	return nil
}

func (s *StartpayWeb3Service) WithdrawOrderUpdate(req *systemReq.UpdateWithdrawOrderRequst) error {
	st, _ := strconv.Atoi(req.Status)
	uwo := &system.UserWithDrawOrder{
		Status: st,
		Memo:   req.Memo,
	}
	uwo.UpdatedAt = time.Now()
	if err := global.GVA_DB.Where("id = ? ", req.Id).Updates(uwo).Error; err != nil {
		return errors.New("更新取现订单失败")
	}

	return nil
}
