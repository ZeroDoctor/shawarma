
<script>
	import Link from '$lib/components/Link.svelte';
	import MenuIcon from '$lib/svg/MenuIcon.svelte';
    import { AppBar, RecursiveTreeView, Tab, TabGroup, TreeView, TreeViewItem } from '@skeletonlabs/skeleton';
    import { Drawer, getDrawerStore, initializeStores } from '@skeletonlabs/skeleton';

    initializeStores();
    const drawerStore = getDrawerStore();

    /** * @type import("@skeletonlabs/skeleton").DrawerSettings */
    let drawerSettings = {
        width: 'w-[280px] md:w-[480px]'
    }

    /** @type {import('./$types').PageData} */
    export let data;

    /** * @type import("@skeletonlabs/skeleton").TreeViewNode[] */
    let tree = data.tree;

    let tabset = 0;
</script>

<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
	<svelte:fragment slot="lead">
        <button class="btn variant-ringed" on:click={() => drawerStore.open(drawerSettings)}> 
            <MenuIcon />
        </button>
    </svelte:fragment>
    Shawarma
	<svelte:fragment slot="trail">(actions)</svelte:fragment>
</AppBar>

<div class="flex flex-row max-w-full">
    <div class="flex justify-center items-center w-2/12">
        <TreeView>
        {#each tree as element}
        <TreeViewItem>
            <Link style="btn vairant-ringed" content={element.content} href={element.lead}/>
        </TreeViewItem>
        {/each}
        </TreeView>
    </div>

    <div class="w-10/12">
        <TabGroup>
            <Tab bind:group={tabset} name="status" value={0}>Status</Tab>
            <Tab bind:group={tabset} name="logs" value={1}>Logs</Tab>
            <Tab bind:group={tabset} name="settings" value={2}>Settings</Tab>
            <svelte:fragment slot="panel">
            {#if tabset === 0}
                Testing
            {/if}
            </svelte:fragment>
        </TabGroup>
    </div>
</div>


<Drawer>
</Drawer>

<style>
</style>