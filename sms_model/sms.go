package sms_model

// SmsSendMessageReq 发送短信请求对象
type SmsSendMessageReq struct {
	BusinessNo   string   `json:"businessNo"  dc:"服务编号"`
	TemplateCode string   `json:"templateCode"  dc:"模版Code"`
	Phones       []string `json:"phones"  dc:"手机号集合"`
	Params       []string `json:"params"  dc:"参数列表"`
	CaptchaType  int      `json:"captchaType" v:"required#参树校验失败" dc:"验证码类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码"`
}

// SmsReceiveSmsReq 接收短信对象
type SmsReceiveSmsReq struct {
	ProviderNo string `json:"providerNo" dc:"渠道商编号" v:"required|in:aliyun,tencent,huawei,qiniu,qyxs#渠道商编号不能为空|渠道商校验失败"`
	BusinessId string `json:"BusinessId"  dc:"服务id"`

	AppId        string `json:"appId"  dc:"应用id"`
	TemplateCode string `json:"templateCode"  dc:"模版Code"`
	SignName     string `json:"signName" dc:"签名名称"`
	Content      string `json:"content" dc:"短信内容"`
	Phone        string `json:"phone" dc:"短信发送者"`
	Type         int    `json:"type"  v:"required|in:1,2,4,8#短信业务类型不能为空|短信业务类型校验失败" dc:"短信业务类型:1发送、2退订、3续费、4充值"`
}

// SmsResponse 短信请求响应体
type SmsResponse struct {
	SmsSendStatus []SmsSendStatus `json:"sms_send_status" dc:"响应集合"`
	RequestId     string          `json:"request_id" dc:"请求ID"`
}

type SmsSendStatus struct {
	Fee     uint64 `json:"fee" dc:"计费条数"`
	Message string `json:"message" dc:"响应信息"`
	Code    string `json:"code" dc:"响应消息"`
	Phone   string `json:"phone" dc:"手机号码"`
}
