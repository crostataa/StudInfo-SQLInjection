package api

//------------- CODICE VULNERABILE -------------//
import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// searchUsers è la funzione che risponderà quando qualcuno visita /api/users
func (rt *_router) searchAppelli(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Diciamo al browser che stiamo restituendo dati in formato JSON
	w.Header().Set("Content-Type", "application/json")

	//accetta richieste da qualunque porta, disabilita il CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//prepariamo la nostra struttura dati (quella creata nel file struct.go)
	var result AppelliList
	result.Appelli = make([]Appello, 0) //inizializziamo una lista vuota

	//step 1:
	//estraiamo ciò che ha scritto l'utente nella barra di ricerca

	queryParam := r.URL.Query().Get("q")

	//step 2:
	//Invece di usare le Prepared Statements sicure (?), incolliamo brutalmente
	//l'input dell'utente dentro la query SQL usando fmt.Sprintf.
	//Usiamo LIKE per permettere la ricerca parziale (es. cerco "Sicur" e trovo "Sicurezza")

	sqlQuery := fmt.Sprintf("SELECT id, insegnamento, docente, data FROM appelli WHERE insegnamento LIKE '%%%s%%'", queryParam)

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

//---------------------------------------------------------------------------------

//------------- CODICE SICURO DA MODIFICARE-------------//
/*

import (
	"encoding/json"
	"net/http"

	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var result SearchResult
	result.Users = make([]User, 0)

	// 1. INPUT: Estraiamo il parametro dall'URL
	usernameParam := r.URL.Query().Get("username")

	// 2. LA DIFESA (Prepared Statement): Usiamo il segnaposto "?"
	// Non ci sono più apici o concatenazioni pericolose!
	sqlQuery := "SELECT id, username, email FROM users WHERE username = ?"

	// 3. ESECUZIONE SICURA: Passiamo la query e il parametro separatamente
	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery, usernameParam)

	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	// 4. ESTRAZIONE
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			rt.baseLogger.WithError(err).Error("Errore durante la lettura di una riga")
			continue
		}
		result.Users = append(result.Users, u)
	}

	// 5. RISPOSTA
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

*/
