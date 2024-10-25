package dinpay

type BaseBalanceRes[T any] struct {
	Code            string `json:"code"`                      // 响应码
	Msg             string `json:"msg"`                       // 返回信息
	SignatureMethod string `json:"signatureMethod,omitempty"` // 签名方式
	Sign            string `json:"sign,omitempty"`            // 签名
	Data            T      `json:"data,omitempty"`            // 响应实体
	MerchantId      string `json:"merchantId"`                // 服务商商户号
}

// MerchantBalanceQueryReq 商户余额查询接口
type MerchantBalanceQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 交易类型 商户账户查询接口
	MerchantId    string `json:"merchantId"`    // 商户编号 商户号由系统分配。注：平台商子商户模式下，查询子商户余额时值为子商户号
}

// MerchantBalanceQueryRes 商户余额查询接口
type MerchantBalanceQueryRes struct {
	InterfaceName    string  `json:"interfaceName"`           // 交易类型 确认支付接口
	MerchantId       string  `json:"merchantId"`              // 商户编号 商户号由系统分配。注：平台商子商户模式下，查询子商户余额时值为子商户号
	AccountStatus    string  `json:"accountStatus"`           // 账户状态 AVAILABLE("正常",1), FROZEN("冻结",2), FREEZE_DEBIT("冻结出帐",3), FREEZE_CREDIT("冻结入账",4), CANCELLED("已注销",5);
	Currency         string  `json:"currency"`                // 币种
	CreateDate       string  `json:"createDate"`              // 创建时间 格式：yyyy-MM-dd HH:mm:ss
	MerchantBalance  string  `json:"merchantBalance"`         // 账户余额 订单金额，以元为单位，最小金额为0.01；
	D0Balance        float64 `json:"d0Balance"`               // D0余额 订单金额，以元为单位，最小金额为0.01
	D1Balance        float64 `json:"d1Balance"`               // D1余额 已经过了一个自然日的资金余额
	T1Balance        string  `json:"t1Balance"`               // T1余额 订单金额，以元为单位，最小金额为0.01
	RechargeBalance  float64 `json:"rechargeBalance"`         // 充值金额 充值金额，根据结算周期清0（T1则次工作日清0，D1则次日清0）
	FrozenBalance    string  `json:"frozenBalance,omitempty"` // 冻结余额 订单金额，以元为单位，最小金额为0.01； 因触犯风控规则而进行冻结的金额；
	SettlableBalance float64 `json:"settlableBalance"`        // 可结算金额 商户可出款金额；
	// SettlableBalance 计算公式：
	// 可结算金额（D1结算周期）= D1余额+充值金额-冻结余额；
	// 可结算金额（T1结算周期）= T1余额+充值金额-冻结余额；
	// 可结算金额（D0结算周期）= 账户余额 – 冻结余额；
	Desc string `json:"desc,omitempty"` // 备注 备注信息
}

// MerchantSettlementReq 商户结算接口
type MerchantSettlementReq struct {
	BizType        string  `json:"P1_bizType" sign:"1"`             // 接口类型 商户结算接口
	OrderId        string  `json:"P2_orderId" sign:"1"`             // 商户订单号 商户系统内部订单号，要求40字符以内，同一商户号下订单号唯一
	CustomerNumber string  `json:"P3_customerNumber" sign:"1"`      // 智付分配商户号
	Amount         float64 `json:"P4_amount" sign:"1"`              // 金额 金额单位为元，最少值5
	Summary        string  `json:"P5_summary,omitempty" sign:"1"`   // 结算备注
	NotifyUrl      string  `json:"P6_notifyUrl,omitempty" sign:"0"` // 结算结果通知地址。
}

// MerchantSettlementRes 商户结算接口
type MerchantSettlementRes struct {
	BizType        string `json:"rt1_bizType" sign:"1"`          // 接口类型 商户结算接口
	RetCode        string `json:"rt2_retCode" sign:"1"`          // 返回码
	RetMsg         string `json:"rt3_retMsg,omitempty" sign:"0"` // 返回信息
	CustomerNumber string `json:"rt4_customerNumber" sign:"1"`   // 商户商编 智付分配商户号
	OrderId        string `json:"rt5_orderId" sign:"1"`          // 商户订单号 商户系统内部订单号，要求40字符以内，同一商户号下订单号唯一
	Sign           string `json:"sign" sign:"0"`                 // SM3WITHSM2 签名结果
}

// MerchantSettlementQueryReq 商户结算查询接口
type MerchantSettlementQueryReq struct {
	BizType        string `json:"P1_bizType" sign:"1"`           // 接口类型 单笔结算查询
	OrderId        string `json:"P2_orderId,omitempty" sign:"1"` // 商户订单号 商户系统内部订单号，要求40字符以内，同一商户号下订单号唯一
	CustomerNumber string `json:"P3_customerNumber" sign:"1"`    // 智付分配商户号
	SettleDate     string `json:"P4_settleDate" sign:"1"`        // 结算日期 通常都是查前一个工作日
}

// MerchantSettlementQueryRes 商户结算查询接口
type MerchantSettlementQueryRes struct {
	BizType       string                     `json:"rt1_bizType" sign:"1"`                 // 接口类型 单笔结算查询
	RetCode       string                     `json:"rt2_retCode" sign:"1"`                 // 返回码
	RetMsg        string                     `json:"rt3_retMsg" sign:"0"`                  // 返回信息
	SettleRecords []MerchantSettlementRecord `json:"rt4_settleRecords,omitempty" sign:"1"` // 结算记录
	Sign          string                     `json:"sign" sign:"0"`                        // SM3WITHSM2 签名结果

}

type MerchantSettlementRecord struct {
	CustomerNumber   string  `json:"customerNumber"`             // 商户系统内部订单号
	OrderId          string  `json:"orderId"`                    // 订单号 同一商户号下订单号唯一
	SettleDate       string  `json:"settleDate"`                 // 结算日期,20180111
	SettlementAmount float64 `json:"settlementAmount"`           // 结算金额
	SettleFee        float64 `json:"settleFee"`                  // 结算手续费
	D0Amount         float64 `json:"d0Amount,omitempty"`         // D0资金 加急资金
	D0Fee            float64 `json:"d0Fee,omitempty"`            // D0手续费 加急手续费
	TodayAmount      float64 `json:"todayAmount,omitempty"`      // Dr资金 当日加急资金
	TodayFee         float64 `json:"todayFee,omitempty"`         // Dr手续费 当日加急手续费
	NonTodayD0Amount float64 `json:"nonTodayD0Amount,omitempty"` // D0-Dr资金 非当日加急资金
	NonTodayD0Fee    float64 `json:"nonTodayD0Fee,omitempty"`    // D0-Dr手续费 非当日加急手续费
	SettleType       string  `json:"settleType"`                 // 结算类型 结算类型（T1/D1）
	OrderStatus      string  `json:"orderStatus"`                // 结算状态,INIT已接收,DOING处理中,DONE成功,FAILED失败,MANUAL人工处理
	Reason           string  `json:"reason,omitempty"`           // 失败原因
	CompleteDate     string  `json:"completeDate,omitempty"`     // 订单完成时间;格式: yyyy-MM-dd HH:mm:ss
}

// MerchantSettlementResultNotifyReq 商户结算结果异步通知验签
type MerchantSettlementResultNotifyReq struct {
	BizType           string                     `json:"rt1_bizType" sign:"1"`        // 接口类型 商户结算接口
	RetCode           string                     `json:"rt2_retCode" sign:"1"`        // 返回码
	RetMsg            string                     `json:"rt3_retMsg" sign:"0"`         // 返回信息
	SettleRecords     []MerchantSettlementRecord `json:"rt4_settleRecords,omitempty"` // 结算记录
	Rt5CustomerNumber string                     `json:"rt5_customerNumber"`          // 智付分配商户号
	Sign              string                     `json:"sign"`                        // SM3WITHSM2 签名结果
}
