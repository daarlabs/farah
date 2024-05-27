import './main.css'

import 'htmx.org'
import LazyLoad from 'vanilla-lazyload'
import * as Stimulus from '@hotwired/stimulus'

import FormController from './controller/form_controller'
import FlashMessageController from "./controller/flash-message_controller";
import MenuController from './controller/menu_controller'
import SearchController from "./controller/search_controller"

const app = Stimulus.Application.start()

document.addEventListener('DOMContentLoaded', () => {
	app.register('flash-message', FlashMessageController)
	app.register('form', FormController)
	app.register('menu', MenuController)
	app.register('search', SearchController)
	new LazyLoad({})
})

document.body.addEventListener('htmx:beforeRequest', () => {})

document.body.addEventListener('htmx:afterRequest', () => {})
