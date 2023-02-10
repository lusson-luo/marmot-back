package main

import (
	_ "marmot/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "marmot/internal/dbinit"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"marmot/internal/cmd"
	_ "marmot/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
