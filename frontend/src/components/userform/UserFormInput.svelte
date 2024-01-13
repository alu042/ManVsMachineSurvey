<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    
    const dispatch = createEventDispatcher();
    function handleFormChange(selectedValue: string) {
        dispatch('update', selectedValue);
    }

    export let inputType:string
    export let label:string
    export let options: string[]
    export let formData: string
	import RadioButtons from "./inputs/RadioButtons.svelte"
    import Select from "./inputs/Select.svelte";
</script>

<div class="flex flex-col md:flex-row justify-between md:justify-start items-start gap-3 md:gap-0 md:h-1/4 w-full">
    <p class="text-primary font-bold w-auto">{label}:</p>
    {#if inputType == "radio"}
        <RadioButtons on:update={(e) => handleFormChange(e.detail)} selected={formData} options={options}/>
    {/if}
    {#if inputType == "select"}
        <Select on:update={(e) => handleFormChange(e.detail)} options={options}/>
    {/if}
</div>
