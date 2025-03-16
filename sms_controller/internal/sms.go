package internal

import (
	"context"
	"errors"

	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/enum"
	"github.com/kysion/sms-library/api/sms_api"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_interface/i_controller"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_enum"
)

// SmsController 部分接口需要根据不同的请求调用不同的短信平台，但是大部分接口直接调用底层dao操作即可。
type SmsController struct {
	i_controller.ISms
	modules sms_interface.IModules
}

// Sms 短信服务
var Sms = func(modules sms_interface.IModules) i_controller.ISms {
	return &SmsController{
		modules: modules,
	}
}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (c *SmsController) SendSms(ctx context.Context, req *sms_api.SendSmsReq) (res *sms_api.SmsResponseRes, err error) {
	ret := &sms_model.SmsResponse{}
	// 准备平台配置信息
	provider, err := c.modules.SmsServiceProviderConfig().GetProviderByPriority(ctx, 1) // 查找优先级别最高的短信配置渠道商
	if err != nil {
		return nil, errors.New("{#error_sms_controller_provider_config_query_failed}")
	}

	// 寻找匹配的短信模版
	template := &sms_model.SmsTemplateConfig{}
	captchaTypes := enum.GetTypes[int, base_enum.CaptchaType](req.CaptchaType, base_enum.Captcha.Type)
	for _, value := range captchaTypes {
		template, err = c.modules.SmsTemplateConfig().GetByProviderNoAndType(ctx, sms_enum.Sms.Type.New(provider.ProviderNo), value.Code())
		if err != nil {
			return nil, errors.New("{#error_sms_controller_template_query_failed}")
		}
		if template != nil {
			isOk := false
			for _, value2 := range captchaTypes {
				if template.Type&value2.Code() == value2.Code() {
					isOk = true
				} else {
					isOk = false
				}
			}
			// 找到符合所有业务场景的短信模版 就退出查找
			if isOk {
				break
			}
		}
	}
	if template == nil || err != nil {
		return nil, errors.New("{#error_sms_controller_template_query_failed}")
	}

	// TODO 后续删掉：如下场景只适合单一类型的验证码业务场景，不支持复合类型
	//template, err := c.modules.SmsTemplateConfig().GetByProviderNoAndType(ctx, sms_enum.Sms.Type.New(provider.ProviderNo), req.CaptchaType)
	//if err != nil {
	//	return nil, errors.New("{#error_sms_controller_template_query_failed}")
	//}

	// 选择对应平台发送短信
	switch provider.ProviderNo {
	// 阿里云发送短信
	case sms_enum.Sms.Type.Aliyun.Code():
		ret, err = c.modules.SmsAliyun().SendSms(ctx, *provider, *template, req.SmsSendMessageReq)

	// 腾讯云发送短信
	case sms_enum.Sms.Type.Tencent.Code():
		ret, err = c.modules.SmsTencent().SendSms(ctx, *provider, *template, req.SmsSendMessageReq)

	// 华为云发送短信
	case sms_enum.Sms.Type.Huawei.Code():

	// 七牛云发送短息
	case sms_enum.Sms.Type.Qiniu.Code():

	// 企业信使发送短信
	default:
		ret, err = c.modules.SmsQyxs().SendSms(ctx, *provider, *template, req.SmsSendMessageReq)
	}
	// 发送短信

	return (*sms_api.SmsResponseRes)(ret), err
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (c *SmsController) ReceiveSms(ctx context.Context, req *sms_api.ReceiveSmsReq) (res sms_api.BoolRes, err error) {
	var ret bool
	// 应用场景，例如：用户需要退订，那么他是需要发送短信，然后我们才能进行确认的，如果接收到用户的退订消息，那么我们就阻止程序不给该用户发送短信业务。
	// 选择对应平台
	switch req.SmsReceiveSmsReq.ProviderNo {
	// 阿里云
	case sms_enum.Sms.Type.Aliyun.Code():
		ret, err = c.modules.SmsAliyun().ReceiveSms(ctx, req.SmsReceiveSmsReq)
	// 腾讯云
	case sms_enum.Sms.Type.Tencent.Code():
		ret, err = c.modules.SmsTencent().ReceiveSms(ctx, req.SmsReceiveSmsReq)

	// 华为云
	case sms_enum.Sms.Type.Huawei.Code():

	// 七牛云
	case sms_enum.Sms.Type.Qiniu.Code():

	// 企业信使
	case sms_enum.Sms.Type.Qyxs.Code():
		ret, err = c.modules.SmsQyxs().ReceiveSms(ctx, req.SmsReceiveSmsReq)

	default:
		return false, errors.New("{#error_sms_controller_not_supported_this_provider}")
	}

	return ret == true, err
}

// GetAppConfigById 根据应用id查询应用
func (c *SmsController) GetAppConfigById(ctx context.Context, req *sms_api.GetAppConfigByIdReq) (*sms_api.SmsAppConfigRes, error) {
	ret, err := c.modules.SmsAppConfig().GetAppConfigById(ctx, req.Id)

	return (*sms_api.SmsAppConfigRes)(ret), err
}

// CreateAppConfig 创建应用 (上下文, 应用编号, 花费数量)
func (c *SmsController) CreateAppConfig(ctx context.Context, req *sms_api.CreateAppConfigReq) (sms_api.BoolRes, error) {
	ret, err := c.modules.SmsAppConfig().CreateAppConfig(ctx, &req.SmsAppConfig)

	return ret == true, err
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func (c *SmsController) GetAppAvailableNumber(ctx context.Context, req *sms_api.GetAppAvailableNumberReq) (sms_api.IntRes, error) {
	ret, err := c.modules.SmsAppConfig().GetAppAvailableNumber(ctx, req.Id)

	return (sms_api.IntRes)(ret), err
}

//
//// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
//func (c *SmsController) UpdateAppNumber(ctx context.Context, req *sms_api.UpdateAppNumberReq) (api_v1.BoolRes, error) {
//	ret, err := c.modules.SmsAppConfig().UpdateAppNumber(ctx, req.AppNo, req.Fee)
//
//	return ret == true, err
//}

// RegisterTemplate 添加短信模版
func (c *SmsController) RegisterTemplate(ctx context.Context, req *sms_api.RegisterTemplateReq) (res *sms_api.SmsTemplateConfigRes, err error) {
	// 根据不同的平台差异操作
	// c.modules.SmsQyxs().RegisterTemplate(ctx, &req.SmsTemplateConfig)   // 企业信使
	// c.modules.SmsAliyun().RegisterTemplate(ctx, &req.SmsTemplateConfig) // 阿里云
	// c.modules.SmsTencent().RegisterTemplate(ctx, &req.SmsTemplateConfig) // 腾讯云

	var ret *sms_model.SmsTemplateConfig
	// 选择对应平台
	switch req.SmsTemplateConfig.ProviderNo {
	// 阿里云
	case sms_enum.Sms.Type.Aliyun.Code():
		ret, err = c.modules.SmsAliyun().RegisterTemplate(ctx, &req.SmsTemplateConfig)
	// 腾讯云
	case sms_enum.Sms.Type.Tencent.Code():
		ret, err = c.modules.SmsTencent().RegisterTemplate(ctx, &req.SmsTemplateConfig)

	// 华为云
	case sms_enum.Sms.Type.Huawei.Code():

	// 七牛云
	case sms_enum.Sms.Type.Qiniu.Code():

	// 企业信使
	case sms_enum.Sms.Type.Qyxs.Code():
		ret, err = c.modules.SmsQyxs().RegisterTemplate(ctx, &req.SmsTemplateConfig)

	default:
		return nil, errors.New("{#error_sms_controller_not_supported_this_provider}")
	}

	//ret, err := c.modules.SmsSignConfig().CreateSign(ctx, &req.SmsSignConfig)

	//ret, err := c.modules.SmsTemplateConfig().CreateTemplate(ctx, &req.SmsTemplateConfig)

	return (*sms_api.SmsTemplateConfigRes)(ret), err
}

// AuditTemplate 短信模版审核
func (c *SmsController) AuditTemplate(ctx context.Context, req *sms_api.AuditTemplateReq) (sms_api.BoolRes, error) {
	ret, err := c.modules.SmsTemplateConfig().AuditTemplate(ctx, req.Id, &req.AuditInfo)

	return ret == true, err
}

// GetByTemplateCode 根据模版编号查询模版信息
func (c *SmsController) GetByTemplateCode(ctx context.Context, req *sms_api.GetByTemplateCodeReq) (*sms_api.SmsTemplateConfigRes, error) {
	ret, err := c.modules.SmsTemplateConfig().GetByTemplateCode(ctx, req.TemplateCode)

	return (*sms_api.SmsTemplateConfigRes)(ret), err
}

// RegisterSign 添加短信签名
func (c *SmsController) RegisterSign(ctx context.Context, req *sms_api.RegisterSignReq) (res *sms_api.SmsSignConfigRes, err error) {
	var ret *sms_model.SmsSignConfig
	// 选择对应平台
	switch req.SmsSignConfig.ProviderNo {
	// 阿里云
	case sms_enum.Sms.Type.Aliyun.Code():
		ret, err = c.modules.SmsAliyun().RegisterSign(ctx, &req.SmsSignConfig)
	// 腾讯云
	case sms_enum.Sms.Type.Tencent.Code():
		ret, err = c.modules.SmsTencent().RegisterSign(ctx, &req.SmsSignConfig)

	// 华为云
	case sms_enum.Sms.Type.Huawei.Code():

	// 七牛云
	case sms_enum.Sms.Type.Qiniu.Code():

	// 企业信使
	case sms_enum.Sms.Type.Qyxs.Code():
		ret, err = c.modules.SmsQyxs().RegisterSign(ctx, &req.SmsSignConfig)

	default:
		return nil, errors.New("{#error_sms_controller_not_supported_this_provider}")
	}

	//ret, err := c.modules.SmsSignConfig().CreateSign(ctx, &req.SmsSignConfig)

	return (*sms_api.SmsSignConfigRes)(ret), err
}

// AuditSign 审核短信签名
func (c *SmsController) AuditSign(ctx context.Context, req *sms_api.AuditSignReq) (sms_api.BoolRes, error) {
	ret, err := c.modules.SmsSignConfig().AuditSign(ctx, req.Id, &req.AuditInfo)

	return ret == true, err
}

// GetSignBySignName 根据签名名称查找签名数据
func (c *SmsController) GetSignBySignName(ctx context.Context, req *sms_api.GetSignBySignNameReq) (*sms_api.SmsSignConfigRes, error) {
	ret, err := c.modules.SmsSignConfig().GetSignBySignName(ctx, req.SignName)

	return (*sms_api.SmsSignConfigRes)(ret), err

}

// CreateProvider 添加渠道商
func (c *SmsController) CreateProvider(ctx context.Context, req *sms_api.CreateProviderReq) (res *sms_api.SmsServiceProviderConfigRes, err error) {
	//// 准备平台配置信息
	//provider := sms_model.SmsServiceProviderConfig{}
	//
	//template := sms_model.SmsTemplateConfig{}

	var ret *sms_model.SmsServiceProviderConfig
	// 选择对应平台
	switch req.SmsServiceProviderConfig.ProviderNo {
	// 阿里云
	case sms_enum.Sms.Type.Aliyun.Code():
		ret, err = c.modules.SmsAliyun().CreateProvider(ctx, &req.SmsServiceProviderConfig)
	// 腾讯云
	case sms_enum.Sms.Type.Tencent.Code():
		ret, err = c.modules.SmsTencent().CreateProvider(ctx, &req.SmsServiceProviderConfig)

	// 华为云
	case sms_enum.Sms.Type.Huawei.Code():

	// 七牛云
	case sms_enum.Sms.Type.Qiniu.Code():

	// 企业信使
	case sms_enum.Sms.Type.Qyxs.Code():
		c.modules.SmsQyxs().CreateProvider(ctx, &req.SmsServiceProviderConfig)

	default:
		return nil, errors.New("{#error_sms_controller_not_supported_this_provider}")
	}

	//ret, err := c.modules.SmsServiceProviderConfig().CreateProvider(ctx, &req.SmsServiceProviderConfig)

	return (*sms_api.SmsServiceProviderConfigRes)(ret), err
}

// QueryProviderByNo 根据No编号获取渠道商列表
func (c *SmsController) QueryProviderByNo(ctx context.Context, req *sms_api.QueryProviderByNoReq) (*sms_model.ServiceProviderConfigListRes, error) {
	ret, err := c.modules.SmsServiceProviderConfig().QueryProviderByNo(ctx, req.No, &req.SearchParams)

	return ret, err
}

// QueryProviderList 获取渠道商列表
func (c *SmsController) QueryProviderList(ctx context.Context, req *sms_api.QueryProviderListReq) (*sms_model.ServiceProviderConfigListRes, error) {
	ret, err := c.modules.SmsServiceProviderConfig().QueryProviderList(ctx, &req.SearchParams, false)

	return ret, err
}

// GetProviderById 根据Id获取渠道商
func (c *SmsController) GetProviderById(ctx context.Context, req *sms_api.GetProviderByIdReq) (*sms_api.SmsServiceProviderConfigRes, error) {
	ret, err := c.modules.SmsServiceProviderConfig().GetProviderById(ctx, req.Id)

	return (*sms_api.SmsServiceProviderConfigRes)(ret), err
}

// CreateBusinessReq 创建业务 (上下文, 业务信息)
func (c *SmsController) CreateBusinessReq(ctx context.Context, req *sms_api.CreateBusinessReq) (*sms_api.SmsBusinessRes, error) {
	ret, err := c.modules.SmsBusinessConfig().CreateBusinessConfig(ctx, &req.SmsBusinessConfig)

	return (*sms_api.SmsBusinessRes)(ret), err
}

// GetBusinessByAppIdReq 根据应用id获取BusinessConfig
func (c *SmsController) GetBusinessByAppIdReq(ctx context.Context, req *sms_api.GetBusinessByAppIdReq) (*sms_api.SmsBusinessRes, error) {
	ret, err := c.modules.SmsBusinessConfig().GetBusinessConfigByAppId(ctx, req.AppId)

	return (*sms_api.SmsBusinessRes)(ret), err
}
