package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "gadmin-backend/internal/packed"

	_ "gadmin-backend/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gadmin-backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
