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
    let is_licensed: string = "Nei"
    let gender: string = "Mann"
    let has_answered_before: string = "Nei"
    let county: string = "Vestland"

    let firstUserQuestion: number = 0

    const handleUserformSubmit = async (age: string, education: string, healthcare_personnel: string, gender: string, has_answered_before: string, county: string) => {
        localStorage.clear()

        const submitDate = new Date().toISOString()

        const response = await postUserformData(age, education, healthcare_personnel, is_licensed, gender, has_answered_before, county, submitDate)
        const userQuestions = await getUserQuestions(response.respondentID)
            
        goto("form/0")
    }
</script>

<div class="flex flex-col md:justify-center gap-6 md:gap-10 h-full md:h-screen">
    <a class="md:ml-32 mt-5" href="/">
        <ArrowBack width="2rem" />
    </a>
    <div class="flex flex-col md:flex-row h-full md:h-4/5 gap-6">
        <div class="flex flex-col justify-center w-full md:w-2/4 gap-4 md:px-32 ">
            <h1 class="text-3xl text-primary font-bold">Deltakerinformasjon</h1>
            <p>Til undersøkelsen ønsker vi opplysninger om din aldersgruppe, om du er student eller jobber innen helse, hva din høyeste fullførte utdanningsgrad er og om du er lege eller medisinstudent med lisens.</p>
            <p>Vi vil igjen minne om at denne undersøkelsen er helt anonym. Se <a href="/personvern"class="text-primary font-bold">Personvern</a> for mer informasjon.</p>  
        </div>
        <div class="flex flex-col gap-6 justify-start items-center md:w-2/4">
            <UserFormInput formData={age} on:update={(e) => age = e.detail} inputType="radio" label="Alder" options={["18-20", "20-30", "30-40", "40-50", "50-60"]}/>
            <UserFormInput formData={county} on:update={(e) => county = e.detail} inputType="select" label="Fylke" options={["Vestland", "Rogaland", "Møre og Romsdal", "Oslo", "Viken", "Nordland", "Trøndelag", "Innlandet", "Troms og Finnmark", "Vestfold og Telemark", "Agder"]}/>
            <UserFormInput formData={education} on:update={(e) => education = e.detail} inputType="select" label="Utdanningsgrad" options={["VGS", "Bachelor", "Master", "PhD"]}/>
            <UserFormInput formData={healthcare_personnel} on:update={(e) => healthcare_personnel = e.detail} inputType="radio" label="Jobber/studerer du innen helse?" options={["Ja", "Nei"]}/>
            <UserFormInput formData={is_licensed} on:update={(e) => is_licensed = e.detail} inputType="radio" label="Er du lege eller medisinstudent med lisens?" options={["Ja", "Nei"]}/>
            <UserFormInput formData={gender} on:update={(e) => gender = e.detail} inputType="radio" label="Kjønn" options={["Mann", "Kvinne", "Annet"]}/>
            <UserFormInput formData={has_answered_before} on:update={(e) => has_answered_before = e.detail} inputType="radio" label="Jeg har svart på denne undersøkelsen tidligere" options={["Ja", "Nei"]}/>
        </div>
    </div>
    <div class="flex justify-center items-center gap-8 text-primary font-bold py-10 md:pb-4">
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