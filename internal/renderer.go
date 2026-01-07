package internal

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
	Title       string
	Description string
	Html        string
	Css         string
	Js          string
	Cover       string
}

type ResourceProps struct {
	Type string

	ContentPath string
	HtmlContent template.HTML
	Content     string
}

type ChildResourceProps struct {
	Path        string
	Title       string
	Description string
	Cover       string
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
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, props)
	if err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}
