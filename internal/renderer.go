package internal

// # отрисовка данных, которые пришли из репозитория
// # работа с шаблонами. логика подготовки данных из renderer для шаблона

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

const (
	templatesDir = "web/templates"
)

type Renderer struct {
	templates map[string]*template.Template
}

type HeaderProps struct {
	Title string
}

type PageProps struct {
	Title       string // from resource name (file or directory)
	Description string // from .meta
	Html        string // is needed?
	Css         string
	Js          string
	Cover       string
}

type ResourceProps struct {
	Type string // use consts // iframe, image, html

	// if image - prerender it to html as string
	// if md or html - prerender it to html as string
	ContentPath string
	HtmlContent template.HTML
	Content     string
}

type ChildResourceProps struct {
	Path        string
	Title       string // filename or dirname
	Description string // prerender it from html or md
	Cover       string // .meta/cover
}

type ResourcePageProps struct {
	Header    HeaderProps
	Page      PageProps
	Resource  ResourceProps
	Resources []ChildResourceProps
}

func NewRenderer() (*Renderer, error) {
	renderer := &Renderer{
		templates: make(map[string]*template.Template),
	}

	// Загружаем все шаблоны
	pattern := filepath.Join(templatesDir, "*.html")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fmt.Printf(file)
		name := filepath.Base(file)
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			return nil, err
		}
		renderer.templates[name] = tmpl
	}

	return renderer, nil
}

// TODO: is html needed?
func (renderer *Renderer) RenderResourcePage(props *ResourcePageProps) (template.HTML, error) {
	fmt.Println("log on renderer")

	tmpl := GetTemplateInstance()
	// tmpl, exists := renderer.templates["index.html"]
	// if !exists {
	// 	return "", fmt.Errorf("template 'index.html' not found")
	// }

	// Создаем буфер для результата
	var buf bytes.Buffer

	// Исполняем шаблон
	err := tmpl.Execute(&buf, props)
	if err != nil {
		return "", err
	}

	// Конвертируем в template.HTML для безопасной вставки
	return template.HTML(buf.String()), nil
}

// err = tmpl.Execute(w, pageData)
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }

// // TODO: to renderer
// mdPath := filepath.Join(metaDirPath, MetaMdName)
// if _, err := os.Stat(mdPath); err == nil {
// 	content, err := os.ReadFile(mdPath)
// 	if err == nil {
// 		parentResourceData.Meta.MdContent = convertMDToHTML(content)
// 	}
// }
