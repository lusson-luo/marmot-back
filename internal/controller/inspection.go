package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/service"
)

var Inspection CInspection = CInspection{}

type CInspection struct {
}

// 1. 查看巡检列表
func (c CInspection) List(ctx context.Context, req *v1.InspectListReq) (res *[]v1.InspectListRes, err error) {
	res, err = service.Inspection().List(ctx)
	return
}

// 2. 巡检选定场景
// 如果没有传场景id，巡检所有场景
func (c CInspection) InspectSelection(ctx context.Context, req *v1.InspectSelectionReq) (res *v1.InspectSelectionRes, err error) {
	service.Inspection().InspectSelection(ctx, req.Ids)
	return
}
