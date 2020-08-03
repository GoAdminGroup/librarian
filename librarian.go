package librarian

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/GoAdminGroup/librarian/controller"
	"github.com/GoAdminGroup/librarian/guard"
	errors "github.com/GoAdminGroup/librarian/modules/error"
	language2 "github.com/GoAdminGroup/librarian/modules/language"
	"github.com/GoAdminGroup/librarian/modules/root"
	"github.com/GoAdminGroup/librarian/modules/theme"
	"github.com/GoAdminGroup/librarian/modules/util"
	"gopkg.in/yaml.v2"
)

type Librarian struct {
	*plugins.Base
	roots          *root.Roots
	theme          string
	buildMenu      bool
	menuUserRoleID int64
	indexURL       string
	prefix         string
	handler        *controller.Handler
	guard          *guard.Guardian
}

func init() {
	plugins.Add(&Librarian{Base: &plugins.Base{PlugName: Name}, prefix: Name, roots: new(root.Roots)})
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
		roots:          &root.Roots{"def": root.Root{Path: rootPath, Title: Name}},
		theme:          "github",
		buildMenu:      true,
		menuUserRoleID: uid,
		prefix:         Name,
	}
}

type Config struct {
	Path           string `json:"path",yaml:"path",ini:"path"`
	Title          string `json:"title",yaml:"title",ini:"title"`
	Theme          string `json:"theme",yaml:"theme",ini:"theme"`
	Prefix         string `json:"prefix",yaml:"prefix",ini:"prefix"`
	BuildMenu      bool   `json:"build_menu",yaml:"build_menu",ini:"build_menu"`
	MenuUserRoleID int64  `json:"menu_user_role_id",yaml:"menu_user_role_id",ini:"menu_user_role_id"`
}

func NewLibrarianWithConfig(cfg Config) *Librarian {

	if cfg.Path == "" {
		panic("librarian: create fail, wrong path")
	}

	if !util.FileExist(cfg.Path) {
		panic("librarian: wrong directory path")
	}

	if cfg.Title == "" {
		cfg.Title = Name
	}

	if cfg.Theme == "" {
		cfg.Theme = "github"
	}

	return &Librarian{
		Base:           &plugins.Base{PlugName: Name},
		roots:          &root.Roots{"def": root.Root{Path: cfg.Path, Title: cfg.Title}},
		theme:          cfg.Theme,
		buildMenu:      cfg.BuildMenu,
		menuUserRoleID: cfg.MenuUserRoleID,
		prefix:         cfg.Prefix,
	}
}

func (l *Librarian) IsInstalled() bool {
	return len(*l.roots) != 0
}

func (l *Librarian) GetIndexURL() string {
	return l.indexURL
}

func (l *Librarian) InitPlugin(srv service.List) {

	// DO NOT DELETE
	l.InitBase(srv, Name)

	if len(*l.roots) == 0 {
		checkExist, _ := l.siteTable().
			Where("key", "=", "librarian_roots").
			First()
		if checkExist != nil {
			_ = json.Unmarshal([]byte(checkExist["value"].(string)), l.roots)
		}
	}

	if l.theme == "" {
		checkThemeExist, _ := l.siteTable().
			Where("key", "=", "librarian_theme").
			First()
		if checkThemeExist != nil {
			l.theme = checkThemeExist["value"].(string)
		}
	}

	l.handler = controller.NewHandler(l.roots, l.theme)
	l.guard = guard.New(l.roots, l.Conn, l.prefix)
	l.App = l.initRouter(srv)
	l.handler.HTML = l.HTMLMenu

	language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)

	if l.buildMenu {
		l.InitMenu()
	}

	errors.Init()

	l.SetInfo(info)
}

var info = plugins.Info{
	Website:     "http://www.go-admin.cn/plugins/detail/TzgE2RAYGqnCgEt7TK",
	Title:       "Librarian",
	Description: "Help you quickly build a online markdown viewer.",
	Version:     "v0.0.5",
	Author:      "Official",
	Url:         "https://github.com/GoAdminGroup/librarian/archive/v0.0.5.zip",
	Cover:       "",
	Agreement:   "",
	Uuid:        "TzgE2RAYGqnCgEt7TK",
	Name:        Name,
	ModulePath:  "github.com/GoAdminGroup/librarian",
	CreateDate:  utils.ParseTime("2020-04-19"),
	UpdateDate:  utils.ParseTime("2020-08-03"),
}

func (l *Librarian) AddRoot(key string, value root.Root) *Librarian {
	l.roots.Add(key, value)
	return l
}

func (l *Librarian) InitMenu() {
	for key, r := range *l.roots {
		navPath := r.Path + "/nav.yml"
		if util.FileExist(navPath) {
			buildMenus, err := l.siteTable().
				Where("key", "=", siteMenuIDsKey(key)).
				First()
			if db.CheckError(err, db.QUERY) {
				logger.Error("librarian build menu error: ", err)
				continue
			}

			checkNavContent, err := l.siteTable().
				Where("key", "=", siteMenuNavKey(key)).
				First()

			if db.CheckError(err, db.QUERY) {
				logger.Error("librarian check menu navs error: ", err)
				continue
			}

			b, err := ioutil.ReadFile(navPath)

			if err != nil {
				logger.Error("librarian check menu navs read files error: ", err)
				continue
			}

			m5 := md5.New()
			m5.Write(b)
			m5res := hex.EncodeToString(m5.Sum(nil))
			if checkNavContent != nil && m5res == checkNavContent["value"].(string) && buildMenus != nil {
				continue
			}

			if buildMenus == nil {
				if err := l.setMenu(b, m5res, key, navPath, false, checkNavContent != nil); err != nil {
					logger.Error("librarian set menu error: ", err)
				}
			} else {

				// clear old menu
				buildMenuIDs := strings.Split(buildMenus["value"].(string), ",")
				buildMenuIDInterfaces := make([]interface{}, len(buildMenuIDs))
				for i := 0; i < len(buildMenuIDs); i++ {
					buildMenuIDInterfaces[i] = buildMenuIDs[i]
				}
				err = l.menuTable().WhereIn("id", buildMenuIDInterfaces).Delete()
				if db.CheckError(err, db.DELETE) {
					logger.Error("librarian clear menu error: ", err)
					continue
				}
				err = l.roleMenuTable().WhereIn("menu_id", buildMenuIDInterfaces).Delete()
				if db.CheckError(err, db.DELETE) {
					logger.Error("librarian clear role menu error: ", err)
					continue
				}
				if err := l.setMenu(b, m5res, key, navPath, true, checkNavContent != nil); err != nil {
					logger.Error("librarian set menu error: ", err)
				}
			}
		}
	}
}

// TODO: add transaction
func (l *Librarian) setMenu(b []byte, m5Str string, prefix, navPath string, has, has2 bool) error {

	var navs = make(map[string]interface{})

	err := yaml.Unmarshal(b, &navs)

	if err != nil {
		return err
	}

	maxOrderMenu, err := l.menuTable().Select("order").OrderBy("order", "desc").First()
	if db.CheckError(err, db.QUERY) {
		logger.Error("librarian insert menu error: ", err)
		return err
	}
	order := int64(1)
	if o, ok := maxOrderMenu["order"].(int64); ok {
		order = o
	}
	ids := make([]string, 0)

	for _, level1 := range navs["nav"].([]interface{}) {
		for key, value := range level1.(map[interface{}]interface{}) {
			if level2, ok := value.([]interface{}); ok {
				level1NavID, err := l.NewMenu(menu.NewMenuData{
					ParentId:   0,
					Order:      order,
					Title:      key.(string),
					Icon:       icon.FileO,
					PluginName: l.Name(),
					Uri:        "",
				})
				if db.CheckError(err, db.INSERT) {
					logger.Error("librarian insert menu error: ", err)
					return err
				}
				ids = append(ids, strconv.Itoa(int(level1NavID)))
				order++
				for _, level2Nav := range level2 {
					for key, value := range level2Nav.(map[interface{}]interface{}) {
						if level3, ok := value.([]interface{}); ok {
							level2NavID, err := l.NewMenu(menu.NewMenuData{
								ParentId:   level1NavID,
								Order:      order,
								Title:      key.(string),
								Icon:       icon.FileO,
								PluginName: l.Name(),
								Uri:        "",
							})
							if db.CheckError(err, db.INSERT) {
								logger.Error("librarian insert menu error: ", err)
								return err
							}
							ids = append(ids, strconv.Itoa(int(level2NavID)))
							order++
							for _, level3Nav := range level3 {
								for key, value := range level3Nav.(map[interface{}]interface{}) {
									// third level
									id, err := l.NewMenu(menu.NewMenuData{
										ParentId:   level2NavID,
										Order:      order,
										Title:      key.(string),
										Icon:       icon.FileO,
										PluginName: l.Name(),
										Uri:        l.menuPath(prefix, value),
									})
									if db.CheckError(err, db.INSERT) {
										logger.Error("librarian insert menu error: ", err)
										return err
									}
									ids = append(ids, strconv.Itoa(int(id)))
									order++
								}
							}
						} else {
							// second level
							id, err := l.NewMenu(menu.NewMenuData{
								ParentId:   level1NavID,
								Order:      order,
								Title:      key.(string),
								Icon:       icon.FileO,
								PluginName: l.Name(),
								Uri:        l.menuPath(prefix, value),
							})
							if db.CheckError(err, db.INSERT) {
								logger.Error("librarian insert menu error: ", err)
								return err
							}
							ids = append(ids, strconv.Itoa(int(id)))
							order++
						}
					}
				}
			} else {
				// first level
				id, err := l.NewMenu(menu.NewMenuData{
					ParentId:   0,
					Order:      order,
					Title:      key.(string),
					Icon:       icon.FileO,
					PluginName: l.Name(),
					Uri:        l.menuPath(prefix, value),
				})
				if db.CheckError(err, db.INSERT) {
					logger.Error("librarian insert menu error: ", err)
					return err
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
				logger.Error("librarian insert menu error: ", err)
				return err
			}
		} else {
			_, err := l.siteTable().Insert(dialect.H{
				"key":   siteMenuIDsKey(prefix),
				"value": strings.Join(ids, ","),
			})
			if db.CheckError(err, db.UPDATE) {
				logger.Error("librarian insert menu error: ", err)
				return err
			}
		}
		if has2 {
			_, err := l.siteTable().Where("key", "=", siteMenuNavKey(prefix)).
				Update(dialect.H{
					"value": m5Str,
				})
			if db.CheckError(err, db.INSERT) {
				logger.Error("librarian insert menu error: ", err)
				return err
			}
		} else {
			_, err := l.siteTable().Insert(dialect.H{
				"key":   siteMenuNavKey(prefix),
				"value": m5Str,
			})
			if db.CheckError(err, db.UPDATE) {
				logger.Error("librarian insert menu error: ", err)
				return err
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
				logger.Error("librarian insert menu error: ", err)
				return err
			}
		}
	}

	return nil
}

type Menu struct {
	Name string
	Path string
}

func (l *Librarian) GetFirstMenu() Menu {
	buildMenus, err := l.siteTable().
		Where("key", "=", siteMenuIDsKey("def")).
		First()
	if db.CheckError(err, db.QUERY) {
		logger.Error("librarian get first menu id error: ", err)
		return Menu{}
	}
	if buildMenus == nil {
		return Menu{}
	}
	firstMenu, err := l.menuTable().Find(strings.Split(buildMenus["value"].(string), ",")[0])
	if db.CheckError(err, db.QUERY) {
		logger.Error("librarian get first menu error: ", err)
		return Menu{}
	}
	return Menu{
		Name: firstMenu["title"].(string),
		Path: firstMenu["uri"].(string),
	}
}

func (l *Librarian) GetSettingPage() table.Generator {
	return func(ctx *context.Context) (fileManagerConfiguration table.Table) {

		cfg := table.DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver)

		message1 := "install"
		message2 := "installation"

		if !l.IsInstalled() {
			cfg = cfg.SetOnlyNewForm()
		} else {

			message1 = "update"
			message2 = "setting"

			cfg = cfg.SetOnlyUpdateForm().SetGetDataFun(func(params parameter.Parameters) ([]map[string]interface{}, int) {

				var m = make([]map[string]interface{}, 1)

				checkRootExist, err := l.siteTable().
					Where("key", "=", "librarian_roots").
					First()

				if db.CheckError(err, db.QUERY) {
					return m, 0
				}

				if checkRootExist == nil {
					return m, 0
				}

				var rootMap = make(root.Roots)

				if err = json.Unmarshal([]byte(checkRootExist["value"].(string)), &rootMap); err != nil {
					return m, 0
				}

				m[0] = make(map[string]interface{})

				names, titles, paths := make([]string, 0), make([]string, 0), make([]string, 0)

				for name, value := range rootMap {
					names = append(names, name)
					titles = append(titles, value.Title)
					paths = append(paths, value.Path)
				}

				m[0]["id"] = "1"
				m[0]["name"] = strings.Join(names, ",")
				m[0]["title"] = strings.Join(titles, ",")
				m[0]["path"] = strings.Join(paths, ",")

				checkThemeExist, err := l.siteTable().
					Where("key", "=", "librarian_theme").
					First()

				if db.CheckError(err, db.QUERY) {
					return m, 0
				}

				if checkThemeExist == nil {
					return m, 0
				}

				m[0]["theme"] = checkThemeExist["value"]
				m[0]["build_menu"] = "0"

				return m, 1
			})
		}

		fileManagerConfiguration = table.NewDefaultTable(cfg)

		formList := fileManagerConfiguration.GetForm().
			AddXssJsFilter().
			HideBackButton().
			HideContinueNewCheckBox().
			HideResetButton()

		formList.AddField(language2.Get("build menu"), "build_menu", db.Varchar, form.Switch).
			FieldOptions(types.FieldOptions{
				{Value: "1", Text: language2.Get("yes")},
				{Value: "0", Text: language2.Get("no")},
			}).FieldDefault("0")

		var ops = make(types.FieldOptions, len(theme.AllThemes))
		for k, th := range theme.AllThemes {
			ops[k] = types.FieldOption{Value: th, Text: language2.Get(th)}
		}

		formList.AddField(language2.Get("theme"), "theme", db.Varchar, form.SelectSingle).
			FieldOptions(ops).FieldDefault("github")

		formList.AddTable(language2.Get("roots"), "roots", func(panel *types.FormPanel) {
			panel.AddField(language2.Get("name"), "name", db.Varchar, form.Text).FieldHideLabel().
				FieldDisplay(func(value types.FieldModel) interface{} {
					return strings.Split(value.Value, ",")
				})
			panel.AddField(language2.Get("title"), "title", db.Varchar, form.Text).FieldHideLabel().
				FieldDisplay(func(value types.FieldModel) interface{} {
					return strings.Split(value.Value, ",")
				})
			panel.AddField(language2.Get("path"), "path", db.Varchar, form.Text).FieldHideLabel().
				FieldDisplay(func(value types.FieldModel) interface{} {
					return strings.Split(value.Value, ",")
				})
		})

		var updateInsertFn = func(values form2.Values) error {

			var rootsMap = make(root.Roots, len(values["name"]))
			for k, name := range values["name"] {
				rootsMap[name] = root.Root{
					Path:  values["path"][k],
					Title: values["title"][k],
				}
			}
			roots, err := json.Marshal(rootsMap)

			if err != nil {
				return err
			}

			checkExist, err := l.siteTable().
				Where("key", "=", "librarian_roots").
				First()

			if db.CheckError(err, db.QUERY) {
				return err
			}

			if checkExist != nil {
				_, _ = l.siteTable().
					Where("key", "=", "librarian_roots").
					Update(dialect.H{
						"value": string(roots),
					})
			} else {
				_, _ = l.siteTable().
					Insert(dialect.H{
						"value": string(roots),
						"key":   "librarian_roots",
					})
			}

			checkThemeExist, err := l.siteTable().
				Where("key", "=", "librarian_theme").
				First()

			if db.CheckError(err, db.QUERY) {
				return err
			}

			if checkThemeExist != nil {
				_, _ = l.siteTable().
					Where("key", "=", "librarian_theme").
					Update(dialect.H{
						"value": values.Get("theme"),
					})
			} else {
				_, _ = l.siteTable().
					Insert(dialect.H{
						"value": values.Get("theme"),
						"key":   "librarian_theme",
					})
			}

			l.roots = &rootsMap
			l.handler.Update(l.roots, values.Get("theme"))
			l.guard.Update(l.roots)

			if values.Get("build_menu") == "1" {
				l.InitMenu()
			}

			return nil
		}

		formList.SetInsertFn(updateInsertFn)
		formList.SetUpdateFn(updateInsertFn)

		formList.EnableAjaxData(types.AjaxData{
			SuccessTitle:   language2.Get(message1 + " success, please restart"),
			ErrorTitle:     language2.Get(message1 + " fail"),
			SuccessJumpURL: config.Url("/info/plugin_librarian/new"),
		}).SetFormNewTitle(language2.GetHTML("librarian " + message2)).
			SetTitle(language2.Get("librarian " + message2)).
			SetFormNewBtnWord(language2.GetHTML(message1))

		return
	}
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

func siteMenuNavKey(prefix string) string {
	return "librarian_build_menu_" + prefix + "_nav"
}

func (l *Librarian) menuPath(prefix string, path interface{}) string {
	p := strings.Replace(path.(string), ".md", "", -1)
	if prefix == "def" {
		if l.prefix != "" {
			return "/" + l.prefix + "/" + p
		}
		return "/" + p
	}
	if l.prefix != "" {
		return "/" + l.prefix + "/" + p + "?__prefix=" + prefix
	}
	return "/" + p + "?__prefix=" + prefix
}
