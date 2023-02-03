package logic

import (
	"context"
	"fmt"
	v1 "marmot/api/v1"
	"marmot/internal/dao"
	"marmot/internal/model/do"
	"marmot/internal/model/entity"
	"marmot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type InspectionLogic struct {
}

// 1. 查看巡检列表
func (InspectionLogic) List(ctx context.Context) (res *[]v1.InspectListRes, err error) {
	res = &[]v1.InspectListRes{}
	var inspections []entity.Inspection
	dao.Inspection.Ctx(ctx).Scan(&inspections)
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
func (InspectionLogic) Inspect(ctx context.Context, id int) {
	inspection := entity.Inspection{}
	err := dao.Inspection.Ctx(ctx).Where("id", id).Scan(&inspection)
	if err != nil {
		g.Log().Errorf(ctx, "err=%v", err)
	}
	if err == nil {
		switch inspection.Name {
		case "mysql":
			MysqlInspector.inspect(ctx, id)
		}
	}
}

// 3. 巡检全部场景
func (logic InspectionLogic) InspectAll(ctx context.Context) {
	var inspections []entity.Inspection
	dao.Inspection.Ctx(ctx).ScanList(&inspections, "inspection")
	for _, inspection := range inspections {
		logic.Inspect(ctx, inspection.Id)
	}
}

// 4. 加载所有巡检处理器到数据库
func registerInspectors() {
	var inspections []entity.Inspection
	ctx := context.TODO()
	dao.Inspection.Ctx(ctx).Scan(&inspections)
	mysqlExist := false
	for _, inspection := range inspections {
		switch inspection.Name {
		case "mysql":
			mysqlExist = true
		}
	}
	if !mysqlExist {
		dao.Inspection.Ctx(ctx).Data(do.Inspection{Name: "mysql"}).Insert()
	}
}

func init() {
	fmt.Println("注册IInspection")
	registerInspectors()
	service.RegisterInspection(InspectionLogic{})
}
