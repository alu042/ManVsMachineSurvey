<script lang="ts">
    import { postBugData } from "../../api/postBugData";
    import { goto } from "$app/navigation";

    let bugText:string = ""

    const handleBugSubmit = async () => {
        if (bugText) {
            await postBugData(bugText)
            goto("javascript:history.back()")
        } else {
            console.log("error");
        }
    }
</script>

<div class="flex flex-col justify-between items-center h-full">
    <div class="flex flex-col gap-4 h-full mt-8">
        <div class="flex flex-col gap-4 px-96 items-center">
            <h1 class="text-3xl text-primary font-bold">Å nei!</h1>
            <p>Har du funnet feil eller mangler, eller er det noe du ønsker å melde i fra om angående spørreundersøkelsen? Skriv det i tekstfeltet under. (Dette er selvfølgelig helt anonymt også!)</p>
        </div>
        
        <div class="flex flex-col gap-4 justify-center items-center">
            <textarea bind:value={bugText} cols="30" rows="8" class="border-solid border-gray-400 border-2 p-3 md:text-l w-1/3" placeholder="Skriv feil her"></textarea>
            <button 
                class="text-primary hover:bg-primary hover:text-bg font-bold border-primary border-2 rounded-full px-8 py-1"
                on:click={handleBugSubmit}
            >
                Send feil
            </button>
        </div>
          
          
    </div>
    <div class="flex justify-center items-center gap-8 text-primary font-bold">
        <!-- svelte-ignore a11y-invalid-attribute -->
        <button><a href="javascript:history.back()">Tilbake</a></button>
    </div>
</div>