package internal

import (
	"fmt"
)

// # основная логика приложения
// # работа с пришедшими с контроллера данными. они должны быть адаптированны для хендлера
// # оркестрация и repository и renderer

type Handler struct {
	renderer   Renderer
	repository Repository
}

func NewHandler(renderer Renderer, repository Repository) *Handler {
	return &Handler{
		renderer:   renderer,
		repository: repository,
	}
}

func (handler *Handler) HandleResource(requestPath string) {
	fmt.Printf("log on handler: %s", requestPath)
	fmt.Println()
}

func (handler *Handler) HandleStatic(requestPath string) {
	fmt.Printf("log on handle static: %s", requestPath)
	fmt.Println()
}
