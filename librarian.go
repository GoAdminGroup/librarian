package librarian

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/librarian/controller"
	"github.com/GoAdminGroup/librarian/guard"
	"github.com/GoAdminGroup/librarian/modules/error"
	language2 "github.com/GoAdminGroup/librarian/modules/language"
	"github.com/GoAdminGroup/librarian/modules/root"
	"github.com/GoAdminGroup/librarian/modules/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

type Librarian struct {
	*plugins.Base

	roots root.Roots
	theme string

	buildMenu bool

	menuUserRoleID int64

	prefix string

	handler *controller.Handler
	guard   *guard.Guardian
}

const Name = "librarian"

func NewLibrarian(rootPath string, menuUserRoleID ...int64) *Librarian {

	if rootPath == "" {
		panic("librarian: create fail, wrong path")
	}

	uid := int64(0)
	if len(menuUserRoleID) > 0 {
		uid = menuUserRoleID[0]
	}
	return &Librarian{
		Base:           &plugins.Base{PlugName: Name},
		roots:          root.Roots{"def": root.Root{Path: rootPath, Title: Name}},
		theme:          "github",
		buildMenu:      true,
		menuUserRoleID: uid,
		prefix:         "librarian",
	}
}

type Config struct {
	Path           string
	Title          string
	Theme          string
	Prefix         string
	BuildMenu      bool
	MenuUserRoleID int64
}

func NewLibrarianWithConfig(cfg Config) *Librarian {

	if cfg.Path == "" {
		panic("librarian: create fail, wrong path")
	}

	if cfg.Title == "" {
		cfg.Title = Name
	}

	if cfg.Theme == "" {
		cfg.Theme = "github"
	}

	if cfg.Prefix == "" {
		cfg.Prefix = "librarian"
	}

	return &Librarian{
		Base:           &plugins.Base{PlugName: Name},
		roots:          root.Roots{"def": root.Root{Path: cfg.Path, Title: cfg.Title}},
		theme:          cfg.Theme,
		buildMenu:      cfg.BuildMenu,
		menuUserRoleID: cfg.MenuUserRoleID,
		prefix:         cfg.Prefix,
	}
}

func (l *Librarian) InitPlugin(srv service.List) {

	// DO NOT DELETE
	l.InitBase(srv)

	l.Conn = db.GetConnection(srv)
	l.handler = controller.NewHandler(l.roots, l.theme)
	l.guard = guard.New(l.roots, l.Conn, l.prefix)
	l.App = l.initRouter(srv)
	l.handler.HTML = l.HTML

	language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)

	if l.buildMenu {
		l.InitMenu()
	}

	errors.Init()
}

func (l *Librarian) AddRoot(key string, value root.Root) *Librarian {
	l.roots.Add(key, value)
	return l
}

func (l *Librarian) InitMenu() {

	for key, r := range l.roots {
		navPath := r.Path + "/nav.yml"
		if util.FileExist(navPath) {
			buildMenus, err := l.siteTable().
				Where("key", "=", siteMenuIDsKey(key)).
				First()
			if db.CheckError(err, db.QUERY) {
				logger.Error("librarian build menu error", err)
				continue
			}

			if buildMenus == nil {
				if err := l.setMenu(key, navPath, false); err != nil {
					logger.Error("librarian set menu error", err)
				}
			} else {
				// clear old menu
				buildMenuIDs := strings.Split(buildMenus["value"].(string), ",")
				buildMenuIDInterfaces := make([]interface{}, len(buildMenuIDs))
				for i := 0; i < len(buildMenuIDs); i++ {
					buildMenuIDInterfaces[i] = buildMenuIDs[i]
				}
				err := l.menuTable().WhereIn("id", buildMenuIDInterfaces).Delete()
				if db.CheckError(err, db.DELETE) {
					logger.Error("librarian clear menu error", err)
					continue
				}
				err = l.roleMenuTable().WhereIn("menu_id", buildMenuIDInterfaces).Delete()
				if db.CheckError(err, db.DELETE) {
					logger.Error("librarian clear role menu error", err)
					continue
				}
				if err := l.setMenu(key, navPath, true); err != nil {
					logger.Error("librarian set menu error", err)
				}
			}
		}
	}
}

// TODO: add transaction
func (l *Librarian) setMenu(prefix, navPath string, has bool) error {
	b, err := ioutil.ReadFile(navPath)

	if err != nil {
		return err
	}

	var navs = make(map[string]interface{})

	err = yaml.Unmarshal(b, &navs)

	if err != nil {
		return err
	}

	maxOrderMenu, err := l.menuTable().Select("order").OrderBy("order", "desc").First()
	if db.CheckError(err, db.QUERY) {
		logger.Fatal(err)
	}
	order := int64(1)
	if o, ok := maxOrderMenu["order"].(int64); ok {
		order = o
	}
	ids := make([]string, 0)

	for _, level1 := range navs["nav"].([]interface{}) {
		for key, value := range level1.(map[interface{}]interface{}) {
			if level2, ok := value.([]interface{}); ok {
				level1NavID, err := l.menuTable().Insert(dialect.H{
					"icon":      "fa-file-o",
					"title":     key.(string),
					"uri":       "",
					"parent_id": 0,
					"order":     order,
				})
				if db.CheckError(err, db.INSERT) {
					logger.Fatal(err)
				}
				ids = append(ids, strconv.Itoa(int(level1NavID)))
				order++
				for _, level2Nav := range level2 {
					for key, value := range level2Nav.(map[interface{}]interface{}) {
						if level3, ok := value.([]interface{}); ok {
							level2NavID, err := l.menuTable().Insert(dialect.H{
								"icon":      "fa-file-o",
								"title":     key.(string),
								"uri":       "",
								"parent_id": level1NavID,
								"order":     order,
							})
							if db.CheckError(err, db.INSERT) {
								logger.Fatal(err)
							}
							ids = append(ids, strconv.Itoa(int(level2NavID)))
							order++
							for _, level3Nav := range level3 {
								for key, value := range level3Nav.(map[interface{}]interface{}) {
									// third level
									id, err := l.menuTable().Insert(dialect.H{
										"icon":      "fa-file-o",
										"title":     key.(string),
										"uri":       l.menuPath(prefix, value),
										"parent_id": level2NavID,
										"order":     order,
									})
									if db.CheckError(err, db.INSERT) {
										logger.Fatal(err)
									}
									ids = append(ids, strconv.Itoa(int(id)))
									order++
								}
							}
						} else {
							// second level
							id, err := l.menuTable().Insert(dialect.H{
								"icon":      "fa-file-o",
								"title":     key.(string),
								"uri":       l.menuPath(prefix, value),
								"parent_id": level1NavID,
								"order":     order,
							})
							if db.CheckError(err, db.INSERT) {
								logger.Fatal(err)
							}
							ids = append(ids, strconv.Itoa(int(id)))
							order++
						}
					}
				}
			} else {
				// first level
				id, err := l.menuTable().Insert(dialect.H{
					"icon":      "fa-file-o",
					"uri":       l.menuPath(prefix, value),
					"title":     key.(string),
					"parent_id": 0,
					"order":     order,
				})
				if db.CheckError(err, db.INSERT) {
					logger.Fatal(err)
				}
				ids = append(ids, strconv.Itoa(int(id)))
				order++
			}
		}
	}

	if len(ids) > 0 {
		if has {
			_, err := l.siteTable().Where("key", "=", siteMenuIDsKey(prefix)).
				Update(dialect.H{
					"value": strings.Join(ids, ","),
				})
			if db.CheckError(err, db.INSERT) {
				logger.Fatal(err)
			}
		} else {
			_, err := l.siteTable().Insert(dialect.H{
				"key":   siteMenuIDsKey(prefix),
				"value": strings.Join(ids, ","),
			})
			if db.CheckError(err, db.UPDATE) {
				logger.Fatal(err)
			}
		}
	}

	if l.menuUserRoleID != int64(0) {
		for _, id := range ids {
			_, err := l.roleMenuTable().Insert(dialect.H{
				"menu_id": id,
				"role_id": l.menuUserRoleID,
			})
			if db.CheckError(err, db.INSERT) {
				logger.Fatal(err)
			}
		}
	}

	return nil
}

func (l *Librarian) menuTable() *db.SQL {
	return db.WithDriver(l.Conn).Table("goadmin_menu")
}

func (l *Librarian) roleMenuTable() *db.SQL {
	return db.WithDriver(l.Conn).Table("goadmin_role_menu")
}

func (l *Librarian) siteTable() *db.SQL {
	return db.WithDriver(l.Conn).Table("goadmin_site")
}

func siteMenuIDsKey(prefix string) string {
	return "librarian_build_menu_" + prefix
}

func (l *Librarian) menuPath(prefix string, path interface{}) string {
	p := strings.Replace(path.(string), ".md", "", -1)
	if prefix == "def" {
		return "/" + l.prefix + "/" + p
	}
	return "/" + l.prefix + "/" + p + "?__prefix=" + prefix
}
