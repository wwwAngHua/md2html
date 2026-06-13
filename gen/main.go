package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	inputDir := flag.String("in", "", "Input directory containing markdown files (required)")
	outputDir := flag.String("out", "./out", "Output directory for generated HTML files")
	assetsDir := flag.String("assets", "", "Directory containing css/js/static assets (default: current directory)")
	siteTitle := flag.String("title", "", "Site name displayed in page titles (e.g. \"My Docs\")")
	lang := flag.String("lang", "zh-CN", "HTML language code for the generated pages (e.g. \"en\", \"zh-CN\", \"ja\")")
	flag.Parse()

	if *inputDir == "" {
		fmt.Println("Error: -in flag is required")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  md2html -in <markdown-dir> [-out <output-dir>] [-assets <assets-dir>] [-title <site-name>] [-lang <lang-code>]")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  md2html -in ./docs")
		fmt.Println("  md2html -in ./docs -out ./dist")
		fmt.Println("  md2html -in ./docs -title \"My Docs\"")
		fmt.Println("  md2html -in ./docs -lang en")
		fmt.Println("  md2html -in ./docs -out ./dist -assets ./theme -title \"My Docs\" -lang ja")
		os.Exit(1)
	}

	assetRoot := *assetsDir
	if assetRoot == "" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get current directory: %v\n", err)
			os.Exit(1)
		}
		assetRoot = resolveAssetsDir(cwd)
	}

	absInput, err := filepath.Abs(*inputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid input path: %v\n", err)
		os.Exit(1)
	}
	absOutput, err := filepath.Abs(*outputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid output path: %v\n", err)
		os.Exit(1)
	}
	absAssets, err := filepath.Abs(assetRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid assets path: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("md2html - Markdown Documentation Generator")
	fmt.Printf("  Input:  %s\n", absInput)
	fmt.Printf("  Output: %s\n", absOutput)
	fmt.Printf("  Assets: %s\n", absAssets)
	if *siteTitle != "" {
		fmt.Printf("  Title:  %s\n", *siteTitle)
	}
	fmt.Printf("  Lang:   %s\n", *lang)
	fmt.Println()

	gen := NewGenerator(absInput, absOutput, absAssets, *siteTitle, *lang)
	if err := gen.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Documentation generated successfully.")
}

// resolveAssetsDir finds the directory containing css/ and js/ assets.
// It checks cwd first, then falls back to the parent directory.
func resolveAssetsDir(cwd string) string {
	if hasAssets(cwd) {
		return cwd
	}
	parent := filepath.Dir(cwd)
	if hasAssets(parent) {
		return parent
	}
	return cwd
}

func hasAssets(dir string) bool {
	_, cssErr := os.Stat(filepath.Join(dir, "css", "tabler.min.css"))
	_, jsErr := os.Stat(filepath.Join(dir, "js", "tabler.min.js"))
	return cssErr == nil && jsErr == nil
}
