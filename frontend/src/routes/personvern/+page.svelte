<script>
	import ButtonComponent from "../../components/ButtonComponent.svelte";

</script>
<div class="flex flex-col justify-between items-center h-full">
    <div class="flex h-full mt-8">
        <div class="flex flex-col gap-4 px-64 items-center">
            <h1 class="text-3xl text-primary font-bold">Personvern</h1>
            <p>Undersøkelsen er helt anonym og informasjonen vil bli brukt til å forbedre vår tjeneste.</p>
        </div>
    </div>
    <div class="flex justify-center items-center gap-8 text-primary font-bold">
        <button><a href="/">Tilbake</a></button>
    </div>
</div>

<style lang="postcss">
    :root {
        color: theme(colors.content);
        
    }
    p {
        font-size: large;
    }
    button {
        font-size: large;
    }
</style>