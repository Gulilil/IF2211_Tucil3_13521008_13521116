/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      color: {
        'themeBlack': "#1A1A1A",
        'themeGray': "#616571",
        'themeBlue': "#93DEFF",
        'themeWhite': "#F7F7F7",
      },
    },
  },
  plugins: [],
};
