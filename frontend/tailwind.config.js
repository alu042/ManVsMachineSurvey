/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		extend: {
			colors: {
				primary: "#08667C",
				secondary: "#CED9DD",
				content: "#121212",
				bg: "#EEF0F3",
				error: "#964747",
			},
		},
	},
	plugins: [],
};
