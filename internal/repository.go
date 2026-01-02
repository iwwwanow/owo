package internal

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	fileTypeImage = "image"
	fileTypeText  = "text"
	fileTypeDir   = "directory"
	fileTypeOther = "other"
)

const (
	// TODO env
	PublicDir        = "/var/www/owwo/shared"
	StaticDir        = "web/static"
	MetaDirName      = ".meta"
	PreviewMaxLength = 50
)

const (
	MetaHtmlName  = "index.html"
	MetaMdName    = "index.md"
	MetaCssName   = "index.css"
	MetaJsName    = "index.js"
	MetaCoverName = "cover"
)

type Repository struct{}

type MetaData struct {
	HtmlPath    string
	CssPath     string
	JsPath      string
	MdPath      string
	CoverPath   string
	Description string
}

type ResourceData struct {
	// TODO: why caps?
	Meta     MetaData
	Path     string
	FullPath string
	Name     string
	Type     string
	// Content  string
	// Preview  string
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository *Repository) GetResourceData(resourcePath string) ResourceData {
	var resourceData ResourceData
	// TODO: publicdirpath
	resourceData.FullPath = filepath.Join("publicdirpath", resourcePath)

	resourceFileInfo, err := os.Stat(resourceData.FullPath)
	if err != nil {
		// TODO: 404 exception
		// http.NotFound(w, r)
		return resourceData
	}

	resourceData.Path = resourcePath
	resourceData.Name = resourceFileInfo.Name()
	resourceData.Type = getFileType(resourceData.Name, resourceFileInfo)

	return resourceData
}

func (repository *Repository) GetResourceMeta(resourceData *ResourceData) MetaData {
	var meta MetaData

	switch resourceData.Type {
	case fileTypeText:
		setResourceMetaDescription(resourceData, &meta)
	case fileTypeOther:
		setResourceMetaDescription(resourceData, &meta)
	case fileTypeDir:
		setDirectoryMeta(resourceData, &meta)
	default:
		setDirectoryMeta(resourceData, &meta)
	}

	return meta
}

func (repository *Repository) GetChildResourceDirs(resourceData *ResourceData) []os.DirEntry {
	files, err := os.ReadDir(resourceData.FullPath)
	if err != nil {
		// TODO: если нет дочерних файлов в директории. что делать?
		// return
	}

	return files
}

func getFileType(filename string, info os.FileInfo) string {

	if info.IsDir() {
		return fileTypeDir
	}

	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp":
		return fileTypeImage
	case ".txt", ".md", ".csv", ".json", ".xml", ".html", ".css", ".js":
		return fileTypeText
	default:
		return fileTypeOther
	}
}

func setResourceMetaDescription(resourceData *ResourceData, meta *MetaData) {
	content, err := os.ReadFile(resourceData.FullPath)
	description := string(content)

	if err == nil {
		meta.Description = description
	}
}

func setDirectoryMeta(resourceData *ResourceData, meta *MetaData) {
	metaDirPath := filepath.Join(resourceData.Path, MetaHtmlName)
	metaDirFullPath := filepath.Join(resourceData.FullPath, MetaHtmlName)

	if _, err := os.Stat(metaDirFullPath); err == nil {
		htmlPath := filepath.Join(metaDirFullPath, MetaHtmlName)
		if _, err := os.Stat(htmlPath); err == nil {
			meta.HtmlPath = filepath.Join(metaDirPath, MetaHtmlName)
		}

		cssPath := filepath.Join(metaDirFullPath, MetaCssName)
		if _, err := os.Stat(cssPath); err == nil {
			meta.CssPath = filepath.Join(metaDirPath, MetaCssName)
		}

		jsPath := filepath.Join(metaDirFullPath, MetaJsName)
		if _, err := os.Stat(jsPath); err == nil {
			meta.JsPath = filepath.Join(metaDirPath, MetaJsName)
		}

		mdPath := filepath.Join(metaDirFullPath, MetaMdName)
		if _, err := os.Stat(mdPath); err == nil {
			meta.MdPath = filepath.Join(metaDirPath, MetaMdName)
		}

		coverPath := findCoverForResource(resourceData.Path, resourceData.FullPath)
		if coverPath != "" {
			meta.CoverPath = coverPath
		}
	}
}

// TODO: refactor
func findCoverForResource(resourcePath, resourceFullPath string) string {
	metaDirPath := filepath.Join(resourceFullPath, MetaDirName)

	if _, err := os.Stat(metaDirPath); err != nil {
		return ""
	}

	coverExtensions := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp"}

	for _, ext := range coverExtensions {
		coverPath := filepath.Join(metaDirPath, MetaCoverName+ext)
		if _, err := os.Stat(coverPath); err == nil {
			return filepath.Join(resourcePath, MetaDirName, MetaCoverName+ext)
		}
	}

	return ""
}
