/*
// ==================== CODICE VULNERABILE ====================

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchPrenotazioni(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var result PrenotazioniList
	result.Prenotazione = make([]Prenotazione, 0)

	queryParam := r.URL.Query().Get("q")
	matricolaUtente := r.URL.Query().Get("matricola")

	// Query vulnerabile: concatenazione parametri senza validazione (soggetta a stacked queries)
	sqlQuery := fmt.Sprintf("SELECT id_prenotazione, id_appello, insegnamento, data_esame, aula, stato FROM prenotazioni_esami WHERE matricola = %s AND insegnamento LIKE '%%%s%%'", matricolaUtente, queryParam)

	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery)
	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Prenotazione
		if err := rows.Scan(&p.Id_prenotazione, &p.Id_appello, &p.Insegnamento, &p.Data_esame, &p.Aula, &p.Stato); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Prenotazione = append(result.Prenotazione, p)
	}

	// Consuma eventuali query accodate (es. DELETE con piggybacking)
	for rows.NextResultSet() {
		for rows.Next() {
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
*/

// ==================== CODICE SICURO ====================

package api

import (
	"encoding/json"
	"net/http"

	"sicurezza/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// searchPrenotazioni gestisce la ricerca delle prenotazioni d'esame per lo studente
func (rt *_router) searchPrenotazioni(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var result PrenotazioniList
	result.Prenotazione = make([]Prenotazione, 0)

	queryParam := r.URL.Query().Get("q")
	matricolaUtente := r.URL.Query().Get("matricola")

	// Query protetta con Prepared Statement
	sqlQuery := "SELECT id_prenotazione, id_appello, insegnamento, data_esame, aula, stato FROM prenotazioni_esami WHERE matricola = ? AND insegnamento LIKE ?"
	parametroRicerca := "%" + queryParam + "%"

	db := rt.db.GetDB()
	rows, err := db.Query(sqlQuery, matricolaUtente, parametroRicerca)

	if err != nil {
		result.Error = err.Error()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Prenotazione
		if err := rows.Scan(&p.Id_prenotazione, &p.Id_appello, &p.Insegnamento, &p.Data_esame, &p.Aula, &p.Stato); err != nil {
			rt.baseLogger.WithError(err).Error("Error scanning row")
			continue
		}
		result.Prenotazione = append(result.Prenotazione, p)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
