package internal

import (
	"fmt"
	"html/template"
	"path/filepath"
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

	resourceData = handler.repository.GetResourceData(resourcePath)
	resourceMetaData := handler.repository.GetResourceMeta(&resourceData)

	childResourceDirs := handler.repository.GetChildResourceDirs(&resourceData)
	for _, childResourceDir := range childResourceDirs {
		childResourcePath := filepath.Join(resourceData.Path, childResourceDir.Name())
		var childResourceData ResourceData

		childResourceData = handler.repository.GetResourceData(childResourcePath)
		childResourceData.Meta = handler.repository.GetResourceMeta(&childResourceData)

		childResourcesData = append(childResourcesData, childResourceData)
	}

	var props ResourcePageProps
	// TODO: title logic
	props.Title = "resource-page-title"
	props.Meta = resourceMetaData
	props.Resource = resourceData
	props.Resources = childResourcesData

	return handler.renderer.RenderResourcePage(&props)
}

func (handler *Handler) HandleStatic(requestPath string) {
	fmt.Printf("log on handle static: %s", requestPath)
	fmt.Println()
}
