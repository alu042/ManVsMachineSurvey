<script lang="ts">
    export let inputType:string
    export let label:string
    export let options: string[]
	import RadioButtons from "./inputs/RadioButtons.svelte"
    import Select from "./inputs/Select.svelte";
</script>

<div class="flex justify-start items-center h-1/4 w-full ">
    <p class="text-primary font-bold w-1/5">{label}:</p>
    {#if inputType == "radio"}
        <RadioButtons options={options}/>
    {/if}
    {#if inputType == "select"}
        <Select options={options}/>
    {/if}
</div>
