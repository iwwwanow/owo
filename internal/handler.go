package internal

import (
	"html/template"
)

// # основная логика приложения
// # работа с пришедшими с контроллера данными. они должны быть адаптированны для хендлера
// # оркестрация и repository и renderer

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
	// TODO: title logic
	props.Header.Title = hostName
	// props.Resource = resourceData
	// props.Resources = childResourcesData

	mapDataToProps(&props, &resourceData, &childResourcesData)

	return handler.renderer.RenderResourcePage(&props)
}

func (handler *Handler) HandleStatic(staticDir string, requestPath string) StaticFileData {
	staticFileData := handler.repository.GetStaticFileData(staticDir, requestPath)
	return staticFileData
}

func mapDataToProps(
	props *ResourcePageProps,
	resourceData *ResourceData,
	childResourcesData *[]ResourceData,
) {
	props.Page.Title = resourceData.Meta.Title
	props.Page.Description = resourceData.Meta.Description
	props.Page.Html = resourceData.Static.HtmlPath
	props.Page.Css = resourceData.Static.CssPath
	props.Page.Js = resourceData.Static.JsPath
	props.Page.Cover = resourceData.Static.CoverPath

	props.Resource.Type = resourceData.Type
	// TODO:
	props.Resource.HtmlContent = resourceData.Static.Content
	if props.Resource.Type == fileTypeImage {
		props.Resource.ContentPath = resourceData.Path
	}

	if props.Resources == nil {
		props.Resources = []ChildResourceProps{}
	}

	if childResourcesData != nil {
		for _, childResourceData := range *childResourcesData {
			var childResourceProps ChildResourceProps
			childResourceProps.Path = childResourceData.Path
			// childResourceProps.Title = childResourceData.Meta.Title
			childResourceProps.Title = childResourceData.Name
			childResourceProps.Description = childResourceData.Meta.Description
			childResourceProps.Cover = childResourceData.Static.CoverPath

			props.Resources = append(props.Resources, childResourceProps)
		}
	}
}
