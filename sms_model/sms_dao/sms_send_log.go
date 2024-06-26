// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sms_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"github.com/kysion/sms-library/sms_model/sms_dao/internal"
)

type SmsSendLogDao = dao_interface.TIDao[internal.SmsSendLogColumns]

func NewSmsSendLog(dao ...dao_interface.IDao) SmsSendLogDao {
	return (SmsSendLogDao)(internal.NewSmsSendLogDao(dao...))
}

var (
	SmsSendLog = NewSmsSendLog()
)
