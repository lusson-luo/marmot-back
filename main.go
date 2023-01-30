package main

import (
	_ "demo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
