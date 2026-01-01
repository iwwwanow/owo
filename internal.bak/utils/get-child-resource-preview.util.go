package utils

import (
	"github.com/iwwwanow/owwo/internal/configs"
	"os"
)

func GetChildResourcePreview(childResourceFullPath string) string {
	content, err := os.ReadFile(childResourceFullPath)
	preview := string(content)
	if err == nil {

		if len(preview) > configs.PreviewMaxLength {
			preview = preview[:configs.PreviewMaxLength] + "..."
		}

	}
	return preview
}
