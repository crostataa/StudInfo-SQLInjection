package api

import (
	"net/http"

	"sicurezza/service/api/reqcontext" // Usiamo il tuo modulo rinominato

	"github.com/julienschmidt/httprouter"
)

// searchUsers è la funzione che risponderà quando qualcuno visita /api/users
func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Diciamo al browser che stiamo restituendo dati in formato JSON
	w.Header().Set("Content-Type", "application/json")

	// Rispondiamo con un codice 200 OK e un finto JSON di test
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "messaggio": "L'API è in ascolto e pronta per il payload SQLi!"}`))
}
