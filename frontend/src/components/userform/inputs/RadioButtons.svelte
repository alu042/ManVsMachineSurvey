<script lang="ts">
    export let options:string[]
</script>

<div class="flex justify-between w-full">
{#each options as data, index (data)}
    <div class="flex flex-col justify-start items-center w-12 text-center">
        <button name={data} class="h-6 w-6 rounded-full border-2 border-primary hover:bg-primary"></button>
        <label class="text-primary text-sm mt-1" for={data}>{data}</label>
    </div>
{/each}
</div>
