// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Inspection is the golang structure of table inspection for DAO operations like Where/Data.
type Inspection struct {
	g.Meta       `orm:"table:inspection, do:true"`
	Id           interface{} //
	Name         interface{} //
	Count        interface{} //
	SuccessCount interface{} //
	FailedCount  interface{} //
	Connection   interface{} //
	Availability interface{} //
	StartTime    *gtime.Time //
	EndTime      *gtime.Time //
}
