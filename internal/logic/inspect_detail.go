package logic

import (
	"context"
	"fmt"
	v1 "marmot/api/v1"
	"marmot/internal/dao"
	"marmot/internal/model/entity"
	"marmot/internal/service"
)

type InspectDetailLogic struct {
}

func (InspectDetailLogic) GetInspectDetail(ctx context.Context, inspectId int) (res *[]v1.DetailRes, err error) {
	res = &[]v1.DetailRes{}
	var inspectionDetails []entity.InspectionDetail
	inspectTaskId := GetCurrentInspectTaskId(ctx, inspectId)
	fmt.Println("inspectTaskId: ", inspectTaskId)
	dao.InspectionDetail.Ctx(ctx).Where("inspection_id", inspectId).Where("inspect_task_id", inspectTaskId).Scan(&inspectionDetails)
	for _, v := range inspectionDetails {
		detail := v1.DetailRes{
			Id:            v.Id,
			InspectTaskId: v.InspectTaskId,
			Name:          v.Name,
			ExecStatus:    v.ExecStatus,
			ErrMsg:        v.ErrMsg,
			StartTime:     v.StartTime,
			EndTime:       v.EndTime,
			InspectionId:  v.InspectionId,
		}

		tmp := append(*res, detail)
		res = &tmp
	}
	err = nil
	return
}

func init() {
	service.RegisterInspectDetail(InspectDetailLogic{})
}
