<script lang="ts">
    import { postEvaluationData } from "../../api/postEvaluationData";
    import { goto } from "$app/navigation";

    let evaluationText:string = ""
    let submittedEval: boolean = false

    const handleEvaluationSubmit = async () => {
        if (evaluationText) {
            await postEvaluationData(evaluationText)
            submittedEval = true
        } else {
            console.log("error");
        }
    }
</script>

<div class="flex flex-col justify-between items-center h-full">
    <div class="flex flex-col gap-14 md:gap-8 h-full mt-8">
        <div class="flex flex-col gap-4 md:px-96 items-center">
            <h1 class="text-3xl text-primary font-bold">Takk for at du tok deg tid!</h1>
            <p>Tusen takk for hjelpen i denne undersøkelsen! Vi setter stor pris på det og håper du får en fin dag videre.</p>
            <p>Har du tid så setter vi veldig pris på om du skrevet en liten tilbakemelding til oss i tekstfeltet under:</p>
        </div>
        
        {#if !submittedEval}
            <div class="flex flex-col gap-12 md:gap-3 justify-center items-center">
                <textarea bind:value={evaluationText} cols="30" rows="8" class="border-solid border-gray-400 border-2 p-3 md:text-l md:w-1/3 rounded-md" placeholder="Skriv evaluering her"></textarea>
                <button 
                    class="text-primary hover:bg-primary hover:text-bg font-bold border-primary border-2 rounded-full px-8 py-1"
                    on:click={handleEvaluationSubmit}
                >
                    Send evaluering
                </button>
            </div>
        {:else}
            <div class="flex justify-center text-center">
                <h1 class="text-xl text-primary font-semibold">Takk for tilbakemeldingen!</h1>
            </div>
        {/if}
        <div class="flex flex-col justify-center items-center md:mt-8">
            <button class="text-bg w-auto bg-primary hover:bg-bg hover:text-primary font-bold border-primary border-2 rounded-xl px-12 py-3" on:click={() => goto("/")}>
                Returner til hjemmesiden
            </button>
        </div>
    </div>
</div>