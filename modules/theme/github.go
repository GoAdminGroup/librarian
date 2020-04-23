package theme

import "html/template"

type Github struct {
	Base
}

func (g *Github) CSS() template.CSS {
	return g.Base.CSS() + `
.markdown-viewer {
	padding-top: 30px;
}
.markdown-viewer .octicon {
  display: inline-block;
  fill: currentColor;
  vertical-align: text-bottom;
}

.markdown-viewer .anchor {
  float: left;
  line-height: 1;
  margin-left: -20px;
  padding-right: 4px;
}

.markdown-viewer .anchor:focus {
  outline: none;
}

.markdown-viewer blockquote {
	font-size: 1.5rem;
}

.markdown-viewer h1 .octicon-link,
.markdown-viewer h2 .octicon-link,
.markdown-viewer h3 .octicon-link,
.markdown-viewer h4 .octicon-link,
.markdown-viewer h5 .octicon-link,
.markdown-viewer h6 .octicon-link {
  color: #1b1f23;
  vertical-align: middle;
  visibility: hidden;
}

.markdown-viewer h1:hover .anchor,
.markdown-viewer h2:hover .anchor,
.markdown-viewer h3:hover .anchor,
.markdown-viewer h4:hover .anchor,
.markdown-viewer h5:hover .anchor,
.markdown-viewer h6:hover .anchor {
  text-decoration: none;
}

.markdown-viewer h1:hover .anchor .octicon-link,
.markdown-viewer h2:hover .anchor .octicon-link,
.markdown-viewer h3:hover .anchor .octicon-link,
.markdown-viewer h4:hover .anchor .octicon-link,
.markdown-viewer h5:hover .anchor .octicon-link,
.markdown-viewer h6:hover .anchor .octicon-link {
  visibility: visible;
}

.markdown-viewer h1:hover .anchor .octicon-link:before,
.markdown-viewer h2:hover .anchor .octicon-link:before,
.markdown-viewer h3:hover .anchor .octicon-link:before,
.markdown-viewer h4:hover .anchor .octicon-link:before,
.markdown-viewer h5:hover .anchor .octicon-link:before,
.markdown-viewer h6:hover .anchor .octicon-link:before {
  width: 16px;
  height: 16px;
  content: ' ';
  display: inline-block;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' version='1.1' width='16' height='16' aria-hidden='true'%3E%3Cpath fill-rule='evenodd' d='M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z'%3E%3C/path%3E%3C/svg%3E");
}.markdown-viewer {
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  line-height: 1.5;
  color: #24292e;
  font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji;
  font-size: 16px;
  line-height: 1.5;
  word-wrap: break-word;
}

.markdown-viewer details {
  display: block;
}

.markdown-viewer summary {
  display: list-item;
}

.markdown-viewer a {
  background-color: initial;
}

.markdown-viewer a:active,
.markdown-viewer a:hover {
  outline-width: 0;
}

.markdown-viewer strong {
  font-weight: inherit;
  font-weight: bolder;
}

.markdown-viewer h1 {
  font-size: 2em;
  margin: .67em 0;
}

.markdown-viewer img {
  border-style: none;
}

.markdown-viewer code,
.markdown-viewer kbd,
.markdown-viewer pre {
  font-family: monospace,monospace;
  font-size: 1em;
}

.markdown-viewer hr {
  box-sizing: initial;
  height: 0;
  overflow: visible;
}

.markdown-viewer input {
  font: inherit;
  margin: 0;
}

.markdown-viewer input {
  overflow: visible;
}

.markdown-viewer [type=checkbox] {
  box-sizing: border-box;
  padding: 0;
}

.markdown-viewer * {
  box-sizing: border-box;
}

.markdown-viewer input {
  font-family: inherit;
  font-size: inherit;
  line-height: inherit;
}

.markdown-viewer a {
  color: #0366d6;
  text-decoration: none;
}

.markdown-viewer a:hover {
  text-decoration: underline;
}

.markdown-viewer strong {
  font-weight: 600;
}

.markdown-viewer hr {
  height: 0;
  margin: 15px 0;
  overflow: hidden;
  background: transparent;
  border: 0;
  border-bottom: 1px solid #dfe2e5;
}

.markdown-viewer hr:after,
.markdown-viewer hr:before {
  display: table;
  content: "";
}

.markdown-viewer hr:after {
  clear: both;
}

.markdown-viewer table {
  border-spacing: 0;
  border-collapse: collapse;
}

.markdown-viewer td,
.markdown-viewer th {
  padding: 0;
}

.markdown-viewer details summary {
  cursor: pointer;
}

.markdown-viewer kbd {
  display: inline-block;
  padding: 3px 5px;
  font: 11px SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
  line-height: 10px;
  color: #444d56;
  vertical-align: middle;
  background-color: #fafbfc;
  border: 1px solid #d1d5da;
  border-radius: 3px;
  box-shadow: inset 0 -1px 0 #d1d5da;
}

.markdown-viewer h1,
.markdown-viewer h2,
.markdown-viewer h3,
.markdown-viewer h4,
.markdown-viewer h5,
.markdown-viewer h6 {
  margin-top: 0;
  margin-bottom: 0;
}

.markdown-viewer h1 {
  font-size: 32px;
}

.markdown-viewer h1,
.markdown-viewer h2 {
  font-weight: 600;
}

.markdown-viewer h2 {
  font-size: 24px;
}

.markdown-viewer h3 {
  font-size: 20px;
}

.markdown-viewer h3,
.markdown-viewer h4 {
  font-weight: 600;
}

.markdown-viewer h4 {
  font-size: 16px;
}

.markdown-viewer h5 {
  font-size: 14px;
}

.markdown-viewer h5,
.markdown-viewer h6 {
  font-weight: 600;
}

.markdown-viewer h6 {
  font-size: 12px;
}

.markdown-viewer p {
  margin-top: 0;
  margin-bottom: 10px;
}

.markdown-viewer blockquote {
  margin: 0;
}

.markdown-viewer ol,
.markdown-viewer ul {
  padding-left: 0;
  margin-top: 0;
  margin-bottom: 0;
}

.markdown-viewer ol ol,
.markdown-viewer ul ol {
  list-style-type: lower-roman;
}

.markdown-viewer ol ol ol,
.markdown-viewer ol ul ol,
.markdown-viewer ul ol ol,
.markdown-viewer ul ul ol {
  list-style-type: lower-alpha;
}

.markdown-viewer dd {
  margin-left: 0;
}

.markdown-viewer code,
.markdown-viewer pre {
  font-family: SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
  font-size: 12px;
}

.markdown-viewer pre {
  margin-top: 0;
  margin-bottom: 0;
}

.markdown-viewer input::-webkit-inner-spin-button,
.markdown-viewer input::-webkit-outer-spin-button {
  margin: 0;
  -webkit-appearance: none;
  appearance: none;
}

.markdown-viewer :checked+.radio-label {
  position: relative;
  z-index: 1;
  border-color: #0366d6;
}

.markdown-viewer .border {
  border: 1px solid #e1e4e8!important;
}

.markdown-viewer .border-0 {
  border: 0!important;
}

.markdown-viewer .border-bottom {
  border-bottom: 1px solid #e1e4e8!important;
}

.markdown-viewer .rounded-1 {
  border-radius: 3px!important;
}

.markdown-viewer .bg-white {
  background-color: #fff!important;
}

.markdown-viewer .bg-gray-light {
  background-color: #fafbfc!important;
}

.markdown-viewer .text-gray-light {
  color: #6a737d!important;
}

.markdown-viewer .mb-0 {
  margin-bottom: 0!important;
}

.markdown-viewer .my-2 {
  margin-top: 8px!important;
  margin-bottom: 8px!important;
}

.markdown-viewer .pl-0 {
  padding-left: 0!important;
}

.markdown-viewer .py-0 {
  padding-top: 0!important;
  padding-bottom: 0!important;
}

.markdown-viewer .pl-1 {
  padding-left: 4px!important;
}

.markdown-viewer .pl-2 {
  padding-left: 8px!important;
}

.markdown-viewer .py-2 {
  padding-top: 8px!important;
  padding-bottom: 8px!important;
}

.markdown-viewer .pl-3,
.markdown-viewer .px-3 {
  padding-left: 16px!important;
}

.markdown-viewer .px-3 {
  padding-right: 16px!important;
}

.markdown-viewer .pl-4 {
  padding-left: 24px!important;
}

.markdown-viewer .pl-5 {
  padding-left: 32px!important;
}

.markdown-viewer .pl-6 {
  padding-left: 40px!important;
}

.markdown-viewer .f6 {
  font-size: 12px!important;
}

.markdown-viewer .lh-condensed {
  line-height: 1.25!important;
}

.markdown-viewer .text-bold {
  font-weight: 600!important;
}

.markdown-viewer .pl-c {
  color: #6a737d;
}

.markdown-viewer .pl-c1,
.markdown-viewer .pl-s .pl-v {
  color: #005cc5;
}

.markdown-viewer .pl-e,
.markdown-viewer .pl-en {
  color: #6f42c1;
}

.markdown-viewer .pl-s .pl-s1,
.markdown-viewer .pl-smi {
  color: #24292e;
}

.markdown-viewer .pl-ent {
  color: #22863a;
}

.markdown-viewer .pl-k {
  color: #d73a49;
}

.markdown-viewer .pl-pds,
.markdown-viewer .pl-s,
.markdown-viewer .pl-s .pl-pse .pl-s1,
.markdown-viewer .pl-sr,
.markdown-viewer .pl-sr .pl-cce,
.markdown-viewer .pl-sr .pl-sra,
.markdown-viewer .pl-sr .pl-sre {
  color: #032f62;
}

.markdown-viewer .pl-smw,
.markdown-viewer .pl-v {
  color: #e36209;
}

.markdown-viewer .pl-bu {
  color: #b31d28;
}

.markdown-viewer .pl-ii {
  color: #fafbfc;
  background-color: #b31d28;
}

.markdown-viewer .pl-c2 {
  color: #fafbfc;
  background-color: #d73a49;
}

.markdown-viewer .pl-c2:before {
  content: "^M";
}

.markdown-viewer .pl-sr .pl-cce {
  font-weight: 700;
  color: #22863a;
}

.markdown-viewer .pl-ml {
  color: #735c0f;
}

.markdown-viewer .pl-mh,
.markdown-viewer .pl-mh .pl-en,
.markdown-viewer .pl-ms {
  font-weight: 700;
  color: #005cc5;
}

.markdown-viewer .pl-mi {
  font-style: italic;
  color: #24292e;
}

.markdown-viewer .pl-mb {
  font-weight: 700;
  color: #24292e;
}

.markdown-viewer .pl-md {
  color: #b31d28;
  background-color: #ffeef0;
}

.markdown-viewer .pl-mi1 {
  color: #22863a;
  background-color: #f0fff4;
}

.markdown-viewer .pl-mc {
  color: #e36209;
  background-color: #ffebda;
}

.markdown-viewer .pl-mi2 {
  color: #f6f8fa;
  background-color: #005cc5;
}

.markdown-viewer .pl-mdr {
  font-weight: 700;
  color: #6f42c1;
}

.markdown-viewer .pl-ba {
  color: #586069;
}

.markdown-viewer .pl-sg {
  color: #959da5;
}

.markdown-viewer .pl-corl {
  text-decoration: underline;
  color: #032f62;
}

.markdown-viewer .mb-0 {
  margin-bottom: 0!important;
}

.markdown-viewer .my-2 {
  margin-bottom: 8px!important;
}

.markdown-viewer .my-2 {
  margin-top: 8px!important;
}

.markdown-viewer .pl-0 {
  padding-left: 0!important;
}

.markdown-viewer .py-0 {
  padding-top: 0!important;
  padding-bottom: 0!important;
}

.markdown-viewer .pl-1 {
  padding-left: 4px!important;
}

.markdown-viewer .pl-2 {
  padding-left: 8px!important;
}

.markdown-viewer .py-2 {
  padding-top: 8px!important;
  padding-bottom: 8px!important;
}

.markdown-viewer .pl-3 {
  padding-left: 16px!important;
}

.markdown-viewer .pl-4 {
  padding-left: 24px!important;
}

.markdown-viewer .pl-5 {
  padding-left: 32px!important;
}

.markdown-viewer .pl-6 {
  padding-left: 40px!important;
}

.markdown-viewer .pl-7 {
  padding-left: 48px!important;
}

.markdown-viewer .pl-8 {
  padding-left: 64px!important;
}

.markdown-viewer .pl-9 {
  padding-left: 80px!important;
}

.markdown-viewer .pl-10 {
  padding-left: 96px!important;
}

.markdown-viewer .pl-11 {
  padding-left: 112px!important;
}

.markdown-viewer .pl-12 {
  padding-left: 128px!important;
}

.markdown-viewer hr {
  border-bottom-color: #eee;
}

.markdown-viewer kbd {
  display: inline-block;
  padding: 3px 5px;
  font: 11px SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
  line-height: 10px;
  color: #444d56;
  vertical-align: middle;
  background-color: #fafbfc;
  border: 1px solid #d1d5da;
  border-radius: 3px;
  box-shadow: inset 0 -1px 0 #d1d5da;
}

.markdown-viewer:after,
.markdown-viewer:before {
  display: table;
  content: "";
}

.markdown-viewer:after {
  clear: both;
}

.markdown-viewer>:first-child {
  margin-top: 0!important;
}

.markdown-viewer>:last-child {
  margin-bottom: 0!important;
}

.markdown-viewer a:not([href]) {
  color: inherit;
  text-decoration: none;
}

.markdown-viewer blockquote,
.markdown-viewer details,
.markdown-viewer dl,
.markdown-viewer ol,
.markdown-viewer p,
.markdown-viewer pre,
.markdown-viewer table,
.markdown-viewer ul {
  margin-top: 0;
  margin-bottom: 16px;
}

.markdown-viewer hr {
  height: .25em;
  padding: 0;
  margin: 24px 0;
  background-color: #e1e4e8;
  border: 0;
}

.markdown-viewer blockquote {
  padding: 0 1em;
  color: #6a737d;
  border-left: .25em solid #dfe2e5;
}

.markdown-viewer blockquote>:first-child {
  margin-top: 0;
}

.markdown-viewer blockquote>:last-child {
  margin-bottom: 0;
}

.markdown-viewer h1,
.markdown-viewer h2,
.markdown-viewer h3,
.markdown-viewer h4,
.markdown-viewer h5,
.markdown-viewer h6 {
  margin-top: 24px;
  margin-bottom: 16px;
  font-weight: 600;
  line-height: 1.25;
}

.markdown-viewer h1 {
  font-size: 2em;
}

.markdown-viewer h1,
.markdown-viewer h2 {
  padding-bottom: .3em;
  border-bottom: 1px solid #eaecef;
}

.markdown-viewer h2 {
  font-size: 1.5em;
}

.markdown-viewer h3 {
  font-size: 1.25em;
}

.markdown-viewer h4 {
  font-size: 1em;
}

.markdown-viewer h5 {
  font-size: .875em;
}

.markdown-viewer h6 {
  font-size: .85em;
  color: #6a737d;
}

.markdown-viewer ol,
.markdown-viewer ul {
  padding-left: 2em;
}

.markdown-viewer ol ol,
.markdown-viewer ol ul,
.markdown-viewer ul ol,
.markdown-viewer ul ul {
  margin-top: 0;
  margin-bottom: 0;
}

.markdown-viewer li {
  word-wrap: break-all;
}

.markdown-viewer li>p {
  margin-top: 16px;
}

.markdown-viewer li+li {
  margin-top: .25em;
}

.markdown-viewer dl {
  padding: 0;
}

.markdown-viewer dl dt {
  padding: 0;
  margin-top: 16px;
  font-size: 1em;
  font-style: italic;
  font-weight: 600;
}

.markdown-viewer dl dd {
  padding: 0 16px;
  margin-bottom: 16px;
}

.markdown-viewer table {
  display: block;
  width: 100%;
  overflow: auto;
}

.markdown-viewer table th {
  font-weight: 600;
}

.markdown-viewer table td,
.markdown-viewer table th {
  padding: 6px 13px;
  border: 1px solid #dfe2e5;
}

.markdown-viewer table tr {
  background-color: #fff;
  border-top: 1px solid #c6cbd1;
}

.markdown-viewer table tr:nth-child(2n) {
  background-color: #f6f8fa;
}

.markdown-viewer img {
  max-width: 100%;
  box-sizing: initial;
  background-color: #fff;
}

.markdown-viewer img[align=right] {
  padding-left: 20px;
}

.markdown-viewer img[align=left] {
  padding-right: 20px;
}

.markdown-viewer code {
  padding: .2em .4em;
  margin: 0;
  font-size: 85%;
  background-color: rgba(27,31,35,.05);
  border-radius: 3px;
}

.markdown-viewer pre {
  word-wrap: normal;
}

.markdown-viewer pre>code {
  padding: 0;
  margin: 0;
  font-size: 100%;
  word-break: normal;
  white-space: pre;
  background: transparent;
  border: 0;
}

.markdown-viewer .highlight {
  margin-bottom: 16px;
}

.markdown-viewer .highlight pre {
  margin-bottom: 0;
  word-break: normal;
}

.markdown-viewer .highlight pre,
.markdown-viewer pre {
  padding: 16px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
  background-color: #f6f8fa;
  border-radius: 3px;
}

.markdown-viewer pre code {
  display: inline;
  max-width: auto;
  padding: 0;
  margin: 0;
  overflow: visible;
  line-height: inherit;
  word-wrap: normal;
  background-color: initial;
  border: 0;
}

.markdown-viewer .commit-tease-sha {
  display: inline-block;
  font-family: SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
  font-size: 90%;
  color: #444d56;
}

.markdown-viewer .full-commit .btn-outline:not(:disabled):hover {
  color: #005cc5;
  border-color: #005cc5;
}

.markdown-viewer .blob-wrapper {
  overflow-x: auto;
  overflow-y: hidden;
}

.markdown-viewer .blob-wrapper-embedded {
  max-height: 240px;
  overflow-y: auto;
}

.markdown-viewer .blob-num {
  width: 1%;
  min-width: 50px;
  padding-right: 10px;
  padding-left: 10px;
  font-family: SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
  font-size: 12px;
  line-height: 20px;
  color: rgba(27,31,35,.3);
  text-align: right;
  white-space: nowrap;
  vertical-align: top;
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.markdown-viewer .blob-num:hover {
  color: rgba(27,31,35,.6);
}

.markdown-viewer .blob-num:before {
  content: attr(data-line-number);
}

.markdown-viewer .blob-code {
  position: relative;
  padding-right: 10px;
  padding-left: 10px;
  line-height: 20px;
  vertical-align: top;
}

.markdown-viewer .blob-code-inner {
  overflow: visible;
  font-family: SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
  font-size: 12px;
  color: #24292e;
  word-wrap: normal;
  white-space: pre;
}

.markdown-viewer .pl-token.active,
.markdown-viewer .pl-token:hover {
  cursor: pointer;
  background: #ffea7f;
}

.markdown-viewer .tab-size[data-tab-size="1"] {
  -moz-tab-size: 1;
  tab-size: 1;
}

.markdown-viewer .tab-size[data-tab-size="2"] {
  -moz-tab-size: 2;
  tab-size: 2;
}

.markdown-viewer .tab-size[data-tab-size="3"] {
  -moz-tab-size: 3;
  tab-size: 3;
}

.markdown-viewer .tab-size[data-tab-size="4"] {
  -moz-tab-size: 4;
  tab-size: 4;
}

.markdown-viewer .tab-size[data-tab-size="5"] {
  -moz-tab-size: 5;
  tab-size: 5;
}

.markdown-viewer .tab-size[data-tab-size="6"] {
  -moz-tab-size: 6;
  tab-size: 6;
}

.markdown-viewer .tab-size[data-tab-size="7"] {
  -moz-tab-size: 7;
  tab-size: 7;
}

.markdown-viewer .tab-size[data-tab-size="8"] {
  -moz-tab-size: 8;
  tab-size: 8;
}

.markdown-viewer .tab-size[data-tab-size="9"] {
  -moz-tab-size: 9;
  tab-size: 9;
}

.markdown-viewer .tab-size[data-tab-size="10"] {
  -moz-tab-size: 10;
  tab-size: 10;
}

.markdown-viewer .tab-size[data-tab-size="11"] {
  -moz-tab-size: 11;
  tab-size: 11;
}

.markdown-viewer .tab-size[data-tab-size="12"] {
  -moz-tab-size: 12;
  tab-size: 12;
}

.markdown-viewer .task-list-item {
  list-style-type: none;
}

.markdown-viewer .task-list-item+.task-list-item {
  margin-top: 3px;
}

.markdown-viewer .task-list-item input {
  margin: 0 .2em .25em -1.6em;
  vertical-align: middle;
}
`
}
