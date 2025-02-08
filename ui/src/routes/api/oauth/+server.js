import { uuidv4 } from '$lib/random';
import { error, redirect } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function GET({ url, cookies }) {
	let data = {
		code: url.searchParams.get('code'),
		state: url.searchParams.get('state') || uuidv4(),
		type: url.searchParams.get('type')
	};

	try {
		let endpoint = import.meta.env.VITE_SERVER_ENDPOINT;
		let resp = await fetch(endpoint + '/v1/register/user', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		});

		if (!resp || !resp.ok) {
			console.error('[fetch] failed to register github user', resp);
		}

		let errorMsg = '';
		switch (resp.status) {
			case 401:
				errorMsg = 'Unauthorized';
				cookies.delete('shawarma_user', { path: '/', domain: import.meta.env.VITE_JWT_DOMAIN });
				return redirect(307, '/');
			case 403:
				errorMsg = 'Forbidden';
				cookies.delete('shawarma_user', { path: '/', domain: import.meta.env.VITE_JWT_DOMAIN });
				return error(resp.status, errorMsg);
		}

		if (resp.status !== 202) {
			return error(500, 'Something went wrong');
		}

		return new Response('', {
			status: 307,
			headers: {
				location: `/projects?state=${data.state}&type=${data.type}&token=true`
			}
		});
	} catch (err) {
		console.error('[fetch] failed to register github user', err);
		return error(500, 'Something went wrong');
	}
}
