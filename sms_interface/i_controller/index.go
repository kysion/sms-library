package i_controller

import (
	"github.com/kysion/sms-library/sms_interface"
)

type iModule interface {
	GetModules() sms_interface.IModules
}
