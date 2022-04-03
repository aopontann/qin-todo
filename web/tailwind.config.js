module.exports = {
  mode: 'jit',
  purge: ['./src/**/*.{js,ts,jsx,tsx}'],
  darkMode: false,
  content: [],
  theme: {
    extend: {
      colors: {
        rose: '#F43F5E',
        orange: '#FB923C',
        yellow: '#FBBF24',
        gray: '#C2C6D2',
        red: '#EF4444',
        surface: '#F1F5F9',
      },
    },
    fontSize: {
      xs: '.75rem',
      sm: '.875rem',
      tiny: '.875rem',
      base: '1rem',
      lg: '1.125rem',
      xl: '1.25rem',
      '2xl': '1.5rem',
      '3xl': '1.875rem',
      '4xl': '2.25rem',
      '5xl': '3rem',
      '6xl': '4rem',
      '7xl': '5rem',
      h2: '1.375rem',
    },
  },
  plugins: [],
};
