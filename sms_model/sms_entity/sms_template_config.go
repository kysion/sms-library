// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsTemplateConfig is the golang structure for table sms_template_config.
type SmsTemplateConfig struct {
	Id                     int64       `json:"id"                     description:"ID"`
	TemplateCode           string      `json:"templateCode"           description:"模版Code"`
	TemplateName           string      `json:"templateName"           description:"模版名称"`
	TemplateContent        string      `json:"templateContent"        description:"模版内容"`
	ThirdPartyTemplateCode string      `json:"thirdPartyTemplateCode" description:"第三方模版Code"`
	ProviderNo             string      `json:"providerNo"             description:"渠道商编号"`
	Remark                 string      `json:"remark"                 description:"备注"`
	Status                 int         `json:"status"                 description:"状态: 0禁用 1正常"`
	AuditUserId            int64       `json:"auditUserId"            description:"审核者UserID 审核者UserID"`
	AuditReplyMsg          string      `json:"auditReplyMsg"          description:"审核回复，仅审核不通过时才有值"`
	AuditAt                *gtime.Time `json:"auditAt"                description:"审核时间"`
	ExtJson                string      `json:"extJson"                description:"拓展字段"`
	UnionMainId            int64       `json:"unionMainId"            description:"关联主体ID"`
	SignName               string      `json:"signName"               description:"签名名称"`
	CreatedAt              *gtime.Time `json:"createdAt"              description:""`
	UpdatedAt              *gtime.Time `json:"updatedAt"              description:""`
	DeletedAt              *gtime.Time `json:"deletedAt"              description:""`
	Type                   int         `json:"type"                   description:"业务场景类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码，32设置邮箱，64忘记用户名&密码"`
}
