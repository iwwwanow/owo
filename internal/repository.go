package internal

// # работа с данными на сервере
// # основное мясо будет здесь

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
	Content  string
	Preview  string
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository *Repository) GetResourceData(resourcePath string, reqursive bool, parentResouceData *ResourceData) (ResourceData, []ResourceData) {
	var resourceData ResourceData
	var childResourcesData []ResourceData
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

	if reqursive {
		switch resourceData.Type {
		case fileTypeDir:
			childResourcesData = getDirectoryData(repository, &resourceData)
		default:
			prepareResourceData(&resourceData)
		}
	}

	if !reqursive {
		switch resourceData.Type {
		// TODO: переделай. все превью. вся работа на превью будет через meta
		case fileTypeText:
			resourceData.Preview = setResourcePreview(&resourceData)
		case fileTypeOther:
			resourceData.Preview = setResourcePreview(&resourceData)
		case fileTypeDir:
			setDirectoryPreview(parentResouceData, &resourceData)
		default:
			prepareResourceData(&resourceData)
		}
	}

	return resourceData, childResourcesData
}

func getDirectoryData(repository *Repository, resourceData *ResourceData) []ResourceData {
	var resourcesData []ResourceData

	// TODO: git pull
	// utils.GitPullIfNeeded(resourceFullPath)

	files, err := os.ReadDir(resourceData.FullPath)
	if err != nil {
		// TODO: если нет дочерних файлов в директории. что делать?
		// return
	}

	for _, childFile := range files {
		childResourcePath := filepath.Join(resourceData.Path, childFile.Name())
		childResourceData, _ := repository.GetResourceData(childResourcePath, false, resourceData)
		resourcesData = append(resourcesData, childResourceData)
	}

	return resourcesData
}

func prepareResourceData(resourceData *ResourceData) {
	content, err := os.ReadFile(resourceData.FullPath)
	if err == nil {
		resourceData.Content = string(content)
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

func setResourcePreview(resourceData *ResourceData) string {
	content, err := os.ReadFile(resourceData.FullPath)
	preview := string(content)

	if err == nil {
		// TODO: move it logic into renderer
		if len(preview) > PreviewMaxLength {
			preview = preview[:PreviewMaxLength] + "..."
		}

	}

	resourceData.Meta.Description = preview
}

func setDirectoryPreview(parentResourceData *ResourceData, childResourceData *ResourceData) {
	switch childResourceData.Name {
	// TODO: refactor
	case MetaDirName:
		metaDirPath := childResourceData.FullPath

		htmlPath := filepath.Join(metaDirPath, MetaHtmlName)
		if _, err := os.Stat(htmlPath); err == nil {
			parentResourceData.Meta.HtmlPath = filepath.Join(childResourceData.Path, MetaHtmlName)
		}

		cssPath := filepath.Join(metaDirPath, MetaCssName)
		if _, err := os.Stat(cssPath); err == nil {
			parentResourceData.Meta.CssPath = filepath.Join(childResourceData.Path, MetaCssName)
		}

		jsPath := filepath.Join(metaDirPath, MetaJsName)
		if _, err := os.Stat(jsPath); err == nil {
			parentResourceData.Meta.JsPath = filepath.Join(childResourceData.Path, MetaJsName)
		}

		mdPath := filepath.Join(metaDirPath, MetaMdName)
		if _, err := os.Stat(mdPath); err == nil {
			parentResourceData.Meta.MdPath = filepath.Join(childResourceData.Path, MetaMdName)
		}
	default:
		coverPath := findCoverForResource(childResourceData.Path, childResourceData.FullPath)
		if coverPath != "" {
			childResourceData.Meta.CoverPath = coverPath
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

// TODO: to renderer
// func convertMDToHTML(mdContent []byte) template.HTML {
// 	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
// 	p := parser.NewWithExtensions(extensions)
//
// 	doc := p.Parse(mdContent)
//
// 	htmlFlags := html.CommonFlags | html.HrefTargetBlank
// 	opts := html.RendererOptions{Flags: htmlFlags}
// 	renderer := html.NewRenderer(opts)
//
// 	htmlContent := markdown.Render(doc, renderer)
//
// 	return template.HTML(htmlContent)
// }

// // TODO: to renderer
// mdPath := filepath.Join(metaDirPath, MetaMdName)
// if _, err := os.Stat(mdPath); err == nil {
// 	content, err := os.ReadFile(mdPath)
// 	if err == nil {
// 		parentResourceData.Meta.MdContent = convertMDToHTML(content)
// 	}
// }
