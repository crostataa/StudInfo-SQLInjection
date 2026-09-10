package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// ==================== CODICE VULNERABILE ===================

// searchAppelli gestisce la ricerca degli appelli per nome dell'insegnamento
func (rt *_router) searchAppelli(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var result AppelliList
	result.Appelli = make([]Appello, 0)

	queryParam := r.URL.Query().Get("q")

	// Query vulnerabile: concatenazione diretta dell'input nella clausola LIKE
	sqlQuery := fmt.Sprintf("SELECT id, insegnamento, docente, data FROM appelli WHERE insegnamento LIKE '%%%s%%'", queryParam)

	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery)

	// Espone l'errore SQL in chiaro al frontend (consente Error-Based SQLi)
	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var a Appello
		if err := rows.Scan(&a.ID, &a.Insegnamento, &a.Docente, &a.Data); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Appelli = append(result.Appelli, a)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

/*
// ==================== CODICE SICURO ====================

func (rt *_router) searchAppelli(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var result AppelliList
	result.Appelli = make([]Appello, 0)

	queryParam := r.URL.Query().Get("q")

	// Query protetta con Prepared Statement: il parametro viene sterilizzato da MySQL
	sqlQuery := "SELECT id, insegnamento, docente, data FROM appelli WHERE insegnamento LIKE ?"
	parametroRicerca := "%" + queryParam + "%"

	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery, parametroRicerca)

	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var a Appello
		if err := rows.Scan(&a.ID, &a.Insegnamento, &a.Docente, &a.Data); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Appelli = append(result.Appelli, a)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
*/
