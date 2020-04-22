package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
)

type FilesParam struct {
	*Base
}

func (g *Guardian) View(ctx *context.Context) {

	relativePath, path, err := g.getPaths(ctx)

	ctx.SetUserValue(viewParamKey, &FilesParam{
		Base: &Base{
			Path:     relativePath,
			FullPath: path,
			Error:    err,
			Prefix:   g.GetPrefix(ctx),
		},
	})
	ctx.Next()
}

func GetViewParam(ctx *context.Context) *FilesParam {
	return ctx.UserValue[viewParamKey].(*FilesParam)
}
