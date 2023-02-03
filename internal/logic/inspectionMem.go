package logic

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/dao"
	"marmot/internal/model/entity"

	"github.com/google/uuid"
)

// 数据存内存
type InspectionMemLogic struct {
}

// 1. 查看巡检列表
func (InspectionMemLogic) List(ctx context.Context) (res *[]v1.InspectListRes, err error) {
	res = &[]v1.InspectListRes{}
	var inspections []*entity.Inspection = dao.Inspection.FindAll()
	for _, v := range inspections {
		inspect := v1.InspectListRes{
			Id:           v.Id,
			Name:         v.Name,
			Count:        v.Count,
			SuccessCount: v.SuccessCount,
			FailedCount:  v.FailedCount,
			Connection:   v.Connection,
			Availability: v.Availability,
			StartTime:    v.StartTime,
			EndTime:      v.EndTime,
		}
		tmp := append(*res, inspect)
		res = &tmp
	}
	err = nil
	return
}

// 2. 巡检单项场景
func (InspectionMemLogic) Inspect(ctx context.Context, id int) {
	inspection, exist := dao.Inspection.FindById(id)
	if exist {
		switch inspection.Name {
		case "mysql":
			MysqlMemInspector.inspect(ctx, id)
		}
	}
}

// 3. 巡检全部场景
func (logic InspectionMemLogic) InspectAll(ctx context.Context) {
	inspections := dao.Inspection.FindAll()
	for _, inspection := range inspections {
		logic.Inspect(ctx, inspection.Id)
	}
}

// 4. 加载所有巡检处理器
func registerInspectorsMem() {
	inspections := dao.Inspection.FindAll()
	mysqlExist := false
	for _, inspection := range inspections {
		switch inspection.Name {
		case "mysql":
			mysqlExist = true
		}
	}
	if !mysqlExist {
		uuid := uuid.New()
		id := uuid.ID()
		inspection := entity.Inspection{
			Id:   int(id),
			Name: "mysql",
		}
		dao.Inspection.Insert(inspection)
	}
}

func init() {
	registerInspectorsMem()
	// service.RegisterInspection(InspectionMemLogic{})
}
