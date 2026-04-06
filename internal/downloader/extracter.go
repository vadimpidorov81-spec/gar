package parser

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(zipPath, dstDir string) error {
	if zipPath == "" {
		return fmt.Errorf("zipPath is empty")
	}
	if dstDir == "" {
		return fmt.Errorf("dstDir is empty")
	}

	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("open zip: %w", err)
	}
	defer reader.Close()

	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("create dst dir: %w", err)
	}

	for _, file := range reader.File {
		if err := extractFile(file, dstDir); err != nil {
			return err
		}
	}

	return nil
}

func extractFile(file *zip.File, dstDir string) error {
	cleanDst := filepath.Clean(dstDir)
	targetPath := filepath.Clean(filepath.Join(cleanDst, file.Name))

	if !strings.HasPrefix(targetPath, cleanDst+string(os.PathSeparator)) && targetPath != cleanDst {
		return fmt.Errorf("invalid zip entry path: %s", file.Name)
	}

	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(targetPath, 0o755); err != nil {
			return fmt.Errorf("create dir: %w", err)
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), 0o755); err != nil {
		return fmt.Errorf("create parent dir: %w", err)
	}

	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("open zip entry: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("copy file: %w", err)
	}

	return nil
}
