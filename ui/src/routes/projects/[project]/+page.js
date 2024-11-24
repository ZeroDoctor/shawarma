
/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    console.log(params.project);
    return {
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
}
