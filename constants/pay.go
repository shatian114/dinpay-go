package constants

const (
	ALIPAY   = "ALIPAY"   // 支付宝
	WXPAY    = "WXPAY"    // 微信支付
	UNIONPAY = "UNIONPAY" // 银联
)

var (

	// PaymentType 支付客户端类型
	PaymentType = map[string]string{
		"ALIPAY":   "支付宝",
		"WXPAY":    "微信支付",
		"UNIONPAY": "银联",
	}

	// PaymentMethods 支付方式
	PaymentMethods = map[string]string{
		"SWIPE":  "被扫",  // 商户扫用户
		"SCAN":   "主扫",  // 用户扫商户
		"WAP":    "WAP", // H5页面支付
		"PUBLIC": "公众号支付",
		"SDK":    "SDK",
		"APPLET": "小程序",
	}

	// RefundStatus 退款状态
	RefundStatus = map[string]string{
		"PRE_REFUND":  "转入退款",
		"PART_REFUND": "部分退款",
		"ALL_REFUND":  "全额退款",
		"FAIL_REFUND": "退款失败",
		"NOT_YET":     "尚未退款",
	}
)
