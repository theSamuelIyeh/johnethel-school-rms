/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*/*/*.{templ,html}", "node_modules/preline/dist/*.js"],
  theme: {
    fontFamily: {
      sans: ["Poppins", "sans-serif"],
    },
    extend: {
      colors: {
        "surface-default": "rgba(var(--surface-default))",
        "surface-primary": "rgba(var(--surface-primary))",
        "surface-secondary": "rgba(var(--surface-secondary))",
        "surface-tertiary": "rgba(var(--surface-tertiary))",
        "surface-brand": "rgba(var(--surface-brand))",
        "surface-accent": "rgba(var(--surface-accent))",
        "surface-invert": "rgba(var(--surface-invert))",
        "surface-warning": "rgba(var(--surface-warning))",
        "surface-error": "rgba(var(--surface-error))",
        "surface-success": "rgba(var(--surface-success))",

        "text-default": "rgba(var(--text-default))",
        "text-primary": "rgba(var(--text-primary))",
        "text-secondary": "rgba(var(--text-secondary))",
        "text-tertiary": "rgba(var(--text-tertiary))",
        "text-brand": "rgba(var(--text-brand))",
        "text-accent": "rgba(var(--text-accent))",
        "text-invert": "rgba(var(--text-invert))",
        "text-warning": "rgba(var(--text-warning))",
        "text-error": "rgba(var(--text-error))",
        "text-success": "rgba(var(--text-success))",

        "border-default": "rgba(var(--border-default))",
        "border-primary": "rgba(var(--border-primary))",
        "border-secondary": "rgba(var(--border-secondary))",
        "border-tertiary": "rgba(var(--border-tertiary))",
        "border-brand": "rgba(var(--border-brand))",
        "border-accent": "rgba(var(--border-accent))",
        "border-invert": "rgba(var(--border-invert))",
        "border-warning": "rgba(var(--border-warning))",
        "border-error": "rgba(var(--border-error))",
        "border-success": "rgba(var(--border-success))",

        // "primary-200": "rgba(var(--primary-200))",
        // "primary-300": "rgba(var(--primary-300))",
        // "accent-200": "rgba(var(--accent-200))",
        // "accent-200": "rgba(var(--accent-200))",
        // "base-200": "rgba(var(--base-200))",
        // "base-300": "rgba(var(--base-300))",
        // "btn-text": "rgba(var(--btn-text))",
        //   primary: "rgba(var(--primary))", // Define your primary color
        //   accent: "rgba(var(--accent))", // Define your accent color
        //   "base-100": "rgba(var(--base-100))", // Define your background color
        //   text: "rgba(var(--text))", // Define your text color
        //   info: "rgba(var(--info))", // Define your info color
        //   error: "rgba(var(--error))", // Define your error color
        //   warning: "rgba(var(--warning))", // Define your warning color
        //   success: "rgba(var(--success))", // Define your success color
      },
    },
  },
  plugins: [
    // require("@tailwindcss/forms"),
    // require("preline/plugin"),
    // require("@tailwindcss/typography"),
    require("daisyui"),
  ],

  daisyui: {
    themes: [
      {
        myLightTheme: {
          // "color-scheme": "light",
          // primary: "#3396d3", // Define your primary color
          // "primary-200": "#67b1de",
          // "primary-300": "#99cbe9",
          // accent: "#8dc63f", // Define your accent color
          // "accent-200": "#a8d56f",
          // "accent-300": "#c6e29e",
          // "base-100": "#f5f3ff", // Define your background color
          // "btn-text": "#161045",
          // text: "#161045", // Define your text color
          // info: "#250de5", // Define your info color
          // error: "#ec1a24", // Define your error color
          // warning: "#fff201", // Define your warning color
          // success: "#8dc63f", // Define your success color
          "--rounded-box": "1rem", // border radius rounded-box utility class, used in card and other large boxes
          "--rounded-btn": "0.5rem", // border radius rounded-btn utility class, used in buttons and similar element
          "--rounded-badge": "1.9rem", // border radius rounded-badge utility class, used in badges and similar
          "--animation-btn": "0.25s", // duration of animation when you click on button
          "--animation-input": "0.2s", // duration of animation for inputs like checkbox, toggle, radio, etc
          "--btn-focus-scale": "0.95", // scale transform of button when you focus on it
          "--border-btn": "1", // border width of buttons
          "--tab-border": "1", // border width of tabs
          "--tab-radius": "0.5rem", // border radius of tabs
        },
      },
      {
        myDarkTheme: {
          // "color-scheme": "dark",
          // primary: "#3396d3", // Define your primary color
          // "primary-200": "#26709f",
          // "primary-300": "#194b6a",
          // accent: "#8dc63f", // Define your accent color
          // "accent-200": "#6a942e",
          // "accent-300": "#476320",
          // "base-100": "#161045", // Define your text color
          // "btn-text": "#161045",
          // text: "#f5f3ff", // Define your background color
          // info: "#250de5", // Define your info color
          // error: "#ec1a24", // Define your error color
          // warning: "#fff201", // Define your warning color
          // success: "#8dc63f", // Define your success color
          "--rounded-box": "1rem", // border radius rounded-box utility class, used in card and other large boxes
          "--rounded-btn": "0.5rem", // border radius rounded-btn utility class, used in buttons and similar element
          "--rounded-badge": "1.9rem", // border radius rounded-badge utility class, used in badges and similar
          "--animation-btn": "0.25s", // duration of animation when you click on button
          "--animation-input": "0.2s", // duration of animation for inputs like checkbox, toggle, radio, etc
          "--btn-focus-scale": "0.95", // scale transform of button when you focus on it
          "--border-btn": "1px", // border width of buttons
          "--tab-border": "1px", // border width of tabs
          "--tab-radius": "0.5rem", // border radius of tabs
        },
      },
      "myLightTheme",
      "myDarkTheme",
    ],
  },
};
