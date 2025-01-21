package dinpay

// MerchantModifyInfo 商户信息变更
func (t *Client) MerchantModifyInfo(reqBody MerchantModifyInfoReq) (res *BaseRes[MerchantModifyInfoRes], err error) {
	const path = "/trx/api/merchantEntryAlteration/modifyMerchantInfo"
	reqBody.InterfaceName = "modifyMerchantInfo"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonMultipartFormPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantModifyInfoRes](baseRes)
}

// MerchantModifyInfoQuery 商户信息变更查询
func (t *Client) MerchantModifyInfoQuery(reqBody MerchantModifyInfoQueryReq) (res *BaseRes[MerchantModifyInfoQueryRes], err error) {
	const path = "/trx/api/merchantCredential/changeOrderQuery"
	reqBody.InterfaceName, reqBody.ChangeType = "changeOrderQuery", "MERCHANT_INFO_SYN"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantModifyInfoQueryRes](baseRes)
}

// MerchantCredentialImageUpload 商户资质图片上传
func (t *Client) MerchantCredentialImageUpload(reqBody MerchantCredentialImageUploadReq) (res *BaseRes[MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUpload"
	reqBody.InterfaceName = "imageUpload"
	var baseRes *BaseRes[string]
	if baseRes, err = t.merchantImageFormUpload(path, reqBody, reqBody.CredentialType, reqBody.GetFileContent); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageUrlUpload 商户资质图片Url上传
func (t *Client) MerchantCredentialImageUrlUpload(reqBody MerchantCredentialImageUrlUploadReq) (res *BaseRes[MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUrlUpload"
	reqBody.InterfaceName = "imageUrlUpload"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageUploadQuery 商户资质图片上传查询
func (t *Client) MerchantCredentialImageUploadQuery(reqBody MerchantCredentialImageUploadQueryReq) (res *BaseRes[MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUploadQuery"
	reqBody.InterfaceName = "imageUploadQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageUrlUploadQuery 商户资质图片Url上传查询
func (t *Client) MerchantCredentialImageUrlUploadQuery(reqBody MerchantCredentialImageUploadQueryReq) (res *BaseRes[MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUrlUploadQuery"
	reqBody.InterfaceName = "imageUrlUploadQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageChangeUpload 商户资质图片变更上传
func (t *Client) MerchantCredentialImageChangeUpload(reqBody MerchantCredentialImageChangeUploadReq) (res *BaseRes[MerchantCredentialImageChangeUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageChangeUpload"
	reqBody.InterfaceName, reqBody.ChangeType = "imageChangeUpload", "MERCHANT_CREDENTIAL"
	var baseRes *BaseRes[string]
	if baseRes, err = t.merchantImageFormUpload(path, reqBody, reqBody.CredentialType, reqBody.GetFileContent); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageChangeUploadRes](baseRes)
}

// MerchantCredentialImageUrlChangeUpload 商户资质图片Url变更上传
func (t *Client) MerchantCredentialImageUrlChangeUpload(reqBody MerchantCredentialImageUrlChangeUploadReq) (res *BaseRes[MerchantCredentialImageChangeUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUrlChangeUpload"
	reqBody.InterfaceName = "imageUrlChangeUpload"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageChangeUploadRes](baseRes)
}

// MerchantCredentialImageChangeUploadQuery 商户资质图片变更上传查询
func (t *Client) MerchantCredentialImageChangeUploadQuery(reqBody MerchantCredentialImageChangeUploadQueryReq) (res *BaseRes[MerchantCredentialImageChangeUploadRes], err error) {
	const path = "/trx/api/merchantCredential/changeOrderQuery"
	reqBody.InterfaceName, reqBody.ChangeType = "changeOrderQuery", "MERCHANT_CREDENTIAL"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantCredentialImageChangeUploadRes](baseRes)
}
