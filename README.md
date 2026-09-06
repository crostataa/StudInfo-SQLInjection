## Introduzione
Progetto per il corso Sicurezza appartenente alla triennale in Informatica de La Sapienza Università di Roma
## Obiettivo
creazione di un sito vulnerabile ad attacchi SQL Injection; successivamente, bisogna rendere il sito sicuro
## To-do
questa è la beta della beta della beta della beta....
le cose da fare sono tante:
### teoria / studio
- cos'è un SQL Injection
- quali sono gli elementi del codice che rendono possibile l'SQL Injection, e quali sono quelli che permettono di proteggere il sistema da questo attacco
- studio **approfondito** del codice !
### progetto
- creazione dell'interfaccia con **vue.js** (adesso è scritta in html puro all'interno del file _main.go_
- creazione di una pagina di login
  - creazione di un database di utenti (DA SALVARE ANCHE SU FILE DI TESTO perché verrà modificato di continuo durante gli attacchi)
  - qui verranno fatti gli attacchi per _entrare_ nel sistema senza credenziali
- creazione di un file di testo con tutti gli script in SQL per attaccare
- creazione di una pagina principale CHE ABBIA UNA BARRA DI RICERCA
  - qui verranno fatti gli attacchi per _modificare_ i database
- (opzionale) implementare il template decaffeinated coffee per portarsi avanti con wasa
- (opzionale) fare l'interfaccia del sito in modo che ricordi infostud (così, sembra divertente)
- to be continued....

### notion
sto tenendo nota dei progressi riguardanti il progetto sul mio notion:
https://app.notion.com/p/SICUREZZA-3110dd71b5c6807890f8e764e39a1c86?source=copy_link

### attacchi possibili:

- utente legittimo:<br>
  `admin`<br>
  dovrebbe apparire una normale tabella con i dati dell'utente digitato (cioè admin)<br>
- error-based SQL Injection<br>
  `'`<br>
  la tabella sparisce e dovrebbe apparire il box rosso: vuol dire che il database è andato in 'syntax   error',<br>
  il che ci conferma che l'input non è sanitizzato contro attacchi di questo tipo!<br>
- tautologia<br>
  `' OR 1=1 #`<br>
  il database riceve la stringa spezzata, valuta la condizione 1=1 (cioè sempre vera) e ignora il<br>
  resto della query grazie a #. L'output è la visione a schermo di una tabella con tutti i dati degli<br>
  utenti presenti nel database<br>
- UNION Injection<br>
  `' UNION SELECT 999, chiave, valore_progetto FROM segreti_aziendali #`<br>
  dal momento che le tabelle che sto unendo devono avere lo stesso
  
