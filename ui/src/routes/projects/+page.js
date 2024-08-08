/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    return {
        tree: [
            {
                id: 'unique-id-1',
                content: 'home',
                lead: '/',
                children: []
            },
            {
                id: 'unique-id-2',
                content: 'youtube',
                lead: 'https://www.youtube.com',
                children: []
            }
        ]
    };
}