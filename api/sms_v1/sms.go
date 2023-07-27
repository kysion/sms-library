package sms_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/sms-library/api/sms_api"
)

type SendCaptchaBySmsReq struct {
	g.Meta ` method:"post" summary:"发送短信验证码" tags:"短信"`

	sms_api.SendCaptchaBySmsReq
}

type SendSmsReq struct {
	g.Meta ` method:"post" summary:"发送短信" tags:"短信"`

	sms_api.SendSmsReq
}

type ReceiveSmsReq struct {
	g.Meta ` method:"post" summary:"接收短信" tags:"短信"`

	sms_api.ReceiveSmsReq
}

type GetAppConfigByIdReq struct {
	g.Meta ` method:"post" summary:"根据Id查询应用" tags:"短信"`

	sms_api.GetAppConfigByIdReq
}

type CreateAppConfigReq struct {
	g.Meta ` method:"post" summary:"创建应用" tags:"短信"`

	sms_api.CreateAppConfigReq
}

type GetAppAvailableNumberReq struct {
	g.Meta ` method:"post" summary:"账户用量统计" tags:"短信"`

	sms_api.GetAppAvailableNumberReq
}

//type UpdateAppNumberReq struct { // 此接口不暴露，短信使用数量只能被动被修改
//	g.Meta ` method:"post" summary:"更新应用使用数量" tags:"短信"`
//	AppNo  string `json:"appNo" v:"required#应用编号校验失败" dc:"应用编号"`
//	Fee    uint64 `json:"fee" v:"required#花费数量不能为空" dc:"花费数量"`
// sms_api.UpdateAppNumberReq
//}

type RegisterTemplateReq struct {
	g.Meta ` method:"post" summary:"添加短信模版" tags:"短信"`

	sms_api.RegisterTemplateReq
}

type AuditTemplateReq struct {
	g.Meta ` method:"post" summary:"审核短信模版" tags:"短信"`

	sms_api.AuditTemplateReq
}

type GetByTemplateCodeReq struct {
	g.Meta ` method:"post" summary:"根据模版Code查询模版信息" tags:"短信"`

	sms_api.GetByTemplateCodeReq
}

type RegisterSignReq struct {
	g.Meta ` method:"post" summary:"添加短信签名" tags:"短信"`

	sms_api.RegisterSignReq
}

type AuditSignReq struct {
	g.Meta ` method:"post" summary:"审核短信签名" tags:"短信"`

	sms_api.AuditSignReq
}

type GetSignBySignNameReq struct {
	g.Meta ` method:"post" summary:"根据签名名称查找签名数据" tags:"短信"`

	sms_api.GetSignBySignNameReq
}

type CreateProviderReq struct {
	g.Meta ` method:"post" summary:"添加渠道商" tags:"短信"`

	sms_api.CreateProviderReq
}

type QueryProviderByNoReq struct {
	g.Meta ` method:"post" summary:"根据No编号获取渠道商" tags:"短信"`

	sms_api.QueryProviderByNoReq
}

type GetProviderByIdReq struct {
	g.Meta ` method:"post" summary:"根据id查询渠道商配置|信息" tags:"短信"`

	sms_api.GetProviderByIdReq
}

type QueryProviderListReq struct {
	g.Meta ` method:"post" summary:"获取渠道商列表" tags:"短信"`

	sms_api.QueryProviderListReq
}

type CreateBusinessReq struct {
	g.Meta ` method:"post" summary:"创建业务" tags:"短信"`

	sms_api.CreateBusinessReq
}

type GetBusinessByAppIdReq struct {
	g.Meta ` method:"post" summary:"根据应用id查询业务" tags:"短信"`

	sms_api.GetBusinessByAppIdReq
}
