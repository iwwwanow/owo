const FRAME_SRC = '.meta/assets/frame-1_245x347.png?static'

export const makeFramesHover = (querySelector) => {
	const gridNode = document.querySelector(querySelector)
	const elementNodes = gridNode.querySelectorAll('a')

	elementNodes.forEach(element => {
		element.addEventListener('mouseenter', mouseEnterHandler)
		element.addEventListener('mouseleave', mouseLeaveHandler)
	});

	// DEBUG:
	drawFramesWrapperLayer()
	// DEBUG:
}

export const mouseEnterHandler = (event) => {
	const elementHref = event.srcElement.getAttribute('href')
	drawFramesWrapperLayer(elementHref)
}

export const mouseLeaveHandler = (event) => {
	// console.log(event)

	const frameWrappers = document.querySelectorAll('.frame-wrapper')
	frameWrappers.forEach(wrapper => wrapper.remove())
	// TODO: disable scroll
}

const drawFramesWrapperLayer = async (elementHref) => {
	// TODO: get it from parent and cycle it
	// const sourceImageLink = elementHref
	const sourceImageLink = '/_frames-script/6104c2143578475.627cf80887bb9.png?static'
	// TODO: how to be in other extentions? jpg, gif, webp?
	const sourceJsonLink = sourceImageLink.replace('.png', '.json');
	console.log(sourceJsonLink)

	const img = new Image();
	img.src = new URL(FRAME_SRC, window.location.origin).href;
	img.onload = () => {console.log(img.naturalWidth, img)}

	const innerDimentions = 0
	const outterDimentions = 0

	const frameWrapper = document.createElement('div')
	const leftWrapper = document.createElement('div')
	const rightWrapper = document.createElement('div')
	const header1 = document.createElement('h1')
	const header2 = document.createElement('h2')
	const header3 = document.createElement('h3')
	const frameImage = document.createElement('img')
	const sourceImage = document.createElement('img')

	frameWrapper.classList.add('frame-wrapper', 'grid')
	leftWrapper.classList.add('left-wrapper')
	rightWrapper.classList.add('right-wrapper')
	frameImage.classList.add('frame-image')
	sourceImage.classList.add('source-image')

	const response = await fetch(sourceJsonLink)
	const json = await response.json()
	console.log(json)

	header1.innerText = json.title
	header2.innerText = json.description
	header3.innerText = json.price
	frameImage.src = FRAME_SRC;
	sourceImage.src = sourceImageLink

	document.body.prepend(frameWrapper);
	frameWrapper.append(leftWrapper);
	frameWrapper.append(rightWrapper);
	leftWrapper.appendChild(header1)
	leftWrapper.appendChild(header2)
	leftWrapper.appendChild(header3)
	rightWrapper.appendChild(frameImage)
	rightWrapper.appendChild(sourceImage)
}
