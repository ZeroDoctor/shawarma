<script>
	import { onDestroy, onMount } from "svelte";

    export let projectId = "";

    onMount(() => {
        console.log('mount project settings =>', projectId)
    });

    onDestroy(() => {
        console.log('destory project settings =>', projectId)
    });
</script>