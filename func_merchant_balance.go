package dinpay

// MerchantBalanceQuery 商户余额查询接口
func (t *Client) MerchantBalanceQuery(reqBody MerchantBalanceQueryReq) (res *BaseRes[MerchantBalanceQueryRes], err error) {
	const path = "/trx/api/merchant/merchantBalanceQuery"
	reqBody.InterfaceName = "merchantBalanceQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return ParseRes[MerchantBalanceQueryRes](baseRes)
}
