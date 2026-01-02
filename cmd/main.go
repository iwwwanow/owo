package main

import (
	"fmt"
	"net/http"

	"github.com/iwwwanow/owo/internal"
)

func main() {

	repository := internal.NewRepository()
	renderer, err := internal.NewRenderer()
	if err != nil {
		fmt.Printf("Renderer not initialized")
		return
	}
	handler := internal.NewHandler(*renderer, *repository)
	controller := internal.NewController(*handler)

	port := internal.GetServerPort()

	http.HandleFunc("/", controller.ProcessRequest())

	internal.LaunchServer(port)
}
