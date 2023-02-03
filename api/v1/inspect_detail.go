package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SingleTaskDetail struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Success   bool        `json:"success"`
	Err       string      `json:"err"`
	StartTime *gtime.Time `json:"startTime"`
	EndTime   *gtime.Time `json:"endTime"`
}

type Detail struct {
	g.Meta `mime:"application/json" example:"string"`
	detail []SingleTaskDetail `json:"all tasks"`
}

type DetailReq struct {
	g.Meta `path:"/api/inspection/detail" tags:"detail" method:"get" summary:"巡检详情"`
	Id     int `json:"id"`
}
