package dinpay

// MerchantAppPaySetting 商户扫码产品开通
func (t *Client) MerchantAppPaySetting(reqBody MerchantAppPaySettingReq) (res *BaseRes[MerchantAppPaySettingRes], err error) {
	const path = "/trx/api/product/appPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "APPPAY"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantAppPaySettingRes](baseRes)
}

// MerchantAppPayQuery 商户扫码产品查询
func (t *Client) MerchantAppPayQuery(reqBody MerchantAppPayQueryReq) (res *BaseRes[MerchantAppPayQueryRes], err error) {
	const path = "/trx/api/product/appPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "APPPAY"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantAppPayQueryRes](baseRes)
}

// MerchantSettlementSetting 商户开通结算产品
func (t *Client) MerchantSettlementSetting(reqBody MerchantSettlementSettingReq) (res *BaseRes[MerchantSettlementSettingRes], err error) {
	const path = "/trx/api/product/settlementSetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "SETTLEMENT"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantSettlementSettingRes](baseRes)
}

// MerchantSettlementProductQuery 商户结算产品查询
func (t *Client) MerchantSettlementProductQuery(reqBody MerchantSettlementProductQueryReq) (res *BaseRes[MerchantSettlementProductQueryRes], err error) {
	const path = "/trx/api/product/settlementQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "SETTLEMENT"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantSettlementProductQueryRes](baseRes)
}

// MerchantTransferSetting 商户开通代付产品
func (t *Client) MerchantTransferSetting(reqBody MerchantTransferSettingReq) (res *BaseRes[MerchantTransferSettingRes], err error) {
	const path = "/trx/api/product/transferSetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "TRANSFER"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantTransferSettingRes](baseRes)
}

// MerchantTransferQuery 商户代付产品查询
func (t *Client) MerchantTransferQuery(reqBody MerchantTransferQueryReq) (res *BaseRes[MerchantTransferQueryRes], err error) {
	const path = "/trx/api/product/transferQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "TRANSFER"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantTransferQueryRes](baseRes)

}

// MerchantQuickPaySetting 商户开通快捷产品
func (t *Client) MerchantQuickPaySetting(reqBody MerchantQuickPaySettingReq) (res *BaseRes[MerchantQuickPaySettingRes], err error) {
	const path = "/trx/api/product/quickPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "QUICKPAY"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantQuickPaySettingRes](baseRes)

}

// MerchantQuickPayQuery 商户快捷产品查询
func (t *Client) MerchantQuickPayQuery(reqBody MerchantQuickPayQueryReq) (res *BaseRes[MerchantQuickPayQueryRes], err error) {
	const path = "/trx/api/product/quickPayQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "QUICKPAY"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantQuickPayQueryRes](baseRes)

}

// MerchantAccountPaySetting 商户开通虚拟账户支付产品
func (t *Client) MerchantAccountPaySetting(reqBody MerchantAccountPaySettingReq) (res *BaseRes[MerchantAccountPaySettingRes], err error) {
	const path = "/trx/api/product/accountPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "ACCOUNT_PAY"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantAccountPaySettingRes](baseRes)

}

// MerchantAccountPayQuery 商户虚拟账户支付产品查询
func (t *Client) MerchantAccountPayQuery(reqBody MerchantAccountPayQueryReq) (res *BaseRes[MerchantAccountPayQueryRes], err error) {
	const path = "/trx/api/product/accountPayQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "ACCOUNT_PAY"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantAccountPayQueryRes](baseRes)
}

// MerchantTransferDepositSetting 商户开通转账充值产品
func (t *Client) MerchantTransferDepositSetting(reqBody MerchantTransferDepositSettingReq) (res *BaseRes[MerchantTransferDepositSettingRes], err error) {
	const path = "/trx/api/product/transferDepositOpen"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "TRANSFERDEPOSIT"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantTransferDepositSettingRes](baseRes)
}

// MerchantTransferDepositQuery 商户转账充值产品查询
func (t *Client) MerchantTransferDepositQuery(reqBody MerchantTransferDepositQueryReq) (res *BaseRes[MerchantTransferDepositQueryRes], err error) {
	const path = "/trx/api/product/transferDepositQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "TRANSFERDEPOSIT"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantTransferDepositQueryRes](baseRes)
}

// MerchantModifyFeeConfig 商户产品手续费配置修改
func (t *Client) MerchantModifyFeeConfig(reqBody MerchantModifyFeeConfigReq) (res *BaseRes[MerchantModifyFeeConfigRes], err error) {
	const path = "/trx/api/product/modifyFeeConfig"
	//configModifi 没写错,智付的历史遗留问题
	reqBody.InterfaceName, reqBody.ModifyType = "configModifi", "FeeCollection"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantModifyFeeConfigRes](baseRes)

}

// MerchantQueryFeeConfig 商户产品手续费配置查询
func (t *Client) MerchantQueryFeeConfig(reqBody MerchantQueryFeeConfigReq) (res *BaseRes[MerchantQueryFeeConfigRes], err error) {
	const path = "/trx/api/product/queryFeeConfig"
	reqBody.InterfaceName, reqBody.QueryType = "configQuery", "FeeCollection"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantQueryFeeConfigRes](baseRes)
}
