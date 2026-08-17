package api

import (
	"encoding/json"
	"net/http"

	"sicurezza/service/api/reqcontext" // Usiamo il tuo modulo rinominato

	"github.com/julienschmidt/httprouter"
)

// searchUsers è la funzione che risponderà quando qualcuno visita /api/users
func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Diciamo al browser che stiamo restituendo dati in formato JSON
	w.Header().Set("Content-Type", "application/json")

	//prepariamo la nostra struttura dati (quella creata nel file struct.go)
	var result SearchResult
	result.Users = make([]User, 0) //inizializziamo una lista vuota

	//step 1: estraiamo il parametro "username" dall'URL (es. ?username=admin)
	usernameParam := r.URL.Query().Get("username")

	//step 2: La vulnerabilitò; concatenazione brutale delle stringhe !

	sqlQuery := "SELECT id, username, email FROM users WHERE username = '" + usernameParam + "'"

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

	//step 5: estrazione; se la query ha successo, leggiamo gli utenti trovati!
	for rows.Next() {
		var u User
		//mappiamo le 3 colonne chieste nella SELECT (cioè id, username e email) nei campi della nostra struct (User)

		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Users = append(result.Users, u)
	}

	//step 6: risposta, inviamo il pacchetto JSON completo al client

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
