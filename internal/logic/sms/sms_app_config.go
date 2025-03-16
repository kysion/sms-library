package sms

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_do"
	"github.com/kysion/sms-library/sms_model/sms_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信应用配置管理
type sAppConfig struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsAppConfig(modules sms_interface.IModules) sms_interface.ISmsAppConfig {
	return &sAppConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetAppConfigByName 根据应用名称获取AppConfig
func (s *sAppConfig) GetAppConfigByName(ctx context.Context, appName string) (*sms_model.SmsAppConfig, error) {
	if appName == "" {
		table := s.modules.Dao().SmsAppConfig.Table()
		return nil, errors.New("{#error_sms_app_config_name_empty}" + table)

	}

	data := sms_entity.SmsAppConfig{}

	err := s.modules.Dao().SmsAppConfig.Ctx(ctx).Where(sms_do.SmsAppConfig{AppName: appName}).Scan(&data)
	if err != nil {
		return nil, errors.New("{#error_sms_app_config_get_by_name_failed}" + s.dao.SmsAppConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsAppConfig](data, &sms_model.SmsAppConfig{})

	return res, nil
}

// GetAppConfigById 根据id获取AppConfig
func (s *sAppConfig) GetAppConfigById(ctx context.Context, id int64) (*sms_model.SmsAppConfig, error) {
	if id == 0 {
		return nil, errors.New("{#error_sms_app_config_id_empty}" + s.dao.SmsAppConfig.Table())

	}

	data := sms_entity.SmsAppConfig{}

	err := s.dao.SmsAppConfig.Ctx(ctx).Where(sms_do.SmsAppConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("{#error_sms_app_config_get_by_id_failed}" + s.dao.SmsAppConfig.Table())
	}

	res := kconv.Struct[*sms_model.SmsAppConfig](data, &sms_model.SmsAppConfig{})

	return res, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用id) (当前应用剩余短信数量)
func (s *sAppConfig) GetAppAvailableNumber(ctx context.Context, id int64) (int, error) {
	if id == 0 {
		return 0, errors.New("{#error_sms_app_config_id_empty}" + s.dao.SmsAppConfig.Table())
	}

	data := sms_entity.SmsAppConfig{}

	err := s.dao.SmsAppConfig.Ctx(ctx).Where(sms_do.SmsAppConfig{Id: id}).Scan(&data)
	if err != nil {
		return 0, errors.New("{#error_sms_app_config_get_by_id_failed}" + s.dao.SmsAppConfig.Table())
	}

	return data.AvailableNumber, nil
}

// CreateAppConfig 创建应用 (上下文, 应用编号, 花费数量)
func (s *sAppConfig) CreateAppConfig(ctx context.Context, config *sms_model.SmsAppConfig) (bool, error) {
	// 应用名称查重
	count, _ := s.dao.SmsAppConfig.Ctx(ctx).Where(sms_do.SmsAppConfig{
		AppName: config.AppName,
	}).Count()

	if count > 0 {
		return false, errors.New("{#error_sms_app_config_name_duplicate}" + s.dao.SmsAppConfig.Table())
	}

	// 生成id
	appConfig := sms_do.SmsAppConfig{}
	gconv.Struct(config, &appConfig)
	appConfig.Id = idgen.NextId()
	appConfig.Status = 1
	// appConfig.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	_, err := s.dao.SmsAppConfig.Ctx(ctx).Insert(appConfig)
	if err != nil {
		return false, errors.New("{#error_sms_app_config_create_failed}" + s.dao.SmsAppConfig.Table())
	}

	return true, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (s *sAppConfig) UpdateAppNumber(ctx context.Context, id int64, fee uint64) (bool, error) {
	if id == 0 {
		return false, errors.New("{#error_sms_app_config_id_empty}" + s.dao.SmsAppConfig.Table())
	}

	// 获取原来的数量
	appConfig, err := s.GetAppConfigById(ctx, id)
	if err != nil {
		return false, err
	}

	newUseNum := appConfig.UseNumber + gconv.Int64(fee)
	newAvailableNum := appConfig.AvailableNumber - gconv.Int64(fee)

	affected, err := daoctl.UpdateWithError(s.dao.SmsAppConfig.Ctx(ctx).
		Data(sms_do.SmsAppConfig{UseNumber: newUseNum, AvailableNumber: newAvailableNum}).
		Where(sms_do.SmsAppConfig{Id: id}))

	return affected > 0, nil
}
