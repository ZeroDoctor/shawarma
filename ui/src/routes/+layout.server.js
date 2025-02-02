/** @type {import('./$types').LayoutServerLoad} */
export async function load({ locals }) {
  console.info('[layout] loading for /', locals);
  return locals
};
