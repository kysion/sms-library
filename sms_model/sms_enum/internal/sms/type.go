package sms

import "github.com/kysion/base-library/utility/enum"

type SmsTypeEnum enum.IEnumCode[string]

type smsType struct {
	Aliyun  SmsTypeEnum
	Huawei  SmsTypeEnum
	Tencent SmsTypeEnum
	Qiniu   SmsTypeEnum
	Qyxs    SmsTypeEnum // 企业信使
}

var Type = smsType{
	Aliyun:  enum.New[SmsTypeEnum]("aliyun", "阿里云"),
	Huawei:  enum.New[SmsTypeEnum]("huawei", "华为云"),
	Tencent: enum.New[SmsTypeEnum]("tencent", "腾讯云"),
	Qiniu:   enum.New[SmsTypeEnum]("qiniu", "七牛云"),
	Qyxs:    enum.New[SmsTypeEnum]("qyxs", "企业信使"),
}

func (e smsType) New(code string) SmsTypeEnum {
	if code == Type.Aliyun.Code() {
		return e.Aliyun
	}
	if code == Type.Huawei.Code() {
		return e.Huawei
	}
	if code == Type.Tencent.Code() {
		return e.Tencent
	}
	if code == Type.Qiniu.Code() {
		return e.Qiniu
	}
	if code == Type.Qyxs.Code() {
		return e.Qyxs
	}

	panic("Sms.Type.New: error")
}
