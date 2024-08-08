const {addIconSelectors} = require('@iconify/tailwind')
module.exports = {
	plugins: [
		require('daisyui'),
		addIconSelectors(['logos', 'carbon'])
	],
	daisyui: {
		themes: ['night']
	},
	theme: {
		extend: {
			animation: {
				'bounce-slow': 'bounce 2s infinite ease-in-out',
			}
		}
	}
};
