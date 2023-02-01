package service

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/logic"
)

var Inspection SInspection = new()

type SInspection interface {
	// 1. 查看巡检列表
	List(ctx context.Context) (*[]v1.InspectListRes, error)
	// 2. 巡检单项场景
	Inspect(ctx context.Context, id int)
	// 3. 巡检全部场景
	InspectAll(ctx context.Context)
}

func new() SInspection {
	return logic.InspectionLogic{}
}