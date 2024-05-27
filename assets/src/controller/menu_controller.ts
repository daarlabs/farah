import { Controller } from '@hotwired/stimulus'
import { useClickOutside } from 'stimulus-use'

import { Class } from './class'

export default class MenuController extends Controller {
	static targets = ['menu', 'chevron']
	declare readonly menuTarget: HTMLElement
	declare readonly chevronTarget: HTMLElement
	declare readonly hasChevronTarget: boolean

	connect() {
		useClickOutside(this)
	}

	clickOutside() {
		this.handleClose()
	}

	handleOpen() {
		this.menuTarget.classList.remove(Class.isInvisible)
		if (!this.menuTarget.classList.contains(Class.isVisible)) {
			this.menuTarget.classList.add(Class.isVisible)
			if (this.hasChevronTarget) {
				this.chevronTarget.style.transform = `rotate(-180deg)`
			}
		}
	}

	handleClose() {
		this.menuTarget.classList.remove(Class.isVisible)
		if (!this.menuTarget.classList.contains(Class.isInvisible)) {
			this.menuTarget.classList.add(Class.isInvisible)
			if (this.hasChevronTarget) {
				this.chevronTarget.style.transform = `rotate(0deg)`
			}
		}
	}
}
