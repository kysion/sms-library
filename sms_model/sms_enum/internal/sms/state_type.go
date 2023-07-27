package sms

import "github.com/kysion/base-library/utility/enum"

type StateEnum enum.IEnumCode[int]

type state struct {
	Reject     StateEnum
	WaitReview StateEnum
	Approve    StateEnum
}

var State = state{
	Reject:     enum.New[StateEnum](-1, "不通过"),
	WaitReview: enum.New[StateEnum](0, "待审核"),
	Approve:    enum.New[StateEnum](1, "通过"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Reject.Code()) == State.Reject.Code() ||
		(code&State.WaitReview.Code()) == State.WaitReview.Code() ||
		(code&State.Approve.Code()) == State.Approve.Code() {
		return enum.New[StateEnum](code, description)
	} else {
		panic("smsState.State.New: error")
	}
}
