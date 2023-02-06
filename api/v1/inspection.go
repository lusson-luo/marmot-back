package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InspectListReq struct {
	g.Meta `path:"/api/inspection/list" tags:"inspect" method:"get" summary:"巡检结果列表"`
}

type InspectListRes struct {
	g.Meta         `mime:"application/json" example:"string"`
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	Count          int         `json:"count"`
	SuccessCount   int         `json:"successCount"`
	FailedCount    int         `json:"failedCount"`
	Connection     bool        `json:"connection"`
	Availability   bool        `json:"availability"`
	InspectLoading bool        `json:"inspectLoading"`
	StartTime      *gtime.Time `json:"startTime"`
	EndTime        *gtime.Time `json:"endTime"`
}

type InspectSelectionReq struct {
	g.Meta `path:"/api/inspection/inspectSelection" tags:"inspectSelection" method:"get" summary:"指定巡检"`
	Ids    []int `json:"ids"`
}

type InspectSelectionRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
