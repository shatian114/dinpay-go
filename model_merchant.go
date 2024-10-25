package dinpay

type BaseMerchantRes[T any] struct {
	Success  bool   `json:"success"`        // 响应码
	Code     string `json:"code"`           // 响应码
	Message  string `json:"message"`        // 返回信息
	Data     T      `json:"data,omitempty"` // 响应实体
	Sign     string `json:"sign,omitempty"` // 签名
	Hostname string `json:"hostname"`       // 处理机别名
}

// MerchantInfoQueryReq 商户信息查询
type MerchantInfoQueryReq struct {
	FirstClassMerchantNo string `json:"firstClassMerchantNo"` // 平台商商编,智付下发的,以C开头的商户编号
	MerchantNo           string `json:"merchantNo"`           // 子商户号
}

// MerchantInfoQueryRes 商户信息查询
type MerchantInfoQueryRes struct {
	MerchantNo       string `json:"merchantNo"`                 // 子商户商编
	SignName         string `json:"signName"`                   // 子商户签约名
	ShowName         string `json:"showName"`                   // 展示名（收银台展示名）
	CreateDate       string `json:"createDate"`                 // 入网时间
	MerchantType     string `json:"merchantType"`               // 子商户类型
	LegalPerson      string `json:"legalPerson"`                // 法人名字
	LegalPersonID    string `json:"legalPersonID"`              // 法人身份证号
	BusinessLicense  string `json:"businessLicense"`            // 营业执照号
	Province         string `json:"province"`                   // 子商户所在省份
	City             string `json:"city"`                       // 子商户所在城市
	Address          string `json:"address"`                    // 通讯地址
	Linkman          string `json:"linkman"`                    // 联系人
	LinkPhone        string `json:"linkPhone"`                  // 联系电话
	Email            string `json:"email"`                      // 联系邮箱
	MerchantStatus   string `json:"merchantStatus"`             // 商户状态
	AccountStatus    string `json:"accountStatus"`              // 账户状态
	MerchantCategory string `json:"merchantCategory"`           // 经营类别
	OrgNum           string `json:"orgNum,omitempty"`           // 组织机构代码
	Region           string `json:"region,omitempty"`           // 子商户所在区
	BindMobile       string `json:"bindMobile,omitempty"`       // 绑定手机
	IdCardStartDate  string `json:"idCardStartDate,omitempty"`  // 身份证有效起始日期
	IdCardEndDate    string `json:"idCardEndDate,omitempty"`    // 身份证有效终止日期
	BusLiceStartDate string `json:"busLiceStartDate,omitempty"` // 营业执照有效起始日期
	BusLiceEndDate   string `json:"busLiceEndDate,omitempty"`   // 营业执照有效终止日期
	PostalAddress    string `json:"postalAddress,omitempty"`    // 注册地址
	ServicePhone     string `json:"servicePhone,omitempty"`     // 客服联系电话
}

// MerchantRegisterReq 商户进件
type MerchantRegisterReq struct {
	FirstClassMerchantNo     string   `json:"firstClassMerchantNo"`             // 平台商商编,智付下发的,以C开头的商户编号
	OrderNo                  string   `json:"orderNo"`                          // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SignName                 string   `json:"signName"`                         // 子商户签约名,营业执照上的工商注册名称(个人商户除外),注:1.个人商户请传“商户_个人姓名”;签约名至少为5个字,不能超过150个字
	ShowName                 string   `json:"showName"`                         // 展示名,用于收银台的展示名,进件测试的可在展示名中添加“测试”字样; 注意:展示名长度不能小于2大于20
	WebSite                  string   `json:"webSite,omitempty"`                // 商户网站地址
	AccessUrl                string   `json:"accessUrl,omitempty"`              // 接入地址,空
	MerchantType             string   `json:"merchantType"`                     // 子商户类型,MerchantType,个人、企业、个体商户交易限额不同,切勿传错
	LegalPerson              string   `json:"legalPerson"`                      // 法人名字
	LegalPersonId            string   `json:"legalPersonID"`                    // 身份证
	OrgNum                   string   `json:"orgNum"`                           // 1.若是个人商户,统一填写身份证号;2.三证合一,则统一填营业执照号
	BusinessLicense          string   `json:"businessLicense"`                  // 执照号
	Province                 string   `json:"province,omitempty"`               // 省份
	City                     string   `json:"city,omitempty"`                   // 城市
	RegionCode               string   `json:"regionCode"`                       // 区县编码
	Address                  string   `json:"address"`                          // 通讯地址
	Linkman                  string   `json:"linkman"`                          // 联系人
	LinkPhone                string   `json:"linkPhone"`                        // 联系电话
	Email                    string   `json:"email"`                            // 各子商户邮箱地址唯一（可不用于接收邮件）
	BindMobile               string   `json:"bindMobile,omitempty"`             // 绑定手机号
	ServicePhone             string   `json:"servicePhone,omitempty"`           // 用户支付后有疑问的,可通过此号码进行咨询,如不填会上送默认电话。如需向银行确认联行号信息,建议必填
	BankCode                 string   `json:"bankCode"`                         // 联行号, 注：需商户电话联系银行核实具体联行号
	AccountName              string   `json:"accountName"`                      // 结算银行卡开户名
	AccountNo                string   `json:"accountNo"`                        // 结算银行卡账号
	SettleBankType           string   `json:"settleBankType"`                   // 结算卡类型,SettleBankType;TOPRIVATE:对私,TOPUBLIC:对公
	SettlementPeriod         string   `json:"settlementPeriod"`                 // 结算类型,SettlementPeriod;T1:工作日隔天结算,D1:自然日隔天结算,D0:当日结算（目前不支持）
	SettlementMode           string   `json:"settlementMode"`                   // 结算方式;SettlementMode;NOTOPEN:不开通;AUTO:自动结算;SELF:自主结算;
	SettlementRemark         string   `json:"settlementRemark,omitempty"`       // 结算备注,自动结算/自主结算所使用的备注信息
	MerchantCategory         string   `json:"merchantCategory"`                 // 经营类别
	IndustryTypeCode         string   `json:"industryTypeCode,omitempty"`       // 行业类型编码
	AuthorizationFlag        bool     `json:"authorizationFlag"`                // 授权使用平台商秘钥
	UnionPayQrCode           string   `json:"unionPayQrCode,omitempty"`         // 子商户若需绑定银联二维码,可填写,注:1.借/贷记卡必须同时配置;区间最大值不低于50000;例如需求区间为1-1000,也必须设置为1-50000
	NeedPosFunction          bool     `json:"needPosFunction"`                  // 是否需要开通 POS 功能
	IdCardStartDate          string   `json:"idCardStartDate,omitempty"`        // 法人身份证开始日期,yyyyMMdd,needPosFunction为true则必传,开通微信产品必传
	IdCardEndDate            string   `json:"idCardEndDate,omitempty"`          // 法人身份证结束日期,yyyyMMdd或者长期有效,needPosFunction为true则必传,开通微信产品必传
	BusinessDateStart        string   `json:"businessDateStart,omitempty"`      // 经营起始日期,yyyyMMdd,needPosFunction为true则必传,开通微信产品必传
	BusinessDateLimit        string   `json:"businessDateLimit,omitempty"`      // 经营期限,yyyyMMdd,needPosFunction为true则必传,开通微信产品必传
	AccountIdCard            string   `json:"accountIdCard,omitempty"`          // 开户人身份证号,needPosFunction为true则必传
	Mcc                      string   `json:"mcc,omitempty"`                    // 银联MCC码,needPosFunction为true则必传
	AgreeProtocol            bool     `json:"agreeProtocol"`                    // 是否同意协议
	CallbackUrl              string   `json:"callbackUrl,omitempty"`            // 人工审核后异步回调地址(返回参数为JSON数据)注:人工审核成功或失败会异步通知;自动审核通过,不会通知
	SettleMode               string   `json:"settleMode"`                       // MERCHANT:按商户结算;MERGE:按结算人结算,同平台商下,结算卡号一致的结算单合并成一笔出款
	SettlementAuth           string   `json:"settlementAuth,omitempty"`         // 对私结算信息鉴权时使用(平台商需开通鉴权产品):STOCK存量,不鉴权;INCREASE新增,鉴权;不传参数默认不鉴权
	PostalAddress            string   `json:"postalAddress,omitempty"`          // 注册地址,营业执照中注册地址
	MicroBizType             string   `json:"microBizType,omitempty"`           // 小微经营类型
	CertType                 string   `json:"certType,omitempty"`               // 证书类型,事业单位必填
	LinkManId                string   `json:"linkManId,omitempty"`              // 联系人身份证号
	NeedAuthorize            bool     `json:"needAuthorize,omitempty"`          // 是否需要认证
	SpecialSignName          bool     `json:"specialSignName,omitempty"`        // 是否需要特殊处理商户名称,根据微信定义的签约名规则去报备只有营业执照名字为“*”，或者个人商户才传true
	IdType                   string   `json:"idType"`                           // 法人证件类型,IdType
	Longitude                string   `json:"longitude,omitempty"`              // 经度
	Latitude                 string   `json:"latitude,omitempty"`               // 纬度
	ElectronicAccountNo      string   `json:"electronicAccountNo,omitempty"`    // 电子账户
	SettleChangeType         string   `json:"settleChangeType,omitempty"`       // 结算到账方式,SettleChangeType
	SettlementIdCardNo       string   `json:"settlementIdCardNo,omitempty"`     // 结算人身份证号
	SettlementPhoneNo        string   `json:"settlementPhoneNo,omitempty"`      // 结算人手机号
	ServiceCodes             []string `json:"serviceCodes,omitempty"`           // 支付宝商户服务类型
	LinkmanType              string   `json:"linkmanType"`                      // 联系人类型,IdHolderType
	LinkmanIdType            string   `json:"linkmanIdType,omitempty"`          // 联系人证件类型,IdType
	LinkmanIdCardStartDate   string   `json:"linkmanIdCardStartDate,omitempty"` // 联系人证件有效期开始时间,yyyyMMdd,联系人类型为经办人必填
	LinkmanIdCardEndDate     string   `json:"linkmanIdCardEndDate,omitempty"`   // 联系人证件有效期结束时间,yyyyMMdd或长期有效,联系人类型为经办人必填
	IdHolderType             string   `json:"idHolderType,omitempty"`           // 证件持有人类型,子商户类型为事业单位时必填,企业商户\个体工商户\个人商户\其他组织默认为经营者/法人
	LegalPersonIdAddress     string   `json:"legalPersonIdAddress,omitempty"`   // 法人证件居住地址,企业商户必填
	EnterpriseOwner          bool     `json:"enterpriseOwner,omitempty"`        // 经营者/法人是否为受益人,企业商户必填
	BenefitLegalPersonIdType string   `json:"benefLegalPersonIdType,omitempty"` // 受益人证件类型,企业商户必填,IdType
	BenefitLegalPerson       string   `json:"benefLegalPerson,omitempty"`       // 受益人证件姓名,企业商户必填
	BenefitAddress           string   `json:"benefAddress,omitempty"`           // 受益人证件居住地址
	BenefitIdCardStartDate   string   `json:"benefIdCardStartDate,omitempty"`   // 受益人证件有效期开始时间,yyyyMMdd
	BenefitIdCardEndDate     string   `json:"benefIdCardEndDate,omitempty"`     // 受益人证件有效期结束时间,yyyyMMdd或长期有效
}

// MerchantRegisterRes 商户进件
type MerchantRegisterRes struct {
	OrderNo     string `json:"orderNo"`              // 订单号 进件时的订单号
	EntryStatus string `json:"entryStatus"`          // 进件状态 INIT:待审核,OVERRULE:申请驳回;AUDITED:审核通过
	Msg         string `json:"msg,omitempty"`        // 进件状态说明 entryStatus为 INIT 时返回
	MerchantNo  string `json:"merchantNo,omitempty"` // 子商户商编 只有进件状态为审核通过才返回
}

// MerchantRegisterQueryReq 商户进件查询
type MerchantRegisterQueryReq struct {
	OrderNo              string `json:"orderNo"`              // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	FirstClassMerchantNo string `json:"firstClassMerchantNo"` // 平台商商编,智付下发的,以C开头的商户编号
}

// MerchantRegisterQueryRes 商户进件查询
type MerchantRegisterQueryRes struct {
	OrderNo           string  `json:"orderNo"`                 // 订单号 进件时的订单号
	Status            string  `json:"status"`                  // 进件状态 见附录5.6(INIT,待审核.OVERRULE,申请驳回.AUDITED,审核通过.)
	Msg               string  `json:"msg,omitempty"`           // 审核信息
	MerchantNo        string  `json:"merchantNo,omitempty"`    // 子商户商编 进件状态为审核通过才返回
	SignName          string  `json:"signName"`                // 子商户签约名 签约名
	ShowName          string  `json:"showName"`                // 展示名(收银台展示名)用于收银台的展示名
	WebSite           string  `json:"webSite,omitempty"`       // 网站网址 商户网站地址
	AccessUrl         string  `json:"accessUrl,omitempty"`     // 接入地址 接入地址，空
	MerchantType      string  `json:"merchantType"`            // 子商户类型 见附录5.7
	LegalPerson       string  `json:"legalPerson"`             // 法人名字 法人名字
	LegalPersonID     string  `json:"legalPersonID"`           // 法人身份证号 身份证
	OrgNum            string  `json:"orgNum"`                  // 组织机构代码 机构号
	BusinessLicense   string  `json:"businessLicense"`         // 营业执照号 执照号
	Province          string  `json:"province"`                // 子商户所在省份 省份
	City              string  `json:"city"`                    // 子商户所在城市 城市
	Address           string  `json:"address"`                 // 通讯地址 通讯地址
	Linkman           string  `json:"linkman"`                 // 联系人 联系人
	LinkPhone         string  `json:"linkPhone"`               // 联系电话 联系电话
	Email             string  `json:"email"`                   // 联系邮箱 用户邮箱
	BindMobile        string  `json:"bindMobile"`              // 绑定手机 绑定手机号
	ReserveDays       int     `json:"reserveDays"`             // 结算周期(天) 结算周期以天为单位
	BankCode          string  `json:"bankCode"`                // 结算卡联行号 联行号
	BankName          string  `json:"bankName"`                // 结算卡银行名称 银行名
	BankBranch        string  `json:"bankBranch"`              // 结算卡银行分行名 分行名
	BankProv          string  `json:"bankProv"`                // 结算卡所属省 省份
	BankCity          string  `json:"bankCity"`                // 结算卡所属市 城市
	AccountName       string  `json:"accountName"`             // 结算账户名 结算银行卡开户名
	AccountNo         string  `json:"accountNo"`               // 结算账号 结算银行卡账号
	SettleBankType    string  `json:"settleBankType"`          // 结算卡类型 见附录5.9
	SettlementPeriod  string  `json:"settlementPeriod"`        // 结算类型 见附录5.10
	SettlementMode    string  `json:"settlementMode"`          // 结算方式 见附录5.15
	WithDrawPrice     float64 `json:"withDrawPrice,omitempty"` // 提现单笔收费(元) 按笔收费
	MerchantCategory  string  `json:"merchantCategory"`        // 经营类别 见附录5.5
	AuthorizationFlag bool    `json:"authorizationFlag"`       // 授权使用平台商秘钥 true代表授权，false代表不授权
	AgreeProtocol     bool    `json:"agreeProtocol"`           // 是否同意协议 true代表同意，false代表不同意
	ServicePhone      string  `json:"servicePhone,omitempty"`  // 客服联系电话 用户支付后有疑问的, 可通过此号码进行咨询
}

// MerchantOpenAppPayProductReq 商户开通扫码产品
type MerchantOpenAppPayProductReq struct {
	ProductType          string                        `json:"productType"`          // 产品类型:APPPAY:扫码产品
	FirstClassMerchantNo string                        `json:"firstClassMerchantNo"` // 平台商编号
	MerchantNo           string                        `json:"merchantNo"`           // 子商户编号
	PayType              string                        `json:"payType,omitempty"`    // 支付类型
	AppPayType           string                        `json:"appPayType"`           // 客户端类型
	Value                float64                       `json:"value,omitempty"`      // 费率
	MinFee               float64                       `json:"minFee,omitempty"`     // 最低费率金额
	AppFeeMode           string                        `json:"appFeeMode,omitempty"` // 费率模式
	FeePurpose           string                        `json:"feePurpose,omitempty"` // 费率类型
	FeeRanges            []MerchantOpenProductFeeRange `json:"feeRanges,omitempty"`  // 银联扫码/wap分段费率
}

type MerchantOpenProductFeeRange struct {
	Start          float64 `json:"start"`          // 开始费率区间值,闭区间
	End            float64 `json:"end"`            // 结束费率区间值,闭区间
	Fee            float64 `json:"fee"`            // 费率值,0.50,代表0.5%（单位,见收费模式）
	MinFee         float64 `json:"minFee"`         // 最低费率值,单位(元)
	MaxFee         float64 `json:"maxFee"`         // 最高费率金额,单位(元),小数后两位
	CalcType       string  `json:"calcType"`       // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	OnlineCardType string  `json:"onlineCardType"` // 卡类型,与平台商一致,DEBIT:借记卡;CREDIT:贷记;银联产品使用,其他产品不需要填写
}

// MerchantOpenAppPayProductRes 商户开通扫码产品
type MerchantOpenAppPayProductRes struct {
	FirstClassMerchantNo string                        `json:"firstClassMerchantNo"` // 平台商编号
	MerchantNo           string                        `json:"merchantNo"`           // 子商户编号
	PayType              string                        `json:"payType,omitempty"`    // 支付类型
	AppPayType           string                        `json:"appPayType"`           // 客户端类型
	Value                float64                       `json:"value,omitempty"`      // 费率,单位(%或元),小数后两位
	MinFee               float64                       `json:"minFee,omitempty"`     // 最低费率金额, 单位(元),小数后两位
	FeeRanges            []MerchantOpenProductFeeRange `json:"feeRanges,omitempty"`  // 银联扫码/wap分段费率
}

// MerchantAppPayProductQueryReq 商户扫码产品查询
type MerchantAppPayProductQueryReq struct {
	ProductType          string `json:"productType"`          // 产品类型:APPPAY:扫码产品
	FirstClassMerchantNo string `json:"firstClassMerchantNo"` // 平台商编号
	MerchantNo           string `json:"merchantNo"`           // 子商户编号
	PayType              string `json:"payType,omitempty"`    // 支付类型
	AppPayType           string `json:"appPayType"`           // 客户端类型
	AppFeeMode           string `json:"appFeeMode,omitempty"` // 费率模式
	FeePurpose           string `json:"feePurpose,omitempty"` // 费率类型
}

// MerchantAppPayProductQueryRes 商户扫码产品查询
type MerchantAppPayProductQueryRes struct {
	FirstClassMerchantNo string                        `json:"firstClassMerchantNo"` // 平台商编号
	MerchantNo           string                        `json:"merchantNo"`           // 子商户编号
	PayType              string                        `json:"payType,omitempty"`    // 支付类型
	AppPayType           string                        `json:"appPayType"`           // 客户端类型
	Value                float64                       `json:"value,omitempty"`      // 费率,单位(%或元),小数后两位
	MinFee               float64                       `json:"minFee,omitempty"`     // 最低费率金额, 单位(元),小数后两位
	FeeRanges            []MerchantOpenProductFeeRange `json:"feeRanges,omitempty"`  // 银联扫码/wap分段费率
}

// MerchantOpenSettlementProductReq 商户开通结算产品
type MerchantOpenSettlementProductReq struct {
	ProductType          string  `json:"productType"`          // 产品类型:SETTLEMENT:结算产品
	FirstClassMerchantNo string  `json:"firstClassMerchantNo"` // 平台商商编:平台商编号
	MerchantNo           string  `json:"merchantNo"`           // 子商户编号:进件审核通过后才有的商户号
	CalcType             string  `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	SettlementPeriod     string  `json:"settlementPeriod"`     // 资金周期:T1,D0,D1
	SettleBankType       string  `json:"settleBankType"`       // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE)
	Value                float64 `json:"value,omitempty"`      // 费率:单位(%或元),小数后两位(非必填)
	Floating             float64 `json:"floating,omitempty"`   // 浮动值:百分比(%),小数后两位(非必填)
	MinFee               float64 `json:"minFee,omitempty"`     // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantOpenSettlementProductRes 商户开通结算产品
type MerchantOpenSettlementProductRes struct {
	FirstClassMerchantNo string  `json:"firstClassMerchantNo"` // 平台商商编
	MerchantNo           string  `json:"merchantNo"`           // 子商户编号
	CalcType             string  `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	SettlementPeriod     string  `json:"settlementPeriod"`     // 资金周期:T1,D0,D1
	SettleBankType       string  `json:"settleBankType"`       // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE)
	Value                float64 `json:"value,omitempty"`      // 费率:单位(%或元),小数后两位(非必填)
	Floating             float64 `json:"floating,omitempty"`   // 浮动值:百分比(%),小数后两位(非必填)
	MinFee               float64 `json:"minFee,omitempty"`     // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantSettlementProductQueryReq 商户结算产品查询
type MerchantSettlementProductQueryReq struct {
	ProductType          string `json:"productType"`          // 产品类型:SETTLEMENT:结算产品
	FirstClassMerchantNo string `json:"firstClassMerchantNo"` // 平台商商编:平台商编号
	MerchantNo           string `json:"merchantNo"`           // 子商户编号:进件审核通过后才有的商户号
	CalcType             string `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	SettlementPeriod     string `json:"settlementPeriod"`     // 资金周期:T1,D0,D1
	SettleBankType       string `json:"settleBankType"`       // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE)
}

// MerchantSettlementProductQueryRes 商户结算产品查询
type MerchantSettlementProductQueryRes struct {
	FirstClassMerchantNo string  `json:"firstClassMerchantNo"` // 平台商商编
	MerchantNo           string  `json:"merchantNo"`           // 子商户编号
	CalcType             string  `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	SettlementPeriod     string  `json:"settlementPeriod"`     // 资金周期:T1,D0,D1
	SettleBankType       string  `json:"settleBankType"`       // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE)
	Value                float64 `json:"value,omitempty"`      // 费率:单位(%或元),小数后两位(非必填)
	Floating             float64 `json:"floating,omitempty"`   // 浮动值:百分比(%),小数后两位(非必填)
	MinFee               float64 `json:"minFee,omitempty"`     // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantOpenAccountPayProductReq 商户开通虚拟账户支付产品
type MerchantOpenAccountPayProductReq struct {
	ProductType          string  `json:"productType"`          // 产品类型:ACCOUNT_PAY:虚拟账户支付产品
	FirstClassMerchantNo string  `json:"firstClassMerchantNo"` // 平台商商编:平台商编号
	MerchantNo           string  `json:"merchantNo"`           // 子商户编号:进件审核通过后才有的商户号
	CalcType             string  `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Value                float64 `json:"value,omitempty"`      // 费率:单位(%或元),小数后两位(非必填)
	MinFee               float64 `json:"minFee,omitempty"`     // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantOpenAccountPayProductRes 商户开通虚拟账户支付产品
type MerchantOpenAccountPayProductRes struct {
	FirstClassMerchantNo string  `json:"firstClassMerchantNo"` // 平台商商编
	MerchantNo           string  `json:"merchantNo"`           // 子商户编号
	CalcType             string  `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Value                float64 `json:"value,omitempty"`      // 费率:单位(%或元),小数后两位(非必填)
	MinFee               float64 `json:"minFee,omitempty"`     // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantAccountPayProductQueryReq 商户虚拟账户支付产品查询
type MerchantAccountPayProductQueryReq struct {
	ProductType          string `json:"productType"`          // 产品类型:ACCOUNT_PAY:虚拟账户支付产品
	FirstClassMerchantNo string `json:"firstClassMerchantNo"` // 平台商商编:平台商编号
	MerchantNo           string `json:"merchantNo"`           // 子商户编号:进件审核通过后才有的商户号
	CalcType             string `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
}

// MerchantAccountPayProductQueryRes 商户虚拟账户支付产品查询
type MerchantAccountPayProductQueryRes struct {
	FirstClassMerchantNo string  `json:"firstClassMerchantNo"` // 平台商商编
	MerchantNo           string  `json:"merchantNo"`           // 子商户编号
	CalcType             string  `json:"calcType"`             // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Value                float64 `json:"value,omitempty"`      // 费率:单位(%或元),小数后两位(非必填)
	MinFee               float64 `json:"minFee,omitempty"`     // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantAppApplyReq APP报备
type MerchantAppApplyReq struct {
	MerchantNo  string `json:"merchantNo"`            // 子商户编号 进件审核通过后才有的商户号
	AppPayType  string `json:"appPayType"`            // 客户端类型 见附录5.2
	OrderNo     string `json:"orderNo"`               // 请求单号
	ApplyType   string `json:"applyType,omitempty"`   // 报备类型 见附录 5.24，不需报名的商户此字段不填
	CallBackUrl string `json:"callBackUrl,omitempty"` // 报名回调通知地址 报名成功/失败回调通知地址
}

// MerchantAppApplyRes APP报备
type MerchantAppApplyRes struct {
	MerchantNo string `json:"merchantNo"` // 子商户编号
	AppPayType string `json:"appPayType"` // 客户端类型
	OrderNo    string `json:"orderNo"`    // 请求单号,请求参数请求单号
	Status     string `json:"status"`     // 报备结果,只有DOING/SUCCESS两种状态
}

// MerchantAppApplyQueryReq APP报备查询
type MerchantAppApplyQueryReq struct {
	MerchantNo string `json:"merchantNo"`          // 子商户编号 进件审核通过后才有的商户号
	AppPayType string `json:"appPayType"`          // 客户端类型 见附录5.2
	OrderNo    string `json:"orderNo"`             // 请求单号
	ApplyType  string `json:"applyType,omitempty"` // 报备类型 见附录 5.24，不需报名的商户此字段不填
}

// MerchantAppApplyQueryRes APP报备查询
type MerchantAppApplyQueryRes struct {
	MerchantNo         string `json:"merchantNo"`         // 子商户编号
	AppPayType         string `json:"appPayType"`         // 客户端类型
	OrderNo            string `json:"orderNo"`            // 请求单号
	Status             string `json:"status"`             // 报备结果
	ReportFailedReason string `json:"reportFailedReason"` // 报备失败原因,当报备结果为DOING、FAIL 时可返回
	WeChantStatus      string `json:"weChantStatus"`      // 报名状态
	// 以下在status为SUCCESS时才返回
	ThreePartnerNoData []ThreePartnerNoData `json:"threePartnerNoData,omitempty"`
	UnionPayRecords    []UnionPayRecord     `json:"unionPayRecords,omitempty"`
}

// ThreePartnerNoData 第三方平台信息
type ThreePartnerNoData struct {
	ChannelMerchantNo  string `json:"channelMerchantNo"`  // 第三方平台渠道号
	ThreePartnerNo     string `json:"threePartnerNo"`     // 第三方交易识别码
	ChannelShortId     string `json:"channelShortId"`     // 渠道
	AuthorizeStatus    string `json:"authorizeStatus"`    // 授权状态
	ConfirmStatus      string `json:"confirmStatus"`      // 认证状态
	QrCodeData         string `json:"qrCodeData"`         // 联系人信息确认二维码,base64
	ReportFailedReason string `json:"reportFailedReason"` // 失败原因
	WxRuleId           string `json:"wxRuleId"`           // 微信结算规则ID
}

// UnionPayRecord 银联二维码信息
type UnionPayRecord struct {
	QrCode        string `json:"qrCode"`        // 银联固码
	SubMerchantNo string `json:"subMerchantNo"` // 二级商编
}

// MerchantWxpayReIdentifyReq 商户微信支付重新发起认证
type MerchantWxpayReIdentifyReq struct {
	FirstClassMerchantNo string `json:"firstClassMerchantNo"`    // 平台商编号
	MerchantNo           string `json:"merchantNo"`              // 子商户号
	ThreePartnerNo       string `json:"threePartnerNo"`          // 通过4.5.2 报备查询接口获取
	NeedAuthorize        *bool  `json:"needAuthorize,omitempty"` // 供商户报备微信成功后补认证
}

// MerchantWxpayReIdentifyRes 商户微信支付重新发起认证
type MerchantWxpayReIdentifyRes struct {
	Status   string `json:"status"`             // 重新认证状态
	ErrorMsg string `json:"errorMsg,omitempty"` // 认证失败返回错误原因
}

// MerchantNotifyReq 商户管理回调基本结构体
type MerchantNotifyReq[T any] struct {
	Success  bool   `json:"success"`        // 响应码
	Code     string `json:"code"`           // 响应码
	Message  string `json:"message"`        // 返回信息
	Data     T      `json:"data,omitempty"` // 响应实体
	Sign     string `json:"sign,omitempty"` // 签名
	Hostname string `json:"hostname"`       // 处理机别名
}

// MerchantRegisterNotifyReqBody 商户进件结果回调Body
type MerchantRegisterNotifyReqBody struct {
	OrderNo          string `json:"orderNo"`                    // 订单号 进件时的订单号
	EntryStatus      string `json:"entryStatus"`                // 进件状态 INIT:待审核;OVERRULE:申请驳回;AUDITED:审核通过
	Msg              string `json:"msg,omitempty"`              // 进件状态说明 当entryStatus为 INIT,OVERRULE 时返回
	MerchantNo       string `json:"merchantNo,omitempty"`       // 子商户商编 只有进件状态为审核通过才返回
	ParentMerchantNo string `json:"parentMerchantNo,omitempty"` // 平台商编 只有进件状态为审核通过才返回
}
