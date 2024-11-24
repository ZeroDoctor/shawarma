<script>
	import AbortIcon from "$lib/svg/AbortIcon.svelte";
	import ErrorIcon from "$lib/svg/ErrorIcon.svelte";
	import PendingIcon from "$lib/svg/PendingIcon.svelte";
	import SuccessIcon from "$lib/svg/SuccessIcon.svelte";
    import { onDestroy, onMount } from "svelte";

    /**
     * @type {PipelineView[]}
     */
    export let pipelines = [];

    /**
     * @type {string}
     */
    export let projectId = "";

    onMount(() => {
        console.log('mount project pipeline =>', projectId)
    });

    onDestroy(() => {
        console.log('destory project pipeline =>', projectId)
    });
</script>

<div class="bg-surface-700">
    {#if !pipelines || pipelines.length <= 0 }
        No pipelines { projectId }
    {:else}
        {#each pipelines as pipeline}
        <dl class="list-dl">
            <div>
                <span class="badge">
                    <a href="{pipeline.url}">
                        {#if pipeline.status === "pending"}
                            <PendingIcon style="fill-warning-500" />
                        {:else if pipeline.status === "failed"}
                            <ErrorIcon style="fill-error-500" />
                        {:else if pipeline.status === "success"}
                            <SuccessIcon style="fill-success-500" />
                        {:else if pipeline.status === "aborted"}
                            <AbortIcon style="fill-tertiary-400" />
                        {/if}
                    </a>
                </span>
                <span class="flex-auto">
                    <dt><a href="{pipeline.url}">#4</a></dt>
                    <dd><a href="{pipeline.remoteCommitUrl}">commit hash: commit message</a></dd>
                </span>
            </div>
        </dl>
        {/each}
    {/if}
</div>