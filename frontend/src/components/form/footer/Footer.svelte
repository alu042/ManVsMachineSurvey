<script lang="ts">
	import ArrowChevron from "../../svg/ArrowChevron.svelte";
    import ButtonComponent from "../../userform/inputs/ButtonComponent.svelte";
    import { goto } from "$app/navigation";
	import { postFormData } from "../../../api/postFormData";

    export let questionNum:number
    export let answeredAll:boolean

    const gotoNextPage = async (questionNum: number) => {        
        goto(`${questionNum + 1}`)
    }

    const gotoPrevPage = async (questionNum: number) => {
        goto(`${questionNum - 1}`)
    }

    const handleFormSubmit = async () => {
        let allFormAnswers:string | null = localStorage.getItem("allFormAnswers")
        let respondentID:string | null = localStorage.getItem("RespondentId")

        if (allFormAnswers && respondentID) {
            await postFormData(Number(respondentID), allFormAnswers)
            goto("/evaluation")
        }
    }

</script>

<div class="flex justify-center items-center gap-8 pb-4 mt-4 md:mb-5">
    <button disabled={questionNum == 0} class={`flex items-center gap-2 text-primary font-semibold ${questionNum == 0 && "opacity-50"}`} on:click={() => gotoPrevPage(questionNum)}>
        <ArrowChevron width=16 direction="left"/>
        Forrige spørsmål
    </button>
    <button disabled={questionNum == 0 || questionNum % 4 != 0} on:click={handleFormSubmit}
    class={`${(questionNum == 0 || questionNum % 4 != 0) || (questionNum % 4 == 0 && !answeredAll) ?  "hidden" : "bg-primary text-bg hover:bg-bg hover:text-primary"} font-bold uppercase border-primary border-2 rounded-full px-2 md:px-8 py-3`}>
            Send inn svar
    </button>
    <button disabled={(questionNum != 0 && questionNum % 4 == 0) || !answeredAll} class={`flex items-center gap-2 text-primary font-semibold ${(questionNum != 0 && questionNum % 4 == 0) || !answeredAll && "opacity-50"} ${questionNum != 0 && questionNum % 4 == 0 && "opacity-50"} ${questionNum % 4 == 0 && questionNum != 0 &&  answeredAll && "hidden"}`} on:click={() => gotoNextPage(questionNum)}>
        Neste spørsmål
        <ArrowChevron width=16 direction="right"/>
    </button>
</div>