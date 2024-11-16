/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    return {
        tree: [
            {
                id: 'unique-id-1',
                content: 'github.com/ZeroDoc-s-Stack/zdapi',
                lead: '#',
                children: []
            },
            {
                id: 'unique-id-2',
                content: 'github.com/ZeroDoc-s-Stack/zdweb',
                lead: '#',
                children: []
            }
        ]
    };
}