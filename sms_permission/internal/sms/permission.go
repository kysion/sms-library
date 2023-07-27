package sms

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kmap"
	"github.com/kysion/sms-library/sms_interface"
)

type PermissionEnum = base_permission.IPermission

type permissionType[T sms_interface.IModules] struct {
	modules T
	enumMap *kmap.HashMap[string, PermissionEnum]

	SendSms          PermissionEnum
	SendCaptchaBySms PermissionEnum
	ReceiveSms       PermissionEnum

	ViewAppConfig   PermissionEnum
	CreateAppConfig PermissionEnum

	CreateTemplate PermissionEnum
	AuditTemplate  PermissionEnum
	ViewTemplate   PermissionEnum

	CreateSign PermissionEnum
	AuditSign  PermissionEnum
	ViewSign   PermissionEnum

	CreateProvider   PermissionEnum
	ViewProvider     PermissionEnum
	ViewProviderList PermissionEnum

	ViewSendLog PermissionEnum
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[sms_interface.IModules]]()
	PermissionType    = func(modules sms_interface.IModules) *permissionType[sms_interface.IModules] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[sms_interface.IModules]{
			modules:          modules,
			enumMap:          kmap.New[string, PermissionEnum](),
			SendSms:          base_permission.NewInIdentifier("SendSms", "发送短信", ""),
			SendCaptchaBySms: base_permission.NewInIdentifier("SendCaptchaBySms", "发送验证码短信", ""),
			ReceiveSms:       base_permission.NewInIdentifier("ReceiveSms", "接收短信", ""),
			ViewAppConfig:    base_permission.NewInIdentifier("ViewAppConfig", "查看应用配置信息", ""),
			CreateAppConfig:  base_permission.NewInIdentifier("CreateAppConfig", "创建应用配置", ""),
			CreateTemplate:   base_permission.NewInIdentifier("CreateTemplate", "创建模板", ""),
			AuditTemplate:    base_permission.NewInIdentifier("AuditTemplate", "审核模板", ""),
			ViewTemplate:     base_permission.NewInIdentifier("ViewTemplate", "查看模板信息", ""),
			CreateSign:       base_permission.NewInIdentifier("CreateSign", "创建签名", ""),
			AuditSign:        base_permission.NewInIdentifier("AuditSign", "审核签名", ""),
			ViewSign:         base_permission.NewInIdentifier("ViewSign", "查看签名", ""),
			CreateProvider:   base_permission.NewInIdentifier("CreateProvider", "创建业务配置", ""),
			ViewProvider:     base_permission.NewInIdentifier("ViewProvider", "查看业务配置信息", ""),
			ViewProviderList: base_permission.NewInIdentifier("ViewProviderList", "查看业务配置列表", ""),
			ViewSendLog:      base_permission.NewInIdentifier("ViewSendLog", "查看短信发送日志", ""),
		})

		for k, v := range gconv.Map(result) {
			result.enumMap.Set(k, v.(PermissionEnum))
		}
		return result
	}
)

// ByCode 通过枚举值取枚举类型
func (e *permissionType[T]) ByCode(identifier string) base_permission.IPermission {
	v, has := e.enumMap.Search(identifier)
	if v != nil && has {
		return v
	}
	return nil
}
