// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sms_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"github.com/kysion/sms-library/sms_model/sms_dao/internal"
)

type SmsTemplateConfig = dao_interface.TIDao[internal.SmsTemplateConfigColumns]

func NewSmsTemplateConfig(dao ...dao_interface.IDao) SmsTemplateConfig {
	return (SmsTemplateConfig)(internal.NewSmsTemplateConfigDao(dao...))
}
