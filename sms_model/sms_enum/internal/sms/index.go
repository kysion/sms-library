package sms

type sms struct {
	State  state
	Type   smsType
	Action action
	//PermissionType func(modules sms_interface.IModules) *sms_permission.permissionType[sms_interface.IModules]
}

var Sms = sms{
	State:  State,
	Type:   Type,
	Action: Action,
	//PermissionType: sms2.PermissionType,
}
