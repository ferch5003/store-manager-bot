import type { Config } from 'tailwindcss'
import daisyui from "daisyui"
import themes from "daisyui/src/theming/themes"
import typography from '@tailwindcss/typography';

export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,vue}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    daisyui,
    typography,
  ],
  daisyui: {
    themes: [{
      "light": {
        ...themes.light,
        primary: "#FBBF24",
        secondary: "#2460FB",
        "info-content": "#D1D5DB"
      },
      "dark": {
        ...themes.dark,
        primary: "#FBBF24",
        secondary: "#2460FB",
        "info-content": "#111827"
      }
    }],
  },
} satisfies Config

