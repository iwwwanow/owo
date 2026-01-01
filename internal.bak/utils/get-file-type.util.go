package utils

import (
	"github.com/iwwwanow/owwo/internal/constants"
	"os"
	"path/filepath"
	"strings"
)

func GetFileType(filename string, info os.FileInfo) string {
	if info.IsDir() {
		return constants.FileTypeDir
	}

	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp":
		return constants.FileTypeImage
	case ".txt", ".md", ".csv", ".json", ".xml", ".html", ".css", ".js":
		return constants.FileTypeText
	default:
		return constants.FileTypeOther
	}
}
