<script lang="ts">
    import CircleExclamation from "../../../components/svg/CircleExclamation.svelte";

    export let formQuestion:string
    export let questionNum:number
</script>

<div class="flex justify-center items-center md:h-20">
    <div class="hidden md:flex w-1/3"/>
    <div class="flex md:justify-center items-center md:w-2/3 gap-4">
        <h1 class="text-xl text-primary font-bold text-center">Spørsmål {Number(questionNum) + 1}</h1>
        <div class="border-r-2 border-primary min-h-24 h-full"></div>
        <p class="text-sm">{formQuestion}</p>
    </div>
    <div class="hidden md:flex justify-end pr-20 items-center w-1/3">
        <a class="-" href="/reportbug">
            <CircleExclamation width="16" height="16" />    
        </a>
    </div>
</div>