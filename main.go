package main

import (
	_ "latest-news/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"latest-news/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
