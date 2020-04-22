package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	_ "github.com/GoAdminGroup/themes/sword"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/librarian"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	e := engine.Default()

	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Driver: config.DriverSqlite,
				File:   "./admin.db",
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:                 language.EN,
		IndexUrl:                 "/librarian/def/view/README",
		Debug:                    true,
		AccessAssetsLogOff:       true,
		HideConfigCenterEntrance: true,
		HideAppInfoEntrance:      true,
		Logo:                     "<b>Li</b>brarian",
		MiniLogo:                 "Li",
		Theme:                    "sword",
		//Animation: config.PageAnimation{
		//	Type: "fadeInUp",
		//},
		CustomHeadHtml: `<style>
.navbar.navbar-static-top, .fa.fa-file-o  {
	display:none;
}
.navigation-box {
   position: fixed;
   width: 260px;
}
</style>`,
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := e.AddConfig(cfg).
		AddNavButtons("Menu", "", action.Jump("/admin/menu")).
		AddNavButtons("", icon.Pencil, action.Jump("/admin/menu")).
		AddPlugins(librarian.NewLibrarian(filepath.Join(dir, "docs"))).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	e.Data("GET", "/admin/librarian", func(ctx *context.Context) {
		conn := e.SqliteConnection()
		user := models.User().SetConn(conn).Find(3)
		_ = auth.SetCookie(ctx, user, conn)
		ctx.Redirect("/admin/librarian/def/view/README")
	}, true)

	go func() {
		_ = r.Run(":9033")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.SqliteConnection().Close()
}
