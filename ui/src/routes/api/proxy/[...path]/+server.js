
/** @type {import('../$types').RequestHandler} */
export async function GET({ url, fetch }) {
    let path_split = url.pathname.split('/api/proxy');
    let path_name = path_split[0]+path_split[1];

    let endpoint = import.meta.env.VITE_SERVER_ENDPOINT;
    console.info(`[proxy] fetching ${endpoint + path_name}`);
    let resp = await fetch(endpoint + path_name, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    console.info(`[proxy] response status=${resp.status}`);

    return new Response(resp.body, resp);
}