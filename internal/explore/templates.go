package explore

import (
	"fmt"
	"text/template"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

var (
	headerTmpl *template.Template
	bodyTmpl   *template.Template
	oauthTmpl  *template.Template
)

func init() {
	headerTmpl = template.Must(template.New("headerTemplate").Parse(headerTemplate))
	bodyTmpl = template.Must(template.New("bodyTemplate").Parse(bodyTemplate))
	oauthTmpl = template.Must(template.New("oauthTemplate").Parse(oauthTemplate))
}

const (
	gcrane     = `<a class="mt" href="https://github.com/google/go-containerregistry/blob/main/cmd/gcrane/README.md">gcrane</a>`
	craneLink  = `<a class="mt" href="https://github.com/google/go-containerregistry/blob/main/cmd/crane/README.md">crane</a>`
	subLinkFmt = `<a class="mt" href="https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane_%s.md">%s</a>`
)

func crane(sub string) string {
	if sub == "" {
		return craneLink
	}

	subLink := fmt.Sprintf(subLinkFmt, sub, sub)
	return craneLink + " " + subLink
}

const (
	landingPage = `
<html>
<body>
<head>
<title>Registry Explorer</title>
<link rel="icon" href="/favicon.svg">
<style>
.mt:hover {
	text-decoration: underline;
}

.mt {
	color: inherit;
	text-decoration: inherit;
}

.link {
	position: relative;
	bottom: .125em;
}

.crane {
	height: 1em;
	width: 1em;
}

.top {
	color: inherit;
	text-decoration: inherit;
}

:root {
  color-scheme: light dark;
}

body {
	font-family: monospace;
	width: fit-content;
	overflow-wrap: anywhere;
	padding: 12px;
}

</style>
</head>
<h1><a class="top" href="/"><img class="crane" src="/favicon.svg"/> <span class="link">Registry Explorer</span></a></h1>
<p>
This beautiful tool allows you to <em>explore</em> the contents of a registry interactively.
</p>
<p>
You can even drill down into layers to explore an image's filesystem.
</p>
<p>
Enter a <strong>public</strong> image, e.g. <tt>"ubuntu:latest"</tt>:
</p>
<form action="/" method="GET" autocomplete="off" spellcheck="false">
<input size="100" type="text" name="image" value="ubuntu:latest"/>
<input type="submit" />
</form>
<p>
<p>
Enter a <strong>public</strong> repository, e.g. <tt>"ubuntu"</tt>:
</p>
<form action="/" method="GET" autocomplete="off" spellcheck="false">
<input size="100" type="text" name="repo" value="ubuntu"/>
<input type="submit" />
</form>
<p>
<h4>Interesting examples</h4>
<ul>
  <li><a href="/?image=cgr.dev/chainguard/static:latest-glibc">cgr.dev/chainguard/static:latest-glibc</a></li>
  <li><a href="/?image=gcr.io/distroless/static">gcr.io/distroless/static:latest</a></li>
  <li><a href="/?repo=ghcr.io/homebrew/core/crane">ghcr.io/homebrew/core/crane</a></li>
  <li><a href="/?repo=registry.k8s.io">registry.k8s.io</a></li>
  <li><a href="/?image=registry.k8s.io/bom/bom:sha256-499bdf4cc0498bbfb2395f8bbaf3b7e9e407cca605aecc46b2ef1b390a0bc4c4.sig">registry.k8s.io/bom/bom:sha256-499bdf4cc0498bbfb2395f8bbaf3b7e9e407cca605aecc46b2ef1b390a0bc4c4.sig</a></li>
  <li><a href="/?image=docker/dockerfile:1.5.1">docker/dockerfile:1.5.1</a></li>
  <li><a href="/?image=pengfeizhou/test-oci:sha256-04eaff953b0066d7e4ea2e822eb5c31be0742fca494561336f0912fabc246760">pengfeizhou/test-oci:sha256-04eaff953b0066d7e4ea2e822eb5c31be0742fca494561336f0912fabc246760</a></li>
  <li><a href="/?image=tianon/true:oci">tianon/true:oci</a></li>
  <li><a href="/?image=ghcr.io/stargz-containers/node:13.13.0-esgz">ghcr.io/stargz-containers/node:13.13.0-esgz</a></li>

</ul>
</p>
<h3>FAQ</h3>
<h4>How does this work?</h4>
<p>
This service lives on <a href="https://cloud.run">Cloud Run</a> and uses <a href="https://github.com/google/go-containerregistry">google/go-containerregistry</a> for registry interactions.
</p>
<h4>Isn't this expensive for the registry?</h4>
<p>
Not really! The first time a layer is accessed, I download and index it. Browsing the filesystem just uses that index, and opening a file uses Range requests to load small chunks of the layer as needed.
</p>
<h4>That can't be true, gzip doesn't support random access!</h4>
<p>
That's not a question.
</p>
<h4>Okay then, how does random access work if the layers are gzipped tarballs?</h4>
<p>Great question! See <a href="https://github.com/madler/zlib/blob/master/examples/zran.c">here</a>.</p>
<p>Tl;dr, you can seek to an arbitrary position in a gzip stream if you know the 32KiB of uncompressed data that comes just before it, so by storing ~1% of the uncompressed layer size, I can jump ahead to predetermined locations and start reading from there rather than reading the entire layer.</p>
<p>Thanks <a href="https://github.com/aidansteele">@aidansteele</a>!</p>
<h4>Is this open source?</h4>
<p>Yes! See <a href="https://github.com/jonjohnsonjr/dag.dev">here</a>.</p>
</body>
</html>
`

	oauthTemplate = `
<html>
<body>
<head>
<title>Registry Explorer</title>
<link rel="icon" href="/favicon.svg">
<style>
.mt:hover {
	text-decoration: underline;
}

.mt {
	color: inherit;
	text-decoration: inherit;
}

.link {
	position: relative;
	bottom: .125em;
}

.crane {
	height: 1em;
	width: 1em;
}

.top {
	color: inherit;
	text-decoration: inherit;
}

:root {
  color-scheme: light dark;
}

body {
	font-family: monospace;
	width: fit-content;
	overflow-wrap: anywhere;
	padding: 12px;
}
</style>
</head>
<h1><a class="top" href="/"><img class="crane" src="/favicon.svg"/> <span class="link">Registry Explorer</span></a></h1>
<p>
It looks like we encountered an auth error:
</p>
<code>
{{.Error}}
</code>
<p>
If you trust <a class="mt" href="https://github.com/jonjohnsonjr">me</a>, click <a href="{{.Redirect}}">here</a> for oauth to use your own credentials.
</p>
</body>
</html>
`

	headerTemplate = `
<html>
<head>
<script type="text/javascript">
function copyToClipboard() {
  var element = document.createElement("textarea");
  element.value = document.getElementById("command").innerText;
  element.select();
  element.setSelectionRange(0, 99999); // For mobile devices
  navigator.clipboard.writeText(element.value);
  setTimeout(function() {
	alert("Command copied to clipboard!\n" + element.value);
  }, 100);
}
</script>
<title>{{.Title}}</title>
<link rel="icon" href="/favicon.svg">
<style>
.mt:hover {
	text-decoration: underline;
}

.mt {
	color: inherit;
	text-decoration: inherit;
}

.link {
	position: relative;
	bottom: .125em;
}

.crane {
	height: 1em;
	width: 1em;
}

.top {
	color: inherit;
	text-decoration: inherit;
}

:root {
  color-scheme: light dark;
}

body {
	font-family: monospace;
	width: fit-content;
	overflow-wrap: anywhere;
	padding: 12px;
}

pre {
	white-space: pre-wrap;
}

.indent {
	margin-left: 2em;
}

.noselect {
	user-select: none;
	-webkit-user-select: none;
	width: fit-content;
	overflow-wrap: none;
	padding-right: 1em;
	text-align: right;
	white-space: nowrap;
	text-decoration: none;
}

td {
	vertical-align: top;
}

// Adapted from https://medium.com/allenhwkim/how-to-build-tabs-only-with-css-844718d7de2f
input + label { display: inline-block } /* show labels in line */
input { display: none; }                /* hide radio buttons */
input ~ .tab { display: none }          /* hide contents */

/* show contents only for selected tab */
#tab1:checked ~ .tab.content1,
#tab2:checked ~ .tab.content2 { display: block; }

input + label {             /* box with rounded corner */
	display: inline-block;
  border: 1px solid #999;
  padding: 4px 12px;
  border-radius: 4px 4px 0 0;
  position: relative;
  top: 1px;
  opacity: 50%;
}
input:checked + label {     /* white background for selected tab */
  opacity: 100%;
}
input ~ .tab {          /* grey line between tab and contents */
  border-top: 1px solid #999;
  padding-top: 0.5em;
}
</style>
</head>
`

	bodyTemplate = `
<body>
<div>
<h1><a class="top" href="/"><img class="crane" src="/favicon.svg"/> <span class="link">Registry Explorer</span></a></h1>
{{ if .Up }}
<h2>{{ if and (ne .Up.Parent "docker.io") (ne .Up.Parent "index.docker.io") }}<a class="mt" href="/?repo={{.Up.Parent}}">{{.Up.Parent}}</a>{{else}}{{.Up.Parent}}{{end}}{{.Up.Separator}}{{if .Handler }}<a class="mt" href="/{{.Handler}}{{.Reference}}{{if .EscapedMediaType}}&mt={{.EscapedMediaType}}{{end}}">{{.Up.Child}}</a>{{else}}{{.Up.Child}}{{end}}{{ range .CosignTags }} (<a href="/?image={{$.Repo}}:{{.Tag}}">{{.Short}}</a>){{end}}{{if .Referrers}} <a href="/?referrers={{$.Repo}}@{{$.Descriptor.Digest}}">(referrers)</a>{{end}}</h2>
{{ else }}
	<h2>{{.Reference}}{{ range .CosignTags }} (<a href="/?image={{$.Repo}}:{{.Tag}}">{{.Short}}</a>){{end}}{{if .Referrers}} <a href="/?referrers={{$.Repo}}@{{$.Descriptor.Digest}}">(referrers)</a>{{end}}</h2>
{{ end }}
{{ if .Descriptor }}
<input type="radio" name="tabs" id="tab1" checked />
<label for="tab1">HTTP</label>
<input type="radio" name="tabs" id="tab2" />
<label for="tab2">OCI</label>
<div class="tab content1">
Content-Type: {{if .MediaTypeLink}}<a class="mt" href="{{.MediaTypeLink}}">{{.Descriptor.MediaType}}</a>{{else}}{{.Descriptor.MediaType}}{{end}}<br>
Docker-Content-Digest: <a class="mt" href="/{{.Handler}}{{$.Repo}}@{{.Descriptor.Digest}}{{if .EscapedMediaType}}&mt={{.EscapedMediaType}}{{end}}&size={{.Descriptor.Size}}">{{.Descriptor.Digest}}</a><br>
Content-Length: {{if .SizeLink}}<a class="mt" href="{{.SizeLink}}">{{.Descriptor.Size}}</a>{{else}}{{.Descriptor.Size}}{{end}}<br>
{{if $.Subject}}OCI-Subject: <a class="mt" href="/?image={{$.Repo}}@{{.Subject}}">{{.Subject}}</a><br>{{end}}
</div>
<div class="tab content2">
{<br>
&nbsp;&nbsp;"mediaType": "{{.Descriptor.MediaType}}",<br>
&nbsp;&nbsp;"digest": "{{.Descriptor.Digest}}",<br>
&nbsp;&nbsp;"size": {{.Descriptor.Size}}<br>
}<br>
</div>

{{end}}
</div>
{{ if .JQ }}
<h4><a class="noselect" onclick="copyToClipboard()">$</a><span id="command">{{.JQ}}</span></h4>
{{ end }}
`

	footer = `
</body>
</html>
`
)

type RepoParent struct {
	Parent    string
	Child     string
	Separator string
}

type OauthData struct {
	Error    string
	Redirect string
}

type TitleData struct {
	Title string
}
type CosignTag struct {
	Tag   string
	Short string
}

type HeaderData struct {
	Repo             string
	CosignTags       []CosignTag
	JQ               string
	Reference        string
	Up               *RepoParent
	Descriptor       *v1.Descriptor
	Handler          string
	EscapedMediaType string
	MediaTypeLink    string
	SizeLink         string
	Referrers        bool
	Subject          string
}
