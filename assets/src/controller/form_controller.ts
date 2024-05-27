import { Controller } from '@hotwired/stimulus'

const FORM_SUBMIT_CLASS = 'form-submit'

export default class FormController extends Controller {
	connect() {
		window.addEventListener('keyup', this.handleEscapePress)
	}

	disconnect() {
		window.removeEventListener('keyup', this.handleEscapePress)
	}

	handleEscapePress = (e: KeyboardEvent) => {
		if (e.key !== 'Escape') {
			return
		}
		this.handleStop()
	}

	handleSubmit() {
		this.element.classList.add(FORM_SUBMIT_CLASS)
	}

	handleStop() {
		this.element.classList.remove(FORM_SUBMIT_CLASS)
	}
}
