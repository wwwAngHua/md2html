# md2html

Welcome to **md2html** — a command-line tool that converts a directory of Markdown files into a beautiful, Tabler UI styled static HTML documentation site.

## Features

- **Tabler UI** — Clean, modern interface based on the [Tabler](https://tabler.io) design system
- **Sidebar Navigation** — Auto-generated from your document headings with Scroll Spy support
- **Syntax Highlighting** — Code blocks are highlighted using the Monokai dark theme
- **Breadcrumb Navigation** — Always know where you are in deeply nested documentation
- **Custom Site Title** — Add your project name as a page title suffix
- **Multi-Language** — Set the HTML `lang` attribute for accessibility and SEO
- **SEO Ready** — Auto-generated meta description and Open Graph tags for better search visibility
- **Mobile Friendly** — Collapsible sidebar on small screens, responsive Tabler layout
- **Favicon Support** — Drop a `favicon.svg` in your docs directory for automatic branding
- **Static Output** — Generates pure HTML files with zero JavaScript build tools

## How It Works

1. Point `md2html` to your Markdown directory
2. It scans all `.md` files recursively
3. Each file is converted to a standalone HTML page
4. CSS and JS assets are copied to the output directory

## Getting Started

See the [Quick Start Guide](./quick-start/index.html) for installation and usage instructions.

## Code Example

Here is a simple Go example with syntax highlighting:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, md2html!")
}
```

## Configuration

| Flag      | Default      | Description                         |
| --------- | ------------ | ----------------------------------- |
| `-in`     | _(required)_ | Input directory with markdown files |
| `-out`    | `./out`      | Output directory for HTML           |
| `-assets` | current dir  | Path to Tabler css/js assets        |
| `-title`  | _(empty)_    | Site name as page title suffix      |
| `-lang`   | `zh-CN`      | HTML language code                  |

## Notes

> **Tip:** Use `index.md` as the filename for directory index pages. The generator automatically maps them to `index.html`.

---

Built with [md2html](https://github.com/wwwanghua/md2html) — a Markdown documentation generator powered by [Goldmark](https://github.com/yuin/goldmark) and [Tabler UI](https://tabler.io).
