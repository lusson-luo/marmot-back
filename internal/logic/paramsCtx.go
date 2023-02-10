package logic

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type lCtx struct {
}

var Ctx lCtx = lCtx{}

const (
	ParamsCtxKey = "paramsCtx"
)

type ParamsCtx struct {
	Username string
}

func (lCtx) Set(r *ghttp.Request, username string) {
	r.SetCtxVar(ParamsCtxKey, ParamsCtx{
		Username: username,
	})
}

func (lCtx) Get(ctx context.Context) *ParamsCtx {
	v, ok := ctx.Value(ParamsCtxKey).(ParamsCtx)
	if ok {
		return &v
	} else {
		return nil
	}
}
