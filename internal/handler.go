package internal

// TODO: большую часть переносить в renderer

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Handler struct {
	renderer   Renderer
	repository Repository
}

func NewHandler(renderer Renderer, repository Repository) *Handler {
	return &Handler{
		renderer:   renderer,
		repository: repository,
	}
}

func (handler *Handler) HandleResource(requestPath string, hostName string) (template.HTML, error) {
	resourcePath := requestPath

	var resourceData ResourceData
	var childResourcesData []ResourceData

	// TODO: git pull
	// utils.GitPullIfNeeded(resourceFullPath)

	handler.repository.SetResourceData(resourcePath, &resourceData)
	handler.repository.SetResourceStaticData(&resourceData, &resourceData.Static)
	handler.repository.SetResourceMetaData(&resourceData, &resourceData.Meta)

	handler.repository.SetChildResourcesData(&resourceData, &childResourcesData)

	var props ResourcePageProps
	props.Header.Title = hostName

	handler.mapDataToProps(&props, &resourceData, &childResourcesData)

	return handler.renderer.RenderResourcePage(&props)
}

func (handler *Handler) HandleStatic(staticDir string, requestPath string) StaticFileData {
	staticFileData := handler.repository.GetStaticFileData(staticDir, requestPath)
	return staticFileData
}

func (handler *Handler) HandleImageResize(requestPath, width, height string) StaticFileData {
	cachedPath, err := handler.repository.GetResizedImagePath(requestPath, width, height)
	if err != nil {
		return handler.repository.GetStaticFileData(UploadsDir, requestPath)
	}
	return StaticFileData{
		Path: cachedPath,
		Ext:  filepath.Ext(cachedPath),
	}
}

func (handler *Handler) mapDataToProps(
	props *ResourcePageProps,
	resourceData *ResourceData,
	childResourcesData *[]ResourceData,
) {
	props.Page.Title = resourceData.Meta.Title
	props.Page.Description = resourceData.Meta.Description
	props.Page.Html = resourceData.Static.HtmlPath
	props.Page.Css = transliteratePathSegments(resourceData.Static.CssPath)
	props.Page.Js = transliteratePathSegments(resourceData.Static.JsPath)
	props.Page.Cover = transliteratePathSegments(resourceData.Static.CoverPath)

	// TODO: для контента, возможно, имеет смысл добавить отдельный объект. нужно перебрать его и на уровне шаблонов
	props.Resource.Type = resourceData.Type
	// TODO:
	props.Resource.Content = resourceData.Static.Content
	if props.Resource.Type == fileTypeImage {
		props.Resource.ContentPath = transliteratePathSegments(resourceData.Path)
	}
	if props.Resource.Type == fileTypeVideo {
		props.Resource.ContentPath = transliteratePathSegments(resourceData.Path)
	}
	// TODO: renderer
	if resourceData.Static.MdPath != "" {
		// TODO: contentType?
		props.Resource.Type = "html"
		mdFullPath := filepath.Join(UploadsDir, resourceData.Static.MdPath)
		content, err := os.ReadFile(mdFullPath)
		if err == nil {
			props.Resource.HtmlContent = convertMDToHTML(content)
		}
	}

	props.Resources = []ChildResourceProps{}
	props.HiddenResources = []ChildResourceProps{}
	props.Sections = []SectionProps{}

	if childResourcesData != nil {
		for _, childResourceData := range *childResourcesData {
			var childResourceProps ChildResourceProps
			childResourceProps.ID = "card-" + strings.NewReplacer("/", "-", " ", "-").Replace(transliteratePathSegments(childResourceData.Path))
			childResourceProps.Path = transliteratePathSegments(childResourceData.Path)
			if childResourceData.Type == fileTypeLink {
				childResourceProps.Title = strings.TrimSuffix(childResourceData.Name, ".link")
			} else {
				childResourceProps.Title = childResourceData.Name
			}
			childResourceProps.Description = childResourceData.Meta.Description
			childResourceProps.Cover = transliteratePathSegments(childResourceData.Static.CoverPath)
			childResourceProps.CoverIsVideo = isVideoExt(childResourceData.Static.CoverPath)
			childResourceProps.IsDir = childResourceData.Type == fileTypeDir

			switch {
			case strings.HasPrefix(childResourceData.Name, "."):
				props.HiddenResources = append(props.HiddenResources, childResourceProps)

			case strings.HasPrefix(childResourceData.Name, "_") && childResourceData.Type == fileTypeDir:
				var grandchildren []ResourceData
				handler.repository.SetChildResourcesData(&childResourceData, &grandchildren)

				label := strings.TrimPrefix(childResourceData.Name, "_")
				section := SectionProps{
					Label:     label,
					Slug:      strings.ToLower(strings.NewReplacer(" ", "-").Replace(TransliterateToLatin(label))),
					Resources: []ChildResourceProps{},
				}
				for _, gc := range grandchildren {
					if strings.HasPrefix(gc.Name, ".") {
						continue
					}
					var gcProps ChildResourceProps
					gcProps.ID = "card-" + strings.NewReplacer("/", "-", " ", "-").Replace(transliteratePathSegments(gc.Path))
					gcProps.Path = transliteratePathSegments(gc.Path)
					if gc.Type == fileTypeLink {
						gcProps.Title = strings.TrimSuffix(gc.Name, ".link")
					} else {
						gcProps.Title = gc.Name
					}
					gcProps.Description = gc.Meta.Description
					gcProps.Cover = transliteratePathSegments(gc.Static.CoverPath)
					gcProps.CoverIsVideo = isVideoExt(gc.Static.CoverPath)
					gcProps.IsDir = gc.Type == fileTypeDir
					section.Resources = append(section.Resources, gcProps)
				}
				props.Sections = append(props.Sections, section)

			default:
				props.Resources = append(props.Resources, childResourceProps)
			}
		}
	}

	if strings.HasPrefix(resourceData.Name, "_") {
		props.HiddenResources = append(props.HiddenResources, props.Resources...)
		props.Resources = []ChildResourceProps{}
	}
}

func isVideoExt(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".mp4" || ext == ".webm" || ext == ".mov"
}

// TODO: to renderer
func convertMDToHTML(mdContent []byte) template.HTML {
	fmt.Println("convertMDToHTML")
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	doc := p.Parse(mdContent)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	htmlContent := markdown.Render(doc, renderer)

	return template.HTML(htmlContent)
}
