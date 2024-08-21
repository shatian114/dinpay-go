package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

// AccountPayOrder 账户支付下单
func (t *Client) AccountPayOrder(reqBody AccountPayOrderReq) (res *BaseAccountPayRes[AccountPayOrderRes], err error) {
	reqBody.InterfaceName = "accountPay"
	var response *req.Response
	if response, err = t.accountPayPost("pay", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayOrderRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayQuery 账户支付订单查询
func (t *Client) AccountPayQuery(reqBody AccountPayQueryReq) (res *BaseAccountPayRes[AccountPayQueryRes], err error) {
	reqBody.InterfaceName = "accountPayQuery"
	var response *req.Response
	if response, err = t.accountPayPost("accountPayQuery", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayQueryRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayGuaranteeConfirm 账户支付担保确认
func (t *Client) AccountPayGuaranteeConfirm(reqBody AccountPayGuaranteeConfirmReq) (res *BaseAccountPayRes[AccountPayGuaranteeConfirmRes], err error) {
	reqBody.InterfaceName = "guaranteeConfirm"
	var response *req.Response
	if response, err = t.accountPayPost("guaranteeConfirm", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayGuaranteeConfirmRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayTransferValidateCode 资金划拨类账户支付获取验证码接口
func (t *Client) AccountPayTransferValidateCode(reqBody AccountPayTransferValidateCodeReq) (res *BaseAccountPayRes[AccountPayTransferValidateCodeRes], err error) {
	reqBody.InterfaceName = "accountPayValidateCode"
	var response *req.Response
	if response, err = t.accountPayPost("accountPayValidateCode", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayTransferValidateCodeRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayTransferConfirm 资金划拨类账户支付确认接口
func (t *Client) AccountPayTransferConfirm(reqBody AccountPayTransferConfirmReq) (res *BaseAccountPayRes[AccountPayTransferConfirmRes], err error) {
	reqBody.InterfaceName = "accountPayConfirmPay"
	var response *req.Response
	if response, err = t.accountPayPost("accountPayConfirmPay", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayTransferConfirmRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayOrderRefund 账户支付退款接口
func (t *Client) AccountPayOrderRefund(reqBody AccountPayOrderRefundReq) (res *BaseAccountPayRes[AccountPayOrderRefundRes], err error) {
	reqBody.InterfaceName = "accountPayRefund"
	var response *req.Response
	if response, err = t.accountPayPost("accountPayRefund", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayOrderRefundRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayOrderRefundQuery 账户支付退款查询接口
func (t *Client) AccountPayOrderRefundQuery(reqBody AccountPayOrderRefundQueryReq) (res *BaseAccountPayRes[AccountPayOrderRefundQueryRes], err error) {
	reqBody.InterfaceName = "accountPayRefund"
	var response *req.Response
	if response, err = t.accountPayPost("accountPayRefundQuery", reqBody); err != nil {
		return nil, err
	}
	res = new(BaseAccountPayRes[AccountPayOrderRefundQueryRes])
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	jsonNode, _ := sonic.Get(response.Bytes(), "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	if res.Sign != t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AccountPayOrderResultNotifyVerify 商户结算结果异步通知验签
func (t *Client) AccountPayOrderResultNotifyVerify(reqBody AccountPayOrderResultNotifyReq, rowBytes []byte) bool {
	jsonNode, _ := sonic.Get(rowBytes, "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	return reqBody.Sign == t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey)
}

// AccountPayOrderRefundNotifyVerify 账户支付回调通知
func (t *Client) AccountPayOrderRefundNotifyVerify(reqBody AccountPayOrderRefundNotifyReq, rowBytes []byte) bool {
	jsonNode, _ := sonic.Get(rowBytes, "data")
	bodyBytes, _ := jsonNode.MarshalJSON()
	return reqBody.Sign == t.SM3WithSM2Sign(bodyBytes, t.accountPaySignKey)
}
