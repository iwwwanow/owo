console.log('frames-script-loaded')

export const makeFramesHover = (elementId) => {
	const elements = document.querySelector(elementId)

	console.log(elements)
	console.log(elements.children)

	const children = elements.children.filter((item) => {
		return item.localName !== 'fieldset'
	})

	console.log(children)
}
