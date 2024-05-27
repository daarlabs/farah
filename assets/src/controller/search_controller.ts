import { Controller } from '@hotwired/stimulus'

export default class SearchController extends Controller {
	static targets = []
	
	connect() {
		const input = (this.element as HTMLInputElement)
		const n = input.value.length
		if (!input.hasAttribute('data-autofocus')) {
			return
		}
		input.focus()
		input.setSelectionRange(n, n)
	}
}
