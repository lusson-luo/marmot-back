package middleware

import (
	"marmot/internal/logic"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Ctx 设置 cookie
func Ctx(r *ghttp.Request) {
	_, username := logic.User.Parse(r.GetCtx(), r)
	r.SetCtxVar(logic.BizCtxKey, logic.Ctx.Init(r.GetCtx(), username))
	r.Middleware.Next()
}

// Auth 查看是否登录
func Auth(r *ghttp.Request) {
	valid := logic.User.IsSignedIn(r.GetCtx(), r)
	if !valid {
		r.SetError(gerror.NewCode(gcode.New(50008, "请重新登录", "非法令牌或还未登录")))
		return
	}
	r.Middleware.Next()
}
