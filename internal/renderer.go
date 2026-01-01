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

// // TODO: base title on domainname
// pageData.Title = "iwwwanowwwwwww"
// pageData.Resource = resource
// pageData.Resources = resources
//
// err = tmpl.Execute(w, pageData)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }

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
