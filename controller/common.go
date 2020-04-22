package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/librarian/modules/constant"
	"github.com/GoAdminGroup/librarian/modules/root"
)

type Handler struct {
	roots root.Roots

	HTML func(ctx *context.Context, panel types.Panel, animation ...bool)
}

func NewHandler(root root.Roots) *Handler {
	return &Handler{
		roots: root,
	}
}

func (h *Handler) Prefix(ctx *context.Context) string {
	prefix := ctx.Query(constant.PrefixKey)
	if prefix == "" {
		return "def"
	}
	return prefix
}
