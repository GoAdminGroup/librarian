package librarian

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"os"
	"path/filepath"
	"strings"
)

func (l *Librarian) initRouter(srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(config.GetUrlPrefix())
	authRoute := route.Group("/", auth.Middleware(l.Conn))

	for _, root := range l.roots {
		if root.Path[0] == '.' {
			root.Path = root.Path[2:]
		}
		replacer := strings.NewReplacer(root.Path, "", ".md", "")

		_ = filepath.Walk(root.Path, func(path string, info os.FileInfo, err error) error {

			if !info.IsDir() && filepath.Ext(path) == ".md" {
				path = replacer.Replace(path)
				authRoute.GET("/"+l.prefix+filepath.ToSlash(path), l.guard.View, l.handler.View)
			}

			return nil
		})
	}

	authRoute.GET("/"+l.prefix+"/write", l.handler.Write)

	return app
}
