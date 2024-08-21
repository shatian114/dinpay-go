package dinpay

// MerchantAgreementContentGet 商户签章合同获取
//func (t *Client) MerchantAgreementContentGet(reqBody models.MerchantAgreementContentGetReq) (res *models.BaseMerchantRes[models.MerchantAgreementContentGetRes], err error) {
//	var response *req.Response
//	if response, err = t.merchantAgreementPost("signContract", reqBody); err != nil {
//		return nil, err
//	}
//	tmpRes := new(models.BaseMerchantRes[string])
//	if err = sonic.Unmarshal(response.Bytes(), tmpRes); err != nil {
//		return nil, err
//	}
//	res = &models.BaseMerchantRes[models.MerchantAgreementContentGetRes]{Success: tmpRes.Success,
//		Code: tmpRes.Code, Message: tmpRes.Message, Sign: tmpRes.Sign, Hostname: tmpRes.Hostname}
//	if tmpRes.Sign != t.MD5Sign([]string{tmpRes.Data}, t.commonSignKey) {
//		err = errors.New("响应内容验签失败")
//	}
//	if !res.Success || res.Code != "0000"{
//		return res, err
//	}
//	if err = sonic.Unmarshal(t.Des3dDecrypt(tmpRes.Data, t.commonEncryptKey), &res.Data); err != nil {
//		return nil, err
//	}
//	return
//}
