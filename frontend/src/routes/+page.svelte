<script>
	import ButtonComponent from "../components/userform/inputs/ButtonComponent.svelte";
</script>
<div class="flex flex-col justify-between items-center h-full md:h-screen w-full">
    <div class="flex flex-col md:flex-row h-full w-full items-center">
        <div class="flex flex-col w-11/12 md:w-3/5 gap-2 md:gap-4">
            <h1 class="text-5xl md:text-7xl text-primary font-bold pb-2">Helseveileder</h1>
            <p>Vi utfører en undersøkelse om kunstig intelligens (AI) i medisin. Vi ønsker å forstå hvordan AI kan fungere som en copilot for leger og pasienter.</p>
            <p>Undersøkelsen er helt anonym.</p>
            <ul>
                <p class="mb-4">I denne undersøkelsen vil vi at du skal vurdere empati, kunnskap og hjelpsomhet i ulike svar på medisinske spørsmål.</p>
                <li><strong>Empati:</strong> Evnen til å oppdage og erkjenne andres følelser og tanker.</li>
                <li><strong>Kunnskap:</strong> Nøyaktighet og relevans av den medisinske informasjonen gitt.</li>
                <li><strong>Hjelpsomhet:</strong> Evnen til å gi nyttig og praktisk råd eller støtte.</li>
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