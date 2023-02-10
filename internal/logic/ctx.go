package logic

import (
	"context"
)

type lCtx struct {
}

var Ctx lCtx = lCtx{}

const (
	BizCtxKey = "bizCtx"
)

type BizCtx struct {
	Username string
}

func (lCtx) Init(ctx context.Context, username string) BizCtx {
	return BizCtx{
		Username: username,
	}
}

func (lCtx) Get(ctx context.Context) *BizCtx {
	v, ok := ctx.Value(BizCtxKey).(BizCtx)
	if ok {
		return &v
	} else {
		return nil
	}
}
