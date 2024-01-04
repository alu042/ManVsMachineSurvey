<script lang="ts">
    import FormHeader from "../../../components/form/header/FormHeader.svelte";
    import AnswerBox from "../../../components/form/answerbox/AnswerBox.svelte";
	import Footer from "../../../components/form/footer/Footer.svelte";

    import { onMount } from "svelte"

    export let data;

    let formQuestion: string = ""
    let questionAnswer1: string = ""
    let questionAnswer2: string = ""

    onMount(async () => {
        const questionNumber = data.slug
        let localstoragequestions = localStorage.getItem("userQuestions");
	    if (localstoragequestions) {
            let questions = JSON.parse(localstoragequestions).questions
            console.log(questions)
            formQuestion = questions[questionNumber].Question.QuestionText
            questionAnswer1 = questions[questionNumber].Answers[0].AnswerText
            questionAnswer2 = questions[questionNumber].Answers[1].AnswerText
            console.log(formQuestion)
            console.log(questionAnswer1)
            console.log(questionAnswer2)
        };  
    })
</script>

<div class="flex flex-col h-full">
    <FormHeader formQuestion={formQuestion}/>
    <div class="flex h-full justify-between gap-12">
        <AnswerBox answerText={questionAnswer1}/>
        <AnswerBox answerText={questionAnswer2}/>
    </div>
    <Footer />
</div>

<style lang="postcss">
</style>