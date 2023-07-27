package sms_model

import (
	"github.com/gogf/gf/v2/database/gdb"
)

type TableName struct {
	SmsSendLog               string `p:"smsSendLog" dc:"短信发送日志表名"`
	SmsAppConfig             string `p:"smsAppConfig" dc:"短信应用配置表名"`
	SmsServiceProviderConfig string `p:"smsServiceProviderConfig" dc:"短信渠道商表名"`
	SmsTemplateConfig        string `p:"smsTemplateConfig" dc:"短信模板表名"`

	SmsSignConfig     string `p:"smsSignConfig" dc:"短信签名表名"`
	SmsBusinessConfig string `p:"smsBusinessConfig" dc:"短信业务配置表名"`
}

type Identifier struct {
	Sms                      string `p:"sms" dc:"短信服务表示符"`
	SmsSendLog               string `p:"smsSendLog" dc:"短信发送日志标识符"`
	SmsAppConfig             string `p:"smsAppConfig" dc:"短信应用配置标识符"`
	SmsServiceProviderConfig string `p:"smsServiceProviderConfig" dc:"短信渠道商标识符"`
	SmsTemplateConfig        string `p:"smsTemplateConfig" dc:"短信模版标识符"`
	SmsSignConfig            string `p:"smsSignConfig" dc:"短信签名标识符"`
	SmsBusinessConfig        string `p:"smsBusinessConfig" dc:"短信业务配置标识符"`
}

type Config struct {
	DB gdb.DB `p:"-" dc:"数据库连接"`
	// AllowEmptyNo                   bool              `p:"allowEmptyNo" dc:"允许员工工号为空" default:"false"`
	// IsCreateDefaultEmployeeAndRole bool              `p:"isCreateDefaultEmployeeAndRole" dc:"是否创建默认员工和角色"`
	HardDeleteWaitAt int64  `p:"hardDeleteWaitAt" dc:"硬删除等待时限,单位/小时" default:"12"`
	KeyIndex         string `p:"keyIndex" dc:"配置索引"`
	RoutePrefix      string `p:"routePrefix" dc:"路由前缀"`
	StoragePath      string `p:"storagePath" dc:"资源存储路径"`
	// UserType                       sys_enum.UserType `p:"userType" dc:"用户类型"`
	I18nName   string     `p:"i18NName" dc:"i18n文件名称"`
	Identifier Identifier `p:"identifier" dc:"标识符"`
	TableName  TableName  `p:"tableName" dc:"模块表名"`
}
