package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/logic"
)

var User CUser = CUser{}

type CUser struct {
}

func (C CUser) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	newToken, err := logic.MyJwt.GenerateToken(ctx, logic.Ctx.Get(ctx).Username)
	res = &v1.RefreshRes{
		Token: newToken,
	}
	return
}
