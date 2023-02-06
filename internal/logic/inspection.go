package logic

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/dao"
	"marmot/internal/model/do"
	"marmot/internal/model/entity"
	"marmot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type InspectionLogic struct {
	// 控制并发，一次只执行1个巡检任务
	structChan     chan struct{}
	inspectLoading map[int]bool
}

// 1. 查看巡检列表
func (l InspectionLogic) List(ctx context.Context) (res *[]v1.InspectListRes, err error) {
	res = &[]v1.InspectListRes{}
	var inspections []entity.Inspection
	dao.Inspection.Ctx(ctx).Scan(&inspections)
	for _, v := range inspections {
		inspect := v1.InspectListRes{
			Id:             v.Id,
			Name:           v.Name,
			Count:          v.Count,
			SuccessCount:   v.SuccessCount,
			FailedCount:    v.FailedCount,
			Connection:     v.Connection,
			Availability:   v.Availability,
			StartTime:      v.StartTime,
			EndTime:        v.EndTime,
			InspectLoading: l.GetInspectLoading(v.Id),
		}
		tmp := append(*res, inspect)
		res = &tmp
	}
	err = nil
	return
}

func (l InspectionLogic) GetInspectLoading(id int) bool {
	if _, ok := l.inspectLoading[id]; ok {
		return l.inspectLoading[id]
	}
	return false
}

func (l InspectionLogic) setInspectLoading(id int, inspectLoading bool) {
	l.inspectLoading[id] = inspectLoading
}

// 巡检单项场景
func (l InspectionLogic) inspect(ctx context.Context, id int) {
	Block(l.structChan)
	defer UnBlock(l.structChan)
	l.setInspectLoading(id, true)
	defer l.setInspectLoading(id, false)
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

func Block(structChan chan struct{}) {
	<-structChan
}

func UnBlock(structChan chan struct{}) {
	structChan <- struct{}{}
}

// 2. 巡检指定场景
// 如果没有传场景id，巡检所有场景
func (l InspectionLogic) InspectSelection(ctx context.Context, ids []int) {
	var inspections []entity.Inspection
	if len(ids) == 0 {
		dao.Inspection.Ctx(ctx).Scan(&inspections)
	} else {
		var idwheres []interface{} = make([]interface{}, 0)
		for _, id := range ids {
			idwheres = append(idwheres, id)
		}
		dao.Inspection.Ctx(ctx).WhereIn("id", idwheres).Scan(&inspections)
	}
	for _, inspection := range inspections {
		l.inspect(ctx, inspection.Id)
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

// GetCurrentInspectTaskId 查询当前最新的 巡检任务 ID
func GetCurrentInspectTaskId(ctx context.Context, inspectId int) int {
	currentInspectTaskId, err := dao.InspectionDetail.Ctx(ctx).Where("inspection_id", inspectId).OrderDesc("inspect_task_id").Limit(1).Value("inspect_task_id")
	if err != nil {
		g.Log().Errorf(ctx, "getCurrentInspectTaskId err: %s", err)
	}
	if currentInspectTaskId.Int() < 10000 {
		return 10000
	} else {
		return currentInspectTaskId.Int()
	}
}

func init() {
	registerInspectors()
	structChan := make(chan struct{}, 1)
	structChan <- struct{}{}
	service.RegisterInspection(InspectionLogic{
		structChan:     structChan,
		inspectLoading: map[int]bool{},
	})
}
