package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/logic"
)

var User CUser = CUser{}

type CUser struct {
}

func (C CUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	role, err := logic.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	tokenString, err := logic.MyJwt.GenerateToken(ctx, req.Username)
	if err != nil {
		return
	}
	res.Token = tokenString
	res.Role = role
	return
}

func (C CUser) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	newToken, err := logic.MyJwt.GenerateToken(ctx, logic.Ctx.Get(ctx).Username)
	res = &v1.RefreshRes{
		Token: newToken,
	}
	return
}
