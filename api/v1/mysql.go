package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MysqlReq struct {
	g.Meta `path:"/mysql" tags:"mysql" method:"get" summary:"You first hello api"`
}
type MysqlRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
