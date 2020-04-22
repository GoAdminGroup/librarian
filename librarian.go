package librarian

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/librarian/controller"
	"github.com/GoAdminGroup/librarian/guard"
	"github.com/GoAdminGroup/librarian/modules/error"
	language2 "github.com/GoAdminGroup/librarian/modules/language"
	"github.com/GoAdminGroup/librarian/modules/root"
)

type Librarian struct {
	*plugins.Base

	roots root.Roots

	handler *controller.Handler
	guard   *guard.Guardian
}

const Name = "librarian"

func NewLibrarian(rootPath string, titles ...string) *Librarian {

	if rootPath == "" {
		panic("librarian: create fail, wrong path")
	}

	title := Name
	if len(titles) > 0 {
		title = titles[0]
	}
	return &Librarian{
		Base:  &plugins.Base{PlugName: Name},
		roots: root.Roots{"def": root.Root{Path: rootPath, Title: title}},
	}
}

type Config struct {
	Path  string
	Title string
}

func NewLibrarianWithConfig(cfg Config) *Librarian {

	if cfg.Path == "" {
		panic("librarian: create fail, wrong path")
	}

	if cfg.Title == "" {
		cfg.Title = Name
	}

	return &Librarian{
		Base:  &plugins.Base{PlugName: Name},
		roots: root.Roots{"def": root.Root{Path: cfg.Path, Title: cfg.Title}},
	}
}

func (l *Librarian) InitPlugin(srv service.List) {

	// DO NOT DELETE
	l.InitBase(srv)

	l.Conn = db.GetConnection(srv)
	l.handler = controller.NewHandler(l.roots)
	l.guard = guard.New(l.roots, l.Conn)
	l.App = l.initRouter(srv)
	l.handler.HTML = l.HTML

	language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)

	errors.Init()
}

func (l *Librarian) AddRoot(key string, value root.Root) *Librarian {
	l.roots.Add(key, value)
	return l
}
