package sms

import (
	"context"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/sms-library/sms_model"
)

// 我需要具备哪些功能？
//  - 发送短信
//  - 接收短信
//  - 查询短信相关信息（账户用量统计、查询短信信息）
//  - 模版管理（添加、删除、修改、查询、审核模版）
//  - 签名管理（同上）
//  - 渠道商管理（不同的短信短信平台配置管理）
//  - 应用管理
//  - 业务管理

type SMSer interface {
	// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error) （不属于接口，属于调用者校验）
	//VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error)

	// SendSms 发送短信 (渠道商, 短信模版,请求内容)
	SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error)

	// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
	ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error)

	// GetAppConfigById 根据应用id查询应用
	GetAppConfigById(ctx context.Context, id int64) (*sms_model.SmsAppConfig, error)

	// CreateAppConfig 创建应用 (上下文, 应用信息)
	CreateAppConfig(ctx context.Context, config *sms_model.SmsAppConfig) (bool, error)

	// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
	GetAppAvailableNumber(ctx context.Context, appId int64) (int, error)

	// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
	UpdateAppNumber(ctx context.Context, appId int64, fee uint64) (bool, error)

	// RegisterTemplate 添加短信模版
	RegisterTemplate(ctx context.Context, signInfo *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error)

	// AuditTemplate 审核短信模版
	AuditTemplate(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error)

	// GetByTemplateCode 根据模版Code查询模版信息
	GetByTemplateCode(ctx context.Context, templateCode string) (*sms_model.SmsTemplateConfig, error)

	// RegisterSign 添加短信签名
	RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error)

	// AuditSign 审核短信签名
	AuditSign(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error)

	// GetSignBySignName 根据签名名称查找签名数据
	GetSignBySignName(ctx context.Context, signName string) (*sms_model.SmsSignConfig, error)

	// CreateProvider 添加渠道商
	CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error)

	// QueryProviderByNo 根据No编号获取渠道商列表
	QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*sms_model.ServiceProviderConfigListRes, error)

	// QueryProviderList 获取渠道商列表
	QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error)

	// CreateBusiness 创建业务
	CreateBusiness(ctx context.Context, info *sms_model.SmsBusinessConfig) (*sms_model.SmsBusinessConfig, error)

	// GetBusinessByAppId 根据应用id查询业务
	GetBusinessByAppId(ctx context.Context, appId int64) (*sms_model.SmsBusinessConfig, error)
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error) （不属于接口，属于调用者校验）
//func VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
//	return true, nil
//}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error) {
	return nil, nil
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error) {
	return true, nil
}

// GetAppConfigById 根据应用id查询应用
func GetAppConfigById(ctx context.Context, id int64) (*sms_model.SmsAppConfig, error) {
	return nil, nil
}

// CreateAppConfig 创建应用 (上下文, 应用信息)
func CreateAppConfig(ctx context.Context, config *sms_model.SmsAppConfig) (bool, error) {
	return false, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func GetAppAvailableNumber(ctx context.Context, appId int64) (int, error) {
	return 0, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func UpdateAppNumber(ctx context.Context, appId int64, fee uint64) (bool, error) {
	return false, nil
}

// RegisterTemplate 添加短信模版
func RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error) {
	return nil, nil
}

// AuditTemplate 短信模版审核
func AuditTemplate(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error) {
	return true, nil
}

// GetByTemplateCode 根据模版编号查询模版信息
func GetByTemplateCode(ctx context.Context, templateCode string) (*sms_model.SmsTemplateConfig, error) {
	return nil, nil
}

// RegisterSign 添加短信签名
func RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error) {
	return nil, nil
}

// AuditSign 审核短信签名
func AuditSign(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error) {
	return false, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func GetSignBySignName(ctx context.Context, signName string) (*sms_model.SmsSignConfig, error) {
	return nil, nil
}

// CreateProvider 添加渠道商
func CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error) {
	return nil, nil
}

// GetProviderByNo 根据No编号获取渠道商
func GetProviderByNo(ctx context.Context, no string) (*sms_model.SmsServiceProviderConfig, error) {
	return nil, nil
}

// QueryProviderList 获取渠道商列表
func QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error) {
	return nil, nil
}

// CreateBusiness 创建业务
func CreateBusiness(ctx context.Context, info *sms_model.SmsBusinessConfig) (*sms_model.SmsBusinessConfig, error) {
	return nil, nil
}

// GetBusinessByAppId 根据应用id查询业务
func GetBusinessByAppId(ctx context.Context, appId int64) (*sms_model.SmsBusinessConfig, error) {
	return nil, nil
}
