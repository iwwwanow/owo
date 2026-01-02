package internal

// обработка входящих с клиента данных (req), подготовка их для хендлерва
// подгготовка отвера для reponse

import (
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
	return func(res http.ResponseWriter, req *http.Request) {
		requestPath := strings.TrimPrefix(req.URL.Path, "/")

		controller.handleResourceRoute(res, req, requestPath)
		controller.handleStaticRoute(res, req, requestPath)
	}
}

func (controller *Controller) handleResourceRoute(res http.ResponseWriter, req *http.Request, requestPath string) {

	resourcePage := controller.handler.HandleResource(requestPath)

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Hello from: " + resourcePage))
}

func (controller *Controller) handleStaticRoute(res http.ResponseWriter, req *http.Request, requestPath string) {
	controller.handler.HandleStatic(requestPath)
}
