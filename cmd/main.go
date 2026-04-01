package main

import (
	"fmt"
	"net/http"

	"github.com/iwwwanow/owo/internal"
)

func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v\n", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

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

	http.HandleFunc("/", recoveryMiddleware(controller.ProcessRequest()))

	internal.LaunchServer(port)
}
