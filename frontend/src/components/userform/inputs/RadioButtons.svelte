<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();
    function handleButtonClick(selectedValue: string) {
        dispatch('update', selectedValue);
    }
    export let options:string[]
    export let selected: string
</script>

<div class="flex justify-between w-full">
{#each options as data, index (data)}
    <button name={data} on:click={() => handleButtonClick(data)} class="flex flex-col justify-start items-center w-12 text-center">
        <div class={`${selected == data && "bg-primary"} h-6 w-6 rounded-full border-2 border-primary hover:bg-primary`}></div>
        <label class="text-primary text-sm mt-1" for={data}>{data}</label>
    </button>
{/each}
</div>
