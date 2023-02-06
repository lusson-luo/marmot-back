package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/service"
)

type CInspectDetail struct {
}

var InspectDetail CInspectDetail = CInspectDetail{}

// 查询详情
func (c CInspectDetail) GetInspectDetail(ctx context.Context, req *v1.DetailReq) (res *[]v1.DetailRes, err error) {
	res, err = service.InspectDetail().GetInspectDetail(ctx, req.InspectionId)
	return
}
