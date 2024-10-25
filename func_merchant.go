package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

// MerchantInfoQuery 商户信息查询
func (t *Client) MerchantInfoQuery(reqBody MerchantInfoQueryReq) (res *BaseMerchantRes[MerchantInfoQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("merchantInfoQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantInfoQueryRes]{Success: tmpRes.Success,
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

// MerchantRegister 商户入驻
func (t *Client) MerchantRegister(reqBody MerchantRegisterReq) (res *BaseMerchantRes[MerchantRegisterRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("register", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantRegisterRes]{Success: tmpRes.Success,
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

// MerchantRegisterQuery 商户入驻查询
func (t *Client) MerchantRegisterQuery(reqBody MerchantRegisterQueryReq) (res *BaseMerchantRes[MerchantRegisterQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("registerQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantRegisterQueryRes]{Success: tmpRes.Success,
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

// MerchantOpenAppPayProduct 商户开通扫码产品
func (t *Client) MerchantOpenAppPayProduct(reqBody MerchantOpenAppPayProductReq) (res *BaseMerchantRes[MerchantOpenAppPayProductRes], err error) {
	reqBody.ProductType = "APPPAY"
	var response *req.Response
	if response, err = t.merchantPost("openProduct", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantOpenAppPayProductRes]{Success: tmpRes.Success,
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

// MerchantAppPayProductQuery 商户扫码产品查询
func (t *Client) MerchantAppPayProductQuery(reqBody MerchantAppPayProductQueryReq) (res *BaseMerchantRes[MerchantAppPayProductQueryRes], err error) {
	reqBody.ProductType = "APPPAY"
	var response *req.Response
	if response, err = t.merchantPost("productQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantAppPayProductQueryRes]{Success: tmpRes.Success,
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

// MerchantOpenSettlementProduct 商户开通结算产品
func (t *Client) MerchantOpenSettlementProduct(reqBody MerchantOpenSettlementProductReq) (res *BaseMerchantRes[MerchantOpenSettlementProductRes], err error) {
	reqBody.ProductType = "SETTLEMENT"
	var response *req.Response
	if response, err = t.merchantPost("openProduct", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantOpenSettlementProductRes]{Success: tmpRes.Success,
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

// MerchantSettlementProductQuery 商户结算产品查询
func (t *Client) MerchantSettlementProductQuery(reqBody MerchantSettlementProductQueryReq) (res *BaseMerchantRes[MerchantSettlementProductQueryRes], err error) {
	reqBody.ProductType = "SETTLEMENT"
	var response *req.Response
	if response, err = t.merchantPost("productQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantSettlementProductQueryRes]{Success: tmpRes.Success,
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

// MerchantOpenAccountPayProduct 商户开通虚拟账户支付产品
func (t *Client) MerchantOpenAccountPayProduct(reqBody MerchantOpenAccountPayProductReq) (res *BaseMerchantRes[MerchantOpenAccountPayProductRes], err error) {
	reqBody.ProductType = "ACCOUNT_PAY"
	var response *req.Response
	if response, err = t.merchantPost("openProduct", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantOpenAccountPayProductRes]{Success: tmpRes.Success,
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

// MerchantAccountPayProductQuery 商户虚拟账户支付产品查询
func (t *Client) MerchantAccountPayProductQuery(reqBody MerchantAccountPayProductQueryReq) (res *BaseMerchantRes[MerchantAccountPayProductQueryRes], err error) {
	reqBody.ProductType = "ACCOUNT_PAY"
	var response *req.Response
	if response, err = t.merchantPost("productQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantAccountPayProductQueryRes]{Success: tmpRes.Success,
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

// MerchantAppApply APP报备
func (t *Client) MerchantAppApply(reqBody MerchantAppApplyReq) (res *BaseMerchantRes[MerchantAppApplyRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("appApply", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantAppApplyRes]{Success: tmpRes.Success,
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

// MerchantAppApplyQuery APP报备查询
func (t *Client) MerchantAppApplyQuery(reqBody MerchantAppApplyQueryReq) (res *BaseMerchantRes[MerchantAppApplyQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("appApplyQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantAppApplyQueryRes]{Success: tmpRes.Success,
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

// MerchantWxpayReIdentify 商户微信支付重新发起认证
func (t *Client) MerchantWxpayReIdentify(reqBody MerchantWxpayReIdentifyReq) (res *BaseMerchantRes[MerchantWxpayReIdentifyRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("reIdentify", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantWxpayReIdentifyRes]{Success: tmpRes.Success,
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

// MerchantRegisterNotify 商户入驻异步通知
func (t *Client) MerchantRegisterNotify(tmpReq MerchantNotifyReq[string]) (res *MerchantNotifyReq[MerchantRegisterNotifyReqBody], err error) {
	res = &MerchantNotifyReq[MerchantRegisterNotifyReqBody]{Success: tmpReq.Success,
		Code: tmpReq.Code, Message: tmpReq.Message, Sign: tmpReq.Sign, Hostname: tmpReq.Hostname}

	if !res.Success || res.Code != "0000" {
		return res, err
	}
	if tmpReq.Sign != t.MD5Sign([]string{tmpReq.Data}, t.commonSignKey) {
		return nil, errors.New("响应内容验签失败")
	}
	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpReq.Data, t.commonEncryptKey), &res.Data); err != nil {
		return nil, err
	}
	return
}
