package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "demo/api/v1"
)

var (
	Mysql = CMysql{}
)

type CMysql struct {
	CreateDb    bool
	CreateTable bool
	Insert      bool
	Update      bool
	Delete      bool
	Select      bool
}

func (c *CMysql) Internal(ctx context.Context, req *v1.MysqlReq) (res *v1.MysqlRes, err error) {

	g.RequestFromCtx(ctx).Response.Writeln("mysql !")
	// 1. 创建数据库
	record, err := g.DB().GetOne(ctx, "create database internal_test")
	if err != nil {
		c.CreateDb = false
		g.Log().Errorf(ctx, "mysql.internal 创建数据库失败,err=%v", err)
		return
	}
	c.CreateDb = false
	g.Log().Debugf(ctx, "mysql.internal=%v", record.Map())
	// 2. 创建表
	// record, err := g.DB().GetOne(ctx, "create database internal_test")
	// if err != nil {
	// 	c.CreateDb = false
	// 	g.Log().Errorf(ctx, "hello.Say() 查询数据库失败,err=%v", err)
	// }
	return
}
