package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
)

// Generator converts a directory of Markdown files into a static HTML documentation site.
type Generator struct {
	InputDir   string
	OutputDir  string
	AssetsDir  string
	SiteTitle  string
	Lang       string
	FaviconURL string

	md goldmark.Markdown
}

// NewGenerator creates a new generator instance.
func NewGenerator(inputDir, outputDir, assetsDir, siteTitle, lang string) *Generator {
	return &Generator{
		InputDir:  inputDir,
		OutputDir: outputDir,
		AssetsDir: assetsDir,
		SiteTitle: siteTitle,
		Lang:      lang,
		md:        NewMarkdownConverter(),
	}
}

// Generate runs the full documentation generation pipeline.
func (g *Generator) Generate() error {
	info, err := os.Stat(g.InputDir)
	if err != nil {
		return fmt.Errorf("input directory error: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("input path is not a directory: %s", g.InputDir)
	}

	if err := os.MkdirAll(g.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	g.detectFavicon()

	fmt.Println("Copying static assets...")
	if err := CopyAssets(g.AssetsDir, g.OutputDir); err != nil {
		return fmt.Errorf("failed to copy assets: %w", err)
	}

	fmt.Println("Scanning markdown files...")
	mdFiles, err := g.collectMarkdownFiles()
	if err != nil {
		return fmt.Errorf("failed to scan markdown files: %w", err)
	}

	if len(mdFiles) == 0 {
		fmt.Println("No markdown files found in input directory.")
		return nil
	}

	fmt.Printf("  Found %d markdown file(s)\n", len(mdFiles))

	fmt.Println("Converting markdown to HTML...")
	hasIndex := false

	for _, relPath := range mdFiles {
		if filepath.ToSlash(relPath) == "index.md" {
			hasIndex = true
		}
		if err := g.convertFile(relPath); err != nil {
			return fmt.Errorf("failed to convert %s: %w", relPath, err)
		}
	}

	if !hasIndex {
		fmt.Println("Generating document index page...")
		if err := g.generateIndex(mdFiles); err != nil {
			return fmt.Errorf("failed to generate index: %w", err)
		}
	}

	return nil
}

// detectFavicon looks for a favicon file in the input directory root and copies it to output.
func (g *Generator) detectFavicon() {
	faviconNames := []string{"favicon.svg", "favicon.ico", "favicon.png", "favicon.jpg", "favicon.jpeg"}
	for _, name := range faviconNames {
		srcPath := filepath.Join(g.InputDir, name)
		if _, err := os.Stat(srcPath); err == nil {
			dstPath := filepath.Join(g.OutputDir, name)
			if err := copyFile(srcPath, dstPath); err == nil {
				g.FaviconURL = name
				fmt.Printf("  Found favicon: %s\n", name)
			}
			return
		}
	}
}

// collectMarkdownFiles recursively collects all .md files from the input directory.
func (g *Generator) collectMarkdownFiles() ([]string, error) {
	var files []string

	err := filepath.Walk(g.InputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") && path != g.InputDir {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			rel, err := filepath.Rel(g.InputDir, path)
			if err != nil {
				return err
			}
			files = append(files, rel)
		}
		return nil
	})

	sort.Strings(files)
	return files, err
}

// convertFile converts a single markdown file to HTML and writes it to the output directory.
func (g *Generator) convertFile(relPath string) error {
	srcPath := filepath.Join(g.InputDir, relPath)
	source, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	result, err := ConvertMarkdown(g.md, source)
	if err != nil {
		return err
	}

	outPath := g.outputPath(relPath)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	assetPrefix := g.computeAssetPrefix(relPath)
	title := g.determineTitle(relPath, result.NavItems)
	breadcrumbs := g.buildBreadcrumbs(relPath, title)

	faviconPath := ""
	if g.FaviconURL != "" {
		faviconPath = assetPrefix + g.FaviconURL
	}

	html, err := RenderPage(PageData{
		Title:       title,
		SiteTitle:   g.SiteTitle,
		Lang:        g.Lang,
		Description: result.Description,
		NavItems:    result.NavItems,
		Breadcrumbs: breadcrumbs,
		FaviconPath: faviconPath,
		Content:     template.HTML(result.HTML),
		AssetPrefix: assetPrefix,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(outPath, []byte(html), 0644); err != nil {
		return err
	}

	relOut, _ := filepath.Rel(g.OutputDir, outPath)
	fmt.Printf("  %s -> %s\n", relPath, relOut)
	return nil
}

// outputPath computes the output HTML path for a given markdown relative path.
func (g *Generator) outputPath(relPath string) string {
	htmlRel := strings.TrimSuffix(relPath, ".md") + ".html"
	return filepath.Join(g.OutputDir, htmlRel)
}

// computeAssetPrefix calculates the relative path prefix for CSS/JS assets based on directory depth.
func (g *Generator) computeAssetPrefix(relPath string) string {
	depth := strings.Count(filepath.ToSlash(relPath), "/")
	if depth == 0 {
		return "./"
	}
	return strings.Repeat("../", depth)
}

// determineTitle returns the page title from the first h1 heading, or falls back to the filename.
func (g *Generator) determineTitle(relPath string, navItems []NavItem) string {
	for _, item := range navItems {
		if item.Level == 1 {
			return item.Title
		}
	}
	base := filepath.Base(relPath)
	name := strings.TrimSuffix(base, ".md")
	if name == "index" {
		dir := filepath.Dir(relPath)
		if dir == "." {
			if g.SiteTitle != "" {
				return g.SiteTitle
			}
			return "Documentation"
		}
		return dirNameToTitle(filepath.Base(dir))
	}
	return dirNameToTitle(name)
}

// dirNameToTitle converts a directory or file name into a human-readable title.
func dirNameToTitle(name string) string {
	title := strings.ReplaceAll(name, "-", " ")
	title = strings.ReplaceAll(title, "_", " ")
	words := strings.Fields(title)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

// buildBreadcrumbs generates breadcrumb navigation for a page based on its directory hierarchy.
func (g *Generator) buildBreadcrumbs(relPath string, pageTitle string) []BreadcrumbItem {
	relPath = filepath.ToSlash(relPath)

	dir := filepath.Dir(relPath)
	if dir == "." {
		return nil
	}

	parts := strings.Split(dir, "/")
	var crumbs []BreadcrumbItem

	upLevels := len(parts)
	homeURL := strings.Repeat("../", upLevels) + "index.html"
	crumbs = append(crumbs, BreadcrumbItem{
		Title: "Home",
		URL:   homeURL,
	})

	for i, part := range parts {
		title := dirNameToTitle(part)
		levelsUp := upLevels - i - 1
		var url string
		if levelsUp > 0 {
			url = strings.Repeat("../", levelsUp) + "index.html"
		} else {
			url = "index.html"
		}
		crumbs = append(crumbs, BreadcrumbItem{
			Title: title,
			URL:   url,
		})
	}

	crumbs = append(crumbs, BreadcrumbItem{
		Title: pageTitle,
		URL:   "",
	})

	return crumbs
}

// generateIndex creates an auto-generated index page listing all documents.
func (g *Generator) generateIndex(mdFiles []string) error {
	var entries []NavItem
	for _, relPath := range mdFiles {
		htmlPath := strings.TrimSuffix(relPath, ".md") + ".html"

		name := strings.TrimSuffix(filepath.Base(relPath), ".md")
		if name == "index" {
			dir := filepath.Dir(relPath)
			if dir == "." {
				name = "Home"
			} else {
				name = filepath.Base(dir)
			}
		}
		title := dirNameToTitle(name)

		entries = append(entries, NavItem{
			ID:    htmlPath,
			Title: title,
		})
	}

	var content strings.Builder
	content.WriteString("<h3>Document Index</h3>\n")
	content.WriteString("<p>The following documents are available:</p>\n")
	content.WriteString("<ul>\n")
	for _, entry := range entries {
		content.WriteString(fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", entry.ID, entry.Title))
	}
	content.WriteString("</ul>\n")

	indexTitle := "Documentation"
	if g.SiteTitle != "" {
		indexTitle = g.SiteTitle
	}

	faviconPath := ""
	if g.FaviconURL != "" {
		faviconPath = "./" + g.FaviconURL
	}

	html, err := RenderIndex(PageData{
		Title:       indexTitle,
		SiteTitle:   g.SiteTitle,
		Lang:        g.Lang,
		NavItems:    entries,
		FaviconPath: faviconPath,
		Content:     template.HTML(content.String()),
		AssetPrefix: "./",
	})
	if err != nil {
		return err
	}

	outPath := filepath.Join(g.OutputDir, "index.html")
	if err := os.WriteFile(outPath, []byte(html), 0644); err != nil {
		return err
	}

	fmt.Printf("  Auto-generated index.html\n")
	return nil
}
