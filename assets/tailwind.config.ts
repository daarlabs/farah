export default {
	content: ['../feature/**/*.go', '../ui/**/*.go', '../web/**/*.go'],
	darkMode: 'class',
	theme: {
		extend: {
			colors: {
				primary: {
					950: '#010024',
					900: '#020037',
					800: '#04006D',
					700: '#050092',
					600: '#0600B6',
					500: '#0700DB',
					400: '#0B02FF',
					300: '#2C24FF',
					200: '#615BFF',
					100: '#A7A4FF',
					50: '#DCDBFF',
				},
				secondary: {
					400: '#15C291',
				},
			},
			boxShadow: {
				focus: '0 0 0 0.25rem rgba(11, 2, 255, 0.25)',
			},
			fontFamily: {
				sora: ['Sora', 'sans-serif'],
			},
		},
	},
	plugins: [],
}
