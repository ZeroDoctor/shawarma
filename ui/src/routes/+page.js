// @ts-nocheck
/** @type {import('./$types').PageLoad} */
export function load() {
	return {
        clientID: import.meta.env.VITE_GITHUB_CLIENT_ID
	};
}