package constants

const (
	ALLOWANCE = "ALLOWANCE" // 活动补贴;基于原交易的虚拟账户支付
	TRANSFER  = "TRANSFER"  // 资金划拨;不基于原交易的虚拟账户支付

	// AccountPayOrderStatusReceive 接收成功
	AccountPayOrderStatusReceive = "RECEIVE" // 接收成功;此状态只出现在需进行短验的 非担保资金划拨类 交易,下单成功之后返回此状态。若不需短验,下单成功后返回DOING
	// AccountPayOrderStatusInEscrow 担保中
	AccountPayOrderStatusInEscrow = "IN_ESCROW" // 担保中;担保资金划拨类 交易用,下单后返回此状态
	// AccountPayOrderStatusEscrowCancelled 担保撤销
	AccountPayOrderStatusEscrowCancelled = "ESCROW_CANCELLED" // 担保撤销
	// AccountPayOrderStatusEscrowDoing 处理中
	//1.活动补贴类交易下单成功 之后返回此状态
	//2.非担保资金划拨类订单确认支付之后成功之后返回此状态
	//3.担保资金划拨类订单，担保确认支付之后返回此状态
	// 订单终态需通过查询接口或回调接口确认
	AccountPayOrderStatusEscrowDoing = "ESCROW_DOING" // 处理中
	// AccountPayOrderStatusSuccess 成功
	AccountPayOrderStatusSuccess = "ESCROW_SUCCESS" // 成功
	// AccountPayOrderStatusFail 失败
	AccountPayOrderStatusFail = "ESCROW_FAIL" // 失败
)
