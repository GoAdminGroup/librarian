package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/librarian/modules/constant"
	"github.com/GoAdminGroup/librarian/modules/root"
	"path/filepath"
	"strings"
)

type Guardian struct {
	conn  db.Connection
	roots root.Roots
}

func New(r root.Roots, c db.Connection) *Guardian {
	return &Guardian{
		roots: r,
		conn:  c,
	}
}

const (
	viewParamKey = "view_param"
)

type Base struct {
	Path     string
	Prefix   string
	FullPath string
	Error    error
}

func (g *Guardian) GetPrefix(ctx *context.Context) string {
	prefix := ctx.Query(constant.PrefixKey)
	if prefix == "" {
		return "def"
	}
	return prefix
}

func (g *Guardian) getPaths(ctx *context.Context) (string, string, error) {
	var (
		err          error
		relativePath = strings.Replace(ctx.Path(), config.Url("/librarian/"+ctx.Query("__prefix")+"/view"), "", -1) + ".md"
		path         = filepath.Join(g.roots.GetPathFromPrefix(ctx), relativePath)
	)

	return relativePath, path, err
}
