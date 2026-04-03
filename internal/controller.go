package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
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

		if requestPath == "backup" {
			controller.handleBackupRoute(res, req)
			return
		}

		if requestPath == "cache/clear" && req.Method == http.MethodPost {
			controller.handleCacheClearRoute(res, req)
			return
		}

		if req.URL.Query().Has("static") {
			resolvedPath := ResolveTransliteratedPath(UploadsDir, requestPath)
			width := req.URL.Query().Get("width")
			height := req.URL.Query().Get("height")
			if width != "" || height != "" {
				controller.handleImageResizeRoute(res, req, resolvedPath, width, height)
				return
			}
			controller.handleStaticRoute(res, req, UploadsDir, resolvedPath)
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

		resolvedPath := ResolveTransliteratedPath(UploadsDir, requestPath)
		if linkTarget := GetResourceLink(resolvedPath); linkTarget != "" {
			decoded, _ := url.PathUnescape(linkTarget)
			http.Redirect(res, req, "/"+transliteratePathSegments(decoded), http.StatusFound)
			return
		}
		controller.handleResourceRoute(res, req, resolvedPath, hostName)
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

func (controller *Controller) handleImageResizeRoute(
	res http.ResponseWriter,
	req *http.Request,
	requestPath, width, height string,
) {
	staticFileData := controller.handler.HandleImageResize(requestPath, width, height)

	if _, err := os.Stat(staticFileData.Path); os.IsNotExist(err) {
		http.NotFound(res, req)
		return
	}

	switch staticFileData.Ext {
	case ".png":
		res.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		res.Header().Set("Content-Type", "image/jpeg")
	case ".webp":
		res.Header().Set("Content-Type", "image/webp")
	}

	http.ServeFile(res, req, staticFileData.Path)
}

func (controller *Controller) handleBackupRoute(res http.ResponseWriter, req *http.Request) {
	filename := fmt.Sprintf("owo-backup-%s.zip", time.Now().Format("20060102-150405"))
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	zw := zip.NewWriter(res)
	defer zw.Close()

	err := filepath.Walk(UploadsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, err := filepath.Rel(UploadsDir, path)
		if err != nil {
			return err
		}
		w, err := zw.Create(rel)
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	})

	if err != nil {
		fmt.Printf("backup error: %v\n", err)
	}
}

func (controller *Controller) handleCacheClearRoute(res http.ResponseWriter, req *http.Request) {
	if err := os.RemoveAll(CacheDir); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusNoContent)
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
