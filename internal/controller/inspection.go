package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/service"
)

var Inspection CInspection = CInspection{}

type CInspection struct {
}

// /api/inspection/list
//       {
//         id: '1',
//         name: 'MariaDB',
//         count: 23000,
//         successCount: 22000,
//         failedCount: 1000,
//         connection: true,
//         availability: true,
//         startTime: '2022-12-28 17:51',
//         endTime: '2022-12-28 17:53',
//       }

// 1. 查看巡检列表
func (c CInspection) List(ctx context.Context, req *v1.InspectListReq) (res *[]v1.InspectListRes, err error) {
	res, err = service.Inspection().List(ctx)
	return
}

// 2. 巡检单项场景
func (c CInspection) Inspect(ctx context.Context, req *v1.InspectReq) (res *v1.InspectRes, err error) {
	service.Inspection().Inspect(ctx, req.Id)
	return
}

// 3. 巡检全部场景
func (c CInspection) InspectAll(ctx context.Context, req *v1.InspectAllReq) (res *v1.InspectAllRes, err error) {
	service.Inspection().InspectAll(ctx)
	return
}
