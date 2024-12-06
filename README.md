# Going NATS

Dette er en presentasjon om:

- Hvorfor Mattilsynet benytter NATS
- Litt om hva NATS er
- Litt om hva vi bruker NATS til.
- Litt teknisk om hvordan man benytter systemet.


## Historikk

Presentasjon ble første gang holdt på en meet up i Hamar i regi av Hamar Digirama: https://www.meetup.com/hamar-digirama/

- https://www.meetup.com/hamar-digirama/events/304396099/?eventOrigin=group_past_events 

## Om "slide-deck"

Gikk litt utenom vanlig verktøykasse her og benyttet en program som heter `slides`: https://github.com/maaslalani/slides

Dette finnes i homebrew og kan installeres på denne måten:

````bash
brew install slides
``````

eventuelt...

```bash
go install github.com/maaslalani/slides@latest
```

Presentasjonen benytter også planuml for illustrasjoner så dette må også installers om du ønsker å titte på noe annet enn markdown direkte på github.

```bash
brew install plantuml
```

## Se på slides

```bash
slides goingnats.md
```

## Eksekvering av codeblocks

Du kan benytte `ctrl+e`for å eksekvere codeblocks som finnes i illustrasjonen.

Da trenger du også go og nats-server som kan installeres på følgende måte:

```bash
brew install golang
brew install nats-server
go mod tidy # på topnivå for å laste ned nødvendig go kode.
```

> Dette er bare nødvendig om du vil eksekvere eksemplene
