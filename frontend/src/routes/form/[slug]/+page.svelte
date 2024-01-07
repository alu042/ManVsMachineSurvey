<script lang="ts">
    import FormHeader from "../../../components/form/header/FormHeader.svelte";
    import AnswerBox from "../../../components/form/answerbox/AnswerBox.svelte";
	import Footer from "../../../components/form/footer/Footer.svelte";
    import { onMount } from "svelte"
    import { browser } from "$app/environment";

    export let data;

    let formQuestion: string = ""
    let questionAnswer1: string = ""
    let questionAnswer2: string = ""
    let questionNumber: number = 0

    // Reactive statement to react on 'data.slug' changes
    $: if (data && data.slug !== undefined && browser) {
        questionNumber = data.slug;
        updateQuestionData();
    }

    // Refactored data updating logic into a function
    function updateQuestionData() {
        let localstoragequestions = localStorage.getItem("userQuestions");
        if (localstoragequestions) {
            let questions = JSON.parse(localstoragequestions).questions;
            formQuestion = questions[questionNumber].Question.QuestionText;
            questionAnswer1 = questions[questionNumber].Answers[0].AnswerText;
            questionAnswer2 = questions[questionNumber].Answers[1].AnswerText;
        }
    }

    // On mount, call the update function
    onMount(() => {
        updateQuestionData();
    });
</script>

<div class="flex flex-col h-full">
    <FormHeader questionNum={questionNumber} formQuestion={formQuestion}/>
    <div class="flex h-full justify-between gap-12">
        <AnswerBox answerNum={1} answerText={questionAnswer1}/>
        <AnswerBox answerNum={2} answerText={questionAnswer2}/>
    </div>
    <Footer questionNum={Number(questionNumber)}/>
</div>