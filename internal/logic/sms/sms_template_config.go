package sms

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_do"
	"github.com/kysion/sms-library/sms_model/sms_entity"
	"github.com/kysion/sms-library/sms_model/sms_enum"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信模板管理
type sTemplateConfig struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsTemplateConfig(modules sms_interface.IModules) sms_interface.ISmsTemplateConfig {
	return &sTemplateConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// CreateTemplate 添加短信模版
func (s *sTemplateConfig) CreateTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error) {
	// 添加短信模版
	data := kconv.Struct(info, &sms_do.SmsTemplateConfig{})

	data.Id = idgen.NextId()
	// 未审核的短信模版是禁用状态
	data.Status = 0
	// data.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	if info.ExtJson != "" {
		data.ExtJson = info.ExtJson
	} else {
		data.ExtJson = nil
	}

	_, err := s.dao.SmsTemplateConfig.Ctx(ctx).OmitNilData().Insert(data)

	if err != nil {
		return nil, errors.New("短信模板添加失败" + s.dao.SmsTemplateConfig.Table())
	}

	return s.GetTemplateById(ctx, gconv.Int64(data.Id))
}

// AuditTemplate 短信模版审核
func (s *sTemplateConfig) AuditTemplate(ctx context.Context, id int64, info *sms_model.AuditInfo) (bool, error) {
	// 判断审核行为，只能是审核通过或者不通过 -1不通过 1通过
	if info.State != sms_enum.Sms.State.Reject.Code() && info.State != sms_enum.Sms.State.Approve.Code() {
		return false, errors.New("审核行为类型错误" + s.dao.SmsTemplateConfig.Table())
	}

	// 审核不通过需要有原因
	if info.State == sms_enum.Sms.State.Reject.Code() && info.ReplyMsg == "" {
		return false, errors.New("审核不通过时必须说明原因" + s.dao.SmsTemplateConfig.Table())
	}

	// 判断模版是否存在
	template, err := daoctl.GetByIdWithError[sms_entity.SmsTemplateConfig](s.dao.SmsTemplateConfig.Ctx(ctx), id)
	if err != nil || template == nil {
		return false, errors.New("短信签名不存在" + s.dao.SmsSignConfig.Table())
	}

	// 改变状态为正常代表审核成功
	_, err = s.dao.SmsTemplateConfig.Ctx(ctx).OmitNilData().Data(sms_do.SmsTemplateConfig{
		AuditUserId:   info.AuditUserId,
		AuditReplyMsg: info.ReplyMsg,
		AuditAt:       gtime.Now(),
		Status:        info.State,
	}).Where(sms_do.SmsTemplateConfig{
		Id: id,
	}).Update()

	if err != nil {
		return false, errors.New("短信签名审核失败：" + err.Error() + s.dao.SmsTemplateConfig.Table())
	}

	return true, nil
}

// GetByTemplateCode 根据模版编号查询模版信息
func (s *sTemplateConfig) GetByTemplateCode(ctx context.Context, templateCode string) (*sms_model.SmsTemplateConfig, error) {
	if templateCode == "" {
		return nil, errors.New("模版编号不能为空" + s.dao.SmsTemplateConfig.Table())
	}

	data := sms_entity.SmsTemplateConfig{}

	err := s.dao.SmsTemplateConfig.Ctx(ctx).Where(sms_do.SmsTemplateConfig{TemplateCode: templateCode}).Scan(&data)
	if err != nil {
		return nil, errors.New("根据模版编号获取模版信息失败：" + err.Error() + s.dao.SmsTemplateConfig.Table())
	}

	result := kconv.Struct[*sms_model.SmsTemplateConfig](data, &sms_model.SmsTemplateConfig{})

	return result, nil
}

// GetTemplateById 根据ID获取模版
func (s *sTemplateConfig) GetTemplateById(ctx context.Context, id int64) (*sms_model.SmsTemplateConfig, error) {
	if id == 0 {
		return nil, errors.New("模板id不能为空" + s.dao.SmsTemplateConfig.Table())
	}

	data := sms_entity.SmsTemplateConfig{}

	err := s.dao.SmsTemplateConfig.Ctx(ctx).Where(sms_do.SmsTemplateConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("根据id获取渠道商信息失败：" + err.Error() + s.dao.SmsTemplateConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsTemplateConfig](data, &sms_model.SmsTemplateConfig{})

	return res, nil
}

// GetByProviderNoAndType 根据业务场景及渠道获取模版
func (s *sTemplateConfig) GetByProviderNoAndType(ctx context.Context, providerNo sms_enum.SmsType, smsType int) (*sms_model.SmsTemplateConfig, error) {
	if providerNo.Code() == "" {
		return nil, errors.New("渠道编号不能为空" + s.dao.SmsTemplateConfig.Table())
	}

	if smsType == 0 {
		return nil, errors.New("请指定业务场景" + s.dao.SmsTemplateConfig.Table())
	}

	data := sms_entity.SmsTemplateConfig{}
	//data, err := daoctl.ScanWithError[sms_model.SmsTemplateConfig](s.dao.SmsTemplateConfig.Ctx(ctx).Where(sms_do.SmsTemplateConfig{ProviderNo: providerNo.Code(), Type: smsType}))

	//err := s.dao.SmsTemplateConfig.Ctx(ctx).Where(sms_do.SmsTemplateConfig{ProviderNo: providerNo.Code()}).Raw("(`type` & %d) = %d", smsType, smsType).Scan(&data)

	model := s.dao.SmsTemplateConfig.Ctx(ctx)
	model = model.Where(sms_do.SmsTemplateConfig{ProviderNo: providerNo.Code()}).Wheref("type & ? = ?", smsType, smsType)
	err := model.Scan(&data)

	if err != nil {
		return nil, errors.New("获取渠道商信息失败：" + err.Error() + s.dao.SmsTemplateConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsTemplateConfig](data, &sms_model.SmsTemplateConfig{})

	return res, nil
}
