package internal

// # отрисовка данных, которые пришли из репозитория
// # работа с шаблонами. логика подготовки данных из renderer для шаблона

import (
	"html/template"
)

type Renderer struct {
}

type ResourcePageData struct {
	Title     string
	Meta      ResourceMetaData
	Resource  ResourceData
	Resources []ResourceData
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

// TODO: is html needed?
func (renderer *Renderer) RenderResourcePage() template.HTML {
	return "resource page"
}
