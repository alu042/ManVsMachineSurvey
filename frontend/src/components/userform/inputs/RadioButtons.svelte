<script lang="ts">
    export let options:string[]
</script>

<div class="flex justify-between">
{#each options as data, index (data)}
    <div class="flex flex-col justify-center items-center w-24 ">
        <button name={data} class="h-7 w-7 rounded-full border-2 border-primary hover:bg-primary"></button>
        <label class="text-primary text-sm mt-1" for={data}>{data}</label>
    </div>
{/each}
</div>
