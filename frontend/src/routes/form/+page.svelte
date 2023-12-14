<script>
    import UserFormInput from "../../components/userform/UserFormInput.svelte";
	import ButtonComponent from "../../components/ButtonComponent.svelte";
    import ArrowChevron from "../../components/svg/ArrowChevron.svelte";
	import CircleExclamation from "../../components/svg/CircleExclamation.svelte";
</script>

<div class="flex flex-col justify-center gap-4 h-full">
    <div class="flex justify-center items-center">
        <div class="flex justify-center items-center w-1/2 gap-4">
            <h1 class="text-xl text-primary font-bold text-center">Spørsmål 3</h1>
            <div class="border-r-2 border-primary h-24"></div>
            <p>Hvis man begynner skabbehandling før man har fått symptomer, men får symptomer iløpet av den uka før man skal smøre seg igjen etter 7 dager, har da ikke første smøringen virket?</p>
        </div>
        <div class="flex justify-center items-center">
            <a class="-" href="/errorreport">
                <CircleExclamation width="24" height="24" />    
            </a>
        </div>
    </div>
    <div class="flex justify-center gap-12">
        <div class="flex flex-col gap-2">
            <h1 class="text-xl text-primary font-bold text-center">Svar 1:</h1>
            <div class="bg-secondary p-4 rounded-xl">
                <p>
                    Hei,
                    Nye symptomer som utslett etter påbegynt skabbehandling er ikke uvanlig og betyr ikke nødvendigvis at behandlingen har feilet. Dette skyldes ofte kroppens reaksjon på skabbmiddene og behandlingen. Det er viktig å gjennomføre den andre kuren som planlagt. Vær oppmerksom på at kløe og utslett kan vedvare i opptil fire uker etter siste behandling, selv når behandlingen har vært effektiv. Hvis du opplever nye utslett mer enn fire uker etter behandlingens slutt, anbefales det å kontakte lege for en videre vurdering.
                </p>
            </div>
        </div>
        <div class="flex flex-col gap-2">
            <h1 class="text-xl text-primary font-bold text-center">Svar 2:</h1>
            <div class="bg-secondary p-4 rounded-xl">
                <p>
                    Hei,
                    Nye symptomer som utslett etter påbegynt skabbehandling er ikke uvanlig og betyr ikke nødvendigvis at behandlingen har feilet. Dette skyldes ofte kroppens reaksjon på skabbmiddene og behandlingen. Det er viktig å gjennomføre den andre kuren som planlagt. Vær oppmerksom på at kløe og utslett kan vedvare i opptil fire uker etter siste behandling, selv når behandlingen har vært effektiv. Hvis du opplever nye utslett mer enn fire uker etter behandlingens slutt, anbefales det å kontakte lege for en videre vurdering.
                </p>
            </div>
        </div>
    </div>
    <div class="flex justify-center h-4/5 gap-12">
        <div class="flex flex-col justify-start items-center gap-4">
            <UserFormInput inputType="radio" label="Kunnskap" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
            <UserFormInput inputType="radio" label="Empati" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
            <UserFormInput inputType="radio" label="Hjelpsomhet" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        </div>
        <div class="flex flex-col justify-start items-center gap-4">
            <UserFormInput inputType="radio" label="Kunnskap" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
            <UserFormInput inputType="radio" label="Empati" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
            <UserFormInput inputType="radio" label="Hjelpsomhet" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        </div>
    </div>
    <div class="flex justify-center items-center gap-8">
        <a href="/">
            <button class="flex items-center gap-2 text-primary font-semibold">
                <ArrowChevron width=16 direction="left"/>
                Forrige spørsmål
            </button>
        </a>
        <ButtonComponent text="Send inn svar" url="/" filled={true} />
        <a href="/">
            <button class="flex items-center gap-2 text-primary font-semibold opacity-50">
                Neste spørsmål
                <ArrowChevron width=16 direction="right"/>
            </button>
        </a>
    </div>
</div>

<style lang="postcss">
    :root {
        color: theme(colors.content);
        
    }
    p {
        font-size: medium;
    }
</style>