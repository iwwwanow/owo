```
// main.go
func main() {
    // 1. Создаем инфраструктурные зависимости
    repo := infra.NewFileSystemRepository("/var/www/owwo/shared")
    renderer := infra.NewHTMLRenderer(templates)

    // 2. Создаем use-case (логика приложения)
    useCase := application.NewRenderPageUseCase(repo, renderer)

    // 3. Создаем контроллер с зависимостью от use-case
    handler := infra.NewHTTPHandler(useCase)

    // 4. Запускаем сервер
    http.Handle("/", handler)
    http.ListenAndServe(":8080", nil)
}
```
