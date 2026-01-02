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

	htmlContent, err := controller.handler.HandleResource(requestPath)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(htmlContent))
}

func (controller *Controller) handleStaticRoute(res http.ResponseWriter, req *http.Request, requestPath string) {
	controller.handler.HandleStatic(requestPath)
}
