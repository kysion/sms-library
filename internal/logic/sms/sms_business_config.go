package sms

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_do"
	"github.com/kysion/sms-library/sms_model/sms_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信业务管理  （业务和应用的区别是啥，我直接使用app应用对接短信， 先有应用然后又业务吗）

type sBusinessConfig struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsBusinessConfig(modules sms_interface.IModules) sms_interface.ISmsBusinessConfig {
	return &sBusinessConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetBusinessConfigById 根据id获取BusinessConfig
func (s *sBusinessConfig) GetBusinessConfigById(ctx context.Context, id int64) (*sms_model.SmsBusinessConfig, error) {
	if id == 0 {
		return nil, errors.New("{#error_sms_business_config_id_empty}" + s.modules.Dao().SmsBusinessConfig.Table())
	}

	data := sms_entity.SmsBusinessConfig{}

	err := s.dao.SmsBusinessConfig.Ctx(ctx).Where(sms_do.SmsBusinessConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("{#error_sms_business_config_get_by_id_failed}: " + err.Error() + s.dao.SmsBusinessConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsBusinessConfig](data, &sms_model.SmsBusinessConfig{})

	return res, nil
}

// CreateBusinessConfig 创建业务 (上下文, 业务信息)
func (s *sBusinessConfig) CreateBusinessConfig(ctx context.Context, config *sms_model.SmsBusinessConfig) (*sms_model.SmsBusinessConfig, error) {
	// 业务名称查重
	count, _ := s.dao.SmsBusinessConfig.Ctx(ctx).Where(sms_do.SmsBusinessConfig{
		BusinessName: config.BusinessName,
	}).Count()

	if count > 0 {
		return nil, errors.New("{#error_sms_business_config_name_duplicate}" + s.dao.SmsBusinessConfig.Table())
	}

	// 生成id
	appConfig := sms_do.SmsBusinessConfig{}
	gconv.Struct(config, &appConfig)
	appConfig.Id = idgen.NextId()
	appConfig.Status = 1
	// appConfig.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	_, err := s.dao.SmsBusinessConfig.Ctx(ctx).Insert(appConfig)
	if err != nil {
		return nil, errors.New("{#error_sms_business_config_create_failed}" + s.dao.SmsBusinessConfig.Table())
	}

	return s.GetBusinessConfigById(ctx, gconv.Int64(appConfig.Id))
}

// GetBusinessConfigByAppId 根据应用id获取BusinessConfig
func (s *sBusinessConfig) GetBusinessConfigByAppId(ctx context.Context, id int64) (*sms_model.SmsBusinessConfig, error) {
	if id == 0 {
		return nil, errors.New("{#error_sms_business_config_app_id_empty}" + s.modules.Dao().SmsBusinessConfig.Table())
	}

	data := sms_entity.SmsBusinessConfig{}

	err := s.dao.SmsBusinessConfig.Ctx(ctx).Where(sms_do.SmsBusinessConfig{AppId: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("{#error_sms_business_config_get_by_id_failed}: " + err.Error() + s.dao.SmsBusinessConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsBusinessConfig](data, &sms_model.SmsBusinessConfig{})

	return res, nil
}
