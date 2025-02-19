import * as cookie from 'cookie';
import { verify_jwt, parse_jwt } from '$lib/js/jwt';
import { redirect, error } from '@sveltejs/kit';

/**
 * @template {any} T
 * @typedef {T | Promise<T>} MaybePromise<T>
 */

/** @typedef {import('@sveltejs/kit').RequestEvent} Event */
/** @typedef {(event: import('@sveltejs/kit').RequestEvent, opt?: import('@sveltejs/kit').ResolveOptions) => MaybePromise<Response>} Resolve */

/** @type {import('@sveltejs/kit').Handle} */
export async function handle({ event, resolve }) {
	console.info(`[handle] hook for page ${event.url.pathname}`);
	let response = await fetch_token(event);
	if (response) {
		return response;
	}

	const jwt =
		cookie.parse(event.request.headers.get('cookie') || '')?.shawarma_user ||
		event.cookies.get('shawarma_user');
	if (!jwt) {
		console.info("[handle] user hasn't logged in");
		event.cookies.delete('shawarma_user', { path: '/', domain: import.meta.env.VITE_JWT_DOMAIN });
		return create_response(event, resolve);
	}

	const verified = await verify_jwt(jwt);
	if (!verified) {
		console.info(`[handle] user tried use unverified token=${jwt}`);
		event.cookies.delete('shawarma_user', { path: '/', domain: import.meta.env.VITE_JWT_DOMAIN });
		return create_response(event, resolve);
	}

	event.locals = parse_jwt(jwt).payload || { verified: false };
	event.locals.verified = true;
	console.debug(`[handle] user=${event.locals.name} jwt token verified=${event.locals.verified}`);

	return create_response(event, resolve);
}

/** @type {import('@sveltejs/kit').HandleServerError} */
export function handleError({ error, message, status }) {
	console.log(`[handleError] [error=${error}] [message=${message}] [status=${status}]`);

	return {
		message: message ?? 'Something went really wrong'
	};
}

/**
 * Creates a response with same-origin policies for Cross-Origin-Opener-Policy and Cross-Origin-Embedder-Policy.
 *
 * @param {Event} event - The event object.
 * @param {Resolve} resolve - The resolve function.
 * @return {Promise<Response>} The response object with updated headers.
 */
async function create_response(event, resolve) {
	if (event.url.pathname !== '/' && event.url.pathname !== '/api/oauth' && !event.locals.verified) {
		return redirect(307, '/');
	}

	const response = await resolve(event);
	response.headers.set('Cross-Origin-Opener-Policy', 'same-origin');
	response.headers.set('Cross-Origin-Embedder-Policy', 'same-origin');
	return response;
}

/**
 * Fetch token using type and state from query params passed in by /api/oauth
 *
 * @param {Event} event - The event object.
 * @return {Promise<Response|undefined>} The response object with updated headers.
 */
async function fetch_token(event) {
	let type = event.url.searchParams.get('type');
	let state = event.url.searchParams.get('state');
	let token = event.url.searchParams.get('token');
	if (!type || !state || !token) {
		return;
	}

	console.info(`[fetch_token] fetching token for type=${type} state=${state}`);
	let endpoint = import.meta.env.VITE_SERVER_ENDPOINT;
	let resp = await event.fetch(endpoint + `/v1/user?type=${type}&state=${state}`, {
		method: 'GET'
	});

	if (!resp || !resp.ok) {
		console.error('[fetch] failed to register github user', resp);
	}

	let errorMsg = '';
	switch (resp.status) {
		case 401:
		case 404:
			errorMsg = 'Unauthorized';
			event.cookies.delete('shawarma_user', {
				path: '/',
				domain: import.meta.env.VITE_JWT_DOMAIN
			});
			return redirect(307, '/');
		case 403:
			errorMsg = 'Forbidden';
			event.cookies.delete('shawarma_user', {
				path: '/',
				domain: import.meta.env.VITE_JWT_DOMAIN
			});
			return error(resp.status, errorMsg);
	}

	if (resp.status !== 202) {
		return error(500, 'Something went wrong');
	}

	let user = await resp.json();
	console.info(`[fetch_token] fetched token=${user.token}`);
	event.cookies.set('shawarma_user', user.token, {
		path: '/',
		domain: import.meta.env.VITE_JWT_DOMAIN
	});
}
