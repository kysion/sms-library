// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsTemplateConfig is the golang structure of table sms_template_config for DAO operations like Where/Data.
type SmsTemplateConfig struct {
	g.Meta                 `orm:"table:sms_template_config, do:true"`
	Id                     interface{} // ID
	TemplateCode           interface{} // 模版Code
	TemplateName           interface{} // 模版名称
	TemplateContent        interface{} // 模版内容
	ThirdPartyTemplateCode interface{} // 第三方模版Code
	ProviderNo             interface{} // 渠道商编号
	Remark                 interface{} // 备注
	Status                 interface{} // 状态: 0禁用 1正常
	AuditUserId            interface{} // 审核者UserID 审核者UserID
	AuditReplyMsg          interface{} // 审核回复，仅审核不通过时才有值
	AuditAt                *gtime.Time // 审核时间
	ExtJson                interface{} // 拓展字段
	UnionMainId            interface{} // 关联主体ID
	SignName               interface{} // 签名名称
	CreatedAt              *gtime.Time //
	UpdatedAt              *gtime.Time //
	DeletedAt              *gtime.Time //
	Type                   interface{} // 业务场景类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码，32设置邮箱，64忘记用户名&密码
}
