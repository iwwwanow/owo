package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// TODO: refactor

var templates = []string{
	"web/templates/index.html.tmpl",
	"web/templates/pages/resource.page.html.tmpl",

	"web/templates/fragments/head.fragment.html.tmpl",
	"web/templates/fragments/header.fragment.html.tmpl",
	"web/templates/fragments/content.fragment.html.tmpl",
	"web/templates/fragments/footer.fragment.html.tmpl",

	"web/templates/components/card.component.html.tmpl",
	"web/templates/components/iframe.component.html.tmpl",
	"web/templates/components/html.component.html.tmpl",
	"web/templates/components/image.component.html.tmpl",
	"web/templates/components/code.component.html.tmpl",
	"web/templates/components/hr.component.html.tmpl",
}

func GetServerPort() int {
	portStr := os.Getenv("PORT")

	if portStr == "" {
		portStr = "3000"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("error portStr formatting:", err)
	}

	return port
}

func GetTemplateInstance() *template.Template {
	tmpl := template.Must(template.ParseFiles(templates...))

	return tmpl
}

func LaunchServer(port int) {
	address := fmt.Sprintf("http://localhost:%d", port)
	fmt.Printf("server listening on %s\n", address)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("error with launch:", err)
	}
}
