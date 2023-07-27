package sms

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_do"
	"github.com/kysion/sms-library/sms_model/sms_enum"

	"github.com/kysion/sms-library/sms_model/sms_dao"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信签名管理

type sSignConfig struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsSignConfig(modules sms_interface.IModules) sms_interface.ISmsSignConfig {
	return &sSignConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// CreateSign 添加短信签名
func (s *sSignConfig) CreateSign(ctx context.Context, info *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error) {
	// 根据渠道商编号和短信签名判断渠道商信息是否已存在
	count, err := s.dao.SmsSignConfig.Ctx(ctx).Where(sms_do.SmsSignConfig{
		ProviderNo: info.ProviderNo,
		SignName:   info.SignName,
	}).Count()

	if err != nil || count > 0 {
		return nil, errors.New("该签名在此渠道商已经存在" + s.dao.SmsSignConfig.Table())
	}

	// 添加短信签名
	data := kconv.Struct(info, &sms_do.SmsSignConfig{})

	data.Id = idgen.NextId()
	// 未审核的短信签名是禁用状态
	data.Status = 0
	//data.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	if info.ExtJson != "" {
		data.ExtJson = info.ExtJson
	} else {
		data.ExtJson = nil
	}

	_, err = s.dao.SmsSignConfig.Ctx(ctx).OmitNilData().Insert(data)

	if err != nil {
		return nil, errors.New("短信签名添加失败" + s.dao.SmsSignConfig.Table())
	}

	return s.GetSignById(ctx, gconv.Int64(data.Id))
}

// AuditSign 审核短信签名, 将短信签名Status状态改变为1
func (s *sSignConfig) AuditSign(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error) {
	// 判断审核行为，只能是审核通过或者不通过 -1不通过 1通过
	if info.State != sms_enum.Sms.State.Reject.Code() && info.State != sms_enum.Sms.State.Approve.Code() {
		return false, errors.New("审核行为类型错误" + s.dao.SmsSignConfig.Table())
	}

	// 审核不通过需要有原因

	if info.State == sms_enum.Sms.State.Reject.Code() && info.ReplyMsg == "" {
		return false, errors.New("审核不通过时必须说明原因" + s.dao.SmsSignConfig.Table())
	}

	// 判断签名是否存在
	sign, err := s.GetSignById(ctx, id)
	if err != nil || sign == nil {
		return false, errors.New("短信签名不存在" + s.dao.SmsSignConfig.Table())
	}

	// 改变状态为正常代表审核成功
	_, err = s.dao.SmsSignConfig.Ctx(ctx).OmitNilData().Data(sms_do.SmsSignConfig{
		AuditUserId:   info.AuditUserId,
		AuditReplyMsg: info.ReplyMsg,
		AuditAt:       gtime.Now(),
		Status:        info.State,
	}).Where(sms_do.SmsSignConfig{
		Id: id,
	}).Update()

	if err != nil {
		return false, errors.New("短信签名审核失败" + s.dao.SmsSignConfig.Table())
	}

	return true, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func (s *sSignConfig) GetSignBySignName(ctx context.Context, signName string) (res *sms_model.SmsSignConfig, err error) {
	err = s.dao.SmsSignConfig.Ctx(ctx).Where(sms_do.SmsSignConfig{
		SignName: signName,
	}).Scan(&res)

	if err != nil {
		return nil, errors.New("该签名不存在" + s.dao.SmsSignConfig.Table())
	}
	return res, nil
}

// GetSignById 根据id查找签名数据
func (s *sSignConfig) GetSignById(ctx context.Context, id int64) (*sms_model.SmsSignConfig, error) {
	if id == 0 {
		return nil, errors.New("签名id不能为空" + s.dao.SmsSignConfig.Table())
	}

	result, err := daoctl.GetByIdWithError[*sms_model.SmsSignConfig](s.dao.SmsSignConfig.Ctx(ctx), id)

	return *result, err
}
