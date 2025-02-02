
/** @type {import('./$types').PageLoad} */
export async function load({ params, fetch }) {

    let endpoint = import.meta.env.VITE_SERVER_ENDPOINT;
	let reposResponse = await fetch(endpoint+"/v1/repos")

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
    }

    data.projects.forEach((project, index) => {
        if(params.project === project.id) {
            data.current = index;
        }
    });

    return data;
}
