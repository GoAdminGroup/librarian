package language

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"html/template"
)

func Get(key string) string {
	return language.GetWithScope(key, "librarian")
}

func GetHTML(key string) template.HTML {
	return template.HTML(language.GetWithScope(key, "librarian"))
}
