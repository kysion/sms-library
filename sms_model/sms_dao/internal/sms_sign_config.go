// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsSignConfigDao is the data access object for table sms_sign_config.
type SmsSignConfigDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SmsSignConfigColumns // columns contains all the column names of Table for convenient usage.
}

// SmsSignConfigColumns defines and stores column names for table sms_sign_config.
type SmsSignConfigColumns struct {
	Id            string // ID
	SignName      string // 短信签名名称
	ProviderNo    string // 渠道商编号
	ProviderName  string // 渠道商名字
	Remark        string // 备注
	Status        string // 状态: -1不通过 0待审核 1正常
	AuditUserId   string // 审核者UserID
	AuditReplyMsg string // 审核回复，仅审核不通过时才有值
	AuditAt       string // 审核时间
	ExtJson       string // 拓展字段
	UnionMainId   string // 关联主体ID
	CreatedAt     string //
	UpdatedAt     string //
	DeletedAt     string //
}

// smsSignConfigColumns holds the columns for table sms_sign_config.
var smsSignConfigColumns = SmsSignConfigColumns{
	Id:            "id",
	SignName:      "sign_name",
	ProviderNo:    "provider_no",
	ProviderName:  "provider_name",
	Remark:        "remark",
	Status:        "status",
	AuditUserId:   "audit_user_id",
	AuditReplyMsg: "audit_reply_msg",
	AuditAt:       "audit_at",
	ExtJson:       "ext_json",
	UnionMainId:   "union_main_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewSmsSignConfigDao creates and returns a new DAO object for table data access.
func NewSmsSignConfigDao(proxy ...dao_interface.IDao) *SmsSignConfigDao {
	var dao *SmsSignConfigDao
	if len(proxy) > 0 {
		dao = &SmsSignConfigDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: smsSignConfigColumns,
		}
		return dao
	}

	return &SmsSignConfigDao{
		group:   "default",
		table:   "sms_sign_config",
		columns: smsSignConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SmsSignConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SmsSignConfigDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SmsSignConfigDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SmsSignConfigDao) Columns() SmsSignConfigColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SmsSignConfigDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SmsSignConfigDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SmsSignConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
