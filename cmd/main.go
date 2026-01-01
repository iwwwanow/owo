package main

import (
	"net/http"

	"github.com/iwwwanow/owwo/internal/configs"
	"github.com/iwwwanow/owwo/internal/handlers"
	"github.com/iwwwanow/owwo/internal/utils"
)

func main() {

	port := utils.GetServerPort()
	tmpl := utils.GetTemplateInstance()

	http.HandleFunc("/", handlers.ResouceHandler(tmpl))
	http.Handle("/static/", http.StripPrefix("/static/", handlers.StaticHandler(configs.StaticDir)))
	http.Handle("/public/", http.StripPrefix("/public/", handlers.StaticHandler(configs.PublicDir)))

	utils.LaunchServer(port)
}
