package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	_ "github.com/GoAdminGroup/themes/sword"

	"github.com/GoAdminGroup/filemanager"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/librarian"
	"github.com/GoAdminGroup/librarian/modules/theme"
	"github.com/gin-gonic/gin"
)

func main() {

	// TODO: installation
	//
	// 1. download the database
	// 2. set librarian.yml

	var (
		dbPath, port, prefix, configPath, filePath, logo, miniLogo, title string
		debug                                                             bool
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.StringVar(&dbPath, "db", "./admin.db", "db path")
	flag.StringVar(&prefix, "prefix", "docs", "url prefix")
	flag.StringVar(&configPath, "config", "", "config path")
	flag.StringVar(&filePath, "path", "", "file path")
	flag.StringVar(&port, "port", "80", "http listen port")
	flag.StringVar(&title, "title", "Librarian", "title")
	flag.StringVar(&logo, "logo", "<b>Li</b>brarian", "logo")
	flag.StringVar(&miniLogo, "mini_logo", "Li", "mini logo")
	flag.Parse()

	if filePath == "" {
		panic("wrong file path")
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.New()
	e := engine.Default()

	const visitorRoleID = int64(3)

	r.Use(func(ctx *gin.Context) {
		_, exist := e.User(ctx)
		if !exist {
			conn := e.SqliteConnection()
			user := models.User().SetConn(conn).Find(visitorRoleID)
			c := context.NewContext(ctx.Request)
			_ = auth.SetCookie(c, user, conn)
			ctx.Header("Set-Cookie", c.Response.Header.Get("Set-Cookie"))
			ctx.Request.Header.Set("Cookie", c.Response.Header.Get("Set-Cookie"))
		}
	})

	var cfg config.Config

	if configPath != "" {
		cfg = config.ReadFromYaml(configPath)
	} else {
		cfg = config.Config{
			Databases: config.DatabaseList{
				"default": {
					Driver: config.DriverSqlite,
					File:   dbPath,
				},
			},
			UrlPrefix: prefix,
			Store: config.Store{
				Path:   "./uploads",
				Prefix: "uploads",
			},
			Language:                      language.EN,
			Debug:                         debug,
			AccessAssetsLogOff:            true,
			HideConfigCenterEntrance:      true,
			HideAppInfoEntrance:           true,
			HideVisitorUserCenterEntrance: true,
			Logo:                          template.HTML(logo),
			MiniLogo:                      template.HTML(miniLogo),
			Theme:                         "sword",
			Title:                         title,
			NoLimitLoginIP:                true,
			ExcludeThemeComponents:        []string{"datatable", "form"},
			//Animation: config.PageAnimation{
			//	Type: "fadeInUp",
			//},
		}
	}

	theme.Set(theme.Config{
		HideNavBar:   true,
		HideMenuIcon: true,
		FixedSidebar: true,
		ChangeTitle:  true,
	})

	li := librarian.NewLibrarianWithConfig(librarian.Config{
		Path:           filePath,
		MenuUserRoleID: visitorRoleID,
		BuildMenu:      true,
	})

	if err := e.AddConfig(cfg).
		AddPlugins(li, filemanager.NewFileManager(filePath)).
		Use(r); err != nil {
		panic(err)
	}

	_ = models.Site().SetConn(e.SqliteConnection()).Update(form.Values{
		"logo":      []string{logo},
		"mini_logo": []string{miniLogo},
		"prefix":    []string{prefix},
		"title":     []string{title},
	})

	//config.GetService(e.Services.Get("config"))

	indexURL := li.GetFirstMenu().Path

	r.Static("/uploads", "./uploads")

	e.Data("GET", "/"+prefix, func(ctx *context.Context) {
		ctx.Redirect(config.Url(indexURL))
	}, true)

	go func() {
		_ = r.Run(":" + port)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.SqliteConnection().Close()
}
