package internal

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	fileTypeImage = "image"
	fileTypeVideo = "video"
	fileTypeText  = "text"
	fileTypeDir   = "directory"
	fileTypeLink  = "link"
	fileTypeOther = "other"
)

const (
	// TODO env
	UploadsDir       = "/var/www/owo/uploads"
	StaticDir        = "/var/www/owo/static"
	CacheDir         = "/var/www/owo/cache"
	MetaDirName      = ".meta"
	PreviewMaxLength = 50
)

const (
	StaticHtmlName  = "index.html"
	StaticMdName    = "index.md"
	StaticCssName   = "index.css"
	StaticJsName    = "index.js"
	StaticCoverName = "cover"
	StaticLinkName  = "index.link"
)

var resizeSemaphore = make(chan struct{}, 3)

type Repository struct{}

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
	Content   string
}

type StaticFileData struct {
	Ext  string
	Path string
}

type ResourceData struct {
	Meta       MetaData
	Static     StaticData
	Path       string
	FullPath   string
	Name       string
	Type       string
	LinkTarget string
}

func NewRepository() *Repository {
	return &Repository{}
}

func GetResourceLink(resourcePath string) string {
	linkFile := filepath.Join(UploadsDir, resourcePath, MetaDirName, StaticLinkName)
	content, err := os.ReadFile(linkFile)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(content))
}

func (repository *Repository) SetResourceData(resourcePath string, resourceData *ResourceData) {
	// TODO: publicdirpath
	resourceData.FullPath = filepath.Join(UploadsDir, resourcePath)

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

func (repository *Repository) SetResourceStaticData(
	resourceData *ResourceData,
	resourceStaticData *StaticData,
) {
	switch resourceData.Type {
	// case fileTypeText:
	// 	setResourceMetaDescription(resourceData, &resourceStaticData)
	case fileTypeImage:
		setImageStaticData(resourceData, resourceStaticData)
	case fileTypeVideo:
		setVideoStaticData(resourceData, resourceStaticData)
	case fileTypeDir:
		setDirectoryStaticData(resourceData, resourceStaticData)
	default:
		setOtherStaticData(resourceData, resourceStaticData)
	}
}

func (repository *Repository) SetResourceMetaData(
	resourceData *ResourceData,
	resourceMetaData *MetaData,
) {
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

func (repository *Repository) SetChildResourcesData(
	resourceData *ResourceData,
	childResourcesData *[]ResourceData,
) {
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

		if link := GetResourceLink(childResourceData.Path); link != "" {
			childResourceData.LinkTarget = link
			if childResourceData.Static.CoverPath == "" {
				decoded, _ := url.PathUnescape(link)
				actualPath := ResolveTransliteratedPath(UploadsDir, decoded)
				var targetData ResourceData
				repository.SetResourceData(actualPath, &targetData)
				repository.SetResourceStaticData(&targetData, &targetData.Static)
				childResourceData.Static.CoverPath = targetData.Static.CoverPath
			}
		}

		if childResourceData.Type == fileTypeLink {
			content, err := os.ReadFile(childResourceData.FullPath)
			if err == nil {
				link := strings.TrimSpace(string(content))
				childResourceData.LinkTarget = link
				decoded, _ := url.PathUnescape(link)
				actualPath := ResolveTransliteratedPath(UploadsDir, decoded)
				var targetData ResourceData
				repository.SetResourceData(actualPath, &targetData)
				repository.SetResourceStaticData(&targetData, &targetData.Static)
				childResourceData.Static.CoverPath = targetData.Static.CoverPath
			}
		}

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
	case ".mp4", ".webm", ".mov":
		return fileTypeVideo
	case ".txt", ".md", ".csv", ".json", ".xml", ".html", ".css", ".js":
		return fileTypeText
	case ".link":
		return fileTypeLink
	default:
		return fileTypeOther
	}
}

func (repository *Repository) GetStaticFileData(
	staticDir string,
	resourcePath string,
) StaticFileData {
	var staticFileData StaticFileData
	staticFileFullPath := filepath.Join(staticDir, filepath.Clean(resourcePath))

	ext := filepath.Ext(staticFileFullPath)

	staticFileData.Path = staticFileFullPath
	staticFileData.Ext = ext

	return staticFileData
}

func setResourceMetaDescription(resourceData *ResourceData, meta *MetaData) {
	content, err := os.ReadFile(resourceData.FullPath)
	description := string(content)

	if err == nil {
		meta.Description = description
	}
}

func setImageStaticData(resourceData *ResourceData, static *StaticData) {
	static.CoverPath = resourceData.Path
}

func setVideoStaticData(resourceData *ResourceData, static *StaticData) {
	static.CoverPath = resourceData.Path
}

func setOtherStaticData(resourceData *ResourceData, static *StaticData) {
	content, err := os.ReadFile(resourceData.FullPath)
	if err != nil {
		// Обработка ошибки чтения файла
		static.Content = ""
		return
	}
	static.Content = string(content)
}

func setDirectoryStaticData(resourceData *ResourceData, static *StaticData) {
	metaDirPath := filepath.Join(resourceData.Path, MetaDirName)
	metaDirFullPath := filepath.Join(resourceData.FullPath, MetaDirName)

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

		coverPath := findDirectoryCover(resourceData.Path, resourceData.FullPath)
		if coverPath != "" {
			static.CoverPath = coverPath
		}
	}
}

func (repository *Repository) GetResizedImagePath(uploadsRelPath, widthStr, heightStr string) (string, error) {
	width, _ := strconv.Atoi(widthStr)
	height, _ := strconv.Atoi(heightStr)

	if width == 0 && height == 0 {
		return filepath.Join(UploadsDir, uploadsRelPath), nil
	}

	dimKey := fmt.Sprintf("%dx%d", width, height)
	cachedPath := filepath.Join(CacheDir, dimKey, uploadsRelPath)

	if _, err := os.Stat(cachedPath); err == nil {
		return cachedPath, nil
	}

	originalPath := filepath.Join(UploadsDir, uploadsRelPath)

	if strings.ToLower(filepath.Ext(uploadsRelPath)) == ".gif" {
		return originalPath, nil
	}

	resizeSemaphore <- struct{}{}
	defer func() { <-resizeSemaphore }()

	src, err := imaging.Open(originalPath, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	resized := imaging.Fill(src, width, height, imaging.Center, imaging.Box)

	if err := os.MkdirAll(filepath.Dir(cachedPath), 0755); err != nil {
		return "", err
	}
	if err := imaging.Save(resized, cachedPath); err != nil {
		return "", err
	}

	return cachedPath, nil
}

func findDirectoryCover(resourcePath, resourceFullPath string) string {
	metaDirPath := filepath.Join(resourceFullPath, MetaDirName)

	if _, err := os.Stat(metaDirPath); err != nil {
		return ""
	}

	files, err := os.ReadDir(metaDirPath)
	if err != nil {
		return ""
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := strings.ToLower(file.Name())

		if !strings.HasPrefix(fileName, strings.ToLower(StaticCoverName)) {
			continue
		}

		// Проверяем расширение
		ext := filepath.Ext(fileName)
		validExts := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp", ".mp4", ".webm", ".mov"}

		if slices.Contains(validExts, ext) {
			return filepath.Join(resourcePath, MetaDirName, file.Name())
		}
	}

	return ""
}
