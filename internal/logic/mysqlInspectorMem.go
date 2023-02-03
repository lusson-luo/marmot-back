package logic

import (
	"context"
	"marmot/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

// 纯内存模式
type mysqlMemInspector struct {
}

var MysqlMemInspector mysqlMemInspector = mysqlMemInspector{}

type MysqlConfigMem struct {
	Enabled bool   `json:"enabled"`
	Url     string `json:"url"`
	Items   []Item `json:"items"`
}

type ItemMem struct {
	Name string   `json:"name"`
	Cmds []string `json:"cmds"`
}

func (MysqlInspector *mysqlMemInspector) inspect(ctx context.Context, id int) {
	if mysqlConfigMem.Enabled {
		inspection, exist := dao.Inspection.FindById(id)
		// inspection := do.Inspection{}
		// err:=dao.Inspection.Ctx(ctx).Where("id",id).Scan(&inspection)
		if exist {
			var err error
			g.Log().Infof(gctx.New(), "mysql 执行")
			db, err := gdb.New(gdb.ConfigNode{
				Link: mysqlConfigMem.Url,
			})
			if db.PingMaster() != nil {
				inspection.Connection = false
			} else {
				inspection.Connection = true
			}
			inspection.StartTime = gtime.NewFromTime(time.Now())
			inspection.Availability = true
			inspection.Count++
			for i, item := range mysqlConfigMem.Items {
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
			}
			if inspection.Availability {
				inspection.SuccessCount++
			} else {
				inspection.FailedCount++
			}
			inspection.EndTime = gtime.NewFromTime(time.Now())
			dao.Inspection.Update(inspection)
		}
	}
}

var mysqlConfigMem *MysqlConfigMem = &MysqlConfigMem{}

func init() {
	g.Cfg().MustGet(gctx.New(), "inspection.mysql").Scan(mysqlConfigMem)
	g.Log().Infof(gctx.New(), "mysqlConfig:%v", mysqlConfigMem)
}
