package sms_module

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/kysion/sms-library/internal/boot"
	"github.com/kysion/sms-library/internal/logic/sms"
	"github.com/kysion/sms-library/internal/logic/sms_aliyun"
	"github.com/kysion/sms-library/internal/logic/sms_qyxs"
	"github.com/kysion/sms-library/internal/logic/sms_tencent"
	"github.com/kysion/sms-library/sms_consts"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
)

type Modules struct {
	conf                     *sms_model.Config
	smsSendLog               sms_interface.ISmsSendLogConfig
	smsAppConfig             sms_interface.ISmsAppConfig
	smsServiceProviderConfig sms_interface.ISmsServiceProviderConfig
	smsTemplateConfig        sms_interface.ISmsTemplateConfig
	smsSignConfig            sms_interface.ISmsSignConfig
	smsBusinessConfig        sms_interface.ISmsBusinessConfig

	smsAliyun  sms_interface.ISmsAliyun
	smsTencent sms_interface.ISmsTencent
	smsQyxs    sms_interface.ISmsQyxs

	i18n *gi18n.Manager
	xDao *sms_dao.XDao
}

func (m *Modules) SmsSendLogConfig() sms_interface.ISmsSendLogConfig {
	return m.smsSendLog
}

func (m *Modules) SmsAppConfig() sms_interface.ISmsAppConfig {
	return m.smsAppConfig
}

func (m *Modules) SmsServiceProviderConfig() sms_interface.ISmsServiceProviderConfig {
	return m.smsServiceProviderConfig
}
func (m *Modules) SmsTemplateConfig() sms_interface.ISmsTemplateConfig {
	return m.smsTemplateConfig
}
func (m *Modules) SmsSignConfig() sms_interface.ISmsSignConfig {
	return m.smsSignConfig
}
func (m *Modules) SmsBusinessConfig() sms_interface.ISmsBusinessConfig {
	return m.smsBusinessConfig
}
func (m *Modules) SmsAliyun() sms_interface.ISmsAliyun {
	return m.smsAliyun
}
func (m *Modules) SmsTencent() sms_interface.ISmsTencent {
	return m.smsTencent
}
func (m *Modules) SmsQyxs() sms_interface.ISmsQyxs {
	return m.smsQyxs
}

func (m *Modules) GetConfig() *sms_model.Config {
	return m.conf
}

func (m *Modules) T(ctx context.Context, content string) string {
	return m.i18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.i18n.TranslateFormat(ctx, format, values...)
}

func (m *Modules) SetI18n(i18n *gi18n.Manager) error {
	if i18n == nil {
		i18n = gi18n.New()
		i18n.SetLanguage("zh-CN")
		//err := i18n.SetPath("i18n/" + gstr.ToLower(m.conf.KeyIndex))
		err := i18n.SetPath("i18n/" + m.conf.I18nName)
		if err != nil {
			return err
		}
	}

	m.i18n = i18n
	return nil
}

func (m *Modules) Dao() *sms_dao.XDao {
	return m.xDao
}

func NewModules(
	conf *sms_model.Config,
	xDao *sms_dao.XDao,
) *Modules {
	module := &Modules{
		conf: conf,
		xDao: xDao,
	}

	// 初始化默认多语言对象
	module.SetI18n(nil)

	module.smsSendLog = sms.NewSmsSendLogConfig(module)
	module.smsAppConfig = sms.NewSmsAppConfig(module)
	module.smsServiceProviderConfig = sms.NewSmsServiceProviderConfig(module)
	module.smsTemplateConfig = sms.NewSmsTemplateConfig(module)
	module.smsSignConfig = sms.NewSmsSignConfig(module)
	module.smsBusinessConfig = sms.NewSmsBusinessConfig(module)
	module.smsAliyun = sms_aliyun.NewSmsAliyun(module)
	module.smsTencent = sms_tencent.NewSmsTencent(module)
	module.smsQyxs = sms_qyxs.NewSmsQyxs(module)

	// 权限树追加权限
	sms_consts.PermissionTree = append(sms_consts.PermissionTree, boot.InitPermission(module, nil)...)

	return module
}
