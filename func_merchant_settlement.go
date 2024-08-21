package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
	"strings"
)

// MerchantBalanceQuery 商户余额查询接口
func (t *Client) MerchantBalanceQuery(reqBody MerchantBalanceQueryReq) (res *BaseBalanceRes[MerchantBalanceQueryRes], err error) {
	reqBody.InterfaceName = "merchantBalanceQuery"
	var response *req.Response
	if response, err = t.merchantBalancePost(reqBody); err != nil {
		return nil, err
	}
	res = new(BaseBalanceRes[MerchantBalanceQueryRes])
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

// MerchantSettlement 商户结算接口
func (t *Client) MerchantSettlement(reqBody MerchantSettlementReq) (res *MerchantSettlementRes, err error) {
	reqBody.BizType = "MerchantSettlement"
	var response *req.Response
	if response, err = t.settlementPost(reqBody); err != nil {
		return nil, err
	}
	res = new(MerchantSettlementRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	signValsBytes := []byte(strings.Join(append(signVals, "SM3WITHSM2"), "&"))
	if res.Sign != t.SM3WithSM2Sign(signValsBytes, t.settlementSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// MerchantSettlementQuery 商户结算查询接口
func (t *Client) MerchantSettlementQuery(reqBody MerchantSettlementQueryReq) (res *MerchantSettlementQueryRes, err error) {
	reqBody.BizType = "MerchantSettlementQuery"
	var response *req.Response
	if response, err = t.settlementPost(reqBody); err != nil {
		return nil, err
	}
	res = new(MerchantSettlementQueryRes)
	if err = sonic.Unmarshal(response.Bytes(), res); err != nil {
		return
	}
	signVals := t.getSignValBySignTag(res, response.Bytes())
	signValsBytes := []byte(strings.Join(append(signVals, "SM3WITHSM2"), "&"))
	if res.Sign != t.SM3WithSM2Sign(signValsBytes, t.settlementSignKey) {
		err = errors.New("响应内容验签失败")
	}
	return
}

// MerchantSettlementResultNotifyVerify 商户结算结果异步通知验签
func (t *Client) MerchantSettlementResultNotifyVerify(reqBody MerchantSettlementResultNotifyReq, rowBytes []byte) bool {
	signVals := t.getSignValBySignTag(reqBody, rowBytes)
	signValsBytes := []byte(strings.Join(append(signVals, "SM3WITHSM2"), "&"))
	return reqBody.Sign == t.SM3WithSM2Sign(signValsBytes, t.settlementSignKey)
}
