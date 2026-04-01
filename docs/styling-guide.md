# Гайд по стилизации страниц в owo

> **Для ИИ-агентов:** этот документ — исчерпывающий контекст для помощи пользователям owo со стилизацией страниц. Owo — это файловый хостинг-портфолио: каждая папка становится страницей. Стили задаются через файл `.meta/index.css` внутри нужной папки. Отвечая на вопросы, всегда предлагай готовый CSS-код, который пользователь просто копирует в этот файл.

---

## Как это работает

Каждая папка в owo может иметь подпапку `.meta/` с файлом `index.css`. Этот файл подключается только к странице своей папки — это твои личные стили для конкретного раздела.

```
uploads/
└── my-portfolio/
    ├── .meta/
    │   ├── index.css   ← стили для этой страницы
    │   ├── index.md    ← текст страницы (markdown)
    │   ├── index.html  ← кастомный HTML (показывается в iframe)
    │   └── cover.jpg   ← обложка карточки
    └── project1/
        └── ...
```

---

## CSS-переменные owo (можно переопределять)

owo использует глобальные переменные. Их можно переопределить в своём `index.css`:

```css
/* Цвета */
--WHITE: #fff
--LIGHT: #e0e0e0    /* фон блоков кода, светлые элементы */
--MEDIUM: #a4a4a4   /* средний серый */
--DARK: #535353     /* основной цвет текста */
--BLACK: #000

--BLUE: #009bff     /* цвет ссылок */
--BLUE-DARK: #0040ff
--GREEN: #00ff9b
--ORANGE: #ff9b00
--RED: #f00

/* Сетка */
--GRID-GAP: 16px    /* расстояние между карточками */
```

---

## Примеры: что и как менять

### Фон страницы

```css
/* Однотонный цвет */
body {
  background-color: #1a1a2e;
}

/* Градиент */
body {
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
}

/* Изображение */
body {
  background-image: url('https://example.com/texture.jpg');
  background-size: cover;
  background-attachment: fixed;
}
```

### Шрифты — через Google Fonts

Зайди на [fonts.google.com](https://fonts.google.com), выбери шрифт, скопируй ссылку `@import`:

```css
/* Шаг 1: подключить шрифт */
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;700&display=swap');

/* Шаг 2: применить */
* {
  font-family: 'Playfair Display', serif;
}
```

Примеры популярных шрифтов:
- `'Inter', sans-serif` — современный, нейтральный
- `'Playfair Display', serif` — элегантный, для арт/дизайн портфолио
- `'Space Mono', monospace` — технический, для разработчиков
- `'Unbounded', sans-serif` — жирный, для дерзкого стиля
- `'Cormorant Garamond', serif` — изысканный, минималистичный

### Цвет текста и ссылок

```css
/* Основной текст */
body {
  color: #f0f0f0;
}

/* Заголовки отдельно */
h1, h2, h3 {
  color: #ffffff;
}

/* Ссылки */
:root {
  --BLUE: #ff6b6b; /* переопределить цвет ссылок */
}
```

### Карточки (`.card__wrapper`)

```css
/* Тёмные карточки */
.card__wrapper {
  background-color: #1e1e2e;
  border-color: #313244;
  color: #cdd6f4;
}

/* Карточки с эффектом при наведении */
.card__wrapper {
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.card__wrapper:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.3);
}

/* Скруглённые карточки */
.card__wrapper {
  border-radius: 12px;
  overflow: hidden;
}
```

### Шапка сайта (`.header`)

```css
/* Фон шапки */
.header {
  background-color: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(10px);
  padding: 16px;
  border-bottom: 1px solid #333;
}

/* Цвет названия в шапке */
.header h5 {
  color: #ffffff;
}
```

### Размер и цвет шрифта заголовков

```css
h1 { font-size: 8em; color: #fff; }
h2 { font-size: 3.2em; color: #ccc; }
h3 { color: #aaa; }
```

### Расстояние между карточками

```css
:root {
  --GRID-GAP: 24px; /* увеличить отступы */
}
```

---

## Тёмная тема — готовый пример

```css
/* .meta/index.css */

body {
  background-color: #0d0d0d;
  color: #e0e0e0;
}

:root {
  --LIGHT: #1a1a1a;
  --DARK: #e0e0e0;
  --BLUE: #64b5f6;
}

.card__wrapper {
  background-color: #1a1a1a;
  border-color: #2a2a2a;
  color: #e0e0e0;
  transition: transform 0.15s ease;
}

.card__wrapper:hover {
  transform: translateY(-3px);
}

.header {
  border-bottom: 1px solid #2a2a2a;
}
```

---

## Структура HTML-классов (для точечной стилизации)

| Класс | Что это |
|---|---|
| `body` | Весь фон страницы |
| `.wrapper` | Центральный контейнер (padding, flex-column) |
| `.header` | Шапка с названием сайта |
| `.grid` | Контейнер карточек (CSS Grid) |
| `.card__wrapper` | Одна карточка |
| `.card__img` | Обложка карточки |
| `.card__text-header` | Заголовок карточки |
| `.grid__right-content` | Блок с основным контентом страницы |

---

## Что ещё можно сделать через `.meta/`

- **`index.md`** — текст страницы в Markdown (превращается в HTML автоматически)
- **`index.html`** — произвольный HTML, отображается в iframe над контентом
- **`index.js`** — JavaScript для анимаций, интерактива
- **`cover.jpg`** / **`cover.png`** / **`cover.webp`** — обложка, которая видна на карточке в родительской папке
- **`index.link`** — ярлык: при открытии папки браузер перенаправляется на другую страницу сайта

### index.link — ярлыки на проекты

Позволяет создать «ярлык» на любой ресурс внутри uploads. Удобно для подборок на главной: карточка живёт в корне, а сам проект — глубоко в структуре по категориям и годам.

Файл содержит путь к таргету относительно uploads (без ведущего `/`):

```
Grafika/_2024/web-project
```

Пример структуры:

```
uploads/
├── featured-work/          ← папка-ярлык (видна как карточка на главной)
│   └── .meta/
│       └── index.link      ← содержит: Grafika/_2024/web-project
└── Grafika/
    └── _2024/
        └── web-project/    ← настоящий проект
            └── .meta/
                └── cover.jpg
```

- Обложка карточки ярлыка берётся автоматически из таргета, если своей нет в `.meta/`
- Можно положить свою `cover.jpg` рядом с `index.link` — она будет в приоритете

---

## Частые вопросы

**Стили применились, но шрифт не загружается**
→ Проверь, что `@import` — первая строка в файле. Импорты должны быть в самом начале.

**Хочу разные стили на разных страницах**
→ Каждая папка имеет свой `.meta/index.css`. Стили не наследуются между страницами.

**Как сделать анимацию появления**
```css
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to   { opacity: 1; transform: translateY(0); }
}

.card__wrapper {
  animation: fadeIn 0.3s ease forwards;
}
```
