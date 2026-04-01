package internal

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

func TransliterateToLatin(s string) string {
	return strings.ReplaceAll(unidecode.Unidecode(s), "'", "")
}

func transliteratePathSegments(path string) string {
	parts := strings.Split(path, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(TransliterateToLatin(p))
	}
	return strings.Join(parts, "/")
}

func resolveTransliteratedSegment(dirFullPath, segment string) string {
	entries, err := os.ReadDir(dirFullPath)
	if err != nil {
		return segment
	}
	for _, entry := range entries {
		if TransliterateToLatin(entry.Name()) == segment {
			return entry.Name()
		}
	}
	return segment
}

func ResolveTransliteratedPath(baseDir, transliteratedPath string) string {
	segments := strings.Split(transliteratedPath, "/")
	resolved := make([]string, 0, len(segments))
	currentDir := baseDir

	for _, seg := range segments {
		if seg == "" {
			continue
		}
		actual := resolveTransliteratedSegment(currentDir, seg)
		resolved = append(resolved, actual)
		currentDir = filepath.Join(currentDir, actual)
	}

	return strings.Join(resolved, "/")
}
