package internal

// обработка входящих с клиента данных (req), подготовка их для хендлерва
// подгготовка отвера для reponse

import (
	"net/http"
	"os"
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

		if req.URL.Query().Get("static") != "" {
			controller.handleStaticRoute(res, req, requestPath)
			return
		}

		if strings.HasPrefix(requestPath, "static/") {
			controller.handleStaticRoute(res, req, strings.TrimPrefix(requestPath, "static/"))
			return
		}

		if isStaticFile(requestPath) {
			controller.handleStaticRoute(res, req, requestPath)
			return
		}

		controller.handleResourceRoute(res, req, requestPath)
	}
}

func (controller *Controller) handleResourceRoute(
	res http.ResponseWriter,
	req *http.Request,
	requestPath string,
) {
	htmlContent, err := controller.handler.HandleResource(requestPath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(htmlContent))
}

func (controller *Controller) handleStaticRoute(
	res http.ResponseWriter,
	req *http.Request,
	requestPath string,
) {
	staticFileData := controller.handler.HandleStatic(requestPath)

	if _, err := os.Stat(staticFileData.Path); os.IsNotExist(err) {
		http.NotFound(res, req)
		return
	}

	switch staticFileData.Ext {
	case ".css":
		res.Header().Set("Content-Type", "text/css")
	case ".js":
		res.Header().Set("Content-Type", "application/javascript")
	case ".png":
		res.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		res.Header().Set("Content-Type", "image/jpeg")
	case ".svg":
		res.Header().Set("Content-Type", "image/svg+xml")
	}

	http.ServeFile(res, req, staticFileData.Path)
}

func isStaticFile(path string) bool {
	staticExtensions := []string{
		".css",
		".js",
		".png",
		".jpg",
		".jpeg",
		".gif",
		".svg",
		".ico",
		".woff",
		".woff2",
		".ttf",
		".webmanifest",
	}

	for _, ext := range staticExtensions {
		if strings.HasSuffix(strings.ToLower(path), ext) {
			return true
		}
	}
	return false
}
