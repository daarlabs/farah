import { Controller } from '@hotwired/stimulus'
import copy from 'copy-to-clipboard'

import { Class } from './class'

const DELAY = 5000

export default class CopyController extends Controller {
	static targets = ['copy', 'check']
	declare readonly copyTarget: HTMLElement
	declare readonly checkTarget: HTMLElement

	handleCopy() {
		copy(this.element.getAttribute('data-value') || '')
		this.copyTarget.classList.add(Class.hidden)
		this.checkTarget.classList.remove(Class.hidden)
		setTimeout(() => {
			this.copyTarget.classList.remove(Class.hidden)
			this.checkTarget.classList.add(Class.hidden)
		}, DELAY)
	}
}
