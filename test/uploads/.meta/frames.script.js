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
}

export const mouseEnterHandler = (event) => {
	console.log(event)
}

export const mouseLeaveHandler = (event) => {
	console.log(event)
}
