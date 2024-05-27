import { Controller } from '@hotwired/stimulus'

export default class FlashMessageController extends Controller {
	#REMOVE_DELAY = 5000

	connect() {
		setTimeout(() => {
			if (!this.element) {
				return
			}
			this.element.parentNode?.removeChild(this.element)
		}, this.#REMOVE_DELAY)
	}

	handleClose(e: Event) {
		const t = e.target as HTMLButtonElement
		const fm = t.closest('[data-controller="flash-message"]')
		if (!fm) {
			console.error(`closest target [data-controller="flash-message"] doesn't exist`)
			return
		}
		fm.parentNode?.removeChild(fm)
	}
}
