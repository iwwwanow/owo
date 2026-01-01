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
	"web/templates/index.html",
	"web/templates/pages/resource.page.html",

	"web/templates/fragments/head.fragment.html",
	"web/templates/fragments/header.fragment.html",
	"web/templates/fragments/content.fragment.html",
	"web/templates/fragments/footer.fragment.html",

	"web/templates/components/resource-card.component.html",
	"web/templates/components/iframe-content.component.html",
	"web/templates/components/markdown-content.component.html",
	"web/templates/components/resource-content.component.html",
	"web/templates/components/resource-content_image.component.html",
	"web/templates/components/resource-content_other.component.html",
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
