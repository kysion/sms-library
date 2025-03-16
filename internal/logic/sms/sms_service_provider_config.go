package sms

import (
	"context"
	"errors"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_do"
	"github.com/kysion/sms-library/sms_model/sms_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// 渠道商管理
type sServiceProviderConfig struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsServiceProviderConfig(modules sms_interface.IModules) sms_interface.ISmsServiceProviderConfig {
	return &sServiceProviderConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// CreateProvider 添加渠道商
func (s *sServiceProviderConfig) CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error) {
	model := s.dao.SmsServiceProviderConfig.Ctx(ctx)

	// 插入渠道商配置信息
	data := sms_do.SmsServiceProviderConfig{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.CreatedAt = gtime.Now()
	// 渠道商默认是可用状态
	data.Status = 1
	if info.ExtJson != "" {
		data.ExtJson = info.ExtJson
	} else {
		data.ExtJson = nil
	}

	_, err := model.OmitNilData().Insert(data)

	if err != nil {
		return nil, errors.New("{#error_sms_provider_add_failed}" + s.dao.SmsServiceProviderConfig.Table())
	}

	return s.GetProviderById(ctx, gconv.Int64(data.Id))
}

// GetProviderById 根据ID获取渠道商
func (s *sServiceProviderConfig) GetProviderById(ctx context.Context, id int64) (*sms_model.SmsServiceProviderConfig, error) {
	if id == 0 {
		return nil, errors.New("{#error_sms_provider_id_empty}" + s.dao.SmsServiceProviderConfig.Table())
	}

	data := sms_entity.SmsServiceProviderConfig{}

	err := s.dao.SmsServiceProviderConfig.Ctx(ctx).Where(sms_do.SmsServiceProviderConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("{#error_sms_provider_get_by_id_failed}: " + err.Error() + s.dao.SmsServiceProviderConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsServiceProviderConfig](data, &sms_model.SmsServiceProviderConfig{})

	return res, nil
}

// GetProviderByPriority 根据优先级获取渠道商
func (s *sServiceProviderConfig) GetProviderByPriority(ctx context.Context, priority int) (*sms_model.SmsServiceProviderConfig, error) {
	if priority == 0 {
		return nil, errors.New("{#error_sms_provider_priority_empty}" + s.dao.SmsServiceProviderConfig.Table())
	}

	data := sms_entity.SmsServiceProviderConfig{}

	err := s.dao.SmsServiceProviderConfig.Ctx(ctx).Where(sms_do.SmsServiceProviderConfig{Priority: priority}).Scan(&data)
	if err != nil {
		return nil, errors.New("{#error_sms_provider_get_by_id_failed}: " + err.Error() + s.dao.SmsServiceProviderConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsServiceProviderConfig](data, &sms_model.SmsServiceProviderConfig{})

	return res, nil
}

// QueryProviderByNo 根据No编号获取渠道商
func (s *sServiceProviderConfig) QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*sms_model.ServiceProviderConfigListRes, error) {
	if no == "" {
		return nil, errors.New("{#error_sms_provider_code_empty}" + s.dao.SmsServiceProviderConfig.Table())
	}

	res, err := daoctl.Query[sms_entity.SmsServiceProviderConfig](s.dao.SmsServiceProviderConfig.Ctx(ctx).Where(
		sms_do.SmsServiceProviderConfig{ProviderNo: no}),
		params,
		false)
	if err != nil {
		return nil, errors.New("{#error_sms_provider_get_by_code_failed}: " + err.Error() + s.dao.SmsServiceProviderConfig.Table())
	}

	ret := kconv.Struct[*sms_model.ServiceProviderConfigListRes](res, &sms_model.ServiceProviderConfigListRes{})

	return ret, nil
}

// QueryProviderList 获取渠道商列表
func (s *sServiceProviderConfig) QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error) {
	result, err := daoctl.Query[*sms_model.SmsServiceProviderConfig](s.dao.SmsServiceProviderConfig.Ctx(ctx), search, isExport)

	return (*sms_model.ServiceProviderConfigListRes)(result), err
}
