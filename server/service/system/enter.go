package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	StartpayWeb3Service
	CasbinService
	InitDBService
	AutoCodeService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	AutoCodeHistoryService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
	SysExportTemplateService
}
