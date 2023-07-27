package sms_controller

import (
	"github.com/kysion/sms-library/sms_controller/internal"
	"github.com/kysion/sms-library/sms_interface/i_controller"
)

type (
	SmsController internal.SmsController
)

type CoController struct {
	Sms i_controller.ISms
}

var (
	Sms = internal.Sms
)
