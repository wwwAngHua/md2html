# Quick Start Guide

This guide will help you get started with md2html.

## Prerequisites

- **Go 1.21+** installed on your system
- A directory of Markdown files
- Basic command line knowledge

## Installation

### From Source

```bash
git clone https://github.com/wwwanghua/md2html.git
cd md2html/gen
go build -o md2html .
```

### Using Go Install

```bash
go install github.com/wwwanghua/md2html/gen@latest
```

## Basic Usage

### Simple Conversion

```bash
./md2html -in ../docs -out ../dist
```

### With Custom Site Title

```bash
./md2html -in ../docs -out ../dist -title "My Project Docs"
```

### With Custom Language

Set the HTML `lang` attribute for your generated pages:

```bash
# English documentation
./md2html -in ../docs -out ../dist -lang en

# Japanese documentation with site title
./md2html -in ../docs -out ../dist -title "My Docs" -lang ja
```

Defaults to `zh-CN` if not specified.

### With Custom Assets

If your Tabler UI assets (css/js directories) are in a different location:

```bash
./md2html -in ../docs -out ../dist -assets ~/my-theme
```

## Directory Structure

Input and output mapping:

```
my-docs/
├── index.md           -> index.html
├── readme.md          -> readme.html
├── favicon.svg        -> favicon.svg (auto-detected)
├── quick-start/
│   └── index.md       -> quick-start/index.html
└── guide/
    ├── index.md       -> guide/index.html
    └── advanced.md    -> guide/advanced.html
```

## Page Titles

- Without `-title`: page titles use the first `h1` heading from each document
- With `-title "My Docs"`: sub-pages get a suffix, e.g. "Quick Start Guide - My Docs"

## Next Steps

- Read the [Advanced Guide](../guide/advanced.html) for customization options
- Explore code syntax highlighting with the Monokai dark theme

> **Note:** All directory paths are preserved exactly as they appear in the input directory.
