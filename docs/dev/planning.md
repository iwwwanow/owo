- [ ] hide hidden

---

- [x] header with desc
- [x] footer with desc
- [x] fix qureyparam ?static to ?upload to serve files from uploads
- [ ] mapper on handler layer
  - pageProps:

  ```go
  // pass to renderer only needed values

  type PageData struct {
      title string // from resource name (file or directory)
      description string // from .meta
      html string // is needed?
      css string
      js string
      cover string
  }

  type ResourceData struct {
      type string // use consts

      // if image - prerender it to html as string
      // if md or html - prerender it to html as string
      content string
  }

  type ResourcesData struct {
      title string // filename or dirname
      description string // prerender it from html or md
      cover string // .meta/cover
  }

  type PageProps struct {
      page PageData
      resource ResourceData
      resources []ResourcesData
  }
  ```

---

- [x] static handler & static repository
- [x] repository refactor - it has todos

- [ ] перебрать весь бэклог
  - нужно разделить все фичи по категориям:
    - настройки конфига
    - стили
- [ ] нужна архитектурная схема. хотябы минимальная и простая
  - domain, app, infra (adapters, presentation/view)
- [ ] доработать работу с гитом
  - добавить в конфиг директории originUrl для git
- [ ] доработать работу с ssh
- [ ] перебери структуру своего сайта.
  - сначала правки на ЖД, затем переноси их на хост
  - не у всех директорий есть обложки
  - наверное лучше выбрать такую структуру `scope_year_name`
