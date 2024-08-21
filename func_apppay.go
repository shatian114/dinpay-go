package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

// AppPayScanOrder 扫码(主扫,用户扫商户)/刷卡(被扫,商户扫用户)下单
func (t *Client) AppPayScanOrder(reqBody AppPayScanOrderReq) (res *AppPayScanOrderRes, err error) {
	reqBody.BizType = "AppPay"
	if reqBody.MarketingRule != nil {
		reqBody.MarketingRuleJson, _ = sonic.MarshalString(reqBody.MarketingRule)
		reqBody.MarketingRuleJson = t.Des3Encrypt([]byte(reqBody.MarketingRuleJson), t.scanEncryptKey)
	}
	reqBody.SplitBillType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitBillType = "FIXED_AMOUNT"
		reqBody.SplitRuleJson, _ = sonic.MarshalString(reqBody.SplitRules)
		reqBody.SplitRuleJson = t.Des3Encrypt([]byte(reqBody.SplitRuleJson), t.scanEncryptKey)
	}
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayScanOrderRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayPublicPreOrder 公众号/JS/服务窗预下单
func (t *Client) AppPayPublicPreOrder(reqBody AppPayPublicPreOrderReq) (res *AppPayPublicPreOrderRes, err error) {
	reqBody.BizType, reqBody.PayType = "AppPayPublic", "PUBLIC"
	if reqBody.MarketingRule != nil {
		reqBody.MarketingRuleJson, _ = sonic.MarshalString(reqBody.MarketingRule)
		reqBody.MarketingRuleJson = t.Des3Encrypt([]byte(reqBody.MarketingRuleJson), t.scanEncryptKey)
	}
	reqBody.SplitBillType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitBillType = "FIXED_AMOUNT"
		reqBody.SplitRuleJson, _ = sonic.MarshalString(reqBody.SplitRules)
		reqBody.SplitRuleJson = t.Des3Encrypt([]byte(reqBody.SplitRuleJson), t.scanEncryptKey)
	}
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayPublicPreOrderRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayAppletPreOrder 小程序预下单
func (t *Client) AppPayAppletPreOrder(reqBody AppPayAppletPreOrderReq) (res *AppPayAppletPreOrderRes, err error) {
	reqBody.BizType, reqBody.PayType = "AppPayApplet", "APPLET"
	if reqBody.MarketingRule != nil {
		reqBody.MarketingRuleJson, _ = sonic.MarshalString(reqBody.MarketingRule)
		reqBody.MarketingRuleJson = t.Des3Encrypt([]byte(reqBody.MarketingRuleJson), t.scanEncryptKey)
	}
	reqBody.SplitBillType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitBillType = "FIXED_AMOUNT"
		reqBody.SplitRuleJson, _ = sonic.MarshalString(reqBody.SplitRules)
		reqBody.SplitRuleJson = t.Des3Encrypt([]byte(reqBody.SplitRuleJson), t.scanEncryptKey)
	}
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayAppletPreOrderRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayOrderClose 交易订单关闭
func (t *Client) AppPayOrderClose(reqBody AppPayOrderCloseReq) (res *AppPayOrderCloseRes, err error) {
	reqBody.BizType = "AppPayClose"
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayOrderCloseRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayOrderCancel 交易订单撤销
func (t *Client) AppPayOrderCancel(reqBody AppPayOrderCancelReq) (res *AppPayOrderCancelRes, err error) {
	reqBody.BizType = "AppPayCancel"
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayOrderCancelRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayOrderQuery 交易订单查询
func (t *Client) AppPayOrderQuery(reqBody AppPayOrderQueryReq) (res *AppPayOrderQueryRes, err error) {
	reqBody.BizType = "AppPayQuery"
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayOrderQueryRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	if len(res.MarketingRuleJson) > 0 {
		bytes := t.Des3dDecrypt(res.MarketingRuleJson, t.scanEncryptKey)
		if err = sonic.Unmarshal(bytes, res.MarketingRule); err != nil {
			return nil, err
		}
	}
	if len(res.SplitRuleJson) > 0 {
		if err = sonic.Unmarshal([]byte(res.SplitRuleJson), &res.SplitRules); err != nil {
			return nil, err
		}
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayOrderRefund 交易订单退款
func (t *Client) AppPayOrderRefund(reqBody AppPayOrderRefundReq) (res *AppPayOrderRefundRes, err error) {
	reqBody.BizType = "AppPayRefund"
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitRuleJson, _ = sonic.MarshalString(reqBody.SplitRules)
		reqBody.SplitRuleJson = t.Des3Encrypt([]byte(reqBody.SplitRuleJson), t.scanEncryptKey)
	}
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayOrderRefundRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// AppPayOrderRefundQuery 交易订单退款查询
func (t *Client) AppPayOrderRefundQuery(reqBody AppPayOrderRefundQueryReq) (res *AppPayOrderRefundQueryRes, err error) {
	reqBody.BizType = "AppPayRefundQuery"
	var response *req.Response
	if response, err = t.appPayPost(reqBody); err != nil {
		return nil, err
	}
	res = new(AppPayOrderRefundQueryRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	} else if res.RetCode != "0000" {
		return res, err
	}
	if len(res.RefundMarketingRuleJson) > 0 {
		bytes := t.Des3dDecrypt(res.RefundMarketingRuleJson, t.scanEncryptKey)
		if err = sonic.Unmarshal(bytes, res.RefundMarketingRule); err != nil {
			return nil, err
		}
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	if res.Sign != t.MD5Sign(signVals, t.scanSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// OrderPayResultNotifyVerify 订单支付结果异步通知验签
func (t *Client) OrderPayResultNotifyVerify(reqBody OrderPayResultNotifyReq, rowBytes []byte) bool {
	if len(reqBody.MarketingRuleJson) > 0 {
		bytes := t.Des3dDecrypt(reqBody.MarketingRuleJson, t.scanEncryptKey)
		_ = sonic.Unmarshal(bytes, reqBody.MarketingRule)
	}
	if len(reqBody.SplitRuleJson) > 0 {
		_ = sonic.Unmarshal([]byte(reqBody.SplitRuleJson), &reqBody.SplitRules)
	}
	signVals := t.getSignValBySignTag(reqBody, rowBytes)
	return reqBody.Sign == t.MD5Sign(signVals, t.scanSignKey)
}

// OrderPayRefundNotifyVerify 订单退款结果异步通知验签
func (t *Client) OrderPayRefundNotifyVerify(reqBody OrderRefundResultNotifyReq, rowBytes []byte) bool {
	if len(reqBody.RefundMarketingRuleJson) > 0 {
		bytes := t.Des3dDecrypt(reqBody.RefundMarketingRuleJson, t.scanEncryptKey)
		_ = sonic.Unmarshal(bytes, reqBody.RefundMarketingRule)
	}
	signVals := t.getSignValBySignTag(reqBody, rowBytes)
	return reqBody.Sign == t.MD5Sign(signVals, t.scanSignKey)
}
