package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/iwwwanow/owwo/internal/configs"
	"github.com/iwwwanow/owwo/internal/constants"
	"github.com/iwwwanow/owwo/internal/interfaces"
	"github.com/iwwwanow/owwo/internal/utils"
)

func ResouceHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// resourcePath := strings.TrimPrefix(r.URL.Path, "/")
		// resourceFullPath := filepath.Join(configs.PublicDir, resourcePath)

		var pageData interfaces.PageStruct
		var resource interfaces.ResourceInfoStruct
		var resources []interfaces.ResourceInfoStruct

		// resourceFileInfo, err := os.Stat(resourceFullPath)
		// if err != nil {
		// 	http.NotFound(w, r)
		// 	return
		// }

		// resource.Path = resourcePath
		// resource.Name = resourceFileInfo.Name()
		// resource.Type = utils.GetFileType(resource.Name, resourceFileInfo)

		if resource.Type == constants.FileTypeDir {
			// utils.GitPullIfNeeded(resourceFullPath)

			// if files, err := os.ReadDir(resourceFullPath); err == nil {
			// for _, childResourceFile := range files {
				childResourceFullPath := filepath.Join(resourceFullPath, childResourceFile.Name())
				childResourceInfo, _ := childResourceFile.Info()

				// var childResource interfaces.ResourceInfoStruct
				// childResource.Path = filepath.Join(resourcePath, childResourceFile.Name())
				// childResource.Name = childResourceFile.Name()
				// childResource.Type = utils.GetFileType(childResourceFile.Name(), childResourceInfo)

				// if childResource.Type == constants.FileTypeText ||
					// childResource.Type == constants.FileTypeOther {
					// childResource.Preview = utils.GetChildResourcePreview(childResourceFullPath)
				} else if childResource.Type == constants.FileTypeDir {
					if childResource.Name == configs.MetaDirName {
						// metaDirPath := childResourceFullPath
						// meta := interfaces.MetaStruct{}
						//
						// htmlPath := filepath.Join(metaDirPath, configs.MetaHtmlName)
						// if _, err := os.Stat(htmlPath); err == nil {
						// 	meta.HtmlPath = filepath.Join(childResource.Path, configs.MetaHtmlName)
						// }
						//
						// cssPath := filepath.Join(metaDirPath, configs.MetaCssName)
						// if _, err := os.Stat(cssPath); err == nil {
						// 	meta.CssPath = filepath.Join(childResource.Path, configs.MetaCssName)
						// }
						//
						// jsPath := filepath.Join(metaDirPath, configs.MetaJsName)
						// if _, err := os.Stat(jsPath); err == nil {
						// 	meta.JsPath = filepath.Join(childResource.Path, configs.MetaJsName)
						// }
						//
						// mdPath := filepath.Join(metaDirPath, configs.MetaMdName)
						// if _, err := os.Stat(mdPath); err == nil {
						// 	content, err := os.ReadFile(mdPath)
						// 	if err == nil {
						// 		meta.MdContent = utils.ConvertMDToHTML(content)
						// 	}
						// }
						//
						// pageData.Meta = meta
					} else {
						// coverPath := utils.FindCoverForResource(childResource.Path, childResourceFullPath)
						// fmt.Printf("%s", coverPath)
						// if coverPath != "" {
						// 	childResource.Cover = coverPath
						// }
					}
				}

				resources = append(resources, childResource)
			}
			// }
		} else {
			// content, err := os.ReadFile(resourceFullPath)
			// if err == nil {
			// 	resource.Content = string(content)
			// }
		}

		// TODO base title on domainname
		pageData.Title = "iwwwanowwwwwww"
		pageData.Resource = resource
		pageData.Resources = resources

		err = tmpl.Execute(w, pageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
