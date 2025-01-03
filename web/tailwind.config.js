/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './index.html'
  ], 
  theme: {
    extend: {
      colors: {
        'myblue': '#007aff',
        'mygray': '#e2e2e2'
      }
    },
  },
  plugins: [],
}

