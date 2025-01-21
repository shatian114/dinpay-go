package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
)

// AccountPayOrder 账户支付下单
func (t *Client) AccountPayOrder(reqBody AccountPayOrderReq) (res *BaseRes[AccountPayOrderRes], err error) {
	orderParameterJsonBytes, err := sonic.Marshal(reqBody.OrderParameter)
	if err != nil {
		return nil, err
	}
	reqBody.OrderParameters = string(orderParameterJsonBytes)

	const path = "/trx/api/accountPay/pay"
	reqBody.InterfaceName = "accountPay"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayOrderRes](baseRes)
}

// AccountPayQuery 账户支付订单查询
func (t *Client) AccountPayQuery(reqBody AccountPayQueryReq) (res *BaseRes[AccountPayQueryRes], err error) {
	const path = "/trx/api/accountPay/accountPayQuery"
	reqBody.InterfaceName = "accountPayQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayQueryRes](baseRes)
}

// AccountPayGuaranteeConfirm 账户支付担保确认
func (t *Client) AccountPayGuaranteeConfirm(reqBody AccountPayGuaranteeConfirmReq) (res *BaseRes[AccountPayGuaranteeConfirmRes], err error) {
	const path = "/trx/api/accountPay/guaranteeConfirm"
	reqBody.InterfaceName = "guaranteeConfirm"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayGuaranteeConfirmRes](baseRes)
}

// AccountPayTransferValidateCode 资金划拨类账户支付获取验证码接口
func (t *Client) AccountPayTransferValidateCode(reqBody AccountPayTransferValidateCodeReq) (res *BaseRes[AccountPayTransferValidateCodeRes], err error) {
	const path = "/trx/api/accountPay/accountPayValidateCode"
	reqBody.InterfaceName = "accountPayValidateCode"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayTransferValidateCodeRes](baseRes)
}

// AccountPayTransferConfirm 资金划拨类账户支付确认接口
func (t *Client) AccountPayTransferConfirm(reqBody AccountPayTransferConfirmReq) (res *BaseRes[AccountPayTransferConfirmRes], err error) {
	const path = "/trx/api/accountPay/accountPayConfirmPay"
	reqBody.InterfaceName = "accountPayConfirmPay"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayTransferConfirmRes](baseRes)

}

// AccountPayOrderRefund 账户支付退款接口
func (t *Client) AccountPayOrderRefund(reqBody AccountPayOrderRefundReq) (res *BaseRes[AccountPayOrderRefundRes], err error) {
	refundDetailJsonBytes, err := sonic.Marshal(reqBody.RefundDetails)
	if err != nil {
		return nil, err
	}
	reqBody.RefundDetail = string(refundDetailJsonBytes)

	const path = "/trx/api/accountPay/accountPayRefund"
	reqBody.InterfaceName = "accountPayRefund"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayOrderRefundRes](baseRes)
}

// AccountPayOrderRefundQuery 账户支付退款查询接口
func (t *Client) AccountPayOrderRefundQuery(reqBody AccountPayOrderRefundQueryReq) (res *BaseRes[AccountPayOrderRefundQueryRes], err error) {
	const path = "/trx/api/accountPay/accountPayRefundQuery"
	reqBody.InterfaceName = "accountPayRefundQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[AccountPayOrderRefundQueryRes](baseRes)
}

// AccountPayOrderResultNotifyVerify 账户支付回调通知验签
func (t *Client) AccountPayOrderResultNotifyVerify(httpBody []byte) (notifyReq *AccountPayOrderResultNotifyReq, err error) {
	var baseRes *NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	if notifyReq, err = ParseNotifyReq[AccountPayOrderResultNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}

// AccountPayOrderRefundNotifyVerify 账户支付退款回调通知
func (t *Client) AccountPayOrderRefundNotifyVerify(httpBody []byte) (notifyReq *AccountPayOrderRefundNotifyReq, err error) {
	var baseRes *NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	if notifyReq, err = ParseNotifyReq[AccountPayOrderRefundNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}
