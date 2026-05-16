const FRAME_SRC = 'assets/frame-1.png?static'

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
	const frameWrapper = document.createElement('div')
	const leftWrapper = document.createElement('div')
	const rightWrapper = document.createElement('div')

	frameWrapper.classList.add('frame-wrapper', 'grid')
	leftWrapper.classList.add('left-wrapper')
	rightWrapper.classList.add('right-wrapper')

	document.body.prepend(frameWrapper);
	frameWrapper.append(leftWrapper);
	frameWrapper.append(rightWrapper);
}
