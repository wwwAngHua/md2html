# Advanced Guide

This guide covers advanced customization and features of md2html.

## Custom Templates

You can customize the HTML template by modifying the `template.go` file in the source.

### Template Variables

The following variables are available in the `PageData` struct:

1. `Title` — Page title (from first h1 or filename)
2. `SiteTitle` — Site name set via `-title` flag
3. `Lang` — HTML language code set via `-lang` flag (default `zh-CN`)
4. `Description` — Meta description auto-extracted from the first paragraph
5. `NavItems` — Sidebar navigation items extracted from headings
6. `Breadcrumbs` — Breadcrumb trail for current page position
7. `FaviconPath` — Path to the favicon file
8. `Content` — Converted HTML content
9. `AssetPrefix` — Relative path prefix for CSS/JS assets

## Code Syntax Highlighting

md2html uses [Goldmark](https://github.com/yuin/goldmark) with the [Chroma](https://github.com/alecthomas/chroma) syntax highlighter and the Monokai dark theme.

### JavaScript

```javascript
function greet(name) {
  return `Hello, ${name}!`;
}

console.log(greet("World"));
```

### Python

```python
def fibonacci(n):
    a, b = 0, 1
    for _ in range(n):
        yield a
        a, b = b, a + b

print(list(fibonacci(10)))
```

### CSS

```css
.markdown {
  font-size: 16px;
  line-height: 1.6;
  color: var(--tblr-body-color);
}
```

## Navigation Customization

### Heading Levels

The sidebar navigation supports headings from `h1` to `h6`:

- **h1** — Main page title (used as page header)
- **h2** — Major sections
- **h3** — Sub-sections
- **h4-h6** — Detail levels

### Breadcrumb Navigation

md2html automatically generates breadcrumb navigation for nested pages:

```
Home / Guide / Advanced Guide
```

## Favicon

Place a `favicon.svg`, `favicon.ico`, or `favicon.png` file in the root of your documentation directory. md2html will automatically detect and include it in all generated pages.

Supported formats: `.svg`, `.ico`, `.png`, `.jpg`, `.jpeg`

## Language

Use the `-lang` flag to set the HTML `lang` attribute for all generated pages. This is important for accessibility, search engines, and proper font rendering.

```bash
# English documentation
./md2html -in ../docs -out ../dist -lang en

# Japanese
./md2html -in ../docs -out ../dist -lang ja

# Default (if not specified)
./md2html -in ../docs -out ../dist  # lang="zh-CN"
```

The `lang` value is placed directly into the `<html lang="...">` tag, so any valid [BCP 47 language tag](https://tools.ietf.org/html/bcp47) is accepted (e.g. `en`, `en-US`, `zh-CN`, `ja`, `ko`, `fr`).

## SEO

md2html automatically generates essential SEO meta tags for every page:

```html
<meta name="description" content="First paragraph text..." />
<meta property="og:title" content="Page Title" />
<meta property="og:description" content="First paragraph text..." />
<meta property="og:type" content="article" />
```

The meta description is auto-extracted from the first paragraph of your markdown content (truncated to ~160 characters). This means every page gets a unique, content-relevant description for search engines and social media previews without any extra configuration.

## Mobile Responsiveness

The sidebar navigation automatically collapses on small screens. A hamburger menu button appears in the page header on mobile devices, allowing users to toggle the navigation. On desktop screens, the sidebar is always visible as a sticky column.

## Performance Tips

- Use `.min.css` and `.min.js` for production — md2html copies only minified assets by default
- Keep images in the `static/` directory
- Avoid deeply nested directories for faster processing

## Troubleshooting

### Common Issues

| Problem             | Solution                                                  |
| ------------------- | --------------------------------------------------------- |
| Missing CSS         | Ensure `css/` directory is present in the assets location |
| No navigation       | Add headings (`#`, `##`, etc.) to your markdown files     |
| Build fails         | Check Go version (`go version`)                           |
| Favicon not showing | Place favicon in the root of your input directory         |

---

For more help, refer to the [md2html repository](https://github.com/wwwanghua/md2html).
