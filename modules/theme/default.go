package theme

import (
	"github.com/GoAdminGroup/librarian/modules/language"
	"html/template"
	"regexp"
	"strconv"
	"strings"
)

type Default struct {
	Base
}

func (*Default) HTML(md []byte) template.HTML {

	reg, _ := regexp.Compile("<h2>(.*)</h2>")

	allHeader := reg.FindAllSubmatch(md, -1)

	boxContent := template.HTML("")

	mdstring := string(md)

	if len(allHeader) > 0 {
		toReplace := make([]string, len(allHeader)*2)
		boxContent = `<label class="navigation-title">` + language.GetHTML("Table of contents") + `</label><ul>`
		for i := 0; i < len(allHeader); i++ {
			toReplace[i*2] = string(allHeader[i][0])
			is := strconv.Itoa(i)
			toReplace[i*2+1] = `<h2 id="_` + is + `">` + string(allHeader[i][1]) + "</h2>"
			boxContent += `<a href="#_` + template.HTML(is) + `"><li>` + template.HTML(allHeader[i][1]) + "</li></a>"
		}
		boxContent += "</ul>"
		replacer := strings.NewReplacer(toReplace...)
		mdstring = replacer.Replace(mdstring)
	}

	return `
<div class="markdown-wrapper">
	<div class="row markdown-row">
		<div class="col-sm-9"> 
			<div class="markdown-viewer">` + template.HTML(mdstring) + `</div>
		</div>
		<div class="col-sm-3">
			<div class="navigation-box">
				` + boxContent + `
			</div>
		</div>
	</div>
</div>
`
}

func (*Default) CSS() template.CSS {
	return `
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
	padding: 20px;
    background-color: #FFFFFF;
    min-height: 500px;
}
.navigation-title {
	display: block;
    padding: 0 .7rem;
    font-weight: 700;
    text-overflow: ellipsis;
    overflow: hidden;
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
	color: black;
}
.navigation-box ul li {
    margin-bottom: 5px;
	padding: 0 .8rem;
}
@media screen and (max-height: 450px) {
	.navigation-box {
		display: none;
	}
}
`
}
