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

// 设置 cookie
func Ctx(r *ghttp.Request) {
	r.SetCtxVar(logic.BizCtxKey, logic.BizCtx{
		Cookie: r.Cookie,
	})
	r.Middleware.Next()
}

// 查看是否登录
func Auth(r *ghttp.Request) {
	valid := logic.User.IsSignedIn(r.GetCtx(), r)
	if !valid {
		r.SetError(gerror.NewCode(gcode.New(50008, "非法令牌", "非法令牌")))
		return
	}
	r.Middleware.Next()
}
