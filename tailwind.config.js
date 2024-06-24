module.exports = {
  content: ["internal/**/*.go", "internal/public/*.js"],

  plugins: [
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography")
  ],
};
