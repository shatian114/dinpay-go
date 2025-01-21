package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
)

// GetMerchantSignContractUrl 获取商户签章合同链接
func (t *Client) GetMerchantSignContractUrl(reqBody GetMerchantSignContractUrlReq) (res *BaseRes[GetMerchantSignContractUrlRes], err error) {
	const path = "/trx/api/merchant/getSignContractUrl"
	reqBody.InterfaceName = "getSignContractUrl"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[GetMerchantSignContractUrlRes](baseRes)

}

// MerchantSignContractQuery 获取商户签章查询
func (t *Client) MerchantSignContractQuery(reqBody MerchantSignContractQueryReq) (res *BaseRes[MerchantSignContractQueryRes], err error) {
	const path = "/trx/api/merchant/contractQuery"
	reqBody.InterfaceName = "contractQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantSignContractQueryRes](baseRes)
}

// MerchantSignContractNotify 商户签章异步通知
func (t *Client) MerchantSignContractNotify(httpBody []byte) (notifyReq *MerchantSignContractNotifyReq, err error) {
	var baseRes *NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	if notifyReq, err = ParseNotifyReq[MerchantSignContractNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}
