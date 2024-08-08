import jwt from '@tsndr/cloudflare-worker-jwt';


/**
 * create sign jwt secrets
 * 
 * @function
 * @name sign_secret
 * @kind function
 * @param {any} secret
 * @returns {Promise<string>}
 * @exports
 */
export function sign_secret(secret) {
    secret['nbf'] = Math.floor(Date.now() / 1000); // Not Before: Now
    secret['exp'] = Math.floor(Date.now() / 1000) + (72 * (60 * 60)); // Expires: Now + 72h
    secret['issuer'] = import.meta.env.VITE_JWT_ISSUER;
    return jwt.sign(
        secret,
        import.meta.env.VITE_JWT_KEY,
        import.meta.env.VITE_JWT_ENC
    )
}

/**
 * Description
 * 
 * @async
 * @function
 * @name verify_secret
 * @kind function
 * @param {any} token
 * @returns {Promise<false | { iss?: string; sub?: string; aud?: string | string[]; exp?: number; nbf?: number; iat?: number; jti?: string; } | undefined>}
 * @exports
 */
export async function verify_secret(token) {
  const valid = await jwt.verify(
    token, 
    import.meta.env.VITE_JWT_KEY
  )

  if (!valid) {
    return false 
  }

  return jwt.decode(token).payload
}
