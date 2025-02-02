// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			verified?: boolean;
			uuid?: int[]?;
			name?: string?;
			session?: int[]?;
			avatar_url?: string?;
			created_at?: string?;
			modified_at?: string?;
			tokens?: {
				github?: string?;
			};
			exp?: int?;
			iat?: int?;
			iss?: string?;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
