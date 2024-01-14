<script lang="ts">
	import ArrowChevron from "../../svg/ArrowChevron.svelte";
    import ButtonComponent from "../../userform/inputs/ButtonComponent.svelte";
    import { goto } from "$app/navigation";
	import { postFormData } from "../../../api/postFormData";

    export let questionNum:number
    export let answeredAll:boolean
    export let answer1ID:number
    export let answer2ID:number

    const gotoNextPage = async (questionNum: number) => {        
        goto(`${questionNum + 1}`)
    }
    
    const gotoPrevPage = async (questionNum: number) => {
        goto(`${questionNum - 1}`)
    }
    
    const skipQuestion = async (questionNum: number) => {
        let allFormAnswers:string | null = localStorage.getItem("allFormAnswers")
        if (allFormAnswers) {
            const convertedAnswers:[] = JSON.parse(allFormAnswers)
            let result = convertedAnswers.filter((item) => item[0] != answer1ID && item[0] != answer2ID)
            const localStorageResult = JSON.stringify(result)
            localStorage.setItem("allFormAnswers", localStorageResult)

            if (questionNum != 0 && questionNum % 4 == 0) {
                handleFormSubmit()
            } else {
                goto(`${questionNum + 1}`)
            }
        } else {
            goto(`${questionNum + 1}`)
        }
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

<div class="flex flex-col gap-8 justify-around items-center pb-6 md:mb-5">
    <div class="flex gap-8">
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
    <div class="flex gap-8">
        <button on:click={() => skipQuestion(questionNum)} class="border-2 border-primary text-primary rounded-3xl hover:bg-primary hover:text-bg px-3 py-2">
            Ønsker ikke vurdere dette spørsmålet
        </button>
        <button on:click={handleFormSubmit} disabled={(questionNum != 0 && questionNum % 4 == 0) || !answeredAll} class={`border-2 border-primary bg-primary text-bg rounded-3xl px-3 py-2 ${(questionNum != 0 && questionNum % 4 == 0) || !answeredAll && "opacity-50"} ${questionNum != 0 && questionNum % 4 == 0 && "opacity-50"} ${questionNum % 4 == 0 && questionNum != 0 &&  answeredAll && "hidden"}`}>
            Avslutt undersøkelsen og send inn svar
        </button>
    </div>
</div>