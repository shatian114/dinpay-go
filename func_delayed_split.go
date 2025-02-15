package dinpay

func (t *Client) DelayedSplit(reqBody DelayedSplitReq) (res *BaseRes[DelayedSplitRes], err error) {
	const path = "/trx/api/delayed/applySplit"
	reqBody.InterfaceName = "delaySplitting"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return ParseRes[DelayedSplitRes](baseRes)
}

func (t *Client) DelayedSplitQuery(reqBody DelayedSplitQueryReq) (res *BaseRes[DelayedSplitQueryRes], err error) {
	const path = "/trx/api/delayed/querySplit"
	reqBody.InterfaceName = "delaySplittingQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return ParseRes[DelayedSplitQueryRes](baseRes)
}

func (t *Client) DelayedSplitBack(reqBody DelayedSplitBackReq) (res *BaseRes[DelayedSplitBackRes], err error) {
	const path = "/trx/api/delayed/backSplit"
	//delaySplittingRfund 没写错,智付的历史遗留问题
	reqBody.InterfaceName = "delaySplittingRfund"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return ParseRes[DelayedSplitBackRes](baseRes)
}

func (t *Client) DelayedSplitBackQuery(reqBody DelayedSplitBackQueryReq) (res *BaseRes[DelayedSplitBackQueryRes], err error) {
	const path = "/trx/api/delayed/queryBackSplit"
	//delaySplittingRfundQuery 没写错,智付的历史遗留问题
	reqBody.InterfaceName = "delaySplittingRfundQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	res = new(BaseRes[DelayedSplitBackQueryRes])
	return ParseRes[DelayedSplitBackQueryRes](baseRes)
}
