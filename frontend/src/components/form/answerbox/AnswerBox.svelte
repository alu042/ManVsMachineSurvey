<script lang="ts">
    import UserFormInput from "../../userform/UserFormInput.svelte";
    import { convertAnswerDataToNumeric, convertAnswerDataToString } from "$lib/convertAnswerData";
    import { onMount } from "svelte";
    import { createEventDispatcher } from 'svelte';
    import surveyConfig from "../../../config/surveyConfig.json";
    
    const dispatch = createEventDispatcher();

    export let answerText: string;
    export let answerNum: number;
    export let answerID: number;
    
    let evaluations = {};
    let isExpanded: boolean = false;

    const handleFormUpdate = () => {
        if (Object.keys(evaluations).length === surveyConfig.evaluationCriteria.length) {
            dispatch("update", true);
        }

        let allFormAnswers = localStorage.getItem("allFormAnswers");

        if (!allFormAnswers) {
            const formAnswer = new Map();
            formAnswer.set(answerID, evaluations);
            localStorage.setItem('allFormAnswers', JSON.stringify(Array.from(formAnswer.entries())));
        } else {
            let existingFormAnswers = new Map(JSON.parse(allFormAnswers));
            existingFormAnswers.set(answerID, evaluations);
            localStorage.setItem('allFormAnswers', JSON.stringify(Array.from(existingFormAnswers.entries())));
        }
    }

    onMount(() => {
        let allFormAnswers = localStorage.getItem("allFormAnswers");
        if (allFormAnswers) {
            let existingFormAnswers = new Map(JSON.parse(allFormAnswers));
            let formAnswers = existingFormAnswers.get(answerID);
            if (formAnswers) {
                evaluations = formAnswers;
            }
        }
    });

    function getTranslation(term: string): string {
        const translations = {
            "Knowledge": "Kunnskap",
            "Empathy": "Empati",
            "Helpfulness": "Hjelpsomhet",
            "Show less": "Vis mindre",
            "Show full answer": "Vis resten av svaret",
            "Answer": "Svar"
        };
        return translations[term] || term;
    }
</script>

<div class="flex flex-col">
    <div class="flex flex-col gap-2 mb-6">
        <h1 class="text-xl text-primary font-bold text-center">{getTranslation("Answer")} {answerNum}:</h1>
        <div class="flex flex-col justify-between bg-secondary rounded-xl px-6 pt-6">     
            <div class={`text-sm ${!isExpanded && "max-h-48"} overflow-hidden`}>
                {answerText}
            </div>
            <button class="py-2 text-primary font-semibold" on:click={() => isExpanded = !isExpanded}>
                {isExpanded ? getTranslation("Show less") : getTranslation("Show full answer")}
            </button>
        </div>
    </div>
    <div class="flex flex-col justify-start items-center gap-6">
        {#each surveyConfig.evaluationCriteria as criteria}
            <UserFormInput 
                formData={evaluations[criteria]} 
                on:update={(e) => {evaluations[criteria] = convertAnswerDataToNumeric(e.detail); handleFormUpdate();}} 
                inputType="radio" 
                label={getTranslation(criteria)} 
                options={surveyConfig.evaluationOptions.map(getTranslation)}
            />
        {/each}
    </div>
</div>