package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// NavItem represents a single entry in the sidebar navigation.
type NavItem struct {
	ID    string
	Title string
	Level int
}

// BreadcrumbItem represents a single level in breadcrumb navigation.
type BreadcrumbItem struct {
	Title string
	URL   string // empty means current page (not a link)
}

// headingIDExtension adds id attributes to heading elements during AST transform.
type headingIDExtension struct{}

func (e *headingIDExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(&headingIDTransformer{}, 100),
		),
	)
}

type headingIDTransformer struct{}

func (t *headingIDTransformer) Transform(doc *ast.Document, reader text.Reader, pc parser.Context) {
	source := reader.Source()
	headings := []*ast.Heading{}
	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if h, ok := n.(*ast.Heading); ok {
			headings = append(headings, h)
		}
		return ast.WalkContinue, nil
	})

	for _, h := range headings {
		var textContent strings.Builder
		for n := h.FirstChild(); n != nil; n = n.NextSibling() {
			collectText(n, &textContent, source)
		}
		id := slugify(textContent.String())
		if id != "" {
			h.SetAttributeString("id", id)
		}
	}
}

// collectText recursively extracts text content from AST nodes.
func collectText(n ast.Node, buf *strings.Builder, source []byte) {
	switch v := n.(type) {
	case *ast.Text:
		buf.Write(v.Segment.Value(source))
	case *ast.String:
		buf.Write(v.Value)
	}
	for child := n.FirstChild(); child != nil; child = child.NextSibling() {
		collectText(child, buf, source)
	}
}

// slugify converts arbitrary text into a URL-safe HTML id.
func slugify(text string) string {
	slug := strings.ToLower(text)
	slug = strings.NewReplacer(
		"：", " ", "。", " ", "，", " ", "；", " ", "！", " ",
		"？", " ", "（", " ", "）", " ", "【", " ", "】", " ",
		"《", " ", "》", " ", "\"", " ", "'", " ", "“", " ",
		"”", " ", "‘", " ", "’", " ",
	).Replace(slug)
	reg := regexp.MustCompile(`[^a-z0-9一-鿿]+`)
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	reg2 := regexp.MustCompile(`-+`)
	slug = reg2.ReplaceAllString(slug, "-")
	return slug
}

// ConvertResult holds the result of a markdown conversion.
type ConvertResult struct {
	HTML        string
	NavItems    []NavItem
	Description string
}

// NewMarkdownConverter creates a configured goldmark converter.
func NewMarkdownConverter() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
			),
			extension.GFM,
			extension.Footnote,
			extension.DefinitionList,
			&headingIDExtension{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
}

// ConvertMarkdown converts markdown source to HTML and extracts navigation.
func ConvertMarkdown(md goldmark.Markdown, source []byte) (*ConvertResult, error) {
	navItems := extractHeadings(md, source)
	description := extractDescription(md, source)

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return nil, fmt.Errorf("markdown conversion failed: %w", err)
	}

	return &ConvertResult{
		HTML:        buf.String(),
		NavItems:    navItems,
		Description: description,
	}, nil
}

// extractHeadings parses the markdown AST and collects all heading nodes.
func extractHeadings(md goldmark.Markdown, source []byte) []NavItem {
	doc := md.Parser().Parse(text.NewReader(source))

	var items []NavItem
	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if h, ok := n.(*ast.Heading); ok {
			var textContent strings.Builder
			for child := h.FirstChild(); child != nil; child = child.NextSibling() {
				collectText(child, &textContent, source)
			}
			title := strings.TrimSpace(textContent.String())
			id := slugify(title)
			if id == "" {
				id = fmt.Sprintf("heading-%d", h.Level)
			}
			items = append(items, NavItem{
				ID:    id,
				Title: title,
				Level: h.Level,
			})
		}
		return ast.WalkContinue, nil
	})

	return items
}

// extractDescription extracts the first paragraph text for use as meta description.
func extractDescription(md goldmark.Markdown, source []byte) string {
	doc := md.Parser().Parse(text.NewReader(source))

	var description string
	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering || description != "" {
			return ast.WalkContinue, nil
		}
		if p, ok := n.(*ast.Paragraph); ok {
			// Skip paragraphs that are inside blockquotes or list items (not "first" content)
			parent := p.Parent()
			if parent != nil {
				if _, isListItem := parent.(*ast.ListItem); isListItem {
					return ast.WalkContinue, nil
				}
				if _, isBlockquote := parent.(*ast.Blockquote); isBlockquote {
					return ast.WalkContinue, nil
				}
			}
			var textContent strings.Builder
			for child := p.FirstChild(); child != nil; child = child.NextSibling() {
				collectText(child, &textContent, source)
			}
			description = strings.TrimSpace(textContent.String())
			return ast.WalkStop, nil
		}
		return ast.WalkContinue, nil
	})

	// Truncate to ~160 chars, breaking at word boundary
	if len([]rune(description)) > 160 {
		runes := []rune(description)
		cutoff := 157
		for cutoff > 120 && runes[cutoff] != ' ' && runes[cutoff] != '.' && runes[cutoff] != '。' {
			cutoff--
		}
		if cutoff <= 120 {
			cutoff = 157
		}
		description = string(runes[:cutoff]) + "..."
	}

	return description
}
