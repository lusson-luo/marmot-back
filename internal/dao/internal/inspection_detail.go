// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InspectionDetailDao is the data access object for table inspection_detail.
type InspectionDetailDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns InspectionDetailColumns // columns contains all the column names of Table for convenient usage.
}

// InspectionDetailColumns defines and stores column names for table inspection_detail.
type InspectionDetailColumns struct {
	Id           string //
	Name         string //
	Success      string //
	ErrMsg       string //
	StartTime    string //
	EndTime      string //
	InspectionId string //
}

// inspectionDetailColumns holds the columns for table inspection_detail.
var inspectionDetailColumns = InspectionDetailColumns{
	Id:           "id",
	Name:         "name",
	Success:      "success",
	ErrMsg:       "err_msg",
	StartTime:    "start_time",
	EndTime:      "end_time",
	InspectionId: "inspection_id",
}

// NewInspectionDetailDao creates and returns a new DAO object for table data access.
func NewInspectionDetailDao() *InspectionDetailDao {
	return &InspectionDetailDao{
		group:   "default",
		table:   "inspection_detail",
		columns: inspectionDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InspectionDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InspectionDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InspectionDetailDao) Columns() InspectionDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InspectionDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InspectionDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InspectionDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
