<script>
	import { uuidv4 } from '$lib/random.js';

    /** @type {import('./$types').PageData} */
    export let data;

    function githubRedirect() {
        var url = new URL("https://github.com/login/oauth/authorize");
        url.searchParams.set('client_id', data.clientID);
        url.searchParams.set('redirect_uri', "http://localhost:3000/api/oauth/github");
        url.searchParams.set('state', uuidv4());
        url.searchParams.set('scope', "read:user read:org");

        // @ts-ignore
        window.location.assign(url.toString());
    }

</script>

<div class="body flex justify-center items-center">
    <button type="button" class="btn variant-filled" on:click={githubRedirect}>Welcome</button>
</div>

<style>
    .body {
        height: 100vh;
    }
</style>
