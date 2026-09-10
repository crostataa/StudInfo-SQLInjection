/*
//------------- CODICE VULNERABILE -------------//

package api

import (
	"encoding/json"
	"net/http"
	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)





import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
	//estraiamo ciò che ha scritto l'utente e la sua matricola
	queryParam := r.URL.Query().Get("q")
	matricolaUtente := r.URL.Query().Get("matricola") // Estratto inviato da Vue

	//step 2: CODICE VULNERABILE (ma filtrato per utente)
	// Rimuoviamo "matricola" dalla SELECT e aggiungiamo WHERE matricola = %s
	sqlQuery := fmt.Sprintf("SELECT id_prenotazione, id_appello, insegnamento, data_esame, aula, stato FROM prenotazioni_esami WHERE matricola = %s AND insegnamento LIKE '%%%s%%'", matricolaUtente, queryParam)

	//step 3: esecuzione
	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery)

	//step 4: error-based SQLi
	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	//step 5: estrazione
	for rows.Next() {
		var p Prenotazione
		if err := rows.Scan(&p.Id_prenotazione, &p.Id_appello, &p.Insegnamento, &p.Data_esame, &p.Aula, &p.Stato); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Prenotazione = append(result.Prenotazione, p)
	}

	// --- INIZIO MODIFICA PER IL TEST (ABILITA STACKED QUERIES) ---
	// Costringiamo il driver a processare anche i comandi successivi (es. la DELETE)
	// ignorando i risultati, ma facendoli eseguire fisicamente al database.
	for rows.NextResultSet() {
		// Consumiamo eventuali righe restituite dai comandi accodati
		for rows.Next() {
			// non facciamo nulla
		}
	}
	// --- FINE MODIFICA PER IL TEST ---

	//step 6: risposta, inviamo il pacchetto JSON completo al client

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

*/

//---------------------------------------------------------------------------------

//------------- CODICE SICURO -------------//

package api

import (
	"encoding/json"
	"net/http"

	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// searchPrenotazioni è la funzione che risponderà quando qualcuno visita /api/prenotazioni
func (rt *_router) searchPrenotazioni(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Diciamo al browser che stiamo restituendo dati in formato JSON
	w.Header().Set("Content-Type", "application/json")

	// Prepariamo la nostra struttura dati
	var result PrenotazioniList
	result.Prenotazione = make([]Prenotazione, 0) // Inizializziamo una lista vuota

	// Step 1: Estraiamo ciò che ha scritto l'utente e la sua matricola
	queryParam := r.URL.Query().Get("q")
	matricolaUtente := r.URL.Query().Get("matricola") // Estratto inviato da Vue

	// Step 2: LA DIFESA (Prepared Statement)
	// Nessun fmt.Sprintf. Usiamo i segnaposto (?) per i dati variabili.
	sqlQuery := "SELECT id_prenotazione, id_appello, insegnamento, data_esame, aula, stato FROM prenotazioni_esami WHERE matricola = ? AND insegnamento LIKE ?"

	// Prepariamo il parametro per il LIKE direttamente in Go
	parametroRicerca := "%" + queryParam + "%"

	// Step 3: Esecuzione sicura
	db := rt.db.GetDB()

	// Passiamo i parametri (matricola e ricerca) direttamente a db.Query
	// Il driver MySQL tratterà questi valori RIGOROSAMENTE come dati e mai come comandi SQL.
	rows, err := db.Query(sqlQuery, matricolaUtente, parametroRicerca)

	// Step 4: Error-based SQLi mitigata (l'errore di query viene comunque gestito)
	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	// Step 5: Estrazione
	for rows.Next() {
		var p Prenotazione
		if err := rows.Scan(&p.Id_prenotazione, &p.Id_appello, &p.Insegnamento, &p.Data_esame, &p.Aula, &p.Stato); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Prenotazione = append(result.Prenotazione, p)
	}

	// Step 6: Risposta, inviamo il pacchetto JSON completo al client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
