package i_controller

import (
	"context"
	"github.com/kysion/sms-library/api/sms_api"
	"github.com/kysion/sms-library/sms_model"
)

type ISms interface {
	iModule

	SendSms(ctx context.Context, req *sms_api.SendSmsReq) (res *sms_api.SmsResponseRes, err error)

	ReceiveSms(ctx context.Context, req *sms_api.ReceiveSmsReq) (sms_api.BoolRes, error)

	GetAppConfigById(ctx context.Context, req *sms_api.GetAppConfigByIdReq) (*sms_api.SmsAppConfigRes, error)

	CreateAppConfig(ctx context.Context, req *sms_api.CreateAppConfigReq) (sms_api.BoolRes, error)

	GetAppAvailableNumber(ctx context.Context, req *sms_api.GetAppAvailableNumberReq) (sms_api.IntRes, error)

	// UpdateAppNumber(ctx context.Context, req *sms_api.UpdateAppNumberReq) (sms_api.BoolRes, error)

	RegisterTemplate(ctx context.Context, req *sms_api.RegisterTemplateReq) (res *sms_api.SmsTemplateConfigRes, err error)

	AuditTemplate(ctx context.Context, req *sms_api.AuditTemplateReq) (sms_api.BoolRes, error)

	GetByTemplateCode(ctx context.Context, req *sms_api.GetByTemplateCodeReq) (*sms_api.SmsTemplateConfigRes, error)

	RegisterSign(ctx context.Context, req *sms_api.RegisterSignReq) (res *sms_api.SmsSignConfigRes, err error)

	AuditSign(ctx context.Context, req *sms_api.AuditSignReq) (sms_api.BoolRes, error)

	GetSignBySignName(ctx context.Context, req *sms_api.GetSignBySignNameReq) (*sms_api.SmsSignConfigRes, error)

	CreateProvider(ctx context.Context, req *sms_api.CreateProviderReq) (res *sms_api.SmsServiceProviderConfigRes, err error)

	QueryProviderByNo(ctx context.Context, req *sms_api.QueryProviderByNoReq) (*sms_model.ServiceProviderConfigListRes, error)

	QueryProviderList(ctx context.Context, req *sms_api.QueryProviderListReq) (*sms_model.ServiceProviderConfigListRes, error)

	GetProviderById(ctx context.Context, req *sms_api.GetProviderByIdReq) (*sms_api.SmsServiceProviderConfigRes, error)

	CreateBusinessReq(ctx context.Context, req *sms_api.CreateBusinessReq) (*sms_api.SmsBusinessRes, error)

	GetBusinessByAppIdReq(ctx context.Context, req *sms_api.GetBusinessByAppIdReq) (*sms_api.SmsBusinessRes, error)
}
