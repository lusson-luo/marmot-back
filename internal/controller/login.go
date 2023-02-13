package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/logic"
)

var Login CLogin = CLogin{}

type CLogin struct {
}

// 登录
func (c CLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	role, tokenString, err := logic.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	res = &v1.LoginRes{
		Token: tokenString,
		Role:  role,
	}
	return
}
