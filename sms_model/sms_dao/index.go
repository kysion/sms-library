package sms_dao

type XDao struct {
	SmsSendLog               SmsSendLogDao
	SmsAppConfig             SmsAppConfigDao
	SmsServiceProviderConfig SmsServiceProviderConfigDao
	SmsTemplateConfig        SmsTemplateConfigDao
	SmsSignConfig            SmsSignConfigDao
	SmsBusinessConfig        SmsBusinessConfigDao
}
