package main

import (
	_ "marmot/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"marmot/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
