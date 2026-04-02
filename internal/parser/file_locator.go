package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFileByPrefix(xmlDir, prefix string) (string, error) {
	if xmlDir == "" {
		return "", fmt.Errorf("dirPath is empty")
	}
	if prefix == "" {
		return "", fmt.Errorf("prefix is empty")
	}

	entries, err := os.ReadDir(xmlDir)
	if err != nil {
		return "", fmt.Errorf("read dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasPrefix(strings.ToUpper(entry.Name()), strings.ToUpper(prefix)) {
			return filepath.Join(xmlDir, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("file with prefix %q not found in %q", prefix, xmlDir)
}
