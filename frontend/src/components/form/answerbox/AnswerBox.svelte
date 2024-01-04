<script lang="ts">
    import UserFormInput from "../../userform/UserFormInput.svelte";
    
    export let answerText:string

</script>

<div class="flex flex-col">
    <div class="flex flex-col gap-2 mb-6">
        <h1 class="text-xl text-primary font-bold text-center">Svar 2:</h1>
        <div class="bg-secondary p-6 rounded-xl text-sm">
            {answerText}
        </div>
    </div>
    <div class="flex flex-col justify-start items-center gap-6">
        <UserFormInput inputType="radio" label="Kunnskap" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput inputType="radio" label="Empati" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput inputType="radio" label="Hjelpsomhet" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
    </div>
</div>