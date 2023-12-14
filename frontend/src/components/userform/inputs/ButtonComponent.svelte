<script lang="ts">
    export let text: string;
    export let url: string;
    export let filled = false;

    let style:string;

    if (!filled) {
        style = "text-primary font-bold uppercase border-primary border-2 rounded-full px-8 py-3 hover:bg-primary hover:text-bg";
    } else {
        style = "bg-primary text-bg font-bold uppercase border-primary border-2 rounded-full px-8 py-3 hover:bg-bg hover:text-primary";
    }
</script>

<a href={url}><button class={style}>{text}</button></a>
