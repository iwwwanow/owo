const FRAME_SRC = '.meta/assets/frame-1.png?static'

export const makeFramesHover = (querySelector) => {
	const gridNode = document.querySelector(querySelector)
	console.log(gridNode)
	const elementNodes = gridNode.querySelectorAll('a')
	console.log(elementNodes)

	elementNodes.forEach(element => {
		element.addEventListener('mouseenter', mouseEnterHandler)
		element.addEventListener('mouseleave', mouseLeaveHandler)
	});

	// DEBUG:
	drawFramesWrapperLayer()
	// DEBUG:
}

export const mouseEnterHandler = (event) => {
	// console.log(event)
	drawFramesWrapperLayer()
}

export const mouseLeaveHandler = (event) => {
	// console.log(event)

	const frameWrappers = document.querySelectorAll('.frame-wrapper')
	console.log(frameWrappers)
	frameWrappers.forEach(wrapper => wrapper.remove())
	// TODO: disable scroll
}

const drawFramesWrapperLayer = () => {
	// TODO: get it from parent and cycle it
	const sourceImageLink = '/_frames-script/6104c2143578475.627cf80887bb9.png?static'

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

	header1.innerText = 'Шаман'
	header2.innerText = 'шаман 1.25x1.25 - акрил, осп-плита, лак'
	header3.innerText = 'price'
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
