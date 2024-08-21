package dinpay

// MerchantWxpayAppApplyReq 商户微信公众号/小程序进件
type MerchantWxpayAppApplyReq struct {
	MerchantNo          string   `json:"merchantNo"`                // 子商户编号 进件审核通过后才有的商户号
	SubscribeAppIds     []string `json:"subscribeAppIds,omitempty"` // 关注公众号列表 关注公众号列表(与关注小程序列表二选一)
	ReceiptAppIds       []string `json:"receiptAppIds,omitempty"`   // 关注小程序列表 关注小程序列表(与关注公众号列表二选一)
	AppIds              []string `json:"appIds"`                    // 支付公众号列表 支付公众号列表
	AuthPayDirs         []string `json:"authPayDirs,omitempty"`     // 支付授权目录列表 支付授权目录列表（url 格式参考微信公众号文档）
	MerchantChannelName string   `json:"merchantChannelName"`       // 渠道名 渠道名（简称）
	OrderNo             string   `json:"orderNo"`                   // 请求单号 请求单号
	OperateType         string   `json:"operateType,omitempty"`     // 操作类型 操作类型，见附录
	ReportId            string   `json:"reportId,omitempty"`        // 唯一标识符 当operateType 非空时此字段必填。平台商下全局唯一，重复报备必须唯一。
	ApplyType           string   `json:"applyType,omitempty"`       // 报备类型 绿洲报名，不需报名的商户此字段不填
}

// MerchantWxpayAppApplyRes 商户微信公众号/小程序进件
type MerchantWxpayAppApplyRes struct {
	MerchantNo          string `json:"merchantNo"`          // 子商户编号
	OrderNo             string `json:"orderNo"`             // 请求单号
	MerchantChannelName string `json:"merchantChannelName"` // 渠道名
	Status              string `json:"status"`              // 结果
}

// MerchantWxpayAppApplyQueryReq 商户微信公众号/小程序进件查询
type MerchantWxpayAppApplyQueryReq struct {
	MerchantNo          string   `json:"merchantNo"`                    // 子商户编号
	SubscribeAppIds     string   `json:"subscribeAppIds,omitempty"`     // 关注公众号列表：注意，此字段值顺序必须跟进件接口值顺序一致
	ReceiptAppIds       []string `json:"receiptAppIds,omitempty"`       // 关注小程序列表 注意，此字段值顺序必须跟进件接口值顺序一致
	MerchantChannelName string   `json:"merchantChannelName,omitempty"` // // 渠道名（简称）
	OrderNo             string   `json:"orderNo"`                       // 请求单号
	ApplyType           string   `json:"applyType,omitempty"`           // 报备类型 蓝海绿洲查询使用参数
}

// MerchantWxpayAppApplyQueryRes 商户微信公众号/小程序进件查询
type MerchantWxpayAppApplyQueryRes struct {
	MerchantNo               string `json:"merchantNo"`                         // 子商户编号
	OrderNo                  string `json:"orderNo"`                            // 请求单号
	MerchantChannelName      string `json:"merchantChannelName"`                // 渠道名
	Status                   string `json:"status"`                             // 结果
	SignUrl                  string `json:"signUrl,omitempty"`                  // 签约链接
	WeChantStatus            string `json:"weChantStatus,omitempty"`            // 报名状态
	ReportFailedReason       string `json:"reportFailedReason,omitempty"`       // 失败原因
	ConfigStatus             string `json:"configStatus"`                       // 公众号支付授权目录,关注appId配置状态
	WxPublicConfigChannelMsg string `json:"wxPublicConfigChannelMsg,omitempty"` // 配置信息
	ChannelShortId           string `json:"channelShortId,omitempty"`           // 渠道
	SubMerchantNo            string `json:"subMerchantNo,omitempty"`            // 二级商编
	AppIds                   string `json:"appIds,omitempty"`                   // 支付公众号
	SubscribeAppIds          string `json:"subscribeAppIds,omitempty"`          // 关注公众号
	ReceiptAppIds            string `json:"receiptAppIds,omitempty"`            // 关注小程序
	AuthPayDirs              string `json:"authPayDirs,omitempty"`              // 授权目录
}

// MerchantAddAuthPayDirsDevConfigReq 商户微信公众号/小程序授权目录配置接口(单独增加单个支付目录)
type MerchantAddAuthPayDirsDevConfigReq struct {
	MerchantNo    string `json:"merchantNo"`          // 子商户编号 进件审核通过后才有的合利宝子商户号
	SubMerchantNo string `json:"subMerchantNo"`       // 微信子商户号 在微信报备的子商户号
	AuthPayDir    string `json:"authPayDir"`          // 支付授权目录 支付授权目录列表（url必须以斜杆结尾, 格式参考微信公众号文档https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_3）
	OrderNo       string `json:"orderNo"`             // 请求单号
	ApplyType     string `json:"applyType,omitempty"` // 报备类型 预留字段
}

// MerchantAddAuthPayDirsDevConfigRes 商户微信公众号/小程序授权目录配置接口(单独增加单个支付目录)
type MerchantAddAuthPayDirsDevConfigRes struct {
	MerchantNo               string `json:"merchantNo"`                         // 合利宝子商户编号
	OrderNo                  string `json:"orderNo"`                            // 请求单号
	Status                   string `json:"status"`                             // 状态
	WxPublicConfigChannelMsg string `json:"wxPublicConfigChannelMsg,omitempty"` // 结果，仅在status为'FAIL'时返回
}
