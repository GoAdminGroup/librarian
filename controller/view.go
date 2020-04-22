package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/librarian/guard"
	"github.com/GoAdminGroup/librarian/modules/theme"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
)

func (h *Handler) View(ctx *context.Context) {

	param := guard.GetViewParam(ctx)

	content, err := ioutil.ReadFile(param.FullPath)

	if err != nil {
		panic(err)
	}

	md := blackfriday.Run(content)

	h.HTML(ctx, types.Panel{
		Content: theme.Get(h.theme).HTML(md),
		CSS:     theme.Get(h.theme).CSS(),
		JS:      theme.Get(h.theme).JS(),
	}, false, true)
}
