package sms_interface

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_enum"
)

type (
	ISmsTemplateConfig interface {
		CreateTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error)
		AuditTemplate(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error)
		GetTemplateById(ctx context.Context, id int64) (*sms_model.SmsTemplateConfig, error)
		GetByTemplateCode(ctx context.Context, templateCode string) (*sms_model.SmsTemplateConfig, error)
		GetByProviderNoAndType(ctx context.Context, providerNo sms_enum.SmsType, smsType int) (*sms_model.SmsTemplateConfig, error)
	}
	ISmsAppConfig interface {
		GetAppConfigByName(ctx context.Context, appName string) (*sms_model.SmsAppConfig, error)
		GetAppConfigById(ctx context.Context, id int64) (*sms_model.SmsAppConfig, error)
		GetAppAvailableNumber(ctx context.Context, id int64) (int, error)
		CreateAppConfig(ctx context.Context, config *sms_model.SmsAppConfig) (bool, error)
		UpdateAppNumber(ctx context.Context, id int64, fee uint64) (bool, error)
	}
	ISmsBusinessConfig interface {
		GetBusinessConfigByAppId(ctx context.Context, id int64) (*sms_model.SmsBusinessConfig, error)
		GetBusinessConfigById(ctx context.Context, id int64) (*sms_model.SmsBusinessConfig, error)
		CreateBusinessConfig(ctx context.Context, config *sms_model.SmsBusinessConfig) (*sms_model.SmsBusinessConfig, error)
	}
	ISmsSendLogConfig interface {
		SaveSmsLog(ctx context.Context, info *sms_model.SmsSendLog) (bool, error)
		GetSmsLogById(ctx context.Context, id int64) (res *sms_model.SmsSendLog, err error)
	}
	ISmsServiceProviderConfig interface {
		CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error)
		GetProviderById(ctx context.Context, id int64) (*sms_model.SmsServiceProviderConfig, error)
		GetProviderByPriority(ctx context.Context, priority int) (*sms_model.SmsServiceProviderConfig, error)
		QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*sms_model.ServiceProviderConfigListRes, error)
		QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error)
	}
	ISmsSignConfig interface {
		CreateSign(ctx context.Context, info *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error)
		AuditSign(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error)
		GetSignBySignName(ctx context.Context, signName string) (res *sms_model.SmsSignConfig, err error)
		GetSignById(ctx context.Context, id int64) (*sms_model.SmsSignConfig, error)
	}

	ISmsAliyun interface {
		SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error)
		ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error)
		RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error)
		RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error)
		CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error)
	}

	ISmsTencent interface {
		SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error)
		ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error)
		RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error)
		RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error)
		CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error)
	}

	ISmsQyxs interface {
		SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error)
		ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error)
		RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error)
		RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error)
		CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error)
	}
)
type IModules interface {
	SmsSendLogConfig() ISmsSendLogConfig
	SmsAppConfig() ISmsAppConfig
	SmsServiceProviderConfig() ISmsServiceProviderConfig
	SmsTemplateConfig() ISmsTemplateConfig
	SmsSignConfig() ISmsSignConfig
	SmsBusinessConfig() ISmsBusinessConfig

	SmsAliyun() ISmsAliyun
	SmsTencent() ISmsTencent
	SmsQyxs() ISmsQyxs

	GetConfig() *sms_model.Config
	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
	Dao() *sms_dao.XDao
}
