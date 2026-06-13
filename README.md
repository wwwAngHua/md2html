# md2html

A command-line tool that converts a directory of Markdown files into a beautiful, self-contained HTML documentation site — styled with [Tabler UI](https://tabler.io).

## Features

- **Tabler UI Design** — Clean, modern interface with responsive layout
- **Sidebar Navigation** — Auto-generated from document headings with Bootstrap Scroll Spy
- **Breadcrumb Trail** — Always know your location in deeply nested documentation
- **Syntax Highlighting** — Code blocks rendered with the Monokai dark theme via [Chroma](https://github.com/alecthomas/chroma)
- **Custom Site Title** — Add your project name as a page title suffix (`Page - Site Name`)
- **Multi-Language** — Configurable `lang` attribute via `-lang` flag
- **SEO Ready** — Auto-generated meta description and Open Graph tags from content
- **Mobile Responsive** — Collapsible sidebar on small screens with hamburger toggle
- **Favicon Support** — Drop a `favicon.svg` in your docs directory for automatic branding
- **Zero Dependencies at Runtime** — Generates pure HTML/CSS/JS files, no build tools needed
- **Directory Structure Preserved** — Input directory layout is mirrored exactly in the output

## Quick Start

```bash
# Clone and build
git clone https://github.com/wwwanghua/md2html.git
cd md2html/gen
go build -o md2html .

# Generate documentation
./md2html -in ../docs -out ../dist -title "My Project"
```

Open `dist/index.html` in your browser.

## Installation

### From Source

```bash
git clone https://github.com/wwwanghua/md2html.git
cd md2html/gen
go build -o md2html .
```

Requires **Go 1.21+**.

### Go Install

```bash
go install github.com/wwwanghua/md2html/gen@latest
```

## Usage

```
md2html -in <markdown-dir> [-out <output-dir>] [-assets <assets-dir>] [-title <site-name>] [-lang <lang-code>]
```

### Flags

| Flag      | Default      | Description                                         |
| --------- | ------------ | --------------------------------------------------- |
| `-in`     | _(required)_ | Input directory containing markdown files           |
| `-out`    | `./out`      | Output directory for generated HTML                 |
| `-assets` | current dir  | Directory containing `css/` and `js/` Tabler assets |
| `-title`  | _(empty)_    | Site name used as page title suffix                 |
| `-lang`   | `zh-CN`      | HTML `lang` attribute for generated pages           |

### Examples

```bash
# Basic usage
./md2html -in ../docs

# Custom output directory
./md2html -in ../docs -out ../dist

# With site title
./md2html -in ../docs -out ../dist -title "t1yOS Developer Docs"

# Custom assets location
./md2html -in ../docs -out ../dist -assets ./tabler-dist

# English documentation
./md2html -in ../docs -out ../dist -lang en

# Japanese documentation with site title
./md2html -in ../docs -out ../dist -title "My Project" -lang ja
```

## Directory Structure

```
my-docs/                        out/
├── index.md        ->          ├── index.html
├── readme.md       ->          ├── readme.html
├── favicon.svg     ->          ├── favicon.svg
├── quick-start/                ├── quick-start/
│   └── index.md    ->          │   └── index.html
└── guide/                      ├── guide/
    ├── index.md    ->          │   ├── index.html
    └── advanced.md ->          │   └── advanced.html
                                ├── css/    (Tabler CSS)
                                └── js/     (Tabler JS)
```

## How It Works

1. Scans the input directory recursively for `.md` files
2. Detects and copies any `favicon.*` file from the input root
3. Converts each Markdown file to HTML using [Goldmark](https://github.com/yuin/goldmark)
4. Extracts headings to build sidebar navigation and breadcrumbs
5. Wraps the HTML content in a Tabler UI page template
6. Copies minified Tabler CSS/JS assets to the output directory
7. Auto-generates an index page if no `index.md` exists at the root

## Page Title Logic

- Pages use the first `h1` heading as the title
- If no `h1` exists, falls back to the filename
- When `-title` is set, sub-pages display as `"Page Title - Site Name"`

## Favicon

Place a `favicon.svg`, `favicon.ico`, or `favicon.png` file in the root of your documentation input directory. md2html automatically detects it, copies it to the output, and links it in every page.

## SEO

Every generated page includes:

- **Meta description** — Auto-extracted from the first paragraph of your content (~160 chars)
- **Open Graph tags** — `og:title`, `og:description`, `og:type` for social media previews
- **Semantic HTML** — Proper heading hierarchy, `<nav>` elements, breadcrumb structured data
- **Language attribute** — `<html lang="...">` configured via `-lang` flag

No extra configuration required — descriptions are derived from your actual content.

## Language

Use the `-lang` flag to set the HTML `lang` attribute:

```bash
./md2html -in ../docs -out ../dist -lang en   # English
./md2html -in ../docs -out ../dist -lang ja   # Japanese
```

Defaults to `zh-CN`. Accepts any [BCP 47 language tag](https://tools.ietf.org/html/bcp47).

## Tech Stack

| Component           | Library                                                        |
| ------------------- | -------------------------------------------------------------- |
| Markdown Parser     | [Goldmark](https://github.com/yuin/goldmark)                   |
| Syntax Highlighting | [Chroma](https://github.com/alecthomas/chroma) (Monokai theme) |
| UI Framework        | [Tabler](https://tabler.io) (Bootstrap-based)                  |
| Template Engine     | Go `html/template`                                             |
| Language            | Go                                                             |

## License

MIT
