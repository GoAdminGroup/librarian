package root

import (
	"github.com/GoAdminGroup/go-admin/context"
	errors "github.com/GoAdminGroup/librarian/modules/error"
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
		if len(r) != 0 {
			for _, v := range r {
				return v
			}
		}
	}
	if rr, ok := r[prefix]; ok {
		return rr
	}
	panic(errors.WrongPrefix)
}
