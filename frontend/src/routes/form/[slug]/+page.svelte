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
            if (Math.random() >= 0.5) { 
                questionAnswer1Text = questions[questionNumber].Answers[0].AnswerText;
                questionAnswer1ID = questions[questionNumber].Answers[0].AnswerID;
                questionAnswer2Text = questions[questionNumber].Answers[1].AnswerText;
                questionAnswer2ID = questions[questionNumber].Answers[1].AnswerID;
            } else {
                questionAnswer1Text = questions[questionNumber].Answers[1].AnswerText;
                questionAnswer1ID = questions[questionNumber].Answers[1].AnswerID;
                questionAnswer2Text = questions[questionNumber].Answers[0].AnswerText;
                questionAnswer2ID = questions[questionNumber].Answers[0].AnswerID;
            }
        }

        let allFormAnswers:string | null = localStorage.getItem("allFormAnswers")

        if (allFormAnswers) {
            let existingFormAnswers = JSON.parse(allFormAnswers);
            let retrievedFormAnswers:Map<number, number[]> = new Map(existingFormAnswers)
            let formAnswer1 = retrievedFormAnswers.get(questionAnswer1ID)
            let formAnswer2 = retrievedFormAnswers.get(questionAnswer2ID)
            if (formAnswer1 && formAnswer2) {
                if (formAnswer1.includes(0) || formAnswer2.includes(0)) {
                    question1Answered = false
                    question2Answered = false
                } else {
                    question1Answered = true
                    question2Answered = true
                }
            } else {
                question1Answered = false
                question2Answered = false
            }
        } 
    }

    // On mount, call the update function
    onMount(() => {
        updateQuestionData();
    });
</script>

<div class="flex flex-col h-full md:h-screen gap-10">
    <FormHeader questionNum={questionNumber} formQuestion={formQuestion}/>
    <div class="flex flex-col md:flex-row h-full justify-between gap-12">
        {#key questionNumber}
            <div class="flex-1">
                <AnswerBox on:update={(e) => question1Answered = e.detail} answerNum={1} answerText={questionAnswer1Text} answerID={questionAnswer1ID}/>
            </div>
            <div class="flex-1">
                <AnswerBox on:update={(e) => question2Answered = e.detail} answerNum={2} answerText={questionAnswer2Text} answerID={questionAnswer2ID}/>
            </div>
        {/key}
    </div>
    {#key questionNumber}
        <Footer answeredAll={question1Answered && question2Answered} questionNum={Number(questionNumber)}/>
    {/key}
</div>