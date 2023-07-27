package sms_model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

// SmsServiceProviderConfig 渠道商来源
type SmsServiceProviderConfig struct {
	ProviderName    string `json:"providerName"  dc:"渠道商名称"`
	ProviderNo      string `json:"providerNo" dc:"渠道商编号" v:"required|in:aliyun,tencent,huawei,qiniu,qyxs#渠道商编号不能为空|渠道商校验失败"`
	AccessKeyId     string `json:"accessKeyId" dc:"身份标识"`
	AccessKeySecret string `json:"accessKeySecret" dc:"身份认证密钥"`
	Endpoint        string `json:"endpoint" dc:"调用域名"`
	SdkAppId        string `json:"sdkAppId" dc:"应用ID"`
	Region          string `json:"region" dc:"地域列表"`
	Remark          string `json:"remark" dc:"备注"`
	Status          int    `json:"status" dc:"状态: 0禁用 1正常"`
	ExtJson         string `json:"extJson"         dc:"拓展字段"`
	Priority        int    `json:"priority"        description:"优先级，使用默认选择优先级最高的"`
	IsDefault       bool   `json:"isDefault"       description:"是否默认：true是、false否 ，默认false"`
}

// SmsTemplateConfig 短信模版
type SmsTemplateConfig struct {
	SignName               string      `json:"signName"               dc:"签名名称"`
	TemplateCode           string      `json:"templateCode" dc:"模版Code"`
	TemplateName           string      `json:"templateName" dc:"模版名称"`
	TemplateContent        string      `json:"templateContent" dc:"模版内容"`
	ThirdPartyTemplateCode string      `json:"thirdPartyTemplateCode" dc:"第三方模版Code"`
	ProviderNo             string      `json:"providerNo" dc:"渠道商编号" v:"required|in:aliyun,tencent,huawei,qiniu,qyxs#渠道商编号不能为空|渠道商校验失败"`
	Remark                 string      `json:"remark" dc:"备注"`
	Status                 int         `json:"status"        dc:"状态: -1不通过 0待审核 1正常"`
	AuditUserId            int64       `json:"auditUserId"   dc:"审核者UserID"`
	AuditReplyMsg          string      `json:"auditReplyMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt                *gtime.Time `json:"auditAt"       dc:"审核时间"`
	ExtJson                string      `json:"extJson"                dc:"拓展字段"`
	UnionMainId            int64       `json:"unionMainId"     dc:"所属主体id"`
	Type                   int         `json:"type"                   description:"业务场景类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码"`
}

// SmsSignConfig 短信签名
type SmsSignConfig struct {
	SignName      string      `json:"signName" dc:"签名名称"`
	ProviderNo    string      `json:"providerNo" dc:"渠道商编号" v:"required|in:aliyun,tencent,huawei,qiniu,qyxs#渠道商编号不能为空|渠道商校验失败"`
	ProviderName  string      `json:"providerName" dc:"渠道商名称"`
	Remark        string      `json:"remark" dc:"备注"`
	Status        int         `json:"status"        dc:"状态: -1不通过 0待审核 1正常"`
	AuditUserId   int64       `json:"auditUserId"   dc:"审核者UserID"`
	AuditReplyMsg string      `json:"auditReplyMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt       *gtime.Time `json:"auditAt"       dc:"审核时间"`
	ExtJson       string      `json:"extJson"       dc:"拓展字段"`
	UnionMainId   int64       `json:"unionMainId"     dc:"所属主体id"`
}

// SmsAppConfig 短信应用
type SmsAppConfig struct {
	Id              int64  `json:"id" dc:"应用ID"`
	AppName         string `json:"appName" dc:"应用名称"`
	AvailableNumber int64  `json:"availableNumber" dc:"可用数量"`
	TotalNumber     int    `json:"totalNumber"     dc:"总条数"`
	UseNumber       int64  `json:"useNumber" dc:"已用数量"`
	Remark          string `json:"remark" dc:"备注"`
	Status          int    `json:"status" dc:"状态: 0禁用 1正常"`
	UnionMainId     int64  `json:"unionMainId"     dc:"所属主体id"`
}

// SmsBusinessConfig 短信业务
type SmsBusinessConfig struct {
	AppId        string `json:"appId" dc:"应用ID"`
	BusinessName string `json:"businessName" dc:"业务名称"`
	BusinessNo   string `json:"businessNo" dc:"业务编号"`
	TemplateCode string `json:"templateCode" dc:"模版Code"`
	BusinessDesc string `json:"businessDesc" dc:"业务说明"`
	Remark       string `json:"remark" dc:"备注"`
	Status       int    `json:"status" dc:"状态: 0禁用 1正常"`
	UnionMainId  int64  `json:"unionMainId"     dc:"所属主体id"`
}

// SmsSendLog 短信发送日志
type SmsSendLog struct {
	AppId       int64  `json:"appId" dc:"应用ID"`
	BusinessNo  string `json:"businessNo" dc:"业务编号"`
	Fee         int    `json:"fee" dc:"条数"`
	PhoneNumber string `json:"phoneNumber" dc:"发送手机号"`
	Message     string `json:"message" dc:"接口响应消息"`
	Code        string `json:"code" dc:"接口响应状态码"`
	Content     string `json:"content" dc:"发送内容"`
	Remark      string `json:"remark" dc:"备注"`
	UnionMainId int64  `json:"unionMainId"     dc:"所属主体id"`
	Form        int64  `json:"form"        dc:"短信来源"`
	Type        int    `json:"type"        dc:"短信类型：1验证、2通知、4业务、8推广"`
	Status      int    `json:"status"      dc:"网关发送状态：0失败、1成功"`
	MetaData    string `json:"metaData"    dc:"网关返回元数据"`
	SignName    string `json:"signName"    dc:"签名名称"`
}

// AuditInfo 审核
type AuditInfo struct {
	State       int    `json:"state" dc:"审核状态"`
	AuditUserId int64  `json:"auditUserId" dc:"审核者UserId"`
	ReplyMsg    string `json:"replyMsg" dc:"审核失败时必填的原因回复"`
}

type ServiceProviderConfigListRes base_model.CollectRes[*SmsServiceProviderConfig]
