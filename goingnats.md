# Going NATS

Hva er dette? Jo, en presentasjon om hva NATS er, hva NATS kan gjøre og hvordan vi kan benytte NATS.

Let's go..... .. . NATS!

---

# Hva er NATS

NATS er et meldingssystem. Det er robust, raskt og om du vil distribuert. 

NATS er et CNCF "incubating project" og har vært det siden 2018.


## Teknisk

- Skrevet i Go
- Klientbiblioteker for 40+ språk.

## Ressurser
- https://nats.io
- https://github.com/nats.io/
- cncf link

---

# Så, hva kan vi benytte NATS til?

 Vi kan benytte NATS for å koble sammen distribuerte systemer og enheter samt forenkle måten de kan kommuniserer på. 


NATS er kommunikasjonsmediumet for prosesser vi har et *tillitsforhold* til.

Egenskapene til NATS gjør det mulig å benytte NATS som et "service" og data mesh". 

---

# Fett, men jeg trenger noe mer konkret!

Skjønner, la oss komme i gang.

```
nats-server --jetstream
```

---

# Subjects

I NATS kommuniserer vi over *subjects*. Dette gir en navnebasert addressering i motsetning til de ulike ip, port og path baserte endepunktene vi vanligvis er nødt til å forholde oss til. 

Eksempel:
- "ordre" er her et subject vi kan benytte for å sende bestillinger.

---

# Hvordan bruker jeg så et subject?

På et subject, i dette tilfelle "ordre", kan vi publisere eller lytte (pub eller sub).

```bash
for i in $(seq 10)
do
    nats pub ordre "{ordreid:${i}"
done
echo "Done..."
```

For et antiklimaks! Men, hva skjedde egentlig der?


---

# Hvordan lytte på et subject?

```bash
for i in $(seq 10)
do
    nats pub ordre "{ordreid:${i}}"
done
echo "Done..."
```

La oss åpne en ny terminal hvor vi også lytter, før vi kjører bash scriptet over.

```
nats sub ordre 
```
---

# Publish / Subscribe 

Den grunnleggende måten å "kommunisere" på i NATS er altså "publish/subscribe."

I utgangspunktet er *subjects* i NATS "ephemeral". De eksisterer så lenge noen publiserer og noen lytter. 

Er det ingen som lytter går meldingen ut i intet.  

---

# Men vent, vi kan gjøre mer med et subject

Et subject er ikke bare en flat struktur i NATS kan det være hierarkisk.

Vi kan utvide "chat" subject benyttet tidligere med meningsfyllt struktur:

- chat
- chat.$room
- chat.$user.dm
- chat.$user.poke

Basert på disse har vi tilført mening og dynamiske subjects. Vår banale "chat klient" har nå ulike subjects å lytte på.

---

# Det var jo litt kult, men hva med...

Vi kan jo ikke være helt "cowboy", vi har *krav!*.

- Vi trenger koordinering av meldinger!
- Hos oss trenger vi persistering!
- Og... 
- Samt...

Slapp av, vi kommer til det.

---

# Kø grupper

I pub/sub blir meldinger levert som *1:N*. Det vil si at alle som lytter vil får meldingen som blir publisert.

Ved å introdusere en "queue group" vil du fortsatt få meldinger etter *1:N* prinsippet, men du har nå muligheten til å koordinere meldinger for alle som lytter med samme kø navn.

Vi får med andre ord en lastbalansering av meldinger og vi kan sørge for at vi bare konsumerer meldingen en gang for et formål.

---

# Run go hello world

```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world!")
}
```

---

# buggy 

becasue of code above?

---

# mermaid

```mermaid
~~~mermaid-ascii -
graph LR
A --> B & C
B --> C & D
D --> C
~~~
```

```
~~~cat
hello world
~~~
```
