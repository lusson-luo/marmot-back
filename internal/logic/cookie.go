package logic

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

type LCookie struct {
}

var Lcookie LCookie = LCookie{}

const (
	BizCtxKey = "bizCtx"
)

type BizCtx struct {
	Cookie *ghttp.Cookie
}

// 把数据放到 cookie 中
func (LCookie) Set(ctx context.Context, key string, value string) (err error) {
	bizCtx, ok := ctx.Value(BizCtxKey).(BizCtx)
	if ok {
		bizCtx.Cookie.Set(key, value)
	} else {
		err = gerror.New("未找到 cookie，是不是没有注册")
	}
	return
}
