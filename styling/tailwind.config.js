/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../views/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],

  daisyui: {
    prefix: "daisy-",
    themes: [
      {
        light: {
          "primary": "#1e3a8a",
          "secondary": "#3b82f6",
          "accent": "#a78bfa",
          "neutral": "#60a5fa",
          "base-100": "#e0f2fe",
          "info": "#eab308",
          "success": "#84cc16",
          "warning": "#e11d48",
          "error": "#7f1d1d",
        },
      },
      {
        dark: {
          "primary": "#bfdbfe",
          "secondary": "#3b82f6",
          "accent": "#a78bfa",
          "neutral": "#bae6fd",
          "base-100": "#1e3a8a",
          "info": "#eab308",
          "success": "#84cc16",
          "warning": "#e11d48",
          "error": "#7f1d1d",
        },
      },
    ],
    darkTheme: "dark",
  }
}

