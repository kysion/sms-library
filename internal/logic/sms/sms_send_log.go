package sms

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_do"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信服务发送日志
type sSendLogConfig struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsSendLogConfig(modules sms_interface.IModules) sms_interface.ISmsSendLogConfig {
	return &sSendLogConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// SaveSmsLog 记录日志
func (s *sSendLogConfig) SaveSmsLog(ctx context.Context, info *sms_model.SmsSendLog) (bool, error) {
	if info.PhoneNumber == "" {
		return false, errors.New("发送手机号不能为空" + s.dao.SmsSendLog.Table())
	}

	if info.SignName == "" {
		return false, errors.New("签名名称不能为空" + s.dao.SmsSendLog.Table())
	}

	if info.Type == 0 {
		return false, errors.New("短信类型不能为空" + s.dao.SmsSendLog.Table())
	}

	data := sms_do.SmsSendLog{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()

	_, err := s.dao.SmsSendLog.Ctx(ctx).Insert(data)
	if err != nil {
		return false, errors.New("短信日志写入失败：" + err.Error() + s.dao.SmsSendLog.Table())
	}

	return true, nil
}

// GetSmsLogById 根据id查询日志
func (s *sSendLogConfig) GetSmsLogById(ctx context.Context, id int64) (res *sms_model.SmsSendLog, err error) {
	err = s.dao.SmsSendLog.Ctx(ctx).Where(sms_do.SmsSendLog{
		Id: id,
	}).Scan(&res)

	if err != nil {
		return nil, errors.New("根据id查询日志失败：" + err.Error() + s.dao.SmsSendLog.Table())
	}

	return res, nil
}
