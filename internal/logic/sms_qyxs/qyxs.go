package sms_qyxs

import (
	"context"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
)

// sSmsQyxs 企业信使短信平台
type sSmsQyxs struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsQyxs(modules sms_interface.IModules) sms_interface.ISmsQyxs {
	return &sSmsQyxs{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error) （不属于接口，属于调用者校验）
//func (t *SmsQyxs) VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
//	return true, nil
//}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (s *sSmsQyxs) SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error) {
	return nil, nil
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (s *sSmsQyxs) ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error) {
	return true, nil
}

// RegisterTemplate 添加短信模版
func (s *sSmsQyxs) RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error) {

	return nil, nil
}

// RegisterSign 添加短信签名
func (s *sSmsQyxs) RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error) {
	return nil, nil
}

// CreateProvider 添加渠道商
func (s *sSmsQyxs) CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error) {
	return nil, nil
}

//// QueryProviderList 获取渠道商列表
//func (s *sSmsQyxs) QueryProviderList(ctx context.Context, search *sys_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error) {
//	return nil, nil
//}
