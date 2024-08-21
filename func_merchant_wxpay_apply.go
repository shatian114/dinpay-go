package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

// MerchantWxpayAppApply 商户微信公众号/小程序进件
func (t *Client) MerchantWxpayAppApply(reqBody MerchantWxpayAppApplyReq) (res *BaseMerchantRes[MerchantWxpayAppApplyRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("wxPublicApply", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantWxpayAppApplyRes]{Success: tmpRes.Success,
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

// MerchantWxpayAppApplyQuery 商户微信公众号/小程序进件查询
func (t *Client) MerchantWxpayAppApplyQuery(reqBody MerchantWxpayAppApplyQueryReq) (res *BaseMerchantRes[MerchantWxpayAppApplyQueryRes], err error) {
	var response *req.Response
	if response, err = t.merchantPost("wxPublicApplyQuery", reqBody); err != nil {
		return nil, err
	}
	tmpRes := new(BaseMerchantRes[string])
	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
		return nil, err
	}
	res = &BaseMerchantRes[MerchantWxpayAppApplyQueryRes]{Success: tmpRes.Success,
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
