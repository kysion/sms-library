package sms_global

import (
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_module"
)

type global struct {
	Modules sms_interface.IModules
}

var (
	PermissionTree []base_permission.IPermission

	Global = global{
		Modules: sms_module.NewModules(
			&sms_model.Config{
				HardDeleteWaitAt: 0,
				KeyIndex:         "Sms",
				I18nName:         "sms",
				RoutePrefix:      "/sms",
				StoragePath:      "./resource/sms",
				Identifier: sms_model.Identifier{
					Sms:                      "sms",
					SmsSendLog:               "smsSendLog",
					SmsAppConfig:             "smsAppConfig",
					SmsServiceProviderConfig: "smsServiceProviderConfig",
					SmsTemplateConfig:        "smsTemplateConfig",
					SmsSignConfig:            "smsSignConfig",
					SmsBusinessConfig:        "smsBusinessConfig",
				},
			},
			&sms_dao.XDao{ // 以下为业务层实例化dao模型，如果不是使用默认模型时需要将自定义dao模型作为参数传进去
				SmsSendLog:               sms_dao.NewSmsSendLog(),
				SmsAppConfig:             sms_dao.NewSmsAppConfig(),
				SmsServiceProviderConfig: sms_dao.NewSmsServiceProviderConfig(),
				SmsTemplateConfig:        sms_dao.NewSmsTemplateConfig(),
				SmsSignConfig:            sms_dao.NewSmsSignConfig(),
				SmsBusinessConfig:        sms_dao.NewSmsBusinessConfig(),
			},
		),
	}
)
