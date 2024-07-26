import { error, redirect } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function GET({ url, cookies }) {
    let data = {
        code: url.searchParams.get('code'),
        state: url.searchParams.get('state'),
        type: "github"
    };

    try {
        let endpoint = import.meta.env.VITE_SERVER_ENDPOINT;
        let resp = await fetch(endpoint+'/v1/register/user', {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        });

        let token = JSON.stringify(await resp.json());
        cookies.set('shawarma_user', token,
            {
                path: '/',
                domain: import.meta.env.VITE_JWT_DOMAIN,
                sameSite: 'strict',
                httpOnly: true
            }
        );
    } catch(e) {
        console.error("[fetch] failed to register github user", e);
        return error(500, "Something went wrong");
    }

    return redirect(302, '/projects');
}