package logic

import (
	"context"
	"errors"
	"fmt"
	"marmot/internal/dao"
	"marmot/internal/model/do"
	"marmot/internal/model/entity"
	"strings"

	"crypto/sha256"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

type lUser struct {
}

var (
	User = lUser{}
)

func (s *lUser) Login(ctx context.Context, username string, password string) (string, error) {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Passport: username,
		Password: password,
	}).Scan(&user)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("账户或者密码错误")
	}
	return user.Role, nil
}

// 查看是否登录
func (s *lUser) IsSignedIn(ctx context.Context, r *ghttp.Request) bool {
	token, exist := s.getToken(r)
	if !exist {
		return false
	}
	valid := MyJwt.Valid(r.Context(), token)
	return valid
}

// 解析 token
func (s *lUser) Parse(ctx context.Context, r *ghttp.Request) (bool, string) {
	token, exist := s.getToken(r)
	if !exist {
		return false, ""
	}
	claims, ok := MyJwt.Parse(r.Context(), token)
	if !ok {
		return false, ""
	}
	return ok, claims.Username
}

// 从 request 的 head 中获得 token string
func (*lUser) getToken(r *ghttp.Request) (string, bool) {
	header := r.GetHeader("Authorization")
	headerList := strings.Split(header, " ")
	if len(headerList) != 2 {
		return "", false
	}
	t := headerList[0]
	token := headerList[1]
	if t != "Bearer" {
		return "", false
	}
	if token == "" {
		return "", false
	}
	return token, true
}

type AdminInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*lUser) MarshalAdminToDb(ctx context.Context) {
	admin := &AdminInfo{}
	g.Cfg().MustGet(ctx, "admin").Scan(admin)
	g.Log().Debugf(ctx, color.RedString("====%v"), admin)
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Passport: admin.Username,
	}).Scan(&user)
	if err != nil {
		g.Log().Infof(ctx, color.RedString("err=%v, dao.User=%v"), err, user)
		panic("查询用户表失败，是否没有连接正确的数据库")
	}
	if user == nil {
		dao.User.Ctx(ctx).Insert(do.User{
			Passport: admin.Username,
			Password: fmt.Sprintf("%x", sha256.Sum256([]byte(admin.Password))),
			Nickname: admin.Username,
			Role:     "admin",
		})
	}
}

func init() {
	ctx := gctx.New()
	User.MarshalAdminToDb(ctx)
}
