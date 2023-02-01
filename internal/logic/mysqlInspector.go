package logic

import (
	"context"
	"marmot/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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
		inspection, exist := dao.Inspection.FindById(id)
		if exist {
			var err error
			g.Log().Infof(gctx.New(), "mysql 执行")
			db, err := gdb.New(gdb.ConfigNode{
				Link: mysqlConfig.Url,
			})
			if db.PingMaster() != nil {
				inspection.Connection = false
			} else {
				inspection.Connection = true
			}
			inspection.StartTime = time.Now()
			inspection.Availability = true
			inspection.Count++
			for i, item := range mysqlConfig.Items {
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
			inspection.EndTime = time.Now()
			dao.Inspection.Update(inspection)
		}
	}
}

var mysqlConfig *MysqlConfig = &MysqlConfig{}

func init() {
	g.Cfg().MustGet(gctx.New(), "inspection.mysql").Scan(mysqlConfig)
	g.Log().Infof(gctx.New(), "mysqlConfig:%v", mysqlConfig)
}
