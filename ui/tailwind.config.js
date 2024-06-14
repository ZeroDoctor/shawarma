
// @ts-check
import { join } from 'path';

import { skeleton } from '@skeletonlabs/tw-plugin';

import { ShawarmaTheme } from './shawarma-theme';

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: [
    "./src/**/*.{html,js,svelte,ts}",
    join(
      require.resolve('@skeletonlabs/skeleton'),
      '../**/*.{html,js,svelte,ts}'
    )
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/typography"),
    skeleton({
      themes: {
        custom: [
          ShawarmaTheme
        ]
      }
    })
  ],
}

