package logic

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt"
)

type lJwt struct {
}

const (
	ExpiresTime = time.Minute * 1
)

var (
	Ljwt      lJwt   = lJwt{}
	JwtSecert []byte = []byte("123456")
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成 jwt 格式 token
func (lJwt) GenerateToken(ctx context.Context, username string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiresTime).Unix(),
		},
	})
	g.Log().Debugf(ctx, "", token)
	tokenString, err = token.SignedString(JwtSecert)
	return
}

// 验证 token 是否合法
func (lJwt) Valid(ctx context.Context, token string) (valided bool) {
	claims := &MyClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecert, nil
	})
	if err != nil {
		valided = false
	} else {
		valided = tkn.Valid
	}
	return
}
