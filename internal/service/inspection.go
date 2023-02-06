package service

import (
	"context"
	v1 "marmot/api/v1"
)

var localInspection IInspection

func Inspection() IInspection {
	if localInspection == nil {
		panic("implement not found for interface IInspection, forgot register?")
	}
	return localInspection
}

type IInspection interface {
	// 1. 查看巡检列表
	List(ctx context.Context) (*[]v1.InspectListRes, error)
	// 2. 巡检指定场景
	InspectSelection(ctx context.Context, ids []int)
}

func RegisterInspection(i IInspection) {
	localInspection = i
}
