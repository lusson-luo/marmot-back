// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Inspection is the golang structure for table inspection.
type Inspection struct {
	Id           int         `json:"id"           ` //
	Name         string      `json:"name"         ` //
	Count        int         `json:"count"        ` //
	SuccessCount int         `json:"successCount" ` //
	FailedCount  int         `json:"failedCount"  ` //
	Connection   bool        `json:"connection"   ` //
	Availability bool        `json:"availability" ` //
	StartTime    *gtime.Time `json:"startTime"    ` //
	EndTime      *gtime.Time `json:"endTime"      ` //
}
