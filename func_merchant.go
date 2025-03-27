package dinpay

import (
	"bytes"
	"errors"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

// MerchantRegister 商户入驻
func (t *Client) MerchantRegister(reqBody MerchantRegisterReq) (res *BaseRes[MerchantRegisterRes], err error) {
	const path = "/trx/api/merchantEntry/registered"
	reqBody.InterfaceName = "register"
	reqBody.ServiceCodesJson = ""
	if len(reqBody.ServiceCodes) > 0 {
		reqBody.ServiceCodesJson, _ = sonic.MarshalString(reqBody.ServiceCodes)
	}
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return nil, err
	}
	return ParseRes[MerchantRegisterRes](baseRes)
}

// MerchantRegisterQuery 商户入驻查询
func (t *Client) MerchantRegisterQuery(reqBody MerchantRegisterQueryReq) (res *BaseRes[MerchantRegisterQueryRes], err error) {
	const path = "/trx/api/merchantEntry/registerQuery"
	reqBody.InterfaceName = "registeredQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return nil, err
	}
	return ParseRes[MerchantRegisterQueryRes](baseRes)
}

// MerchantRegisterNotify 商户入驻异步通知
func (t *Client) MerchantRegisterNotify(httpBody []byte) (notifyReq *MerchantRegisterNotifyReq, err error) {
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
	if notifyReq, err = ParseNotifyReq[MerchantRegisterNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}

// MerchantStatusAlterNotify 商户状态变更异步通知
func (t *Client) MerchantStatusAlterNotify(httpBody []byte) (notifyReq *MerchantStatusAlterNotifyReq, err error) {
	var baseRes *NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	if notifyReq, err = ParseNotifyReq[MerchantStatusAlterNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}
