package system

import (
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type StartpayWeb3Api struct{}

func (b *StartpayWeb3Api) CreateProject(c *gin.Context) {
	var r systemReq.CreateProject
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Error("test", zap.Any("CreateProject", r))
	err = utils.Verify(r, utils.CreateProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	UserId := utils.GetUserID(c)

	project := &system.SysProject{AssembleChain: r.AssembleChain, ProName: r.Name, SettleCurrency: r.SettleCurrency, AssembleAddress: r.AssembleAddress}
	project.UserId = UserId

	global.GVA_LOG.Error("test", zap.Any("project", project))

	ProjectReturn, err := StartpayWeb3Service.CreateProject(*project)

	global.GVA_LOG.Error("test", zap.Any("ProjectReturn", ProjectReturn))

	if err != nil {
		global.GVA_LOG.Error("创建项目失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysProjectResponse{Project: ProjectReturn}, "创建项目失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysProjectResponse{Project: ProjectReturn}, "创建项目成功", c)
}

// func (b*StartpayWeb3Api) GetWalletList(c *gin.Context, user system.SysUser) {
func (b *StartpayWeb3Api) GetProjectList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("test", zap.Any("GetProjectList", r),
			zap.Any("err", err.Error()),
		)
		//response.FailWithMessage(err.Error(), c)
		//return
	}

	global.GVA_LOG.Error("test", zap.Any("GetProjectList", r))

	err = utils.Verify(r, utils.GetProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("test", zap.Any("Verify", err.Error()))
		return
	}

	userId := utils.GetUserID(c)

	ProjectReturn, err := StartpayWeb3Service.GetProjectList(userId, r.Page, r.PageSize)
	if err != nil {
		global.GVA_LOG.Error("获取项目失败!", zap.Error(err))
		response.FailWithDetailed(ProjectReturn.Data, "获取项目失败", c)
		return
	}

	response.OkWithDetailed(ProjectReturn.Data, "获取项目成功", c)
}

func (b *StartpayWeb3Api) GetWalletList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindJSON(&r)
	if err != nil {
		//response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("test welcome", zap.Any("err", err.Error()))
		//return
	}
	global.GVA_LOG.Error("test", zap.Any("GetWalletList", r))
	err = utils.Verify(r, utils.GetWalletVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)

	global.GVA_LOG.Error("test", zap.Any("userId111", userId))

	ProjectReturn, err := StartpayWeb3Service.GetProjectList(userId, r.Page, r.PageSize)
	if err != nil {
		global.GVA_LOG.Error("获取钱包失败!", zap.Error(err))
		response.FailWithDetailed(ProjectReturn.Message, "获取钱包失败", c)
		return
	}
	WalletResp := systemRes.GetWalletRespons{}

	WalletResp.Page = ProjectReturn.Data.Page
	WalletResp.PageSize = ProjectReturn.Data.PageSize
	WalletResp.Last = ProjectReturn.Data.Last
	WalletResp.TotalPages = ProjectReturn.Data.TotalPages
	WalletResp.Total = ProjectReturn.Data.Total

	for _, conv := range ProjectReturn.Data.Content {
		wlc := systemRes.WalletInfo{}
		wlc.Id = conv.Id
		wlc.Address = conv.AssembleAddress
		wlc.Chain = conv.AssembleChain
		wlc.MerchantAddressId = userId
		wlc.ProjectId = &conv.Id
		wlc.ProjectName = &conv.Name
		wlc.CreateTime = time.Unix(conv.CreateTime, 0)
		WalletResp.Content = append(WalletResp.Content, wlc)
	}

	response.OkWithDetailed(WalletResp, "获取钱包成功", c)
}

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *StartpayWeb3Api) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (b *StartpayWeb3Api) GetAccountInfo(c *gin.Context) {

	AccountReturn, err := StartpayWeb3Service.GetAccountInfo()
	if err != nil {
		global.GVA_LOG.Error("获取account失败!", zap.Error(err))
		response.FailWithDetailed(AccountReturn.Message, "获取account失败", c)
		return
	}
	//AccountInfoResp := systemRes.GetAccountInfoRespons{}
	response.OkWithDetailed(AccountReturn.Data, "获取account成功", c)
}

func (b *StartpayWeb3Api) GetChainListInfo(c *gin.Context) {
	ChainReturn, err := StartpayWeb3Service.GetChainListInfo()
	if err != nil {
		global.GVA_LOG.Error("获取链列表失败!", zap.Error(err))
		response.FailWithDetailed(ChainReturn.Message, "获取链列表失败!", c)
		return
	}
	response.OkWithDetailed(ChainReturn.Data, "获取链列表成功", c)
}

func (b *StartpayWeb3Api) GetDepositAddress(c *gin.Context) {

	//userId := utils.GetUserID(c)
}
func (b *StartpayWeb3Api) GetAdepositOrder(c *gin.Context) {
	//userId := utils.GetUserID(c)
}

func (b *StartpayWeb3Api) GetbankAccountList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)
	list, _, err := StartpayWeb3Service.GetbankAccountList(userId, r.Page, r.PageSize)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	blist := make([]systemRes.UserBankRespons, 0)

	for _, uwvalue := range list {
		ubs := systemRes.UserBankRespons{}
		ubs.ReceiverNumber = uwvalue.ReceiverNumber
		ubs.Id = strconv.FormatInt(int64(uwvalue.ID), 10)
		ubs.MerchantId = strconv.FormatInt(int64(uwvalue.MerchantId), 10)
		ubs.BankTitle = uwvalue.BankTitle
		ubs.BankCode = uwvalue.BankCode
		ubs.Region = uwvalue.Region
		ubs.FedWire = uwvalue.FedWire
		ubs.EnterpriseTitle = uwvalue.EnterpriseTitle
		ubs.ReceiverAddress = uwvalue.ReceiverAddress
		ubs.ReceiverName = uwvalue.ReceiverName
		ubs.CreateTime = uwvalue.CreatedAt
		ubs.ReferenceField = uwvalue.ReferenceField
		ubs.RemittanceType = uwvalue.RemittanceType
		ubs.UpdateTime = uwvalue.UpdatedAt
		blist = append(blist, ubs)
	}

	response.OkWithDetailed(blist, "获取成功", c)
}

func (b *StartpayWeb3Api) BankAccountCreate(c *gin.Context) {
	var r systemReq.CreateBank
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)

	userBank := &system.UserBank{BankTitle: r.BankTitle, ReceiverName: r.ReceiverName, ReceiverAddress: r.ReceiverAddress, BankCode: r.BankCode,
		EnterpriseTitle: r.EnterpriseTitle, FedWire: r.FedWire, ReceiverNumber: r.ReceiverNumber, Region: r.Region, RemittanceType: r.RemittanceType}
	userBank.MerchantId = int64(userId)

	err = StartpayWeb3Service.BankAccountCreate(userBank)
	if err != nil {
		global.GVA_LOG.Error("获取 xxx 失败!", zap.Error(err))
		response.FailWithDetailed("false", "添加 bank 失败", c)
		return
	}
	response.OkWithDetailed("true", "添加bank成功", c)
}

func (b *StartpayWeb3Api) BankAccountDelete(c *gin.Context) {
	var r systemReq.DeleteDBank
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("BankAccountDelete web3 db before", zap.Any("BankAccountDelete", r))
	userbank := &system.UserBank{}
	u64, err := strconv.ParseUint(r.Id, 10, 64)
	if err != nil {
		global.GVA_LOG.Error("删除bank id错误!", zap.Error(err))
		response.FailWithDetailed("false", "删除bank id错误", c)
		return
	}
	userbank.ID = uint(u64)

	err = StartpayWeb3Service.BankAccountDelete(userbank)
	if err != nil {
		global.GVA_LOG.Error("删除bank失败!", zap.Error(err))
		response.FailWithDetailed("false", "删除bank 失败", c)
		return
	}
	response.OkWithDetailed("true", "删除bank成功", c)
}

func (b *StartpayWeb3Api) UserContactList(c *gin.Context) {
	var r systemReq.GetCommonPageInfo
	r.Page = 1
	r.PageSize = 20

	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	global.GVA_LOG.Error("GetbankAccountList web3 db before", zap.Any("GetbankAccountList", r))
	userId := utils.GetUserID(c)
	list, total, err := StartpayWeb3Service.UserContactList(userId, r.Page, r.PageSize)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	uws := systemRes.UserWalletRespons{}

	for _, uwvalue := range list {
		uds := systemRes.UserDEsposit{}
		uds.Chain = uwvalue.Chain
		uds.Name = uwvalue.Name
		uds.Id = strconv.FormatInt(int64(uwvalue.ID), 10)
		uds.Address = uwvalue.Address
		uds.IsInternal = false
		uws.Content = append(uws.Content, uds)
	}
	uws.Page = r.Page
	uws.PageSize = r.PageSize
	uws.Total = total
	uws.TotalPages = int(total)/r.PageSize + 1
	uws.Last = false
	response.OkWithDetailed(uws, "获取成功", c)
}
func (b *StartpayWeb3Api) UserContactCreate(c *gin.Context) {
	var r systemReq.CreateDespoitAddress
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("UserContactCreate web3 db before", zap.Any("UserContactCreate", r))
	userId := utils.GetUserID(c)

	userAddress := &system.Userwallet{Address: r.Address, Name: r.Name, Chain: r.Chain}
	userAddress.MerchantId = int64(userId)

	err = StartpayWeb3Service.UserContactCreate(userAddress)
	if err != nil {
		global.GVA_LOG.Error("获取 xxx 失败!", zap.Error(err))
		response.FailWithDetailed("false", "添加 bank 失败", c)
		return
	}
	response.OkWithDetailed("true", "添加bank成功", c)
}
func (b *StartpayWeb3Api) UserContactDelete(c *gin.Context) {
	var r systemReq.DeleteDespoitAddress
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("xxx ShouldBindJSON fail", zap.Any("err", err.Error()))
	}

	/*err = utils.Verify(r, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/

	global.GVA_LOG.Error("UserContactDelete web3 db before", zap.Any("UserContactDelete", r))

	userAddress := &system.Userwallet{}
	u64, err := strconv.ParseUint(r.Id, 10, 64)
	if err != nil {
		global.GVA_LOG.Error("删除收款地址id错误!", zap.Error(err))
		response.FailWithDetailed("false", "删除收款地址id错误", c)
		return
	}
	userAddress.ID = uint(u64)

	err = StartpayWeb3Service.UserContactDelete(userAddress)
	if err != nil {
		global.GVA_LOG.Error("删除收款地址失败!", zap.Error(err))
		response.FailWithDetailed("false", "删除收款地址k 失败", c)
		return
	}
	response.OkWithDetailed("true", "删除收款地址成功", c)
}
