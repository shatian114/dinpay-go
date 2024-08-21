package constants

const (
	ALIPAY       = "ALIPAY"       // 支付宝
	WXPAY        = "WXPAY"        // 微信支付
	UNIONPAY     = "UNIONPAY"     // 银联固码/聚合码
	QQPAY        = "QQPAY"        // QQ钱包
	JDPAY        = "JDPAY"        // 京东
	UNIONPAY_ALL = "UNIONPAY_ALL" // 银联主被扫支付
	EACHWALLET   = "EACHWALLET"   // 互通钱包
	HELIFU       = "HELIFU"       // 合利付APP
)

var (

	// AppPayTypes 客户端支付类型
	AppPayTypes = map[string]string{
		"ALIPAY":       "支付宝",
		"WXPAY":        "微信支付",
		"UNIONPAY":     "银联固码/聚合码",
		"QQPAY":        "QQ钱包",
		"JDPAY":        "京东",
		"UNIONPAY_ALL": "银联主被扫支付", // 报备查询时区分
		"EACHWALLET":   "互通钱包",    // 新增如:网联互联互通
		"HELIFU":       "合利付APP",
	}

	// PayType 支付类型
	PayType = map[string]string{
		"SWIPE":  "刷卡",
		"SCAN":   "扫码",
		"WAP":    "WAP",
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
