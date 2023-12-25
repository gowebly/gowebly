/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
		"**/*.{html,templ}",
		"./node_modules/flowbite/**/*.js"
	],
    theme: {
        extend: {},
    },
    plugins: [
        require("@tailwindcss/typography"),
		require('flowbite/plugin')
    ],
}
