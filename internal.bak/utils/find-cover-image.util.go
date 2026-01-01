package utils

import (
	"github.com/iwwwanow/owwo/internal/configs"
	"os"
	"path/filepath"
)

func FindCoverForResource(resourcePath, resourceFullPath string) string {
	metaDirPath := filepath.Join(resourceFullPath, configs.MetaDirName)

	if _, err := os.Stat(metaDirPath); err != nil {
		return ""
	}

	coverExtensions := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp"}

	for _, ext := range coverExtensions {
		coverPath := filepath.Join(metaDirPath, configs.MetaCoverName+ext)
		if _, err := os.Stat(coverPath); err == nil {
			return filepath.Join(resourcePath, configs.MetaDirName, configs.MetaCoverName+ext)
		}
	}

	return ""
}
