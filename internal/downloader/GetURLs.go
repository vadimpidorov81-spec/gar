package downloader

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const lastDownloadInfoURL = "https://fias.nalog.ru/WebServices/Public/GetLastDownloadFileInfo"

type LastDownloadInfo struct {
	VersionID      int64  `json:"VersionId"`
	GarXMLFullURL  string `json:"GarXMLFullURL"`
	GarXMLDeltaURL string `json:"GarXMLDeltaURL"`
	Date           string `json:"Date"`
}

func GetArchiveURLs(ctx context.Context, client *http.Client) (fullURL string, deltaURL string, err error) {
	if client == nil {
		client = &http.Client{}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lastDownloadInfoURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var info LastDownloadInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return "", "", fmt.Errorf("decode response: %w", err)
	}

	if info.GarXMLFullURL == "" {
		return "", "", fmt.Errorf("empty GarXMLFullURL")
	}
	if info.GarXMLDeltaURL == "" {
		return "", "", fmt.Errorf("empty GarXMLDeltaURL")
	}

	return info.GarXMLFullURL, info.GarXMLDeltaURL, nil
}
