package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

// MerchantInfoModify 商户信息变更V2
func (t *Client) MerchantInfoModify(reqBody MerchantInfoModifyReq) (res *BaseMerchantRes[MerchantInfoModifyRes], err error) {
	var response *req.Response
	if response, err = t.merchantUploadPost("modifyMerchantInfoV2", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantInfoModifyRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantAlteration 商户信息变更
func (t *Client) MerchantAlteration(reqBody MerchantAlterationReq) (res *BaseMerchantRes[MerchantAlterationRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("infoAlteration", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantAlterationRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantAlterationQuery 商户变更查询
func (t *Client) MerchantAlterationQuery(reqBody MerchantAlterationQueryReq) (res *BaseMerchantRes[MerchantAlterationQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("queryAlteration", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantAlterationQueryRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantSettlementCardAlteration 商户结算卡信息变更
func (t *Client) MerchantSettlementCardAlteration(reqBody MerchantSettlementCardAlterationReq) (res *BaseMerchantRes[MerchantSettlementCardAlterationRes], err error) {
	var response *req.Response
	if response, err = t.merchantUploadPost("settlementCardAlteration", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantSettlementCardAlterationRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantUploadCredential 商户资质上传接口
func (t *Client) MerchantUploadCredential(reqBody MerchantUploadCredentialReq) (res *BaseMerchantRes[MerchantUploadCredentialRes], err error) {
	var response *req.Response
	if response, err = t.merchantUploadPost("uploadCredential", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantUploadCredentialRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantUploadCredentialQuery 商户资质上传查询接口
func (t *Client) MerchantUploadCredentialQuery(reqBody MerchantUploadCredentialQueryReq) (res *BaseMerchantRes[MerchantUploadCredentialQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("credentialQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantUploadCredentialQueryRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantUploadAlterationCredential 商户资质上传变更接口
func (t *Client) MerchantUploadAlterationCredential(reqBody MerchantUploadAlterationCredentialReq) (res *BaseMerchantRes[MerchantUploadAlterationCredentialRes], err error) {
	var response *req.Response
	if response, err = t.merchantUploadPost("uploadAlterationAptitude", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantUploadAlterationCredentialRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantUploadCredentialUrl 商户资质Url上传接口
func (t *Client) MerchantUploadCredentialUrl(reqBody MerchantUploadCredentialUrlReq) (res *BaseMerchantRes[MerchantUploadCredentialUrlRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("uploadImageUrl", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantUploadCredentialUrlRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantUploadCredentialUrlQuery 商户资质Url上传查询接口
func (t *Client) MerchantUploadCredentialUrlQuery(reqBody MerchantUploadCredentialUrlQueryReq) (res *BaseMerchantRes[MerchantUploadCredentialUrlQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("imageUrlQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantUploadCredentialUrlQueryRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantUploadAlterationCredentialUrl 商户资质Url上传变更接口
func (t *Client) MerchantUploadAlterationCredentialUrl(reqBody MerchantUploadAlterationCredentialUrlReq) (res *BaseMerchantRes[MerchantUploadAlterationCredentialUrlRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("imageUrlAlteration", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantUploadAlterationCredentialUrlRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantModifyProductConfig 商户产品手续费收取方式修改接口
func (t *Client) MerchantModifyProductConfig(reqBody MerchantModifyProductConfigReq) (res *BaseMerchantRes[MerchantModifyProductConfigRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("modifyProductConfig", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantModifyProductConfigRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}

// MerchantModifyProductConfigQuery 商户产品手续费收取方式查询接口
func (t *Client) MerchantModifyProductConfigQuery(reqBody MerchantModifyProductConfigQueryReq) (res *BaseMerchantRes[MerchantModifyProductConfigQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("getProductConfig", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantModifyProductConfigQueryRes]{Success: tmpRes.Success,
		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
		err = errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}
