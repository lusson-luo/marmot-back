package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type InspectListReq struct {
	g.Meta `path:"/api/inspection/list" tags:"inspect" method:"get" summary:"巡检结果列表"`
}

type InspectListRes struct {
	g.Meta       `mime:"application/json" example:"string"`
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Count        int       `json:"count"`
	SuccessCount int       `json:"successCount"`
	FailedCount  int       `json:"failedCount"`
	Connection   bool      `json:"connection"`
	Availability bool      `json:"availability"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}

type InspectReq struct {
	g.Meta `path:"/api/inspection/inspect" tags:"inspect" method:"get" summary:"巡检结果列表"`
	Id     int `json:"id"`
}

type InspectRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type InspectAllReq struct {
	g.Meta `path:"/api/inspection/inspectAll" tags:"inspect" method:"get" summary:"巡检结果列表"`
	Id     int `json:"id"`
}

type InspectAllRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
