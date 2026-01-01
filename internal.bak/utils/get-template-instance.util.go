package utils

import (
	"github.com/iwwwanow/owwo/internal/configs"
	"html/template"
)

func GetTemplateInstance() *template.Template {
	tmpl := template.Must(template.ParseFiles(configs.Templates...))

	return tmpl
}
