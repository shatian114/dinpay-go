package dinpay

// MerchantInfoModifyReq 商户信息修改
type MerchantInfoModifyReq struct {
	OrderNo          string            `json:"orderNo"`                    // 订单号 变更请求订单号
	MerchantNo       string            `json:"merchantNo"`                 // 子商户编号
	LegalPersonID    string            `json:"legalPersonID,omitempty"`    // 法人身份证号(需人工审核)
	LegalPerson      string            `json:"legalPerson,omitempty"`      // 法人姓名(需人工审核)
	BusinessLicense  string            `json:"businessLicense,omitempty"`  // 营业执照号(需人工审核)
	SignName         string            `json:"signName,omitempty"`         // 签约名(需人工审核)
	IDCardStartDate  string            `json:"idCardStartDate,omitempty"`  // 身份证有效期开始
	IDCardEndDate    string            `json:"idCardEndDate,omitempty"`    // 身份证有效期结束
	BusLiceStartDate string            `json:"busLiceStartDate,omitempty"` // 营业执照有效期开始
	BusLiceEndDate   string            `json:"busLiceEndDate,omitempty"`   // 营业执照有效期结束
	RegionCode       string            `json:"regionCode,omitempty"`       // 地区码
	Address          string            `json:"address,omitempty"`          // 地址
	MCC              string            `json:"mcc,omitempty"`              // mcc码
	ShowName         string            `json:"showName,omitempty"`         // 商户简称
	ServicePhone     string            `json:"servicePhone,omitempty"`     // 客服电话
	LinkPhone        string            `json:"linkPhone,omitempty"`        // 联系电话
	Linkman          string            `json:"linkman,omitempty"`          // 联系人
	MerchantCategory string            `json:"merchantCategory,omitempty"` // 经营类别
	IndustryTypeCode string            `json:"industryTypeCode,omitempty"` // 行业类型编码
	MicroBizType     string            `json:"microBizType,omitempty"`     // 小微经营类型
	CertType         string            `json:"certType,omitempty"`         // 证书类型
	LinkManId        string            `json:"linkManId,omitempty"`        // 联系人证件号
	SpecialSignName  string            `json:"specialSignName,omitempty"`  // 是否需要特殊处理商户名称
	SynChannel       bool              `json:"synChannel,omitempty"`       // 是否同步通道到修改
	AppPayType       string            `json:"appPayType,omitempty"`       // 支付类型
	FileUrlMap       map[string]string `json:"fileUrlMap,omitempty"`       // 文件URL
	FileSigns        map[string]string `json:"fileSigns,omitempty"`        // 文件签名列表
	IDType           string            `json:"idType,omitempty"`           // 法人证件类型
	ServiceCodes     string            `json:"serviceCodes,omitempty"`     // 支付宝商户服务类型
}

// MerchantInfoModifyRes 商户信息修改
type MerchantInfoModifyRes struct {
	OrderNo          string `json:"orderNo"`          // 变更请求的订单号
	MerchantNo       string `json:"merchantNo"`       // 子商户编号
	Result           string `json:"result"`           // 审核状态结果
	AlterationStatus string `json:"alterationStatus"` // 审核状态
}

// MerchantAlterationReq 商户信息变更
type MerchantAlterationReq struct {
	OrderNo                     string `json:"orderNo"`                     // 订单号 变更请求订单号
	MerchantNo                  string `json:"merchantNo"`                  // 子商户编号 子商户编号
	MerchantEntryAlterationType string `json:"merchantEntryAlterationType"` // 变更类型 见附录5.18
	UpdateShowName              string `json:"updateShowName"`              // 展示名
	UpdateLinkPhone             string `json:"updateLinkPhone"`             // 客服电话
	Linkman                     string `json:"linkman,omitempty"`           // 联系人 仅修改支付宝报备信息
	LinkPhone                   string `json:"linkPhone,omitempty"`         // 联系电话（手机）仅修改支付宝报备信息
	BusinessLicense             string `json:"businessLicense,omitempty"`   // 营业执照 仅修改支付宝报备信息
	LegalPersonID               string `json:"legalPersonID,omitempty"`     // 身份证号 仅修改支付宝报备信息
	Address                     string `json:"address,omitempty"`           // 地址 仅修改支付宝报备信息
	MerchantCategory            string `json:"merchantCategory,omitempty"`  // 经营类别 线下零售 仅修改支付宝报备信息见附录5.5
	RegionCode                  string `json:"regionCode,omitempty"`        // 区县编码 仅修改支付宝报备信息（如未传区县编码，则在省市信息中随机选取一个区县编码）
	Province                    string `json:"province,omitempty"`          // 省 仅修改支付宝报备信息
	City                        string `json:"city,omitempty"`              // 市 仅修改支付宝报备信息
}

// MerchantAlterationRes 商户信息变更
type MerchantAlterationRes struct {
	OrderNo          string `json:"orderNo"`                    // 订单号 变更请求的订单号
	MerchantNo       string `json:"merchantNo"`                 // 子商户编号 子商户编号
	Result           string `json:"result"`                     // 审核状态结果 见附录5.19
	AlterationStatus string `json:"alterationStatus,omitempty"` // 审核状态 见附录5.19
}

// MerchantAlterationQueryReq 商户变更查询
type MerchantAlterationQueryReq struct {
	OrderNo                     string `json:"orderNo"`                     // 订单号 变更请求订单号
	MerchantNo                  string `json:"merchantNo"`                  // 子商户编号 子商户编号
	MerchantEntryAlterationType string `json:"merchantEntryAlterationType"` // 变更类型 见附录5.18
}

// MerchantAlterationQueryRes 商户变更查询
type MerchantAlterationQueryRes struct {
	OrderNo          string `json:"orderNo"`          // 订单号 变更请求的订单号
	MerchantNo       string `json:"merchantNo"`       // 子商户编号 子商户编号
	Result           string `json:"result"`           // 审核状态结果 见附录5.19
	AlterationStatus string `json:"alterationStatus"` // 审核状态 见附录5.19
	Remark           string `json:"remark,omitempty"` // 备注

}

// MerchantSettlementCardAlterationReq 商户结算卡信息变更
type MerchantSettlementCardAlterationReq struct {
	OrderNo                     string                                    `json:"orderNo"`                             // 订单号 变更请求订单号
	MerchantNo                  string                                    `json:"merchantNo"`                          // 子商户编号 进件审核通过后才有的子商户编号
	AccountName                 string                                    `json:"accountName"`                         // 开户人名称 开户人名称
	UpdateAccountName           string                                    `json:"updateAccountName,omitempty"`         // 变更后开户人名称 变更后开户人名称，输入该字段需上传结算账户指定书
	AccountNo                   string                                    `json:"accountNo"`                           // 原结算卡号 原结算卡号
	UpdateAccountNo             string                                    `json:"updateAccountNo,omitempty"`           // 变更后结算卡号 变更后结算卡号
	SettleBankType              string                                    `json:"settleBankType"`                      // 结算卡类型 见附录5.9
	UpdateSettleBankType        string                                    `json:"updateSettleBankType,omitempty"`      // 变更后结算卡类型 见附录5.9
	BankCode                    string                                    `json:"bankCode"`                            // 结算卡联行号 结算卡联行号
	UpdateBankCode              string                                    `json:"updateBankCode,omitempty"`            // 变更后结算卡联行号 变更后结算卡联行号
	MerchantEntryAlterationType string                                    `json:"merchantEntryAlterationType"`         // 变更类型 见附录5.18
	FileSigns                   MerchantSettlementCardAlterationFileSigns `json:"fileSigns,omitempty"`                 // 文件 HASH 映射 参见以下关于结算卡变更接口文件Hash映射补充说明
	SettlementIdCardStartDate   string                                    `json:"settlementIdCardStartDate,omitempty"` // 结算人身份证开始日期 yyyyMMdd
	SettlementIdCardEndDate     string                                    `json:"settlementIdCardEndDate,omitempty"`   // 结算人身份证结束日期 yyyyMMdd 或者长期有效
	SettlementIdCardNo          string                                    `json:"settlementIdCardNo,omitempty"`        // 结算人身份证号
	SettlementMode              string                                    `json:"settlementMode,omitempty"`            // 结算方式 见附录5.15
	ElectronicAccountNo         string                                    `json:"electronicAccountNo,omitempty"`       // 电子账户
	SettleChangeType            string                                    `json:"settleChangeType,omitempty"`          // 结算到账方式 具体值见附录5.38
	SettlementPhoneNo           string                                    `json:"settlementPhoneNo,omitempty"`         // 结算人手机号
	SettlementPeriod            string                                    `json:"settlementPeriod,omitempty"`          // 结算周期类型 见附录5.10
	SettleMode                  string                                    `json:"settleMode,omitempty"`                // 结算模式 商户合并结算使用，见附录 5.27[合并结算使用]
}

type MerchantSettlementCardAlterationFileSigns struct {
	AuthorizationForSettlement string `json:"authorizationForSettlement,omitempty"` // 结算账户指定书上传时必填，文件 MD5 校验码
	FrontOfIdCard              string `json:"frontOfIdCard,omitempty"`              // 持卡人身份证正面照上传时必填，文件 MD5 校验码
	BackOfIdCard               string `json:"backOfIdCard,omitempty"`               // 持卡人身份证反面照上传时必填，文件 MD5 校验码
	HandheldOfIdCard           string `json:"handheldOfIdCard,omitempty"`           // 持卡人手持身份证照上传时必填，文件 MD5 校验码
	HandheldOfBankCard         string `json:"handheldOfBankCard,omitempty"`         // 持卡人手持银行卡照上传时必填，文件 MD5 校验码
	AccountOpeningCertificate  string `json:"accountOpeningCertificate,omitempty"`  // 银行开户证明上传时必填，文件 MD5 校验码
	SubleaseCertificate        string `json:"subleaseCertificate,omitempty"`        // 转租证明上传时必填，文件 MD5 校验码
	PermitForBankAccount       string `json:"permitForBankAccount,omitempty"`       // 开户许可证上传时必填，文件 MD5 校验码
	BankCard                   string `json:"bankCard,omitempty"`                   // 结算卡上传时必填，文件 MD5 校验码
}

// MerchantSettlementCardAlterationRes 商户结算卡信息变更
type MerchantSettlementCardAlterationRes struct {
	OrderNo          string `json:"orderNo"`          // 订单号 变更请求订单号
	MerchantNo       string `json:"merchantNo"`       // 子商户编号 进件审核通过后才有的子商户编号
	Result           string `json:"result"`           // 审核状态结果
	AlterationStatus string `json:"alterationStatus"` // 审核状态
}

// MerchantUploadCredentialReq 商户资质上传接口
type MerchantUploadCredentialReq struct {
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 进件审核通过后才有的商户号
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 此参数需与进件的orderNo一致，待审核商户才需填写此参数
	CredentialType string `json:"credentialType"`       // 资质类型 见附录5.21
	FileSign       string `json:"fileSign"`             // 资质文件 HASH 值 文件 MD5 校验码
}

// MerchantUploadCredentialRes 商户资质上传接口
type MerchantUploadCredentialRes struct {
	MerchantNo       string `json:"merchantNo,omitempty"` // 子商户编号 (非必填)
	OrderNo          string `json:"orderNo,omitempty"`    // 请求单号 (非必填)
	CredentialType   string `json:"credentialType"`       // 资质类型 (必填) 见附录5.21
	CredentialStatus string `json:"credentialStatus"`     // 上传结果 (必填) 见附录5.29
}

// MerchantUploadCredentialQueryReq 商户资质上传查询接口
type MerchantUploadCredentialQueryReq struct {
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 进件审核通过后才有的商户号
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 此参数需与进件的orderNo一致，待审核商户才需填写此参数
	CredentialType string `json:"credentialType"`       // 资质类型 见附录5.21
}

// MerchantUploadCredentialQueryRes 商户资质上传查询接口
type MerchantUploadCredentialQueryRes struct {
	MerchantNo       string `json:"merchantNo,omitempty"` // 子商户编号 (非必填)
	OrderNo          string `json:"orderNo,omitempty"`    // 请求单号 (非必填)
	CredentialType   string `json:"credentialType"`       // 资质类型 (必填) 见附录5.21
	CredentialStatus string `json:"credentialStatus"`     // 上传结果 (必填) 见附录5.29
}

// MerchantUploadAlterationCredentialReq 商户资质变更上传接口
type MerchantUploadAlterationCredentialReq struct {
	OrderNo                     string `json:"orderNo,omitempty"`           // 请求单号 此参数需与进件的orderNo一致，待审核商户才需填写此参数
	MerchantNo                  string `json:"merchantNo,omitempty"`        // 子商户编号 进件审核通过后才有的商户号
	CredentialType              string `json:"credentialType"`              // 资质类型 见附录5.21
	MerchantEntryAlterationType string `json:"merchantEntryAlterationType"` // 资质类型 见附录5.21
	FileSign                    string `json:"fileSign"`                    // 资质文件HASH值,文件MD5校验码
}

// MerchantUploadAlterationCredentialRes 商户资质变更上传接口
type MerchantUploadAlterationCredentialRes struct {
	OrderNo          string `json:"orderNo"`          // 请求单号
	MerchantNo       string `json:"merchantNo"`       // 子商户编号
	Result           string `json:"result"`           // 审核状态结果,见附录5.19
	AlterationStatus string `json:"alterationStatus"` // 审核状态,见附录5.19
	EffectTime       string `json:"effectTime"`       // 生效时间,审核通过时间
}

// MerchantUploadCredentialUrlReq 商户资质Url上传接口
type MerchantUploadCredentialUrlReq struct {
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 进件审核通过后才有的商户号
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 此参数需与进件的orderNo一致，待审核商户才需填写此参数
	CredentialType string `json:"credentialType"`       // 资质类型 见附录5.21
	CredentialUrl  string `json:"credentialUrl"`        // 资质文件地址
}

// MerchantUploadCredentialUrlRes 商户资质Url上传接口
type MerchantUploadCredentialUrlRes struct {
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 (非必填)
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 (非必填)
	CredentialType string `json:"credentialType"`       // 资质类型 (必填) 见附录5.21
	Status         string `json:"status"`               // 上传结果 (必填) 见附录5.29
}

// MerchantUploadCredentialUrlQueryReq 商户资质Url上传查询接口
type MerchantUploadCredentialUrlQueryReq struct {
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 进件审核通过后才有的商户号
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 此参数需与进件的orderNo一致，待审核商户才需填写此参数
	CredentialType string `json:"credentialType"`       // 资质类型 见附录5.21
}

// MerchantUploadCredentialUrlQueryRes 商户资质Url上传查询接口
type MerchantUploadCredentialUrlQueryRes struct {
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 (非必填)
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 (非必填)
	CredentialType string `json:"credentialType"`       // 资质类型 (必填) 见附录5.21
	Status         string `json:"status"`               // 上传结果 (必填) 见附录5.29
	Remark         string `json:"remark"`               // 备注信息
}

// MerchantUploadAlterationCredentialUrlReq 商户资质Url变更上传接口
type MerchantUploadAlterationCredentialUrlReq struct {
	OrderNo        string `json:"orderNo,omitempty"`    // 请求单号 此参数需与进件的orderNo一致，待审核商户才需填写此参数
	MerchantNo     string `json:"merchantNo,omitempty"` // 子商户编号 进件审核通过后才有的商户号
	CredentialType string `json:"credentialType"`       // 资质类型 见附录5.21
	CredentialUrl  string `json:"credentialUrl"`        // 资质文件地址
}

// MerchantUploadAlterationCredentialUrlRes 商户资质Url变更上传接口
type MerchantUploadAlterationCredentialUrlRes struct {
	OrderNo        string `json:"orderNo"`        // 请求单号
	MerchantNo     string `json:"merchantNo"`     // 子商户编号
	CredentialType string `json:"credentialType"` // 资质类型 见附录5.21
	Status         string `json:"status"`         // 上传结果 (必填) 见附录5.29
}

// MerchantModifyProductConfigReq 商户产品手续费收取方式修改接口
type MerchantModifyProductConfigReq struct {
	MerchantNo  string `json:"merchantNo"`  // 子商户编号
	Type        string `json:"type"`        // 类型
	ProductType string `json:"productType"` // 产品类型
	Value       string `json:"value"`       // 收取方式
}

// MerchantModifyProductConfigRes 商户产品手续费收取方式修改接口
type MerchantModifyProductConfigRes struct {
	MerchantNo string `json:"merchantNo"` // 子商户编号
	Status     string `json:"status"`     // SUCCESS:成功;FAIL:失败
}

// MerchantModifyProductConfigQueryReq 商户产品手续费收取方式查询接口
type MerchantModifyProductConfigQueryReq struct {
	MerchantNo  string `json:"merchantNo"`  // 子商户编号
	Type        string `json:"type"`        // 类型
	ProductType string `json:"productType"` // 产品类型
}

// MerchantModifyProductConfigQueryRes 商户产品手续费收取方式查询接口
type MerchantModifyProductConfigQueryRes struct {
	MerchantNo string `json:"merchantNo"` // 子商户编号
	Value      string `json:"value"`      // 收取方式,交易手续费收取方式(自身资金账户、平台商手续费账号)
}
