package main

import (
	"net/http"

	"github.com/iwwwanow/owo/internal"
)

func main() {

	repository := internal.NewRepository()
	renderer := internal.NewRenderer()
	handler := internal.NewHandler(*renderer, *repository)
	controller := internal.NewController(*handler)

	port := internal.GetServerPort()
	// tmpl := internal.GetTemplateInstance()

	http.HandleFunc("/", controller.ProcessRequest())

	internal.LaunchServer(port)
}
