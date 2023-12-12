<script lang="ts">
    export let options:string[]
</script>

<select class="ml-6 pl-2 pr-16 py-2 rounded-xl text-primary border-primary border-2 text-start hover:cursor-pointer ">
    {#each options as data, index (data)}
        <option id={data}>{data}</option>
    {/each}
</select>
