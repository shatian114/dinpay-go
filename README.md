# 智付SDK V2

### 初始化

```
platformMerchantId := "D10000000006688"
merchantSM2PrivatePassword := "8TkDZmcn6bLcDNohPBU607ph67IOnzME"
merchantSM2PrivateKeyPath := "dir/D10000000006688_75afbaa493c2727d4cfa7aaa9b3ab9224c434397_privateKey.pfx"
platformSM2PublicKeyPath := "dir/dinpayPublicKey.cer"
devMode := false // devMode 为true会打印接口请求和响应内容, 依然是请求智付的生产环境，请悉知
client, err := dinpay.NewClient(platformMerchantId, merchantSM2PrivatePassword,
merchantSM2PrivateKeyPath, platformSM2PublicKeyPath, devMode)
```

### 打赏
赞赏多少都是您的心意，感谢您的支持！

<img src="./image/微信收款码.jpg" height="300"> <img src="./image/支付宝收款码.jpg" height="300">
