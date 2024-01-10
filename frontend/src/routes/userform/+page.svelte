<script lang="ts">
    import UserFormInput from "../../components/userform/UserFormInput.svelte";
    import ArrowBack from "../../components/svg/ArrowBack.svelte";
	import ButtonComponent from "../../components/userform/inputs/ButtonComponent.svelte";
	import { postUserformData } from "../../api/postUserformData";
	import { getUserQuestions } from "../../api/getUserQuestions";
    import { goto } from "$app/navigation";


    let age: string = "18-20"
    let education: string = "PhD"
    let healthcare_personnel: string = "Ja"
    let gender: string = "Mann"
    let has_answered_before: string = "Nei"
    let county: string = "Vestland"

    let firstUserQuestion: number = 0

    const handleUserformSubmit = async (age: string, education: string, healthcare_personnel: string, gender: string, has_answered_before: string, county: string) => {
        localStorage.clear()

        const submitDate = new Date().toISOString()

        const response = await postUserformData(age, education, healthcare_personnel, gender, has_answered_before, county, submitDate)
        const userQuestions = await getUserQuestions(response.respondentID)
            
        goto("form/0")
    }
</script>

<div class="flex flex-col justify-center gap-20 h-full ">
    <a class="ml-32" href="/">
        <ArrowBack width="2rem" />
    </a>
    <div class="flex h-4/5 ">
        <div class="flex flex-col justify-center w-2/4 gap-4 px-32 ">
            <h1 class="text-3xl text-primary font-bold">Informasjon om deltaker</h1>
            <p>Til undersøkelsen trenger vi opplysninger om din aldersgruppe, utdanningsgrad, kjønn og bekreftelse på at du er helsepersonell.</p>
            <p>Vi vil igjen minne om at denne undersøkelsen er helt anonym. Se <a href="/personvern"class="text-primary font-bold">Personvern</a> for mer informasjon.</p>  
        </div>
        <div class="flex flex-col justify-start items-center w-2/4">
            <UserFormInput formData={age} on:update={(e) => age = e.detail} inputType="radio" label="Alder" options={["18-20", "20-30", "30-40", "40-50", "50-60"]}/>
            <UserFormInput formData={education} on:update={(e) => education = e.detail} inputType="select" label="Utdanningsgrad" options={["PhD", "Master", "Fagbrev", "Bachelor"]}/>
            <UserFormInput formData={county} on:update={(e) => county = e.detail} inputType="select" label="Fylke" options={["Vestland", "Rogaland", "Møre og Romsdal", "Oslo", "Viken", "Nordland", "Trøndelag", "Innlandet", "Troms og Finnmark", "Vestfold og Telemark", "Agder"]}/>
            <UserFormInput formData={healthcare_personnel} on:update={(e) => healthcare_personnel = e.detail} inputType="radio" label="Helsepersonell" options={["Ja", "Nei"]}/>
            <UserFormInput formData={gender} on:update={(e) => gender = e.detail} inputType="radio" label="Kjønn" options={["Mann", "Kvinne", "Annet"]}/>
            <UserFormInput formData={has_answered_before} on:update={(e) => has_answered_before = e.detail} inputType="radio" label="Jeg har svart på denne undersøkelsen tidligere" options={["Ja", "Nei"]}/>
        </div>
    </div>
    <div class="flex justify-center items-center gap-8 text-primary font-bold">
        <ButtonComponent text="Start undersøkelse" filled={true} onclick={() => handleUserformSubmit(age, education, healthcare_personnel, gender, has_answered_before, county)} />
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
    p {
        font-size: larger;
    }
</style>