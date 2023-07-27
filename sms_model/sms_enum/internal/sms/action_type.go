package sms

import "github.com/kysion/base-library/utility/enum"

type ActionEnum enum.IEnumCode[int]

type action struct {
	Send    ActionEnum
	Refund  ActionEnum
	Renewal ActionEnum
	TopUp   ActionEnum
}

var Action = action{
	Send:    enum.New[ActionEnum](1, "发送"),
	Refund:  enum.New[ActionEnum](2, "退订"),
	Renewal: enum.New[ActionEnum](4, "续费"),
	TopUp:   enum.New[ActionEnum](8, "充值"),
}

func (e action) New(code int, description string) ActionEnum {
	if (code&Action.Send.Code()) == Action.Send.Code() ||
		(code&Action.Refund.Code()) == Action.Refund.Code() ||
		(code&Action.Renewal.Code()) == Action.Renewal.Code() ||
		(code&Action.TopUp.Code()) == Action.TopUp.Code() {
		return enum.New[ActionEnum](code, description)
	} else {
		panic("smsState.State.New: error")
	}
}
