<script lang="ts">
    import UserFormInput from "../../components/userform/UserFormInput.svelte";
    import ArrowBack from "../../components/svg/ArrowBack.svelte";
    import ButtonComponent from "../../components/userform/inputs/ButtonComponent.svelte";
    import { postUserformData } from "../../api/postUserformData";
    import { getUserQuestions } from "../../api/getUserQuestions";
    import { goto } from "$app/navigation";


    let age: string = ""
    let education: string = ""
    let healthcare_personnel: string = ""
    let is_licensed: string = ""
    let gender: string = ""
    let has_answered_before: string = ""
    let county: string = ""

    let firstUserQuestion: number = 0

    const handleUserformSubmit = async (age: string, education: string, healthcare_personnel: string, gender: string, has_answered_before: string, county: string) => {
        localStorage.clear()

        const submitDate = new Date().toISOString()

        const response = await postUserformData(age, education, healthcare_personnel, is_licensed, gender, has_answered_before, county, submitDate)
        const userQuestions = await getUserQuestions(response.respondentID)
            
        goto("form/0")
    }
</script>

<div class="flex flex-col md:justify-center gap-6 md:gap-0 h-full md:h-screen">
    <a class="md:ml-32 mt-5" href="/">
        <ArrowBack width="2rem" />
    </a>
    <div class="flex flex-col md:flex-row h-full md:h-4/5 gap-6">
        <div class="flex flex-col justify-center w-full md:w-2/4 gap-4 md:px-32 ">
            <h1 class="text-3xl text-primary font-bold">Deltakerinformasjon</h1>
            <p>Til undersøkelsen ønsker vi opplysninger om din aldersgruppe, om du er student eller jobber innen helse, hva din høyeste fullførte utdanningsgrad er og om du er lege eller medisinstudent med lisens.</p>
            <p>Vi vil igjen minne om at denne undersøkelsen er helt anonym. Se <a href="/personvern" class="text-primary font-bold">Personvern</a> for mer informasjon.</p>  
        </div>
        <div class="flex flex-col gap-4 justify-start items-center md:w-2/4">
            <UserFormInput formData={age} on:update={(e) => age = e.detail} inputType="radio" label="Alder" options={["18-20", "21-30", "31-40", "41-50", "51-60", "61-70", "71+"]}/>

            <div class="flex flex-col md:flex-row justify-between md:justify-start items-start gap-3 md:gap-0 md:h-1/4 w-full">
                <div class="w-full md:w-3/12">
                    <p class="text-primary font-bold w-auto">Hvilket fylke befinner du deg i?</p>
                </div>
                <div class="w-full md:w-5/6">
                    <select class="md:ml-6 pl-2 pr-16 py-2 rounded-xl text-primary border-primary border-2 text-start hover:cursor-pointer" 
                            on:change={(e) => county = e.target.value}
                    >
                        {#each ["", "Vestland", "Rogaland", "Møre og Romsdal", "Oslo", "Viken", "Nordland", "Trøndelag", "Innlandet", "Troms og Finnmark", "Vestfold og Telemark", "Agder"] as data}
                            <option value={data}>{data}</option>
                        {/each}
                    </select>
                </div>
            </div>

            <div class="flex flex-col md:flex-row justify-between md:justify-start items-start gap-3 md:gap-0 md:h-1/4 w-full">
                <div class="w-full md:w-3/12">
                    <p class="text-primary font-bold w-auto">Høyeste fullførte utdanningsgrad:</p>
                </div>
                <div class="w-full md:w-5/6">
                    <select class="md:ml-6 pl-2 pr-16 py-2 rounded-xl text-primary border-primary border-2 text-start hover:cursor-pointer" 
                            on:change={(e) => education = e.target.value}
                    >
                        {#each ["", "VGS", "Bachelor", "Master", "Profesjonsstudium", "PhD"] as data}
                            <option value={data}>{data}</option>
                        {/each}
                    </select>
                </div>
            </div>
                

            <UserFormInput formData={healthcare_personnel} on:update={(e) => healthcare_personnel = e.detail} inputType="radio" label="Jobber/studerer du innen helse, eller har du bakgrunn fra helsevesenet?" options={["Ja", "Nei"]}/>
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
</style>
