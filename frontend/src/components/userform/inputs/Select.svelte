<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();
    function handleButtonClick(selectedValue: string) {
        dispatch('update', selectedValue);
    }
    export let options:string[]
</script>

<select class="ml-6 pl-2 pr-16 py-2 rounded-xl text-primary border-primary border-2 text-start hover:cursor-pointer ">
    {#each options as data, index (data)}
        <option on:click={() => handleButtonClick(data)} id={data}>{data}</option>
    {/each}
</select>
