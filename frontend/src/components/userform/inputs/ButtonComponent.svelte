<script lang="ts">
    export let text: string;
    export let url: string;
    export let filled = false;
    export let onclick: () => {};

    let style:string;

    style = `${!filled ? "text-primary hover:bg-primary hover:text-bg" 
    : "bg-primary text-bg hover:bg-bg hover:text-primary"} 
    font-bold uppercase border-primary border-2 rounded-full px-8 py-3`
</script>

<a href={url}><button class={style} on:click={onclick}>{text}</button></a>
