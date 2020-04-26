package root

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/librarian/modules/error"
)

type Root struct {
	Path  string
	Title string
}

type Roots map[string]Root

func (r Roots) Add(key string, value Root) {
	r[key] = value
}

func (r Roots) GetPathFromPrefix(ctx *context.Context) string {
	return r.GetFromPrefix(ctx).Path
}

func (r Roots) GetTitleFromPrefix(ctx *context.Context) string {
	return r.GetFromPrefix(ctx).Title
}

func (r Roots) GetFromPrefix(ctx *context.Context) Root {
	prefix := ctx.Query("__prefix")
	if prefix == "" {
		prefix = "def"
	}
	if root, ok := r[prefix]; ok {
		return root
	}
	panic(errors.WrongPrefix)
}
