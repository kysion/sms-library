package controller

import (
	"context"
	"github.com/kysion/sms-library/api/sms_api"
	"github.com/kysion/sms-library/api/sms_v1"
	"github.com/kysion/sms-library/sms_controller"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_interface/i_controller"
	"github.com/kysion/sms-library/sms_model"
)

type SmsController struct {
	i_controller.ISms
}

var Sms = func(modules sms_interface.IModules) *SmsController {
	return &SmsController{
		sms_controller.Sms(modules),
	}
}

func (c *SmsController) GetModules() sms_interface.IModules {
	return c.ISms.GetModules()
}

func (c *SmsController) SendSms(ctx context.Context, req *sms_v1.SendSmsReq) (res *sms_api.SmsResponseRes, err error) {
	return c.ISms.SendSms(ctx, &req.SendSmsReq)
}

func (c *SmsController) ReceiveSms(ctx context.Context, req *sms_v1.ReceiveSmsReq) (sms_api.BoolRes, error) {
	return c.ISms.ReceiveSms(ctx, &req.ReceiveSmsReq)
}

func (c *SmsController) GetAppConfigById(ctx context.Context, req *sms_v1.GetAppConfigByIdReq) (*sms_api.SmsAppConfigRes, error) {
	return c.ISms.GetAppConfigById(ctx, &req.GetAppConfigByIdReq)
}

func (c *SmsController) CreateAppConfig(ctx context.Context, req *sms_v1.CreateAppConfigReq) (sms_api.BoolRes, error) {
	return c.ISms.CreateAppConfig(ctx, &req.CreateAppConfigReq)
}

func (c *SmsController) GetAppAvailableNumber(ctx context.Context, req *sms_v1.GetAppAvailableNumberReq) (sms_api.IntRes, error) {
	return c.ISms.GetAppAvailableNumber(ctx, &req.GetAppAvailableNumberReq)
}

//func (c *SmsController) UpdateAppNumber(ctx context.Context, req *sms_v1.UpdateAppNumberReq) (sms_api.BoolRes, error) {
//return c.ISmsUpdateAppNumberQueryProviderList(ctx,&req.)}UpdateAppNumberReqr

func (c *SmsController) RegisterTemplate(ctx context.Context, req *sms_v1.RegisterTemplateReq) (res *sms_api.SmsTemplateConfigRes, err error) {
	return c.ISms.RegisterTemplate(ctx, &req.RegisterTemplateReq)
}

func (c *SmsController) AuditTemplate(ctx context.Context, req *sms_v1.AuditTemplateReq) (sms_api.BoolRes, error) {
	return c.ISms.AuditTemplate(ctx, &req.AuditTemplateReq)
}

func (c *SmsController) GetByTemplateCode(ctx context.Context, req *sms_v1.GetByTemplateCodeReq) (*sms_api.SmsTemplateConfigRes, error) {
	return c.ISms.GetByTemplateCode(ctx, &req.GetByTemplateCodeReq)
}

func (c *SmsController) RegisterSign(ctx context.Context, req *sms_v1.RegisterSignReq) (res *sms_api.SmsSignConfigRes, err error) {
	return c.ISms.RegisterSign(ctx, &req.RegisterSignReq)
}

func (c *SmsController) AuditSign(ctx context.Context, req *sms_v1.AuditSignReq) (sms_api.BoolRes, error) {
	return c.ISms.AuditSign(ctx, &req.AuditSignReq)
}

func (c *SmsController) GetSignBySignName(ctx context.Context, req *sms_v1.GetSignBySignNameReq) (*sms_api.SmsSignConfigRes, error) {
	return c.ISms.GetSignBySignName(ctx, &req.GetSignBySignNameReq)
}

func (c *SmsController) CreateProvider(ctx context.Context, req *sms_v1.CreateProviderReq) (res *sms_api.SmsServiceProviderConfigRes, err error) {
	return c.ISms.CreateProvider(ctx, &req.CreateProviderReq)
}

func (c *SmsController) QueryProviderByNo(ctx context.Context, req *sms_v1.QueryProviderByNoReq) (*sms_model.ServiceProviderConfigListRes, error) {
	return c.ISms.QueryProviderByNo(ctx, &req.QueryProviderByNoReq)
}

func (c *SmsController) QueryProviderList(ctx context.Context, req *sms_v1.QueryProviderListReq) (*sms_model.ServiceProviderConfigListRes, error) {
	return c.ISms.QueryProviderList(ctx, &req.QueryProviderListReq)
}

func (c *SmsController) GetProviderById(ctx context.Context, req *sms_v1.GetProviderByIdReq) (*sms_api.SmsServiceProviderConfigRes, error) {
	return c.ISms.GetProviderById(ctx, &req.GetProviderByIdReq)
}
func (c *SmsController) CreateBusinessReq(ctx context.Context, req *sms_v1.CreateBusinessReq) (*sms_api.SmsBusinessRes, error) {
	return c.ISms.CreateBusinessReq(ctx, &req.CreateBusinessReq)
}

func (c *SmsController) GetBusinessByAppIdReq(ctx context.Context, req *sms_v1.GetBusinessByAppIdReq) (*sms_api.SmsBusinessRes, error) {
	return c.ISms.GetBusinessByAppIdReq(ctx, &req.GetBusinessByAppIdReq)
}
