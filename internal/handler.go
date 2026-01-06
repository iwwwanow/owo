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

func (handler *Handler) HandleResource(requestPath string) (template.HTML, error) {
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
	// props.Title = "resource-page-title"
	props.Resource = resourceData
	props.Resources = childResourcesData

	return handler.renderer.RenderResourcePage(&props)
}

func (handler *Handler) HandleStatic(requestPath string) StaticFileData {
	staticFileData := handler.repository.GetStaticFileData(requestPath)
	return staticFileData
}
