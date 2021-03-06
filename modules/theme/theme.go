package theme

import (
	"html/template"
	"regexp"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/librarian/modules/language"
)

type Theme interface {
	HTML(md []byte) template.HTML
	CSS() template.CSS
	JS() template.JS
}

type Base struct{}

func (base *Base) HTML(md []byte) template.HTML {

	reg, _ := regexp.Compile("<h2>(.*)</h2>")

	allHeader := reg.FindAllSubmatch(md, -1)

	boxContent := template.HTML("")

	mdstring := string(md)

	if len(allHeader) > 0 {
		boxContent = `
<div class="col-sm-3">
<div class="navigation-box">
<label class="navigation-title">` + language.GetHTML("Table of contents") + `</label><ul>
`
		toReplace := make([]string, len(allHeader)*2)
		for i := 0; i < len(allHeader); i++ {
			toReplace[i*2] = string(allHeader[i][0])
			is := strconv.Itoa(i)
			toReplace[i*2+1] = `<h2 id="_` + is + `">` + string(allHeader[i][1]) + "</h2>"
			boxContent += `<a href="#_` + template.HTML(is) + `"><li>` + template.HTML(allHeader[i][1]) + "</li></a>"
		}
		boxContent += "</ul></div></div>"
		replacer := strings.NewReplacer(toReplace...)
		mdstring = replacer.Replace(mdstring)
	}

	return `
<div class="markdown-wrapper">
	<div class="row markdown-row">
		<div class="col-sm-9"> 
			<div class="markdown-viewer">` + template.HTML(mdstring) + `</div>
		</div>
		` + boxContent + `
	</div>
</div>
`
}
func (base *Base) CSS() template.CSS {
	css := template.CSS("")

	if config.HideMenuIcon {
		css += `
.sidebar-menu .fa.fa-file-o  {
	display:none;
}
`
	}

	if config.HideNavBar {
		css += `
.navbar.navbar-static-top {
	display:none;
}
`
	}
	return css + `
.content {
    padding: 0px;
}
.markdown-wrapper {
	padding-top: 20px; 
	width: 100%;
	padding-bottom: 20px;
}
.markdown-row {
    width: 98%;
    margin: auto;
}
.markdown-viewer {
	margin: auto;
	padding: 10px 35px 20px 35px;
    background-color: #FFFFFF;
    min-height: 500px;
}
.navigation-title {
	display: block;
    padding: 0 .7rem;
    font-weight: 700;
    text-overflow: ellipsis;
    overflow: hidden;
	margin-bottom: 9px;
}
.markdown-viewer img {
	max-width: 100%;
}
.navigation-box {
    padding: 20px 6px 20px 6px;
    background-color: #fff;
}
.navigation-box ul {
	margin: 0;
    padding: 0;
    list-style: none;
}
.navigation-box ul a {
	color: #656565;
}
.navigation-box ul li {
    margin-bottom: 5px;
	padding: 0 .9rem;
}
.navigation-box ul li:hover {
    color: #4190ff;
}
@media screen and (max-height: 450px) {
	.navigation-box {
		display: none;
	}
}`
}

func (base *Base) JS() template.JS {
	js := template.JS(``)

	if config.ChangeTitle {
		js += template.JS(`
let titleH1 = $(".markdown-viewer h1");
if (titleH1.length > 0 && $(titleH1[0]).text() !== "") {
	document.title = $(titleH1[0]).text();
}
`)
	}

	if config.FixedSidebar {
		return js + `
$('.main-sidebar').css('position', 'fixed');
$('.main-header .logo').css('position', 'fixed');
`
	}
	return js
}

var themes = map[string]Theme{
	"default": new(Default),
	"github":  new(Github),
}

var AllThemes = []string{"default", "github"}

func Get(name string) Theme {
	return themes[name]
}

type Config struct {
	HideNavBar   bool
	HideMenuIcon bool
	FixedSidebar bool
	ChangeTitle  bool
}

var config Config

func Set(c Config) {
	config = c
}
