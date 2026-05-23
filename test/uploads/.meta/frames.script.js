// TODO: поправить мобильные разрешения
// TODO: добавить задержку и плавность
// TODO: подумать, как это будет работать на мобильных

const FRAME_SRC = '.meta/assets/frame-1_245x347.png?static'
const IMAGE_EXTS = /\.(jpg|jpeg|png|gif|webp|bmp)(\?.*)?$/i
const CYCLE_INTERVAL_MS = 2000

// DEBUG: раскомментируй чтобы тестировать рамку без hover
const DEBUG_HREF = '/_frames-script/frame-image-1'

let _cycleIntervalId = null

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

const getDirectoryImages = async (href) => {
	const res = await fetch(href)
	const text = await res.text()
	const doc = new DOMParser().parseFromString(text, 'text/html')
	return Array.from(doc.querySelectorAll('a.card__wrapper'))
		.map(a => a.getAttribute('href'))
		.filter(href => IMAGE_EXTS.test(href))
		.map(href => href + '?static')
}

export const makeFramesHover = (querySelector) => {
	const gridNode = document.querySelector(querySelector)
	gridNode.querySelectorAll('a').forEach(el => {
		el.addEventListener('mouseenter', mouseEnterHandler)
		el.addEventListener('mouseleave', mouseLeaveHandler)
	})

	// DEBUG: использует DEBUG_HREF если раскомментирована выше
	drawFramesWrapperLayer(typeof DEBUG_HREF !== 'undefined' ? DEBUG_HREF : undefined)
}

export const mouseEnterHandler = (event) => {
	drawFramesWrapperLayer(event.srcElement.getAttribute('href'))
}

export const mouseLeaveHandler = () => {
	clearExistingFrames()
}

const clearExistingFrames = () => {
	if (_cycleIntervalId !== null) {
		clearInterval(_cycleIntervalId)
		_cycleIntervalId = null
	}
	document.querySelectorAll('.frame-wrapper').forEach(el => {
		el._resizeObserver?.disconnect()
		el.remove()
	})
}

const drawFramesWrapperLayer = async (elementHref) => {
	clearExistingFrames()

	const frameImg = await loadImage(FRAME_SRC)
	const { naturalWidth, naturalHeight } = frameImg

	const inner = parseInnerDimensions(FRAME_SRC)
	if (!inner) return

	const padH = (naturalWidth - inner.w) / 2
	const padV = (naturalHeight - inner.h) / 2

	const frameWrapper = document.createElement('div')
	const leftWrapper = document.createElement('div')
	const rightWrapper = document.createElement('div')
	const frameImage = document.createElement('img')
	const sourceImage = document.createElement('img')
	const h1 = document.createElement('h1')
	const h2 = document.createElement('h2')
	const h3 = document.createElement('h3')

	frameWrapper.classList.add('frame-wrapper', 'grid')
	leftWrapper.classList.add('left-wrapper')
	rightWrapper.classList.add('right-wrapper')
	frameImage.classList.add('frame-image')
	sourceImage.classList.add('source-image')

	frameImage.src = new URL(FRAME_SRC, window.location.origin).href

	document.body.prepend(frameWrapper)
	frameWrapper.append(leftWrapper, rightWrapper)
	leftWrapper.append(h1, h2, h3)
	rightWrapper.append(frameImage, sourceImage)

	// object-fit: contain; object-position: left top — рамка вписывается в right-wrapper
	// с сохранением пропорций, начиная с (0, 0). Считаем реальный rendered-размер:
	const applySourceImagePosition = () => {
		const { width: rwWidth, height: rwHeight } = rightWrapper.getBoundingClientRect()
		const scale = (rwWidth / rwHeight > naturalWidth / naturalHeight)
			? rwHeight / naturalHeight  // ограничено по высоте
			: rwWidth / naturalWidth    // ограничено по ширине
		Object.assign(sourceImage.style, {
			left: `${padH * scale}px`,
			top: `${padV * scale}px`,
			width: `${inner.w * scale}px`,
			height: `${inner.h * scale}px`,
		})
	}

	// Двойной rAF гарантирует, что браузер успел сделать layout
	// и getBoundingClientRect() вернёт правильные размеры
	await new Promise(resolve => requestAnimationFrame(() => requestAnimationFrame(resolve)))
	applySourceImagePosition()

	const resizeObserver = new ResizeObserver(applySourceImagePosition)
	resizeObserver.observe(rightWrapper)
	frameWrapper._resizeObserver = resizeObserver

	if (elementHref) {
		try {
			const jsonRes = await fetch(elementHref + '/.meta/data.json?static')
			if (jsonRes.ok) {
				const json = await jsonRes.json()
				h1.innerText = json.title ?? ''
				h2.innerText = json.description ?? ''
				h3.innerText = json.price ?? ''
			}
		} catch {}
	}

	if (elementHref) {
		try {
			const images = await getDirectoryImages(elementHref)
			if (images.length > 0) {
				sourceImage.src = images[0]
				if (images.length > 1) {
					let idx = 0
					_cycleIntervalId = setInterval(() => {
						idx = (idx + 1) % images.length
						sourceImage.src = images[idx]
					}, CYCLE_INTERVAL_MS)
				}
			}
		} catch {}
	}
}
