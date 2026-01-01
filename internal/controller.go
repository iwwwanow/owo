package internal

// обработка входящих с клиента данных (req), подготовка их для хендлерва
// подгготовка отвера для reponse

import (
	"fmt"
	"net/http"
	"strings"
)

type Controller struct {
	handler Handler
}

func NewController(handler Handler) *Controller {
	return &Controller{
		handler: handler,
	}
}

func (controller *Controller) ProcessRequest() http.HandlerFunc {
	fmt.Println("log on controller 1")
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("log on controller 2")

		requestPath := strings.TrimPrefix(req.URL.Path, "/")

		controller.handler.HandleResource(requestPath)
		controller.handler.HandleStatic(requestPath)
	}
}

// func Controller(res http.ResponseWriter, req *http.Request, handler HandlerType) {
// 	resourcePath := strings.TrimPrefix(req.URL.Path, "/")
// 	handler(resourcePath)
// }
