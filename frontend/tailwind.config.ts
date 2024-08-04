import type { Config } from 'tailwindcss'
import daisyui from "daisyui"
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
    themes: ["light", "dark", "cupcake"],
  },
} satisfies Config

