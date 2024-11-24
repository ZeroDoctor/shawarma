
<script>
	import ProjectLogs from '$lib/components/ProjectLogs.svelte';
	import ProjectPipeline from '$lib/components/ProjectPipeline.svelte';
	import ProjectSettings from '$lib/components/ProjectSettings.svelte';
	import MenuIcon from '$lib/svg/MenuIcon.svelte';
    import { AppBar, Tab, TabGroup } from '@skeletonlabs/skeleton';
    import { Drawer, getDrawerStore, initializeStores } from '@skeletonlabs/skeleton';

    initializeStores();
    const drawerStore = getDrawerStore();

    /** @type import("@skeletonlabs/skeleton").DrawerSettings */
    let drawerSettings = {
        width: 'w-[280px] md:w-[480px]'
    }

    /** @type {import('./$types').PageData} */
    export let data;

    /** @type ProjectNav[] */
    let projects = data.projects;
    let currentProject = data.current;
    let tabset = 0;
</script>

<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end" background="bg-surface-900 shadow-2xl">
	<svelte:fragment slot="lead">
        <button class="btn bg-initial" on:click={() => drawerStore.open(drawerSettings)}> 
            <MenuIcon style="fill-primary-100" />
        </button>
    </svelte:fragment>
    <span class="text-primary-100"> Shawarma </span>
	<svelte:fragment slot="trail">
        (actions)
    </svelte:fragment>
</AppBar>

<div class="flex flex-row max-w-full bg-surface-900 shadow-2xl">
    <nav class="hidden lg:block lg:w-3/12 list-nav p-3">
        <ul class="list">
            {#each projects as project, index}
            <li class="">
                <div class="flex items-center break-words">
                    <span>-</span>
                    <a class="btn bg-initial w-full whitespace-normal" href={'/projects/'+project.id}
                        on:click={() => {currentProject = index}}>
                        {project.content}
                    </a>
                </div>
            </li>
            {/each}
        </ul>
    </nav>

    <div class="w-full p-3">
        <TabGroup>
            <Tab bind:group={tabset} name="pipeline" value={0}> Pipeline </Tab>
            <Tab bind:group={tabset} name="logs" value={1}> Logs </Tab>
            <Tab bind:group={tabset} name="settings" value={2}> Settings </Tab>

            <svelte:fragment slot="panel">
            {#if tabset === 0}
                <ProjectPipeline projectId={data.projects[currentProject].id} />
            {:else if tabset === 1}
                <ProjectLogs projectId={data.projects[currentProject].id} />
            {:else if tabset === 2}
                <ProjectSettings projectId={data.projects[currentProject].id} />
            {/if}
            </svelte:fragment>
        </TabGroup>
    </div>
</div>


<Drawer>
    <nav class="block list-nav mt-8 ml-3">
        <ul class="list">
            {#each projects as project, index}
            <li>
                <div class="flex items-center">
                    <span>-</span>
                    <a class="btn bg-initial w-full whitespace-normal" href={project.url} 
                        on:click={() => {currentProject = index}}>
                        {project.content}
                    </a>
                </div>
            </li>
            {/each}
        </ul>
    </nav>
</Drawer>
