package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//------------- CODICE VULNERABILE -------------//
import (
	"fmt"
	_ "fmt"
	"sicurezza/service/api/reqcontext"
)

// searchUsers è la funzione che risponderà quando qualcuno visita /api/users
func (rt *_router) searchPrenotazioni(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Diciamo al browser che stiamo restituendo dati in formato JSON
	w.Header().Set("Content-Type", "application/json")

	//prepariamo la nostra struttura dati (quella creata nel file struct.go)
	var result PrenotazioniList
	result.Prenotazione = make([]Prenotazione, 0) //inizializziamo una lista vuota

	//step 1:
	//estraiamo ciò che ha scritto l'utente nella barra di ricerca

	queryParam := r.URL.Query().Get("q")

	//step 2:
	//Invece di usare le Prepared Statements sicure (?), incolliamo brutalmente
	//l'input dell'utente dentro la query SQL usando fmt.Sprintf.
	//Usiamo LIKE per permettere la ricerca parziale (es. cerco "Sicur" e trovo "Sicurezza")

	sqlQuery := fmt.Sprintf("SELECT id_prenotazione, matricola, id_appello, insegnamento, data_esame, aula, stato FROM prenotazioni_esami WHERE insegnamento LIKE '%%%s%%'", queryParam)

	//step 3: esecuzione; chiediamo la connessione grezza e lanciamo la query

	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery)

	//step 4: error-based SQLi; se il database va in errore, non lo nascondiamo,
	//ma inviamo l'errore direttamente al frontend dentro al campo "error" del JSON

	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	//step 5: estrazione; se la query ha successo, leggiamo gli appello trovati!
	for rows.Next() {
		var p Prenotazione
		//mappiamo le 4 colonne chieste nella SELECT (cioè id, insegnamento, docente e data) nei campi della nostra struct (Appello)

		if err := rows.Scan(&p.Id_prenotazione, &p.Matricola, &p.Id_appello, &p.Insegnamento, &p.Data_esame, &p.Aula, &p.Stato); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Prenotazione = append(result.Prenotazione, p)
	}

	//step 6: risposta, inviamo il pacchetto JSON completo al client

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

//---------------------------------------------------------------------------------

//------------- CODICE SICURO -------------//
/*

import (
	_ "fmt"
	"sicurezza/service/api/reqcontext"
)

// searchUsers è la funzione che risponderà quando qualcuno visita /api/users
func (rt *_router) searchAppelli(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Diciamo al browser che stiamo restituendo dati in formato JSON
	w.Header().Set("Content-Type", "application/json")


	//prepariamo la nostra struttura dati (quella creata nel file struct.go)
	var result AppelliList
	result.Appelli = make([]Appello, 0) //inizializziamo una lista vuota

	//step 1:
	//estraiamo ciò che ha scritto l'utente nella barra di ricerca

	queryParam := r.URL.Query().Get("q")

	// step 2: LA DIFESA (Prepared Statement)
	// Al posto di fmt.Sprintf, mettiamo un segnaposto (?) nella query
	sqlQuery := "SELECT id, insegnamento, docente, data FROM appelli WHERE insegnamento LIKE ?"

	// Aggiungiamo i simboli % direttamente alla variabile in Go, non in SQL
	parametroRicerca := "%" + queryParam + "%"

	// step 3: Esecuzione sicura
	db := rt.db.GetDB()

	// Passiamo 'parametroRicerca' come secondo argomento.
	// Go e MySQL lavoreranno insieme per sterilizzare questo dato!
	rows, err := db.Query(sqlQuery, parametroRicerca)

	//step 4: error-based SQLi; se il database va in errore, non lo nascondiamo,
	//ma inviamo l'errore direttamente al frontend dentro al campo "error" del JSON

	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	//step 5: estrazione; se la query ha successo, leggiamo gli appello trovati!
	for rows.Next() {
		var a Appello
		//mappiamo le 4 colonne chieste nella SELECT (cioè id, insegnamento, docente e data) nei campi della nostra struct (Appello)

		if err := rows.Scan(&a.ID, &a.Insegnamento, &a.Docente, &a.Data); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Appelli = append(result.Appelli, a)
	}

	//step 6: risposta, inviamo il pacchetto JSON completo al client

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

*/
