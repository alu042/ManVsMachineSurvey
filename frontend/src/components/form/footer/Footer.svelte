<script>
	import ArrowChevron from "../../svg/ArrowChevron.svelte";
    import ButtonComponent from "../../userform/inputs/ButtonComponent.svelte";
</script>

<div class="flex justify-center items-center gap-8">
    <a href="/">
        <button class="flex items-center gap-2 text-primary font-semibold">
            <ArrowChevron width=16 direction="left"/>
            Forrige spørsmål
        </button>
    </a>
    <ButtonComponent text="Send inn svar" url="/" filled={true} />
    <a href="/">
        <button class="flex items-center gap-2 text-primary font-semibold opacity-50">
            Neste spørsmål
            <ArrowChevron width=16 direction="right"/>
        </button>
    </a>
</div>