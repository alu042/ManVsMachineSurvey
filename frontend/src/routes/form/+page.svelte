<script>
    import FormHeader from "../../components/form/header/FormHeader.svelte";
    import AnswerBox from "../../components/form/answerbox/AnswerBox.svelte";
	import Footer from "../../components/form/footer/Footer.svelte";
</script>

<div class="flex flex-col h-full">
    <FormHeader />
    <div class="flex h-full justify-between gap-12">
        <AnswerBox />
        <AnswerBox />
    </div>
    <Footer />
</div>

<style lang="postcss">
    :root {
        color: theme(colors.content);
        
    }
    p {
        font-size: medium;
    }
</style>