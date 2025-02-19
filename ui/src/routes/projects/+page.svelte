
<script>
	import MenuIcon from '$lib/svg/MenuIcon.svelte';
    import { AppBar } from '@skeletonlabs/skeleton';
    import { Drawer, getDrawerStore, initializeStores } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';

    initializeStores();
    const drawerStore = getDrawerStore();

    onMount(async () => {
        let reposResponse = await fetch("/api/proxy/v1/repos");
        let repos = await reposResponse.json();
        console.info(`repos found: ${repos.length}`);
    });

    /** @type import("@skeletonlabs/skeleton").DrawerSettings */
    let drawerSettings = {
        width: 'w-[280px] md:w-[480px]'
    }
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
    <h4>Projects Not Found</h4>
</div>


<Drawer>
    <h4>Projects Not Found</h4>
</Drawer>
