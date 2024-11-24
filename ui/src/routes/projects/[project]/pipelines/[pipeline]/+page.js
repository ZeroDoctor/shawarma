/** @type {import('./$types').PageLoad} */
export function load() {
	return {
		pipeline: {
			url: '',
			status: '',
			remoteCommitUrl: '',
            stages: [
				{
					name: "build",
					commands: ["go build -o main .", "echo done."],
					status: "executing",
					child: ["test"]
				}
			]
		}
	};
}
