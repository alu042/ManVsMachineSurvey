<script lang="ts">
    import FormHeader from "../../../components/form/header/FormHeader.svelte";
    import AnswerBox from "../../../components/form/answerbox/AnswerBox.svelte";
	import Footer from "../../../components/form/footer/Footer.svelte";
    import { onMount } from "svelte"
    import { browser } from "$app/environment";

    export let data;

    let formQuestion: string = ""
    let questionAnswer1Text: string = ""
    let questionAnswer1ID: number = 0
    let questionAnswer2Text: string = ""
    let questionAnswer2ID: number = 0
    let questionNumber: number = 0

    let question1Answered: boolean = false
    let question2Answered: boolean = false
    
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
            questionAnswer1Text = questions[questionNumber].Answers[0].AnswerText;
            questionAnswer1ID = questions[questionNumber].Answers[0].AnswerID;
            questionAnswer2Text = questions[questionNumber].Answers[1].AnswerText;
            questionAnswer2ID = questions[questionNumber].Answers[1].AnswerID;
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
        {#key questionNumber}     
            <AnswerBox on:update={(e) => question1Answered = e.detail} answerNum={1} answerText={questionAnswer1Text} answerID={questionAnswer1ID}/>
            <AnswerBox on:update={(e) => question2Answered = e.detail} answerNum={2} answerText={questionAnswer2Text} answerID={questionAnswer2ID}/>
        {/key}
    </div>
    {#key questionNumber}
        <Footer answeredAll={question1Answered && question2Answered} questionNum={Number(questionNumber)}/>
    {/key}
</div>