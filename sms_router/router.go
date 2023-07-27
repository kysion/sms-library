package sms_router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kysion/sms-library/sms_controller"
	"github.com/kysion/sms-library/sms_interface"
)

func ModulesGroup(modules sms_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	SmsGroup(modules, group)

	return group
}

func SmsGroup(modules sms_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/sms"
	controller := sms_controller.Sms(modules)

	//group.POST(routePrefix+"/sendCaptchaBySms", controller.)

	group.POST(routePrefix+"/sendSms", controller.SendSms)
	group.POST(routePrefix+"/receiveSms", controller.ReceiveSms)
	group.POST(routePrefix+"/getAppConfigById", controller.GetAppConfigById)
	group.POST(routePrefix+"/createAppConfig", controller.CreateAppConfig)
	group.POST(routePrefix+"/getAppAvailableNumber", controller.GetAppAvailableNumber)

	// group.POST(routePrefix+"/updateAppNumber", controller.UpdateAppNumber)
	group.POST(routePrefix+"/registerTemplate", controller.RegisterTemplate)
	group.POST(routePrefix+"/auditTemplate", controller.AuditTemplate)
	group.POST(routePrefix+"/getByTemplateCode", controller.GetByTemplateCode)
	group.POST(routePrefix+"/registerSign", controller.RegisterSign)
	group.POST(routePrefix+"/auditSign", controller.AuditSign)
	group.POST(routePrefix+"/getSignBySignName", controller.GetSignBySignName)

	group.POST(routePrefix+"/createProvider", controller.CreateProvider)
	group.POST(routePrefix+"/queryProviderByNo", controller.QueryProviderByNo)
	group.POST(routePrefix+"/queryProviderList", controller.QueryProviderList)
	group.POST(routePrefix+"/getProviderById", controller.GetProviderById)

	group.POST(routePrefix+"/createBusinessReq", controller.CreateBusinessReq)
	group.POST(routePrefix+"/getBusinessByAppIdReq", controller.GetBusinessByAppIdReq)

	return group
}
