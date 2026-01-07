- [ ] hide hidden
- [ ] выпустить v1
- [ ] сформировать roadmap до v2
- [ ] настроить версионирование при мердже в мастер
- [ ] curl backup to archive

- [ ] разделять директжории и файлы
  - добавить иконку обязательно
  - либо проставить между ними отступ в 2 строки
  - либо проставить HR

- [ ] git pull должен кешироваться. и не делаться раньше, чем раз в 5 минут
- [ ] если директория - гит репозиторй, нужно както стилистически её выделять
  - и на кровне карточки
  - и на уровне самой директории

- [ ] добавить изображениям urlParam - для получения спецефичного размера
- [ ] добавить кеширование изображений с переодической чисткой
  - раз в 5 дней?

- [ ] если создание ярлыков (ссылок) не будет работать - можно ввести спецефичный нейминг для копирования ресурса. чтобы не дублировать место на диске
- [ ] by owwo - исправь. добавить просто owo или чтот около того.
- [ ] нужна возможность расширять футер
- [ ] `.` директории и файлы будут скрыты
- [ ] `!` директории и файлы в самом верху. может быть их даже расширить
- [ ] `#` директории и файлы в самом низу. укороченный вариант карточки

- [ ] api: добавить ручку на получение всех айтемов со страницы
- [ ] config: добавить параметр hideItems/hide_items: boolean
- [ ] при помощи js и webcomponents отрисовывать айтемы по категориям
  - ref: https://timurnovikov.com/en
- [ ] нужна витрина. с перелинковкой и золотыми рамами

- [ ] запаблить и резюме и cv на owwo.
- [ ] нужно перенести всю эту движуху на iwwwanow.ru домен
- [ ] если в мета не заданна обложка выбирать первую картинку
- [ ] обязательно добавить перелинковку изображений и директорий

- [ ] разделить получение страницы и получение файла, как статики ЛИБО при помощи `queryParam` ЛИБО при помощи браузерных средств (могу ли я понять инициатора запроса)?
- [ ] при сборке образа, нужно копировать в него ssh-ключи, чтобы при повторном подключении не высвечивалось предупреждении об изменении ключей. тоесть нужно зафиксировать постоянный ключ в образе
- [ ] render error message; on git failed, например
- [ ] render 404 and other error pages
- [ ] config on page
- [ ] save host keys on update images
- [ ] curl, or scripts, or exe, or bit to fast install on server and localhost
- [ ] git repo
  - .git direcotry in folder
  - auto pull on page request
  - render index.html
  - how to get .js & .css files?
  - meeting-with-the-goddess-of-flowers-1_interactive refactor
- [ ] images list on infra (config.yaml maybe)
- [ ] scripts.sh on infra (pull, clean & deploy.sh)
- [ ] identical layout on mobile. just 2 columns
  - without hovers
- [ ] config.(json/yaml) in root
- [ ] render hidden resources
- [ ] show-hidden-resources - config option
- [ ] render directory
  - iosevka nerd font
- [ ] ssh connection to docker
- [ ] just black/inversion font on cards
- [ ] favicon
- [ ] metainfo
- [ ] render hidden folders
- [ ] цифровая путь к храму
- [ ] вертикальная путь к храму
- [ ] after rebuild image cannot connet. need to rm ssh auth key and reconnect
- [ ] owwo workflows
- [ ] cache
- [ ] videos
- [ ] interfactive meeting
- [ ] wrap on resource card

---

- [x] подключить стили в директорию public. передавать их независимо от шаблонов
- [x] - рендер сетки в директории
  - [x] - выбор directory || text || image || other для шаблона resouce-card
  - [x] - для type=text сделать подготовку preview
    - для image можно заталкивать туда ссылку
- [x] - рендер страницы с ресурсом
  - [x] - get resource content
  - [x] - full-size-image
- [x] - docker, docker-compose for deploy, or docker only for deploy. it will be better
  - [x] - ssh-connection
  - [x] - linux-enveronment
  - [x] - remote-server-usage
- [x] - dns setup for domain
- [x] - docker compose up error on workflow deploy run
- [x] - ssl sertificate
  - [x] - сайт перестал отвечать на http
  - [x] - проверить наличие сертификата на сервере. создается ли он вообще
  - [x] - перепроверить сайт с наличием сертификата
  - [x] - проверить http-config в nginx. если неверный, вернуть как было
- [x] - latest owwo image not deployed

- покачто делаю без зарезервированных имен, типа кавера, resource-info.json и пр
- также делаю без даты типа дата создания, дата регистрации и пр
- просто рендер изобрежений и текста
- после рендера страницы настраивать деплой

ltf:

- [x] - css, js, html, md
- [ ] - sftp users on docker
- [ ] directory resource type card design. должен отличаться
- [ ] - md text to html
- [ ] - баннер на главной
- [ ] resource page
  - [ ] - css link to page
  - [ ] - js link to page
  - [ ] - custom json data
- [ ] resource full page render:
  - [ ] - cover
  - [ ] - other resource data
  - [ ] - root layout. как это реализовать в шаблоне? просто встроить стили в страницу?
  - [ ] - hr для футера. 2 режима отображения - с текстом и без
  - [ ] - content. 2 режима отображения - grid и full-size (image & text)
  - [ ] - изображения
  - [ ] - текст
  - [ ] - прочие файлы. при нажатии на них - скачивание, прямо из сетки. пре переходе на страницу - тоже скачивание

---

- [x] базовая верстка, без контента
  - просто перенести верстку из предыдущих версий
- [ ] рендер контента
  - енв, который принимает путь до директории с файлами

---

- https://chat.deepseek.com/a/chat/s/d2e60272-d8be-4724-9925-723ecae87d50

- [ ] - get mime type
- [ ] - get cover path
- [ ] - pass it to renders

---

- [ ] - нужно переделать структуру. сначала получение верстки в виде строк, затем рендер

- [x] - чтение директории
- [x] - вынести отрисовку child-resources в ui
- [x] - рендирить актуальные файлы со ссылкой
  - [x] - перебрал функцию getfilenames. нужно передавать в неё готовый struct
- [ ] - рендерить информацию о текущей директории
  - [ ] - вынести uploads-dir-path в env-переменную
  - [ ] - лепить строку из uploads-dir-path и uri
  - [ ] - передавать её в get-resource-data
  - [ ] - возвращаемым значение должнен быть массив из resources. в которых будет mimetype, path, name, description, cover
    - все маршруты будут относительными путями до соответствующих файлов
- [ ] - переходить по вложенным директориям и рендерить их информацию
- [ ] - можно покачто просто рендерить названия файлов. без вожений, типо список
- [ ] - дальше углубляться искать covers, (descriptions?) and etc
- [ ] - чтение js-файла, который отвечает за reload_js
  - https://chat.deepseek.com/a/chat/s/acf46b4a-7284-430b-8ebd-a3ade4de4577

---

- [ ] - нужно организовать отправку собранных стилей в проде
- [ ] - нужно добавить nginx в докер, обрубить его на сервере
- [ ] - структура репозитория какбудто просится разделиться на site и system
  - скрипты поместить в site
  - docker у каждого приложения свой
  - owwo-deploy можно вынести в собственный репозиторий. тамже прописать и настройки nginx

- [x] - tsc
- [x] - move static folder out of src
- [x] - static folder to env
- [x] - check required env on startup
- [x] - test docker on local
  - [x] - build docker
  - [x] - env with uploads path to docker
  - [x] - external uploads dir
  - [x] - system docker build
  - [x] - system docker connect via sftp
  - [x] - system docker build with my keys (will be in secrets later)
  - [x] - system docker connect via sftp
  - [x] - secrets to docks
  - [x] - secrets to github actions
  - [x] - compose test
- [x] - refactor system docker
  - [x] - connect only by key
  - [x] - connect via root (full system available)
  - [x] - connect via effectivnayarabota1 (era1 directory only)
- [ ] - workflows:
    x если в докер будут зашиты приватные ключи (пароль для аутентификации), то как я могу публиковать докер в публичное пространство?
    - в докер зашивается только публичный ключ, поэтому пофиг
  - [x] - release (docker) on merge master
  - [x] - pull docker on hosting on merge master (deploy)
    - 1. - создать ключ для прямого подключения к хосту (с локалки)
    - 2. - создать ключ для подключения через гитхаб экшенс
    - 3. - добавить authorized keys на хост (оба)
    - 4. - удалить изначальный ключ на хосте
    - 5. - прописать компоуз на хосте
    - 6. - прописать .env на хосте
    - 7. - смотри пример в naked science, вроде все достаточно просто
- [ ] - FEAT stage branch; stage subdomain

- [ ] - FEAT cover for md, css and other files; nerd fonts or icons?
- [ ] - FEAT html to iframe
  - [ ] - md or html - optional
  - [ ] - if html - render it in iframe
  - [ ] - может нахуй гельветику спиратить?
- [ ] - FEAT отрисовка страниц-файлов. (extended-page)

- [x] - dont show reserved page
- [x] - dont show resources if it empty
- [x] - dont show reserverd names
- [x] - /effectivnayarabota1/Screenshot%20from%202025-02-18%2021-12-12.png - didnt work
- [x] - 404 if resource not found
- [x] - favicon
- [x] - test page with md and inner images
- [x] - about page
- [x] - rename Text component to Content component
- [x] - rm unused from header

- [x] - show file content on html md css and js and other
- [x] - connect html, css, js into page

- [x] - getCoverByPath - file cover
- [x] - not render image on cart if it not exits or false

- [x] - repository refactor
- [x] - in the start of repository define type of resource and next steps depends on it

- [x] - что делать с пробелами?
- [x] - каждую карточку сделать идентичного размера. изображения ровняются по верхнему краю. рамка
