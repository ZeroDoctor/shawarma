import { redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */
export function load({ url }) {
	console.log(url);

	let data = {
		current: 0,
		projects: [
			{
				id: 'unique-id-1',
				content: 'github.com/ZeroDoc-s-Stack/zdapi',
				url: '#',
				children: []
			},
			{
				id: 'unique-id-2',
				content: 'github.com/ZeroDoc-s-Stack/zdweb',
				url: '#',
				children: []
			}
		]
	};

    if(data.projects.length > 0) {
        redirect(307, '/projects/'+data.projects[0].id);
    }


	return data;
}