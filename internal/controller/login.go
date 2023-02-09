package controller

import (
	"context"
	v1 "marmot/api/v1"
	"marmot/internal/logic"
)

var Login Clogin = Clogin{}

type Clogin struct {
}

func (C Clogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
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
