module.exports = {
  content: [
    'internal/*.html',
    'internal/**/*.html',
    "internal/public/*.js",  
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}