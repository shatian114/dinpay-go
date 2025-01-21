package dinpay

// MerchantApplyAppReport 商户上游报备
func (t *Client) MerchantApplyAppReport(reqBody MerchantApplyAppReportReq) (res *BaseRes[MerchantApplyAppReportRes], err error) {
	const path = "/trx/api/product/applyAppReport"
	reqBody.InterfaceName = "reporting"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantApplyAppReportRes](baseRes)

}

// MerchantAppApplyQuery 商户上游报备查询
func (t *Client) MerchantAppApplyQuery(reqBody MerchantAppApplyQueryReq) (res *BaseRes[MerchantAppApplyQueryRes], err error) {
	const path = "/trx/api/product/appApplyReportQuery"
	reqBody.InterfaceName = "reportingQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantAppApplyQueryRes](baseRes)

}

// MerchantReIdentify 商户重新发起认证
func (t *Client) MerchantReIdentify(reqBody MerchantReIdentifyReq) (res *BaseRes[MerchantReIdentifyRes], err error) {
	const path = "/trx/api/product/reIdentify"
	reqBody.InterfaceName = "reCertification"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantReIdentifyRes](baseRes)

}

// MerchantWxPublicApply 商户微信交易配置
func (t *Client) MerchantWxPublicApply(reqBody MerchantWxPublicApplyReq) (res *BaseRes[MerchantWxPublicApplyRes], err error) {
	const path = "/trx/api/product/wxPublicApply"
	reqBody.InterfaceName = "wxAppidConfig"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantWxPublicApplyRes](baseRes)

}

// MerchantWxPayDirsApply 商户微信配置支付目录
func (t *Client) MerchantWxPayDirsApply(reqBody MerchantWxPayDirsApplyReq) (res *BaseRes[MerchantWxPayDirsApplyRes], err error) {
	const path = "/trx/api/product/weiXinAddAuthPayDirs"
	reqBody.InterfaceName = "wxAddConfig"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantWxPayDirsApplyRes](baseRes)

}

// MerchantWxPublicApplyQuery 商户微信交易配置查询
func (t *Client) MerchantWxPublicApplyQuery(reqBody MerchantWxPublicApplyQueryReq) (res *BaseRes[MerchantWxPublicApplyQueryRes], err error) {
	const path = "/trx/api/product/wxPublicApplyQuery"
	reqBody.InterfaceName = "wxAppidConfigQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantWxPublicApplyQueryRes](baseRes)

}
