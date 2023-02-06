package service

import (
	"context"
	v1 "marmot/api/v1"
)

var localInspectDetail IInspectDetail

func InspectDetail() IInspectDetail {
	if localInspectDetail == nil {
		panic("implement not found for interface IInspectDetail, forgot register?")
	}
	return localInspectDetail
}

type IInspectDetail interface {
	// 查询详情
	GetInspectDetail(ctx context.Context, inspectId int) (*[]v1.DetailRes, error)
}

func RegisterInspectDetail(i IInspectDetail) {
	localInspectDetail = i
}
