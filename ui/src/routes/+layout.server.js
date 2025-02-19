/** @type {import('./$types').LayoutServerLoad} */
export async function load({ locals, url }) {
	console.info(`[layout] loading locals for path ${url.pathname}`);
	return locals;
}
