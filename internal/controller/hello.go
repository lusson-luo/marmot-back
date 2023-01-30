package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "demo/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Say(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
