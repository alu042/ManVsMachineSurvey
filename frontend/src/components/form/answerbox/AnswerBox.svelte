<script>
    import UserFormInput from "../../userform/UserFormInput.svelte";
</script>

<div class="flex flex-col">
    <div class="flex flex-col gap-2 mb-6">
        <h1 class="text-xl text-primary font-bold text-center">Svar 2:</h1>
        <div class="bg-secondary p-6 rounded-xl text-sm">
            Hei,
            Nye symptomer som utslett etter påbegynt skabbehandling er ikke uvanlig og betyr ikke nødvendigvis at behandlingen har feilet. Dette skyldes ofte kroppens reaksjon på skabbmiddene og behandlingen. Det er viktig å gjennomføre den andre kuren som planlagt. Vær oppmerksom på at kløe og utslett kan vedvare i opptil fire uker etter siste behandling, selv når behandlingen har vært effektiv. Hvis du opplever nye utslett mer enn fire uker etter behandlingens slutt, anbefales det å kontakte lege for en videre vurdering.
        </div>
    </div>
    <div class="flex flex-col justify-start items-center gap-6">
        <UserFormInput inputType="radio" label="Kunnskap" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput inputType="radio" label="Empati" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
        <UserFormInput inputType="radio" label="Hjelpsomhet" options={["Veldig dårlig", "Dårlig", "Nøytral", "Bra", "Veldig bra"]}/>
    </div>
</div>