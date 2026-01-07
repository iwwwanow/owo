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
	"/var/www/owo/templates/index.html.tmpl",
	"/var/www/owo/templates/pages/resource.page.html.tmpl",

	"/var/www/owo/templates/fragments/head.fragment.html.tmpl",
	"/var/www/owo/templates/fragments/header.fragment.html.tmpl",
	"/var/www/owo/templates/fragments/content.fragment.html.tmpl",
	"/var/www/owo/templates/fragments/footer.fragment.html.tmpl",

	"/var/www/owo/templates/components/card.component.html.tmpl",
	"/var/www/owo/templates/components/iframe.component.html.tmpl",
	"/var/www/owo/templates/components/html.component.html.tmpl",
	"/var/www/owo/templates/components/image.component.html.tmpl",
	"/var/www/owo/templates/components/code.component.html.tmpl",
	"/var/www/owo/templates/components/hr.component.html.tmpl",
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
