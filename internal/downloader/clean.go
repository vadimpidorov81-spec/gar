package downloader

import (
	"fmt"
	"os"
)

func DeleteFolder(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("папка не найдена: %s", path)
		}
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("это не папка: %s", path)
	}

	return os.RemoveAll(path)
}
