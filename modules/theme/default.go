package theme

import (
	"html/template"
)

type Default struct {
	Base
}

func (*Default) JS() template.JS {
	return `
window.addEventListener("scroll", function () {
	// var scroH = $(document).scrollTop();
	console.log("scroll....");
});
let titleH1 = $(".markdown-viewer h1");
if (titleH1.length > 0 && $(titleH1[0]).text() !== "") {
	document.title = $(titleH1[0]).text();
}
`
}

func (d *Default) CSS() template.CSS {

	return d.Base.CSS() + `
.markdown-viewer table {
  	padding: 0;
  	word-break: initial;
	margin-bottom: 13px;
}
.markdown-viewer table tr {
  border-top: 1px solid #dadfe6;
  margin: 0;
  padding: 0
}
.markdown-viewer table.md-table tr:nth-child(2n) {
  background-color: #fafbfc
}
.markdown-viewer table tr th {
  font-weight: 400;
  border: 1px solid #dadfe6;
  text-align: left;
  margin: 0;
  padding: 6px 13px
}
.markdown-viewer table tr td {
  border: 1px solid #dadfe6;
  text-align: left;
  margin: 0;
  padding: 6px 13px
}
.markdown-viewer table tr td:first-child,
.markdown-viewer table tr th:first-child {
  margin-top: 0
}
.markdown-viewer table tr td:last-child,
.markdown-viewer table tr th:last-child {
  margin-bottom: 0
}
`
}
