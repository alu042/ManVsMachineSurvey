<script lang="ts">
    import UserFormInput from "../../userform/UserFormInput.svelte";
    import { convertAnswerDataToNumeric, convertAnswerDataToString } from "$lib/convertAnswerData";
	import { onMount } from "svelte";
    import { createEventDispatcher } from 'svelte';
    
    const dispatch = createEventDispatcher();

    export let answerText:string
    export let answerNum:number
    export let answerID:number
    
    let knowledge:string
    let empathy:string
    let helpfulness:string

    let isExpanded:boolean = false

    const handleFormUpdate = () => {
        if (knowledge && empathy && helpfulness) {
            dispatch("update", true)
        }

        let knowledgeNumeric = convertAnswerDataToNumeric(knowledge)
        let empathyNumeric = convertAnswerDataToNumeric(empathy)
        let helpfulnessNumeric = convertAnswerDataToNumeric(helpfulness)

        let allFormAnswers = localStorage.getItem("allFormAnswers")

        // First time submitting answer. Create map-array and add it to localstorage
        if (!allFormAnswers) {
            const formAnswer = new Map()
            formAnswer.set(answerID, [knowledgeNumeric, empathyNumeric, helpfulnessNumeric])
            let allFormAnswers = Array.from(formAnswer.entries())
            localStorage.setItem('allFormAnswers', JSON.stringify(allFormAnswers));
        } else {
            let existingFormAnswers = JSON.parse(allFormAnswers);
            let retrievedFormAnswers = new Map(existingFormAnswers)
            retrievedFormAnswers.set(answerID, [knowledgeNumeric, empathyNumeric, helpfulnessNumeric])
            let newFormAnswers = Array.from(retrievedFormAnswers.entries())
            localStorage.setItem('allFormAnswers', JSON.stringify(newFormAnswers));
        }
    }

    onMount(() => {
        let allFormAnswers:string | null = localStorage.getItem("allFormAnswers")

        if (allFormAnswers) {
            let existingFormAnswers = JSON.parse(allFormAnswers);
            let retrievedFormAnswers:Map<number, number[]> = new Map(existingFormAnswers)
            let formAnswers = retrievedFormAnswers.get(answerID)
            if (formAnswers) {
                knowledge = convertAnswerDataToString(formAnswers[0])
                empathy = convertAnswerDataToString(formAnswers[1])
                helpfulness = convertAnswerDataToString(formAnswers[2])
            }
        } else {
            knowledge=""
            empathy=""
            helpfulness=""
        }
    })

</script>

<div class="flex flex-col">
    <div class="flex flex-col gap-2 mb-6">
        <h1 class="text-xl text-primary font-bold text-center">Svar {answerNum}:</h1>
        <div class="flex flex-col justify-between bg-secondary rounded-xl px-6 pt-6">     
            <div class={`text-sm ${!isExpanded && "max-h-48"} overflow-hidden`}>
                {answerText}
            </div>
            <button class="py-2 text-primary font-semibold" on:click={() => isExpanded = !isExpanded}>
                {isExpanded ? "Vis mindre" : "Vis resten av svaret"}
            </button>
        </div>
    </div>
    <div class="flex flex-col justify-start items-center gap-6">
        <UserFormInput formData={knowledge} on:update={(e) => {knowledge = e.detail; handleFormUpdate()}} inputType="radio" label="Kunnskap" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra", "Vet ikke"]}/>
        <UserFormInput formData={empathy} on:update={(e) => {empathy = e.detail; handleFormUpdate()}} inputType="radio" label="Empati" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput formData={helpfulness} on:update={(e) => {helpfulness = e.detail; handleFormUpdate()}} inputType="radio" label="Hjelpsomhet" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
    </div>
</div>