package theme

import "html/template"

type Theme interface {
	HTML(md []byte) template.HTML
	CSS() template.CSS
	JS() template.JS
}

type Base struct{}

func (base *Base) HTML(md []byte) template.HTML { return "" }
func (base *Base) CSS() template.CSS            { return "" }
func (base *Base) JS() template.JS              { return "" }

var themes = map[string]Theme{
	"default": new(Default),
}

func Get(name string) Theme {
	return themes[name]
}
