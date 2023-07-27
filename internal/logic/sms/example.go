package sms

import (
	"context"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/sms-library/sms_model"
)

// SmsTest 企业信使
type SmsTest struct {
	// 只要写了这个匿名字段，那么就可以实现接口的部分方法
	SMSer
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error) （不属于接口，属于调用者校验）
//func (t *SmsTest) VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
//	return true, nil
//}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (t *SmsTest) SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error) {
	return nil, nil
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (t *SmsTest) ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error) {
	return true, nil
}

// GetAppConfigById 根据应用id查询应用
func (t *SmsTest) GetAppConfigById(ctx context.Context, id int64) (*sms_model.SmsAppConfig, error) {
	return nil, nil
}

// CreateAppConfig 创建应用 (上下文, 应用信息)
func (t *SmsTest) CreateAppConfig(ctx context.Context, config *sms_model.SmsAppConfig) (bool, error) {
	return false, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func (t *SmsTest) GetAppAvailableNumber(ctx context.Context, appid int64) (int, error) {
	return 0, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (t *SmsTest) UpdateAppNumber(ctx context.Context, appId int64, fee uint64) (bool, error) {
	return false, nil
}

// RegisterTemplate 添加短信模版
func (t *SmsTest) RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error) {
	return nil, nil
}

// AuditTemplate 短信模版审核
func (t *SmsTest) AuditTemplate(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error) {
	return true, nil
}

// GetByTemplateCode 根据模版Code查询模版信息
func (t *SmsTest) GetByTemplateCode(ctx context.Context, templateCode string) (*sms_model.SmsTemplateConfig, error) {
	return nil, nil
}

// RegisterSign 添加短信签名
func (t *SmsTest) RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error) {
	return nil, nil
}

// AuditSign 审核短信签名
func (t *SmsTest) AuditSign(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error) {
	return false, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func (t *SmsTest) GetSignBySignName(ctx context.Context, signName string) (*sms_model.SmsSignConfig, error) {
	return nil, nil
}

// CreateProvider 添加渠道商
func (t *SmsTest) CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error) {
	return nil, nil
}

// QueryProviderByNo 根据No编号获取渠道商
func (t *SmsTest) QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*sms_model.ServiceProviderConfigListRes, error) {
	return nil, nil
}

// QueryProviderList 获取渠道商列表
func (t *SmsTest) QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error) {
	return nil, nil
}

// CreateBusiness 创建业务
func (t *SmsTest) CreateBusiness(ctx context.Context, info *sms_model.SmsBusinessConfig) (*sms_model.SmsBusinessConfig, error) {
	return nil, nil
}

// GetBusinessByAppId 根据应用id查询业务
func (t *SmsTest) GetBusinessByAppId(ctx context.Context, appId int64) (*sms_model.SmsBusinessConfig, error) {
	return nil, nil
}
