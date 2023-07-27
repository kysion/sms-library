package sms_api

import (
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/sms-library/sms_model"
)

type SendCaptchaBySmsReq struct {
	//CaptchaType int `json:"captchaType" v:"required|in:1,2,4,8#验证码类型错误|参路校验失败" dc:"验证码类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码"`
	CaptchaType int `json:"captchaType" v:"required#参数校验失败" dc:"验证码类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码"`
}

type SendSmsReq struct {
	CaptchaType int `json:"captchaType" v:"required#参树校验失败" dc:"验证码类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码"`

	sms_model.SmsSendMessageReq // 发送短信数据
}

type ReceiveSmsReq struct {
	sms_model.SmsReceiveSmsReq
}

type GetAppConfigByIdReq struct {
	Id int64 `json:"id" v:"required#应用id校验错误" dc:"应用id" `
}

type CreateAppConfigReq struct {
	sms_model.SmsAppConfig
}

type GetAppAvailableNumberReq struct {
	Id int64 `json:"id" v:"required#应用id校验失败" dc:"应用ID"`
}

//type UpdateAppNumberReq struct { // 此接口不暴露，短信使用数量只能被动被修改
//	g.Meta `path:"/updateAppNumber" method:"post" summary:"更新应用使用数量" tags:"短信"`
//	AppNo  string `json:"appNo" v:"required#应用编号校验失败" dc:"应用编号"`
//	Fee    uint64 `json:"fee" v:"required#花费数量不能为空" dc:"花费数量"`
//}

type RegisterTemplateReq struct {
	sms_model.SmsTemplateConfig
}

type AuditTemplateReq struct {
	Id int64 `json:"id" v:"required#模板id不能为空" dc:"模板id" `
	sms_model.AuditInfo
}

type GetByTemplateCodeReq struct {
	TemplateCode string `json:"templateCode" v:"required#模板Code号校验失败" dc:"模板Code"`
}

type RegisterSignReq struct {
	sms_model.SmsSignConfig
}

type AuditSignReq struct {
	Id int64 `json:"id" v:"required#短信签名id不能为空" dc:"短信签名id"`
	sms_model.AuditInfo
}

type GetSignBySignNameReq struct {
	SignName string `json:"signName" v:"required#签名名称不能为空" dc:"签名名称"`
}

type CreateProviderReq struct {
	sms_model.SmsServiceProviderConfig
}

type QueryProviderByNoReq struct {
	No string `json:"no" v:"required#渠道商编号不能为空" dc:"渠道商编号"`
	base_model.SearchParams
}

type GetProviderByIdReq struct {
	Id int64 `json:"id" v:"required#渠道商id不能为空" dc:"渠道商id"`
}

type QueryProviderListReq struct {
	base_model.SearchParams
}

type CreateBusinessReq struct {
	sms_model.SmsBusinessConfig
}

type GetBusinessByAppIdReq struct {
	AppId int64 `json:"appId" v:"required#应用id不能为空" dc:"应用id"`
}

type SmsResponseRes sms_model.SmsResponse

type SmsTemplateConfigRes sms_model.SmsTemplateConfig

type SmsSignConfigRes sms_model.SmsSignConfig

type SmsServiceProviderConfigRes sms_model.SmsServiceProviderConfig

type SmsAppConfigRes sms_model.SmsAppConfig

type SmsBusinessRes sms_model.SmsBusinessConfig

type BoolRes bool

type IntRes int
