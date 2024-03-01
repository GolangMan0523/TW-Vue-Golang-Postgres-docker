const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  mode: 'jit',
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class', // or 'media' or 'class'
  theme: {
    fontFamily: {
      sans: ['Helvetica Neue var', ...defaultTheme.fontFamily.sans],
    },
    extend: {
      colors: {
        blue: '#6768CB',
        darkblue: '#2795D9',
        lightblue: '#DBDBF3',
        darkest: '#202327',
        darker: '#1C1F23',
        dark: '#2f3336',
        gray: '#D9D9D9',
        light: '#9797DB',
        lighter: '#DBDBF3',
        lightest: '#DFDFFF',
        success: '#17BF63',
        danger: '#E0245E',
      },
    },
    fill: (theme) => ({
      current: 'currentColor',
      primary: theme('colors.primary'),
    }),
  },
  variants: {
    extend: {
      display: ['group-hover'],
    },
  },
  plugins: [require('@tailwindcss/ui')],
}
