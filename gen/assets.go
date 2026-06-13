package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// assetsToCopy lists the exact Tabler asset files referenced by the HTML templates.
var assetsToCopy = []string{
	"css/tabler.min.css",
	"css/tabler-flags.min.css",
	"css/tabler-marketing.min.css",
	"css/tabler-payments.min.css",
	"css/tabler-socials.min.css",
	"css/tabler-themes.min.css",
	"css/tabler-vendors.min.css",
	"js/tabler.min.js",
}

// CopyAssets copies the required static assets to the output directory.
func CopyAssets(projectRoot, outDir string) error {
	for _, relPath := range assetsToCopy {
		srcPath := filepath.Join(projectRoot, relPath)
		dstPath := filepath.Join(outDir, relPath)

		if _, err := os.Stat(srcPath); os.IsNotExist(err) {
			return fmt.Errorf("required asset not found: %s", srcPath)
		}

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}

		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to copy %s: %w", relPath, err)
		}
	}

	fmt.Printf("  Copied %d asset file(s)\n", len(assetsToCopy))
	return nil
}

// copyFile copies a single file from src to dst.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
