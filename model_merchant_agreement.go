package dinpay

// MerchantAgreementContentGetReq 商户签章合同接口
type MerchantAgreementContentGetReq struct {
	InterfaceName string `json:"interfaceName"`         // 接口名称
	MerchantNo    string `json:"merchantNo"`            // 平台商编号
	Email         string `json:"email"`                 // 申请合同的商户邮箱
	Phone         string `json:"phone"`                 // 申请合同的商户手机号
	LegalPersonId string `json:"legalPersonID"`         // 申请合同的商户身份证号
	LegalPerson   string `json:"legalPerson"`           // 申请合同的商户法人姓名
	SignName      string `json:"signName"`              // 公司名称
	Address       string `json:"address"`               // 商户地址
	CallBackUrl   string `json:"callBackUrl,omitempty"` // 签章完成回调地址 (非必填)
}
