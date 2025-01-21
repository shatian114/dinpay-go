package dinpay

import (
	"bytes"
	"errors"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

// AppPayOrderClose 交易订单关闭
func (t *Client) AppPayOrderClose(reqBody AppPayOrderCloseReq) (res *BaseRes[AppPayOrderCloseRes], err error) {
	const path = "/trx/api/appPay/payClose"
	reqBody.InterfaceName = "AppPayClose"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	return ParseRes[AppPayOrderCloseRes](baseRes)
}

// AppPayOrderCancel 交易订单撤销
func (t *Client) AppPayOrderCancel(reqBody AppPayOrderCancelReq) (res *BaseRes[AppPayOrderCancelRes], err error) {
	const path = "/trx/api/appPay/payCancel"
	reqBody.InterfaceName = "AppPayCancelOrder"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	return ParseRes[AppPayOrderCancelRes](baseRes)

}

// AppPayOrderRefund 交易订单退款
func (t *Client) AppPayOrderRefund(reqBody AppPayOrderRefundReq) (res *BaseRes[AppPayOrderRefundRes], err error) {
	const path = "/trx/api/appPay/payRefund"
	reqBody.InterfaceName = "AppPayRefund"
	if len(reqBody.RefundSplitRules) >= 1 {
		reqBody.RefundSplitRuleJson, _ = sonic.MarshalString(reqBody.RefundSplitRules)
	}
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	return ParseRes[AppPayOrderRefundRes](baseRes)
}

// AppPayOrderRefundQuery 交易订单退款查询
func (t *Client) AppPayOrderRefundQuery(reqBody AppPayOrderRefundQueryReq) (res *BaseRes[AppPayOrderRefundQueryRes], err error) {
	const path = "/trx/api/appPay/payRefundQuery"
	reqBody.InterfaceName = "AppPayRefundQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	res, err = ParseRes[AppPayOrderRefundQueryRes](baseRes)
	if err != nil {
		return nil, err
	}
	if len(res.Data.RefundMarketingRulesJson) > 0 {
		if err = sonic.Unmarshal([]byte(res.Data.RefundMarketingRulesJson), res.Data.RefundMarketingRules); err != nil {
			return nil, err
		}
	}
	if len(res.Data.RefundSplitRulesJson) > 0 {
		if err = sonic.Unmarshal([]byte(res.Data.RefundSplitRulesJson), res.Data.RefundSplitRules); err != nil {
			return nil, err
		}
	}
	return
}

// OrderPayRefundNotifyVerify 订单退款结果异步通知验签
func (t *Client) OrderPayRefundNotifyVerify(httpBody []byte) (notifyReq *OrderRefundResultNotifyReq, err error) {
	reqReq := http.Request{Method: "POST", Body: io.NopCloser(bytes.NewBuffer(httpBody)), Header: http.Header{}}
	reqReq.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
	if err = reqReq.ParseForm(); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(reqReq.Form.Get("data")), reqReq.Form.Get("sign")) {
		return nil, errors.New("响应内容验签失败")
	}
	httpBodyMap := make(map[string]string, len(reqReq.Form))
	for key, strings := range reqReq.Form {
		httpBodyMap[key] = strings[0]
	}
	httpBody, err = sonic.Marshal(httpBodyMap)
	if err != nil {
		return nil, err
	}

	var baseRes *NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if notifyReq, err = ParseNotifyReq[OrderRefundResultNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	if len(notifyReq.Data.RefundMarketingRulesJson) > 0 {
		_ = sonic.Unmarshal([]byte(notifyReq.Data.RefundMarketingRulesJson), &notifyReq.Data.RefundMarketingRules)
	}
	if len(notifyReq.Data.RefundSplitRulesJson) > 0 {
		_ = sonic.Unmarshal([]byte(notifyReq.Data.RefundSplitRulesJson), &notifyReq.Data.RefundSplitRules)
	}
	return
}
