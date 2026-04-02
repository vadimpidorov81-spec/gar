package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type Downloader struct {
	client *http.Client
}

func New(client *http.Client) *Downloader {
	if client == nil {
		client = &http.Client{}
	}

	return &Downloader{
		client: client,
	}
}

func (d *Downloader) Download(ctx context.Context, fileURL, dstDir string) (string, error) {
	if fileURL == "" {
		return "", fmt.Errorf("package downloader fileURL is empty")
	}
	if dstDir == "" {
		return "", fmt.Errorf("package downloader DST_DIR is empty")
	}

	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return "", fmt.Errorf("create dst dir: %w", err)
	}

	fileName := path.Base(fileURL)
	if fileName == "." || fileName == "/" || fileName == "" {
		fileName = "archive.zip"
	}

	dstPath := filepath.Join(dstDir, fileName)
	tmpPath := dstPath + ".part"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	file, err := os.Create(tmpPath)
	if err != nil {
		return "", fmt.Errorf("create file: %w", err)
	}

	_, copyErr := io.Copy(file, resp.Body)
	closeErr := file.Close()

	if copyErr != nil {
		_ = os.Remove(tmpPath)
		return "", fmt.Errorf("write file: %w", copyErr)
	}

	if closeErr != nil {
		_ = os.Remove(tmpPath)
		return "", fmt.Errorf("close file: %w", closeErr)
	}

	if err := os.Rename(tmpPath, dstPath); err != nil {
		_ = os.Remove(tmpPath)
		return "", fmt.Errorf("rename file: %w", err)
	}

	return dstPath, nil
}
