module.exports = {
  content: [
    'internal/web/*.html',
    'internal/web/**/*.html',
    "internal/web/public/*.js",  
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}