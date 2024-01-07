<script lang="ts">
    import UserFormInput from "../../userform/UserFormInput.svelte";
    import { convertAnswerData } from "$lib/convertAnswerData";
	import { onMount } from "svelte";

    export let answerText:string
    export let answerNum:number
    let answeredAll:boolean
    
    let knowledge:string
    let empathy:string
    let helpfulness:string

    const handleFormUpdate = (category:string, value:string) => {
        if (knowledge != "" && empathy != "" && helpfulness != "") {
            answeredAll = true
        }

        let numericValue = convertAnswerData(value)
        switch (category) {
            case "knowledge":
                knowledge = numericValue
        }

        return ""
    }

    onMount(() => {
        knowledge=""
        empathy=""
        helpfulness=""
    })

</script>

<div class="flex flex-col">
    <div class="flex flex-col gap-2 mb-6">
        <h1 class="text-xl text-primary font-bold text-center">Svar {answerNum}:</h1>
        <div class="bg-secondary p-6 rounded-xl text-sm">
            {answerText}
        </div>
    </div>
    <div class="flex flex-col justify-start items-center gap-6">
        <UserFormInput formData={knowledge} on:update={(e) => knowledge = e.detail} inputType="radio" label="Kunnskap" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput formData={empathy} on:update={(e) => empathy = e.detail} inputType="radio" label="Empati" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput formData={helpfulness} on:update={(e) => helpfulness = e.detail} inputType="radio" label="Hjelpsomhet" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
    </div>
</div>