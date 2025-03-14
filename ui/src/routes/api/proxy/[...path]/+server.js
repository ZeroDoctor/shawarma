
/** @type {import('../$types').RequestHandler} */
export async function GET({url, fetch, request}) {
    return proxy(url, fetch, request);
}

/** @type {import('../$types').RequestHandler} */
export async function POST({url, fetch, request}) {
    return proxy(url, fetch, request);
}

/** @type {import('../$types').RequestHandler} */
export async function PUT({url, fetch, request}) {
    return proxy(url, fetch, request);
}

/** @type {import('../$types').RequestHandler} */
export async function DELETE({url, fetch, request}) {
    return proxy(url, fetch, request);
}

/** @type {import('../$types').RequestHandler} */
export async function PATCH({url, fetch, request}) {
    return proxy(url, fetch, request);
}

/**
 * Proxies a request to the specified URL.
 *
 * @param {URL} url - The URL to proxy the request to.
 * @param {fetch} fetch - The fetch function to use for the request.
 * @param {Request} request - The request to proxy.
 * @return {Promise<Response>} The response from the proxied request.
 */
async function proxy(url, fetch, request) {
    let path_split = url.pathname.split('/api/proxy');
    let path_name = path_split[0]+path_split[1];

    let endpoint = import.meta.env.VITE_SERVER_ENDPOINT;
    console.info(`[proxy] fetching ${endpoint + path_name}`);
    let resp = await fetch(endpoint + path_name, request);
    console.info(`[proxy] response status=${resp.status}`);

    return new Response(resp.body, resp);
}