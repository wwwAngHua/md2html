package main

import (
	"html/template"
	"strings"
)

// PageData holds all data passed to the HTML template for rendering.
type PageData struct {
	Title       string
	SiteTitle   string
	Lang        string
	Description string
	NavItems    []NavItem
	Breadcrumbs []BreadcrumbItem
	FaviconPath string
	Content     template.HTML
	AssetPrefix string
}

// FullTitle returns the complete page title, including the site name suffix when set.
func (p PageData) FullTitle() string {
	if p.SiteTitle == "" {
		return p.Title
	}
	if p.Title == p.SiteTitle {
		return p.SiteTitle
	}
	return p.Title + " - " + p.SiteTitle
}

// commonHead returns the shared <head> content for all pages.
const commonHead = `    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, viewport-fit=cover"
    />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>{{.FullTitle}}</title>
    {{if .FaviconPath}}<link rel="icon" href="{{$.FaviconPath}}" />{{end}}
    {{if .Description}}<meta name="description" content="{{.Description}}" />{{end}}
    <meta property="og:title" content="{{.FullTitle}}" />
    {{if .Description}}<meta property="og:description" content="{{.Description}}" />{{end}}
    <meta property="og:type" content="article" />
    <link href="{{.AssetPrefix}}css/tabler.min.css" rel="stylesheet" />
    <link href="{{.AssetPrefix}}css/tabler-flags.min.css" rel="stylesheet" />
    <link href="{{.AssetPrefix}}css/tabler-socials.min.css" rel="stylesheet" />
    <link href="{{.AssetPrefix}}css/tabler-payments.min.css" rel="stylesheet" />
    <link href="{{.AssetPrefix}}css/tabler-vendors.min.css" rel="stylesheet" />
    <link href="{{.AssetPrefix}}css/tabler-marketing.min.css" rel="stylesheet" />
    <link href="{{.AssetPrefix}}css/tabler-themes.min.css" rel="stylesheet" />`

// commonStyles contains the shared CSS for both page and index templates.
const commonStyles = `      @import url("https://rsms.me/inter/inter.css");

      .markdown :not(pre) > code {
        background-color: #f0f0f0;
        color: #d63384;
        padding: 2px 6px;
        border-radius: 4px;
        font-size: 0.875em;
      }

      .markdown pre {
        border-radius: 8px;
        padding: 16px;
        overflow-x: auto;
        margin: 1rem 0;
        background-color: #272822;
      }

      .markdown pre code {
        font-family: "JetBrains Mono", "Fira Code", "Cascadia Code", "SF Mono", Consolas, monospace;
        font-size: 0.875em;
        line-height: 1.6;
        background: none !important;
        padding: 0;
      }

      .markdown pre[style] {
        background-color: #272822 !important;
      }

      .breadcrumb {
        margin: 0;
        padding: 0;
        list-style: none;
        display: flex;
        flex-wrap: wrap;
      }
      .breadcrumb-item {
        font-size: 0.875rem;
        color: var(--tblr-muted, #656d77);
      }
      .breadcrumb-item + .breadcrumb-item::before {
        content: "/";
        margin: 0 0.5rem;
        color: var(--tblr-muted, #656d77);
      }
      .breadcrumb-item a {
        color: var(--tblr-muted, #656d77);
        text-decoration: none;
      }
      .breadcrumb-item a:hover {
        color: var(--tblr-primary, #206bc4);
        text-decoration: underline;
      }
      .breadcrumb-item.active {
        color: var(--tblr-body-color, #1e293b);
      }

      /* Mobile sidebar toggle */
      .sidebar-toggle {
        display: none;
      }
      @media (max-width: 575.98px) {
        .sidebar-toggle {
          display: inline-flex;
          align-items: center;
          gap: 0.25rem;
        }
        .sidebar-collapse {
          margin-bottom: 1rem;
        }
      }`

// commonFooter contains the shared end-of-body content.
const commonFooter = `    <script src="{{.AssetPrefix}}js/tabler.min.js" defer></script>`

// pageTemplate is the main document page template based on Tabler UI.
const pageTemplate = `<!doctype html>
<html lang="{{.Lang}}">
  <head>` + commonHead + `
    <style>` + commonStyles + `
    </style>
  </head>
  <body>
    <div class="page">
      <div class="page-wrapper">
        <div class="page-header d-print-none" aria-label="Page header">
          <div class="container-xl">
            <div class="row g-2 align-items-center">
              <div class="col">
                <h2 class="page-title">{{.Title}}</h2>
                {{if .Breadcrumbs}}
                <ol class="breadcrumb" aria-label="breadcrumb">
                  {{range $i, $bc := .Breadcrumbs}}
                  {{if $bc.URL}}
                  <li class="breadcrumb-item"><a href="{{$bc.URL}}">{{$bc.Title}}</a></li>
                  {{else}}
                  <li class="breadcrumb-item active" aria-current="page">{{$bc.Title}}</li>
                  {{end}}
                  {{end}}
                </ol>
                {{end}}
              </div>
              {{if .NavItems}}
              <div class="col-auto ms-auto d-sm-none">
                <button class="btn btn-icon sidebar-toggle" type="button" data-bs-toggle="collapse" data-bs-target="#sidebar-menu" aria-expanded="false" aria-controls="sidebar-menu" aria-label="Toggle menu">
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/></svg>
                </button>
              </div>
              {{end}}
            </div>
          </div>
        </div>
        <div class="page-body">
          <div class="container-xl">
            <div class="row g-5">
              {{if .NavItems}}
              <div class="col-sm-2">
                <div class="collapse d-sm-block sidebar-collapse" id="sidebar-menu">
                  <div class="sticky-top">
                    <nav class="nav nav-vertical nav-pills" id="pills">
                      {{range .NavItems}}
                      <a class="nav-link nav-link-level-{{.Level}}" href="#{{.ID}}">{{.Title}}</a>
                      {{end}}
                    </nav>
                  </div>
                </div>
              </div>
              {{end}}
              <div
                class="col-sm"
                data-bs-spy="scroll"
                data-bs-target="#pills"
                data-bs-offset="0"
              >
                <div class="card card-lg">
                  <div class="card-body markdown">
                    {{.Content}}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>` + commonFooter + `
  </body>
</html>`

// indexTemplate is the auto-generated index page template.
const indexTemplate = `<!doctype html>
<html lang="{{.Lang}}">
  <head>` + commonHead + `
    <style>` + commonStyles + `
    </style>
  </head>
  <body>
    <div class="page">
      <div class="page-wrapper">
        <div class="page-header d-print-none" aria-label="Page header">
          <div class="container-xl">
            <div class="row g-2 align-items-center">
              <div class="col">
                <h2 class="page-title">{{.Title}}</h2>
                {{if .Breadcrumbs}}
                <ol class="breadcrumb" aria-label="breadcrumb">
                  {{range $i, $bc := .Breadcrumbs}}
                  {{if $bc.URL}}
                  <li class="breadcrumb-item"><a href="{{$bc.URL}}">{{$bc.Title}}</a></li>
                  {{else}}
                  <li class="breadcrumb-item active" aria-current="page">{{$bc.Title}}</li>
                  {{end}}
                  {{end}}
                </ol>
                {{end}}
              </div>
              {{if .NavItems}}
              <div class="col-auto ms-auto d-sm-none">
                <button class="btn btn-icon sidebar-toggle" type="button" data-bs-toggle="collapse" data-bs-target="#sidebar-menu" aria-expanded="false" aria-controls="sidebar-menu" aria-label="Toggle menu">
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/></svg>
                </button>
              </div>
              {{end}}
            </div>
          </div>
        </div>
        <div class="page-body">
          <div class="container-xl">
            <div class="row g-5">
              {{if .NavItems}}
              <div class="col-sm-2">
                <div class="collapse d-sm-block sidebar-collapse" id="sidebar-menu">
                  <div class="sticky-top">
                    <nav class="nav nav-vertical nav-pills" id="pills">
                      {{range .NavItems}}
                      <a class="nav-link" href="{{.ID}}">{{.Title}}</a>
                      {{end}}
                    </nav>
                  </div>
                </div>
              </div>
              {{end}}
              <div class="col-sm">
                <div class="card card-lg">
                  <div class="card-body markdown">
                    {{.Content}}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>` + commonFooter + `
  </body>
</html>`

var (
	parsedPageTemplate  *template.Template
	parsedIndexTemplate *template.Template
)

func init() {
	funcMap := template.FuncMap{}

	var err error
	parsedPageTemplate, err = template.New("page").Funcs(funcMap).Parse(pageTemplate)
	if err != nil {
		panic("failed to parse page template: " + err.Error())
	}
	parsedIndexTemplate, err = template.New("index").Funcs(funcMap).Parse(indexTemplate)
	if err != nil {
		panic("failed to parse index template: " + err.Error())
	}
}

// RenderPage renders a document page using the Tabler UI template.
func RenderPage(data PageData) (string, error) {
	var buf strings.Builder
	if err := parsedPageTemplate.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderIndex renders the index page template.
func RenderIndex(data PageData) (string, error) {
	var buf strings.Builder
	if err := parsedIndexTemplate.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
