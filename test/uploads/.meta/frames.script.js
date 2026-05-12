const FRAME_SRC = '.meta/assets/frame-1.png?static'

// Отступы рамы — доля от итогового размера (0..1).
// Подбирать визуально под конкретный PNG.
const FRAME_INSET = {
  top: 0.13,
  right: 0.13,
  bottom: 0.13,
  left: 0.13,
}

// Задержка перед показом при ховере (мс)
const HOVER_DELAY = 175

// Фаза 1: константы
const CARD_DATA = {
  description: 'Описание работы',
  priceOld: '30 000 ₽',
  priceNew: '18 000 ₽',
}

// Пропорции рамы (заполняются после preload)
const frameRatio = { w: 1, h: 1 }

let overlayEl = null
let frameWrapEl = null
let frameImgEl = null
let innerImgEl = null
let textTitleEl = null
let hoverTimer = null

function preloadFrame() {
  return new Promise((resolve) => {
    const img = new Image()
    img.onload = () => {
      frameRatio.w = img.naturalWidth
      frameRatio.h = img.naturalHeight
      resolve()
    }
    img.onerror = resolve
    img.src = FRAME_SRC
  })
}

function createOverlay() {
  overlayEl = document.createElement('div')
  overlayEl.id = 'frame-overlay'

  // ── Текстовый блок ──────────────────────────────────────
  const text = document.createElement('div')
  text.className = 'frame-overlay__text'

  textTitleEl = document.createElement('h1')

  const desc = document.createElement('p')
  desc.textContent = CARD_DATA.description

  const price = document.createElement('div')
  price.className = 'frame-overlay__price'

  const priceOld = document.createElement('h2')
  priceOld.className = 'frame-overlay__price-old'
  priceOld.textContent = CARD_DATA.priceOld

  const priceNew = document.createElement('h2')
  priceNew.className = 'frame-overlay__price-new'
  priceNew.textContent = CARD_DATA.priceNew

  price.append(priceOld, priceNew)
  text.append(textTitleEl, desc, price)

  // ── Frame wrap ──────────────────────────────────────────
  frameWrapEl = document.createElement('div')
  frameWrapEl.className = 'frame-overlay__frame-wrap'

  innerImgEl = document.createElement('img')
  innerImgEl.className = 'frame-overlay__inner-img'
  innerImgEl.alt = ''

  frameImgEl = document.createElement('img')
  frameImgEl.className = 'frame-overlay__frame-img'
  frameImgEl.src = FRAME_SRC
  frameImgEl.alt = ''

  // inner под рамой (z-index: 1 < 2)
  frameWrapEl.append(innerImgEl, frameImgEl)

  overlayEl.append(text, frameWrapEl)
  document.body.appendChild(overlayEl)
}

function getCardImgSrc(card) {
  const img = card.querySelector('picture.card__cover-container img')
  if (img) return img.src
  const video = card.querySelector('video.card__cover-container')
  if (video) return video.src
  return null
}

function getGridCols(grid) {
  const cols = getComputedStyle(grid).gridTemplateColumns
  return cols.split(' ').filter(Boolean).length
}

function isNarrowMode(grid) {
  if (grid.classList.contains('grid--hidden')) return true
  return getGridCols(grid) <= 2
}

function computeFrameWidth(card, grid) {
  if (isNarrowMode(grid)) {
    // вертикальный режим: ширина = viewport минус паддинги wrapper
    const wrapperPad = parseInt(getComputedStyle(document.querySelector('.wrapper')).paddingLeft, 10) || 0
    return document.documentElement.clientWidth - wrapperPad * 2
  }
  return card.getBoundingClientRect().width
}

function positionInnerImg(frameW, frameH) {
  const top    = frameH * FRAME_INSET.top
  const right  = frameW * FRAME_INSET.right
  const bottom = frameH * FRAME_INSET.bottom
  const left   = frameW * FRAME_INSET.left

  innerImgEl.style.top    = top + 'px'
  innerImgEl.style.left   = left + 'px'
  innerImgEl.style.width  = frameW - left - right + 'px'
  innerImgEl.style.height = frameH - top - bottom + 'px'
}

function showOverlay(card, grid) {
  const src = getCardImgSrc(card)
  if (!src) return

  const title = card.querySelector('.card__text-header')?.textContent.trim() ?? CARD_DATA.title
  textTitleEl.textContent = title
  innerImgEl.src = src

  const frameW = computeFrameWidth(card, grid)
  const frameH = frameW * (frameRatio.h / frameRatio.w)

  frameWrapEl.style.width  = frameW + 'px'
  frameWrapEl.style.height = frameH + 'px'

  positionInnerImg(frameW, frameH)

  overlayEl.classList.add('frame-overlay--visible')
  document.body.classList.add('frame-blur')
}

function hideOverlay() {
  clearTimeout(hoverTimer)
  overlayEl.classList.remove('frame-overlay--visible')
  document.body.classList.remove('frame-blur')
}

export const makeFramesHover = async (elementId) => {
  const grid = document.querySelector(elementId)
  if (!grid) return

  await preloadFrame()
  createOverlay()

  const cards = grid.querySelectorAll('.card__wrapper')
  const isTouch = 'ontouchstart' in window

  if (isTouch) {
    cards.forEach((card) => {
      card.addEventListener('touchstart', (e) => {
        e.stopPropagation()
        if (overlayEl.classList.contains('frame-overlay--visible')) {
          hideOverlay()
        } else {
          showOverlay(card, grid)
        }
      })
    })

    // Тап вне оверлея и вне карточки — закрыть
    document.addEventListener('touchstart', (e) => {
      if (!e.target.closest('#frame-overlay') && !e.target.closest('.card__wrapper')) {
        hideOverlay()
      }
    })
  } else {
    cards.forEach((card) => {
      card.addEventListener('mouseenter', () => {
        console.log('mouse enter')
        clearTimeout(hoverTimer)
        hoverTimer = setTimeout(() => showOverlay(card, grid), HOVER_DELAY)
      })
      card.addEventListener('mouseleave', () => {
        clearTimeout(hoverTimer)
        hideOverlay()
      })
    })
  }
}
