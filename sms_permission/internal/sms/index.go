package sms

import (
	"github.com/kysion/sms-library/sms_interface"
)

type sms struct {
	PermissionType func(modules sms_interface.IModules) *permissionType[sms_interface.IModules]
}

var Sms = sms{
	PermissionType: PermissionType,
	//PermissionType: sms2.PermissionType,	//PermissionType: sms2.PermissionType,

}
