const FRAME_ASSETS_HREF = '/.meta/assets'
const FRAME_NAME_PATTERN = /^frame-.+_\d+x\d+\.(png|jpg|jpeg|webp)$/i
const IMAGE_EXTS = /\.(jpg|jpeg|png|gif|webp|bmp)(\?.*)?$/i
const CYCLE_INTERVAL_MS = 2000

// DEBUG: раскомментируй чтобы тестировать рамку без hover
const DEBUG_HREF = '/_frames-script/frame-image-1'

const HOVER_DELAY_MS = 250

let _cycleIntervalId = null
let _hoverTimeoutId = null
let _touchStartX = 0
let _touchStartY = 0
let _frameHrefs = null

const loadImage = (src) => new Promise((resolve, reject) => {
	const img = new Image()
	img.onload = () => resolve(img)
	img.onerror = reject
	img.src = new URL(src, window.location.origin).href
})

const parseInnerDimensions = (src) => {
	const filename = src.split('/').pop().split('?')[0]
	const match = filename.match(/_(\d+)x(\d+)/)
	if (!match) return null
	return { w: parseInt(match[1]), h: parseInt(match[2]) }
}

const loadFrameHrefs = async () => {
	if (_frameHrefs !== null) return _frameHrefs
	const res = await fetch(FRAME_ASSETS_HREF)
	const text = await res.text()
	const doc = new DOMParser().parseFromString(text, 'text/html')
	_frameHrefs = Array.from(doc.querySelectorAll('a.card__wrapper'))
		.map(a => a.getAttribute('href'))
		.filter(href => FRAME_NAME_PATTERN.test(href.split('/').pop()))
	return _frameHrefs
}

const getRandomFrameSrc = async () => {
	const hrefs = await loadFrameHrefs()
	if (hrefs.length === 0) return null
	const href = hrefs[Math.floor(Math.random() * hrefs.length)]
	return href + '?static'
}

const getDirectoryImages = async (href) => {
	const res = await fetch(href)
	const text = await res.text()
	const doc = new DOMParser().parseFromString(text, 'text/html')
	return Array.from(doc.querySelectorAll('a.card__wrapper'))
		.map(a => a.getAttribute('href'))
		.filter(href => IMAGE_EXTS.test(href))
}

export const makeFramesHover = (querySelector) => {
	// DEBUG: в режиме отладки не вешаем листенеры — рама остаётся активной
	if (typeof DEBUG_HREF !== 'undefined') {
		drawFramesWrapperLayer(DEBUG_HREF)
		return
	}

	const gridNode = document.querySelector(querySelector)
	gridNode.querySelectorAll('a').forEach(el => {
		el.addEventListener('mouseenter', mouseEnterHandler)
		el.addEventListener('mouseleave', mouseLeaveHandler)
		el.addEventListener('touchstart', touchStartHandler, { passive: true })
		el.addEventListener('touchend', touchEndHandler)
	})

	window.addEventListener('scroll', clearExistingFrames, { passive: true })
}

export const mouseEnterHandler = (event) => {
	const href = event.srcElement.getAttribute('href')
	clearTimeout(_hoverTimeoutId)
	_hoverTimeoutId = setTimeout(() => drawFramesWrapperLayer(href), HOVER_DELAY_MS)
}

export const touchStartHandler = (event) => {
	_touchStartX = event.touches[0].clientX
	_touchStartY = event.touches[0].clientY
}

export const touchEndHandler = (event) => {
	const dx = Math.abs(event.changedTouches[0].clientX - _touchStartX)
	const dy = Math.abs(event.changedTouches[0].clientY - _touchStartY)
	if (dx > 10 || dy > 10) return
	event.preventDefault()
	const href = event.currentTarget.getAttribute('href')
	clearExistingFrames()
	drawFramesWrapperLayer(href, true)
}

export const mouseLeaveHandler = () => {
	clearTimeout(_hoverTimeoutId)
	_hoverTimeoutId = null
	clearExistingFrames()
}

const clearExistingFrames = () => {
	if (_hoverTimeoutId !== null) {
		clearTimeout(_hoverTimeoutId)
		_hoverTimeoutId = null
	}
	if (_cycleIntervalId !== null) {
		clearInterval(_cycleIntervalId)
		_cycleIntervalId = null
	}
	document.querySelectorAll('.frame-wrapper').forEach(el => {
		el._resizeObserver?.disconnect()
		el.classList.remove('visible')
		el.addEventListener('transitionend', (e) => { if (e.target === el) el.remove() })
	})
}

const drawFramesWrapperLayer = async (elementHref, isTouchMode = false) => {
	clearExistingFrames()

	const frameSrc = await getRandomFrameSrc()
	if (!frameSrc) return

	const frameImg = await loadImage(frameSrc)
	const { naturalWidth, naturalHeight } = frameImg

	const inner = parseInnerDimensions(frameSrc)
	if (!inner) return

	const padH = (naturalWidth - inner.w) / 2
	const padV = (naturalHeight - inner.h) / 2

	const frameWrapper = document.createElement('div')
	const leftWrapper = document.createElement('div')
	const rightWrapper = document.createElement('div')
	const frameImage = document.createElement('img')
	const sourceImage = document.createElement('img')
	const sourceImageAlt = document.createElement('img')
	const h1 = document.createElement('h1')
	const h2 = document.createElement('h2')
	const p = document.createElement('p')

	frameWrapper.classList.add('frame-wrapper', 'grid')
	leftWrapper.classList.add('left-wrapper')
	rightWrapper.classList.add('right-wrapper')
	frameImage.classList.add('frame-image')
	sourceImage.classList.add('source-image')
	sourceImageAlt.classList.add('source-image')
	sourceImageAlt.style.opacity = '0'

	frameImage.src = new URL(frameSrc, window.location.origin).href

	document.body.prepend(frameWrapper)
	requestAnimationFrame(() => requestAnimationFrame(() => frameWrapper.classList.add('visible')))

	if (isTouchMode) {
		frameWrapper.classList.add('touch-mode')
		const topBar = document.createElement('div')
		const goLink = document.createElement('a')
		const closeLink = document.createElement('a')
		const goLabel = document.createElement('h5')
		const closeLabel = document.createElement('h5')
		topBar.classList.add('frame-top-bar')
		goLabel.textContent = 'ПЕРЕЙТИ'
		closeLabel.textContent = 'ЗАКРЫТЬ'
		goLink.href = elementHref
		goLink.append(goLabel)
		closeLink.append(closeLabel)
		closeLink.addEventListener('click', (e) => { e.preventDefault(); clearExistingFrames() })
		topBar.append(goLink, closeLink)
		frameWrapper.append(topBar, leftWrapper, rightWrapper)
	} else {
		frameWrapper.append(leftWrapper, rightWrapper)
	}

	leftWrapper.append(h2, h1, p)
	rightWrapper.append(frameImage, sourceImageAlt, sourceImage)

	// object-fit: contain; object-position: left top — рамка вписывается в right-wrapper
	// с сохранением пропорций, начиная с (0, 0). Считаем реальный rendered-размер:
	const applySourceImagePosition = () => {
		const { width: rwWidth, height: rwHeight } = rightWrapper.getBoundingClientRect()
		const scale = (rwWidth / rwHeight > naturalWidth / naturalHeight)
			? rwHeight / naturalHeight  // ограничено по высоте
			: rwWidth / naturalWidth    // ограничено по ширине
		const bleed = 8
		const pos = {
			left: `${padH * scale - bleed}px`,
			top: `${padV * scale - bleed}px`,
			width: `${inner.w * scale + bleed * 2}px`,
			height: `${inner.h * scale + bleed * 2}px`,
		}
		Object.assign(sourceImage.style, pos)
		Object.assign(sourceImageAlt.style, pos)
	}

	// Двойной rAF гарантирует, что браузер успел сделать layout
	// и getBoundingClientRect() вернёт правильные размеры
	await new Promise(resolve => requestAnimationFrame(() => requestAnimationFrame(resolve)))
	applySourceImagePosition()

	// Размер rightWrapper — одинаков для всех рам при одном viewport,
	// не зависит от конкретной рамы и гарантированно ненулевой после layout
	const { width: displayW, height: displayH } = rightWrapper.getBoundingClientRect()
	const imageSrc = (href, dpr = 1) =>
		`${href}?static&width=${Math.round(displayW * dpr)}&height=${Math.round(displayH * dpr)}`

	const resizeObserver = new ResizeObserver(applySourceImagePosition)
	resizeObserver.observe(rightWrapper)
	frameWrapper._resizeObserver = resizeObserver

	if (elementHref) {
		try {
			const jsonRes = await fetch(elementHref + '/.meta/data.json?static')
			if (jsonRes.ok) {
				const json = await jsonRes.json()
				h2.innerText = json.title ?? ''
				h1.innerText = `${json.price}₽` ?? ''
				p.innerText = json.description ?? ''
			}
		} catch {}
	}

	if (elementHref) {
		try {
			const images = await getDirectoryImages(elementHref)
			if (images.length > 0) {
				const dpr = Math.min(Math.ceil(window.devicePixelRatio || 1), 2)
				await Promise.all(images.map(href => loadImage(imageSrc(href, dpr))))

				sourceImage.src = imageSrc(images[0])
				sourceImage.srcset = `${imageSrc(images[0], 2)} 2x`
				if (images.length > 1) {
					let idx = 0
					let front = sourceImage
					let back = sourceImageAlt
					_cycleIntervalId = setInterval(() => {
						idx = (idx + 1) % images.length
						back.src = imageSrc(images[idx])
						back.srcset = `${imageSrc(images[idx], 2)} 2x`
						back.style.opacity = '1'
						front.style.opacity = '0'
						;[front, back] = [back, front]
					}, CYCLE_INTERVAL_MS)
				}
			}
		} catch {}
	}
}
