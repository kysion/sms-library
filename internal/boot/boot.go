package boot

import (
	"context"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_permission"
)

// InitPermission 初始化权限树
func InitPermission(module sms_interface.IModules, factory base_permission.IPermission) []base_permission.IPermission {

	result := []base_permission.IPermission{
		// 资质
		factory.SetId(5947986066667973).
			SetName(module.T(context.TODO(), "{#SmsName}")).
			SetIdentifier(module.GetConfig().Identifier.Sms).
			SetType(1).
			SetIsShow(1).
			SetItems([]base_permission.IPermission{
				sms_permission.Sms.PermissionType(module).SendSms,
				sms_permission.Sms.PermissionType(module).SendCaptchaBySms,
				sms_permission.Sms.PermissionType(module).ReceiveSms,
				sms_permission.Sms.PermissionType(module).ViewAppConfig,
				sms_permission.Sms.PermissionType(module).CreateAppConfig,
				sms_permission.Sms.PermissionType(module).CreateTemplate,
				sms_permission.Sms.PermissionType(module).AuditTemplate,
				sms_permission.Sms.PermissionType(module).ViewTemplate,
				sms_permission.Sms.PermissionType(module).CreateSign,
				sms_permission.Sms.PermissionType(module).AuditSign,
				sms_permission.Sms.PermissionType(module).ViewSign,
				sms_permission.Sms.PermissionType(module).CreateProvider,
				sms_permission.Sms.PermissionType(module).ViewProvider,
				sms_permission.Sms.PermissionType(module).ViewProviderList,
				sms_permission.Sms.PermissionType(module).ViewSendLog,
			}),
	}
	return result
}
