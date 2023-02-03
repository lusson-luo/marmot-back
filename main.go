package main

import (
	_ "marmot/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"marmot/internal/cmd"
	_ "marmot/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
