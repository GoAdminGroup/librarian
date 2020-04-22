package main

import (
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
		Language:           language.EN,
		IndexUrl:           "/librarian/def/view/README.md",
		Debug:              true,
		AccessAssetsLogOff: true,
		Theme:              "sword",
		Animation: config.PageAnimation{
			Type: "fadeInUp",
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := e.AddConfig(cfg).
		AddPlugins(librarian.
			NewLibrarian(filepath.Join(dir, "docs")),
		).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	go func() {
		_ = r.Run(":9033")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.SqliteConnection().Close()
}
