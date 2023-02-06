package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DetailRes struct {
	g.Meta        `mime:"application/json" example:"string"`
	Id            int         `json:"id"`
	InspectTaskId int         `json:"inspectTaskId"`
	Name          string      `json:"name"`
	ExecStatus    bool        `json:"execStatus"`
	ErrMsg        string      `json:"errMsg"`
	StartTime     *gtime.Time `json:"startTime"`
	EndTime       *gtime.Time `json:"endTime"`
	InspectionId  int         `json:"inspectionId" `
}

type DetailReq struct {
	g.Meta       `path:"/api/inspection/detail" tags:"detail" method:"get" summary:"巡检详情"`
	InspectionId int `json:"inspectionId"`
}
