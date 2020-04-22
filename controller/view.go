package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/librarian/guard"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
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
		Content: `
<div class="markdown-viewer" style="
    padding: 20px;
    margin: auto;
    width: 94%;
    background-color: #FFFFFF;
    min-height: 800px;
">` + template.HTML(md) + `</div>
`,
	}, false, true)
}
