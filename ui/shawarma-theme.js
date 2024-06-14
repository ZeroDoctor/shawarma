export const ShawarmaTheme = {
    name: 'shawarma-theme',
    properties: {
		// =~= Theme Properties =~=
		"--theme-font-family-base": `Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'`,
		"--theme-font-family-heading": `Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'`,
		"--theme-font-color-base": "0 0 0",
		"--theme-font-color-dark": "255 255 255",
		"--theme-rounded-base": "12px",
		"--theme-rounded-container": "8px",
		"--theme-border-base": "2px",
		// =~= Theme On-X Colors =~=
		"--on-primary": "255 255 255",
		"--on-secondary": "0 0 0",
		"--on-tertiary": "255 255 255",
		"--on-success": "var(--color-tertiary-900)",
		"--on-warning": "var(--color-tertiary-900)",
		"--on-error": "255 255 255",
		"--on-surface": "255 255 255",
		// =~= Theme Colors  =~=
		// primary | #8A5A44 
		"--color-primary-50": "237 230 227", // #ede6e3
		"--color-primary-100": "232 222 218", // #e8deda
		"--color-primary-200": "226 214 208", // #e2d6d0
		"--color-primary-300": "208 189 180", // #d0bdb4
		"--color-primary-400": "173 140 124", // #ad8c7c
		"--color-primary-500": "138 90 68", // #8A5A44
		"--color-primary-600": "124 81 61", // #7c513d
		"--color-primary-700": "104 68 51", // #684433
		"--color-primary-800": "83 54 41", // #533629
		"--color-primary-900": "68 44 33", // #442c21
		// secondary | #FFFFFF 
		"--color-secondary-50": "255 255 255", // #ffffff
		"--color-secondary-100": "255 255 255", // #ffffff
		"--color-secondary-200": "255 255 255", // #ffffff
		"--color-secondary-300": "255 255 255", // #ffffff
		"--color-secondary-400": "255 255 255", // #ffffff
		"--color-secondary-500": "255 255 255", // #FFFFFF
		"--color-secondary-600": "230 230 230", // #e6e6e6
		"--color-secondary-700": "191 191 191", // #bfbfbf
		"--color-secondary-800": "153 153 153", // #999999
		"--color-secondary-900": "125 125 125", // #7d7d7d
		// tertiary | #5B3B0C 
		"--color-tertiary-50": "230 226 219", // #e6e2db
		"--color-tertiary-100": "222 216 206", // #ded8ce
		"--color-tertiary-200": "214 206 194", // #d6cec2
		"--color-tertiary-300": "189 177 158", // #bdb19e
		"--color-tertiary-400": "140 118 85", // #8c7655
		"--color-tertiary-500": "91 59 12", // #5B3B0C
		"--color-tertiary-600": "82 53 11", // #52350b
		"--color-tertiary-700": "68 44 9", // #442c09
		"--color-tertiary-800": "55 35 7", // #372307
		"--color-tertiary-900": "45 29 6", // #2d1d06
		// success | #2ECC71 
		"--color-success-50": "224 247 234", // #e0f7ea
		"--color-success-100": "213 245 227", // #d5f5e3
		"--color-success-200": "203 242 220", // #cbf2dc
		"--color-success-300": "171 235 198", // #abebc6
		"--color-success-400": "109 219 156", // #6ddb9c
		"--color-success-500": "46 204 113", // #2ECC71
		"--color-success-600": "41 184 102", // #29b866
		"--color-success-700": "35 153 85", // #239955
		"--color-success-800": "28 122 68", // #1c7a44
		"--color-success-900": "23 100 55", // #176437
		// warning | #E9B506 
		"--color-warning-50": "252 244 218", // #fcf4da
		"--color-warning-100": "251 240 205", // #fbf0cd
		"--color-warning-200": "250 237 193", // #faedc1
		"--color-warning-300": "246 225 155", // #f6e19b
		"--color-warning-400": "240 203 81", // #f0cb51
		"--color-warning-500": "233 181 6", // #E9B506
		"--color-warning-600": "210 163 5", // #d2a305
		"--color-warning-700": "175 136 5", // #af8805
		"--color-warning-800": "140 109 4", // #8c6d04
		"--color-warning-900": "114 89 3", // #725903
		// error | #C0392B 
		"--color-error-50": "246 225 223", // #f6e1df
		"--color-error-100": "242 215 213", // #f2d7d5
		"--color-error-200": "239 206 202", // #efceca
		"--color-error-300": "230 176 170", // #e6b0aa
		"--color-error-400": "211 116 107", // #d3746b
		"--color-error-500": "192 57 43", // #C0392B
		"--color-error-600": "173 51 39", // #ad3327
		"--color-error-700": "144 43 32", // #902b20
		"--color-error-800": "115 34 26", // #73221a
		"--color-error-900": "94 28 21", // #5e1c15
		// surface | #1f2b3e 
		"--color-surface-50": "221 223 226", // #dddfe2
		"--color-surface-100": "210 213 216", // #d2d5d8
		"--color-surface-200": "199 202 207", // #c7cacf
		"--color-surface-300": "165 170 178", // #a5aab2
		"--color-surface-400": "98 107 120", // #626b78
		"--color-surface-500": "31 43 62", // #1f2b3e
		"--color-surface-600": "28 39 56", // #1c2738
		"--color-surface-700": "23 32 47", // #17202f
		"--color-surface-800": "19 26 37", // #131a25
		"--color-surface-900": "15 21 30", // #0f151e
		
	}
}