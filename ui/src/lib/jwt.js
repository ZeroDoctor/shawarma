import jwt from '@tsndr/cloudflare-worker-jwt';

/**
 * verify jwt token
 *
 * @async
 * @function
 * @name verify_jwt
 * @kind function
 * @param {any} token
 * @returns {Promise<false | { iss?: string; sub?: string; aud?: string | string[]; exp?: number; nbf?: number; iat?: number; jti?: string; } | undefined>}
 * @exports
 */
export async function verify_jwt(token) {
	try {
		const valid = await jwt.verify(
			token,
			import.meta.env.VITE_JWT_SECRET,
			import.meta.env.VITE_JWT_ENC
		);

		if (!valid) {
			return false;
		}
	} catch (error) {
		console.error(`[verify_jwt] failed to verify token=${token}`, error);
		return false;
	}

	return jwt.decode(token).payload;
}

/**
 * Parse a JWT token.
 *
 * @async
 * @function
 * @name parse_jwt
 * @kind function
 * @param {string} token - The JWT token to parse.
 * @returns {import('@tsndr/cloudflare-worker-jwt').JwtData} The parsed JWT token payload, or undefined if parsing fails.
 */
export function parse_jwt(token) {
  return jwt.decode(token);
}