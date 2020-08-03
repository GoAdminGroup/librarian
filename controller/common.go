package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/librarian/modules/constant"
	"github.com/GoAdminGroup/librarian/modules/root"
)

type Handler struct {
	roots *root.Roots
	theme string

	HTML func(ctx *context.Context, panel types.Panel, options ...template.ExecuteOptions)
}

func NewHandler(root *root.Roots, theme string) *Handler {
	return &Handler{
		roots: root,
		theme: theme,
	}
}

func (h *Handler) Prefix(ctx *context.Context) string {
	prefix := ctx.Query(constant.PrefixKey)
	if prefix == "" {
		return "def"
	}
	return prefix
}

func (h *Handler) Update(root *root.Roots, theme string) {
	h.roots = root
	h.theme = theme
}
