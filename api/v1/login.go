package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/api/user/login" tags:"login" method:"POST" summary:"登录"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Token  string `json:"token"`
	// user or admin
	Role string `json:"role"`
}
