package logic

import (
	"context"
	"fmt"
	"marmot/internal/dao"
	"marmot/internal/model/do"
	"marmot/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

type mysqlInspector struct {
}

var MysqlInspector mysqlInspector = mysqlInspector{}

type MysqlConfig struct {
	Enabled bool   `json:"enabled"`
	Url     string `json:"url"`
	Items   []Item `json:"items"`
}

type Item struct {
	Name string   `json:"name"`
	Cmds []string `json:"cmds"`
}

func (MysqlInspector *mysqlInspector) inspect(ctx context.Context, id int) {
	if mysqlConfig.Enabled {
		// inspection, exist := dao.Inspection.FindById(id)
		inspection := entity.Inspection{}
		err := dao.Inspection.Ctx(ctx).Where("id", id).Scan(&inspection)
		if err != nil {
			g.Log().Errorf(ctx, "err=", err)
		}
		if err == nil {
			var err error
			g.Log().Infof(gctx.New(), "mysql 开始巡检")
			db, err := gdb.New(gdb.ConfigNode{
				Link: mysqlConfig.Url,
			})
			if db.PingMaster() != nil {
				inspection.Connection = false
			} else {
				inspection.Connection = true
			}
			inspection.StartTime = gtime.NewFromTime(time.Now())
			inspection.Availability = true
			inspection.Count++

			inspectTaskId := GetCurrentInspectTaskId(ctx, inspection.Id)

			for i, item := range mysqlConfig.Items {
				startTime := gtime.NewFromTime(time.Now())
				for _, cmd := range item.Cmds {
					_, err = db.GetOne(ctx, cmd)
					if err != nil {
						break
					}
				}
				if err == nil {
					g.Log().Infof(gctx.New(), "\t%d. %s 执行成功", i, item.Name)
				} else {
					g.Log().Infof(gctx.New(), "\t%d. %s 执行失败", i, item.Name)
					g.Log().Debugf(gctx.New(), "\t%d. %s 执行失败原因=%v", i, item.Name, err)
					inspection.Availability = false
				}
				inspectionDetail := do.InspectionDetail{
					Name:          item.Name,
					InspectTaskId: inspectTaskId + 1,
					ExecStatus:    err == nil,
					ErrMsg:        err,
					StartTime:     startTime,
					EndTime:       gtime.NewFromTime(time.Now()),
					InspectionId:  inspection.Id,
				}
				result, err := dao.InspectionDetail.Ctx(ctx).Data(inspectionDetail).Insert()
				if err != nil {
					fmt.Sprintf("insert inspectionDetail failed: %s", err)
				} else {
					fmt.Sprintf("insert inspectionDetail successed: %s", result)
				}
			}
			if inspection.Availability {
				inspection.SuccessCount++
			} else {
				inspection.FailedCount++
			}
			inspection.EndTime = gtime.NewFromTime(time.Now())
			_, err = dao.Inspection.Ctx(ctx).Data(inspection).Where("id", inspection.Id).Update()
			if err != nil {
				g.Log().Errorf(gctx.New(), "mysql 巡检结果保存失败=%v", err)
			}
		}
	}
}

var mysqlConfig *MysqlConfig = &MysqlConfig{}

func init() {
	g.Cfg().MustGet(gctx.New(), "inspection.mysql").Scan(mysqlConfig)
	g.Log().Infof(gctx.New(), "mysqlConfig:%v", mysqlConfig)
}
