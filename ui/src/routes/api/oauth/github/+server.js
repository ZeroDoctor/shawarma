import { error, redirect } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function GET({ url }) {
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
    } catch(e) {
        console.error("[fetch] failed to register github user", e);
    }


    return redirect(302, '/project')
}