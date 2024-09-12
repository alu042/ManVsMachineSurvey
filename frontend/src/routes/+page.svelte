<script>
	import ButtonComponent from "../components/userform/inputs/ButtonComponent.svelte";
	import surveyConfig from "../config/surveyConfig.json";

</script>

<div class="flex flex-col justify-between items-center h-full md:h-screen w-full">
    <div class="flex flex-col md:flex-row h-full w-full items-center">
        <div class="flex flex-col w-11/12 md:w-3/5 gap-2 md:gap-4">
            <h1 class="text-5xl md:text-7xl text-primary font-bold pb-2">{surveyConfig.surveyTitle}</h1>
            <p>{surveyConfig.introduction}</p>
            <p>{surveyConfig.surveyDescription}</p>
            <p>{surveyConfig.surveyContext}</p>
            <p>{surveyConfig.surveyInstructions}</p>
            <ul>
                {#each surveyConfig.evaluationCriteria as criteria}
                    <li><strong>{criteria.name}:</strong> {criteria.description}</li>
                {/each}
            </ul>               
        </div>
        <div class="flex justify-center items-center md:w-2/5 mt-8 md:mt-0">
            <a href="/userform">
                <ButtonComponent text="Start undersøkelse" filled={false} />
            </a>
        </div>
    </div>
    <div class="flex justify-center items-center gap-20 md:gap-20 text-primary font-bold mt-10 md:mb-3">
        <button class="w-20"><a href="/personvern">Personvern</a></button>
        <button class="w-20"><a href="/info">Mer info</a></button>
    </div>
</div>

<style lang="postcss">
    :root {
        color: theme(colors.content);
    }
    li {
        list-style-type: disc;
        margin-left: 2rem;
        font-size: large;
    }
    p, li{
        font-size: larger;
    }
    button {
        font-size: large;
    }
</style>