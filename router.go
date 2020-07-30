package librarian

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
)

func (l *Librarian) initRouter(srv service.List) *context.App {

	app := context.NewApp()

	authRoute := app.Group("/", auth.Middleware(l.Conn))
	count := 0

	for _, root := range *l.roots {
		if root.Path[0] == '.' {
			root.Path = root.Path[2:]
		}
		replacer := strings.NewReplacer(root.Path, "", ".md", "")

		_ = filepath.Walk(root.Path, func(path string, info os.FileInfo, err error) error {

			if !info.IsDir() && filepath.Ext(path) == ".md" {
				path = filepath.ToSlash(replacer.Replace(path))
				authRoute.GET(path, l.guard.View, l.handler.View)
				if count == 0 {
					l.indexURL = config.Url("/" + l.prefix + path)
				}
				count++
			}

			return nil
		})
	}

	authRoute.GET("/"+l.prefix+"/write", l.handler.Write)

	return app
}
