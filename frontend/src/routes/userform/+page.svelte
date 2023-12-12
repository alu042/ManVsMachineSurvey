<script>
    import UserFormInput from "../../components/userform/UserFormInput.svelte";
    import ArrowBack from "../../components/svg/ArrowBack.svelte";
</script>

<div class="flex flex-col justify-center gap-20 h-full ">
    <a class="ml-32" href="/">
        <ArrowBack width="2rem" />
    </a>
    <div class="flex h-4/5 ">
        <div class="flex flex-col justify-center w-2/4 gap-4 px-32 ">
            <h1 class="text-3xl text-primary font-bold">Medisinsk AI Samtale Undersøkelse</h1>
            <p>For å kunne få strukturert undersøkelsen så må det samles inn litt informasjon.
            <p>Her trengs det alder, utdanningsgrad, om du er helsepersonell og kjønn.</p>
            <p>Vi vil igjen nevne at denne undersøkelsen er helt anonym og informasjonen vil kun bli brukt til å forbedre vår tjeneste.</p>  
        </div>
        <div class="flex flex-col justify-start items-center w-2/4">
            <UserFormInput inputType="radio" label="Alder" options={["18-20", "20-30", "30-40", "40-50", "50-60"]}/>
            <UserFormInput inputType="select" label="Utdanningsgrad" options={["PhD", "Master", "Fagbrev", "Bachelor"]}/>
            <UserFormInput inputType="radio" label="Helsepersonell" options={["Ja", "Nei"]}/>
            <UserFormInput inputType="radio" label="Kjønn" options={["Mann", "Kvinne", "Annet"]}/>
        </div>
    </div>
    <div class="flex justify-center items-center gap-8 text-primary font-bold">
        <button class="font-bold uppercase border-primary border-2 rounded-xl px-8 py-3 bg-primary text-bg text-xs"><a href="/form">Start undersøkelse</a></button>
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
        font-size: large;
    }
</style>