package internal

import (
	"fmt"
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
		hostName := req.Host
		requestPath := strings.TrimPrefix(req.URL.Path, "/")

		if req.URL.Query().Has("static") {
			fmt.Println("handle static")
			controller.handleStaticRoute(res, req, UploadsDir, requestPath)
			return
		}

		if strings.HasPrefix(requestPath, "static/") {
			controller.handleStaticRoute(
				res,
				req,
				StaticDir,
				strings.TrimPrefix(requestPath, "static/"),
			)
			return
		}

		controller.handleResourceRoute(res, req, requestPath, hostName)
	}
}

func (controller *Controller) handleResourceRoute(
	res http.ResponseWriter,
	req *http.Request,
	requestPath string,
	hostName string,
) {
	htmlContent, err := controller.handler.HandleResource(requestPath, hostName)
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
	staticPath string,
	requestPath string,
) {
	staticFileData := controller.handler.HandleStatic(staticPath, requestPath)

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
