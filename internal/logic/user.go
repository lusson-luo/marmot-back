package logic

import (
	"context"
	"errors"
	"marmot/internal/dao"
	"marmot/internal/model/do"
	"marmot/internal/model/entity"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
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

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *lUser) IsSignedIn(ctx context.Context, r *ghttp.Request) bool {
	header := r.GetHeader("Authorization")
	headerList := strings.Split(header, " ")
	if len(headerList) != 2 {
		return false
	}
	t := headerList[0]
	token := headerList[1]
	if t != "Bearer" {
		return false
	}
	if token == "" {
		return false
	}
	valid := MyJwt.Valid(r.Context(), token)
	return valid
}
