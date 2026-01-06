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
	StaticHtmlName  = "index.html"
	StaticMdName    = "index.md"
	StaticCssName   = "index.css"
	StaticJsName    = "index.js"
	StaticCoverName = "cover"
)

type Repository struct {
}

type MetaData struct {
	Title       string
	Description string
}

type StaticData struct {
	HtmlPath  string
	MdPath    string
	CssPath   string
	JsPath    string
	CoverPath string
}

type ResourceData struct {
	// TODO: why caps?
	Meta     MetaData
	Static   StaticData
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

func (repository *Repository) SetResourceData(resourcePath string, resourceData *ResourceData) {
	// TODO: publicdirpath
	resourceData.FullPath = filepath.Join("publicdirpath", resourcePath)

	resourceFileInfo, err := os.Stat(resourceData.FullPath)
	if err != nil {
		// TODO: 404 exception
		// http.NotFound(w, r)
		return
	}

	resourceData.Path = resourcePath
	resourceData.Name = resourceFileInfo.Name()
	resourceData.Type = getFileType(resourceData.Name, resourceFileInfo)
}

func (repository *Repository) SetResourceStaticData(resourceData *ResourceData, resourceStaticData *StaticData) {

	switch resourceData.Type {
	// case fileTypeText:
	// 	setResourceMetaDescription(resourceData, &resourceStaticData)
	// case fileTypeOther:
	// 	setResourceMetaDescription(resourceData, &resourceStaticData)
	case fileTypeDir:
		setDirectoryStatic(resourceData, resourceStaticData)
	default:
		setDirectoryStatic(resourceData, resourceStaticData)
	}
}

func (repository *Repository) SetResourceMetaData(resourceData *ResourceData, resourceMetaData *MetaData) {

	switch resourceData.Type {
	case fileTypeText:
		setResourceMetaDescription(resourceData, resourceMetaData)
	case fileTypeOther:
		setResourceMetaDescription(resourceData, resourceMetaData)
		// case fileTypeDir:
		// 	setDirectoryStatic(resourceData, resourceStaticData)
		// default:
		// 	setDirectoryStatic(resourceData, resourceStaticData)
	}
}

func (repository *Repository) SetChildResourcesData(resourceData *ResourceData, childResourcesData *[]ResourceData) {
	childResourceDirs, err := os.ReadDir(resourceData.FullPath)
	if err != nil {
		// TODO: если нет дочерних файлов в директории. что делать?
		// return
	}

	for _, childResourceDir := range childResourceDirs {
		childResourcePath := filepath.Join(resourceData.Path, childResourceDir.Name())
		var childResourceData ResourceData

		repository.SetResourceData(childResourcePath, &childResourceData)
		repository.SetResourceMetaData(&childResourceData, &childResourceData.Meta)
		repository.SetResourceStaticData(&childResourceData, &childResourceData.Static)

		*childResourcesData = append(*childResourcesData, childResourceData)
	}
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

func setDirectoryStatic(resourceData *ResourceData, static *StaticData) {
	metaDirPath := filepath.Join(resourceData.Path, StaticHtmlName)
	metaDirFullPath := filepath.Join(resourceData.FullPath, StaticHtmlName)

	if _, err := os.Stat(metaDirFullPath); err == nil {
		htmlPath := filepath.Join(metaDirFullPath, StaticHtmlName)
		if _, err := os.Stat(htmlPath); err == nil {
			static.HtmlPath = filepath.Join(metaDirPath, StaticHtmlName)
		}

		cssPath := filepath.Join(metaDirFullPath, StaticCssName)
		if _, err := os.Stat(cssPath); err == nil {
			static.CssPath = filepath.Join(metaDirPath, StaticCssName)
		}

		jsPath := filepath.Join(metaDirFullPath, StaticJsName)
		if _, err := os.Stat(jsPath); err == nil {
			static.JsPath = filepath.Join(metaDirPath, StaticJsName)
		}

		mdPath := filepath.Join(metaDirFullPath, StaticMdName)
		if _, err := os.Stat(mdPath); err == nil {
			static.MdPath = filepath.Join(metaDirPath, StaticMdName)
		}

		coverPath := findCoverForResource(resourceData.Path, resourceData.FullPath)
		if coverPath != "" {
			static.CoverPath = coverPath
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
		coverPath := filepath.Join(metaDirPath, StaticCoverName+ext)
		if _, err := os.Stat(coverPath); err == nil {
			return filepath.Join(resourcePath, MetaDirName, StaticCoverName+ext)
		}
	}

	return ""
}
