<h1> Systemarkitektur </h1>

Programmet lytter til port :8080, og kjøres lokalt i en nettleser. Det blir hentet informasjon fra en databaseserver, basert på det brukeren ønsker å finne. 

<b> Kort introduksjon til programvaren: </b>
- Vi har en main funksjon som kobler programmet opp med port :8080, og kjører funksjonene “Open” og “Web” på forskjellige stier. 
- Vi har en type som heter “Poke”, som lager en struct basert på API hentet fra nettet. 
- Funksjonen “Open” får opp en nettside med inputfelt når du starter programmet, og scanner hva du skriver her. Vi kaller på HTML-filen “Start.html” i direktoratet “Template”. 
- Funksjonen “Web” tar brukeren videre til en ny lokal nettside. Vi bruker variabelen “URL” til å få API-en vår. Det brukeren har skrevet som input (som er navn eller ID-nummer til en pokémon) vil bli lagt på i slutten av URL-en, som da spesifiserer hvor informasjonen skal hentes fra. Videre unmarshler vi json-data, og gjennom HTML-filen “Dex.html” sendes dataen til brukerens nettleser. 
- Funksjonen “getResp” sjekker hvilken generasjon pokémonen hører til ved hjelp av ID-nummeret, og printer en melding om den generasjonen. Her sjekker den også om det ikke har kommet opp noe pokémon, og printer da en feilmelding.
- Vi har også laget en egen fil kalt template_test.go, som tester om templatene er tilgjengelig. 
