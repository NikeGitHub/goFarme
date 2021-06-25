package tool

import (
	"code.shomes.cn/pkg/bootstrap/tool"
	"code.shomes.cn/pkg/db"
	"golang.org/x/net/context"
)

func Smart(ctx context.Context) *db.DB {
	//return db.Ctx(ctx, tool.DBCtx+"smart")
	return db.Ctx(ctx, tool.DBCtx+"smartgy")
}
