<script>
    import UserFormInput from "../../components/userform/UserFormInput.svelte";
	import ButtonComponent from "../../components/ButtonComponent.svelte";
    import ArrowChevronLeft from "../../components/svg/ArrowChevronLeft.svelte";
</script>

<div class="flex flex-col justify-center gap-4 h-full">
    <div class="flex justify-center items-center mx-auto w-1/2 gap-4">
        <h1 class="text-xl text-primary font-bold text-center">Spørsmål 3</h1>
        <div class="border-r-2 border-primary h-24"></div>
        <p>Hvis man begynner skabbehandling før man har fått symptomer, men får symptomer iløpet av den uka før man skal smøre seg igjen etter 7 dager, har da ikke første smøringen virket?</p>
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
    <div class="flex justify-center items-center gap-8 text-primary font-bold">
        <a href="">
            <button><ArrowChevronLeft width=16 />
                Forrige spørsmål
            </button>
        </a>
        <ButtonComponent text="Send inn svar" url="/" filled={true} />
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
        font-size: medium;
    }
</style>