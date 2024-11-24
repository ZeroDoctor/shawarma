<script>
	import MenuIcon from '$lib/svg/MenuIcon.svelte';
    import { AppBar } from '@skeletonlabs/skeleton';
    import { Drawer, getDrawerStore, initializeStores } from '@skeletonlabs/skeleton';

    initializeStores();
    const drawerStore = getDrawerStore();

    /** @type import("@skeletonlabs/skeleton").DrawerSettings */
    let drawerSettings = {
        width: 'w-[280px] md:w-[480px]'
    }

    /** @type {import('./$types').PageData} */
    export let data;
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

<Drawer>
    <nav class="block list-nav mt-8 ml-3">
        <ul class="list">
        </ul>
    </nav>
</Drawer>
