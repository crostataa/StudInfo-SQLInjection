// --------- CODICE VULNERABILE ---------

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sicurezza/service/api/reqcontext"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// 1. GESTIONE LOGIN (Ora impostato come VULNERABILE per testare la Tautologia)
func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	var res LoginResponse

	// Leggiamo matricola e password inviate da Vue
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = "Dati non validi"
		json.NewEncoder(w).Encode(res)
		return
	}

	db := rt.db.GetDB()

	// --- INIZIO CODICE VULNERABILE ---
	// Incolliamo la password direttamente nella stringa SQL usando fmt.Sprintf
	query := fmt.Sprintf("SELECT matricola, nome, cognome FROM students WHERE matricola = %d AND password = '%s'", req.Matricola, req.Password)

	// Eseguiamo la query grezza (senza passare i parametri separatamente)
	err := db.QueryRow(query).Scan(&res.Matricola, &res.Nome, &res.Cognome)
	// --- FINE CODICE VULNERABILE ---

	/*
		// --- INIZIO CODICE SICURO ---
		query := "SELECT matricola, nome, cognome FROM students WHERE matricola = ? AND password = ?"
		err := db.QueryRow(query, req.Matricola, req.Password).Scan(&res.Matricola, &res.Nome, &res.Cognome)
		// --- FINE CODICE SICURO ---
	*/

	if err != nil {
		res.Success = false
		res.Error = "Matricola o password errate"
	} else {
		res.Success = true
	}

	json.NewEncoder(w).Encode(res)
}

// 2. GESTIONE PROFILO (Sicuro)
func (rt *_router) getProfilo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	matricolaStr := ps.ByName("matricola")
	matricola, _ := strconv.Atoi(matricolaStr)

	var s Studente
	var res StudentiList

	db := rt.db.GetDB()
	query := "SELECT matricola, nome, cognome, data_di_nascita, indirizzo, email FROM students WHERE matricola = ?"

	err := db.QueryRow(query, matricola).Scan(&s.Matricola, &s.Nome, &s.Cognome, &s.Data_di_nascita, &s.Indirizzo, &s.Email)

	if err != nil {
		res.Error = "Studente non trovato"
	} else {
		res.Studente = append(res.Studente, s)
	}

	json.NewEncoder(w).Encode(res)
}

// 3. GESTIONE LIBRETTO (Sicuro e ripristinato)
func (rt *_router) getLibretto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	matricolaStr := ps.ByName("matricola")
	matricola, _ := strconv.Atoi(matricolaStr)

	var res LibrettoList
	res.Libretto = make([]Libretto, 0)

	db := rt.db.GetDB()

	// Ripristinata la query corretta per il libretto
	query := "SELECT matricola, id_appello, insegnamento, data_registrazione, voto, CFU FROM libretto_esami WHERE matricola = ?"

	rows, err := db.Query(query, matricola)

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var l Libretto
		if err := rows.Scan(&l.Matricola, &l.Id_appello, &l.Insegnamento, &l.Data_registrazione, &l.Voto, &l.CFU); err != nil {
			fmt.Println("Errore lettura database", err)
			continue
		}
		res.Libretto = append(res.Libretto, l)
	}

	json.NewEncoder(w).Encode(res)
}

/*
// -------- CODICE SICURO ---------------
package api

import (
"encoding/json"
"fmt"
"net/http"
"sicurezza/service/api/reqcontext"
"strconv"

"github.com/julienschmidt/httprouter"
)

// 1. GESTIONE LOGIN (Messo in sicurezza con Prepared Statement)
func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	var res LoginResponse

	// Leggiamo matricola e password inviate da Vue
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = "Dati non validi"
		json.NewEncoder(w).Encode(res)
		return
	}

	db := rt.db.GetDB()

	// --- INIZIO CODICE SICURO ---
	// Usiamo i segnaposto (?) per separare la struttura della query dai dati dell'utente.
	// Questo neutralizza completamente gli attacchi di Tautologia.
	query := "SELECT matricola, nome, cognome FROM students WHERE matricola = ? AND password = ?"

	// Passiamo i parametri separatamente nella funzione QueryRow
	err := db.QueryRow(query, req.Matricola, req.Password).Scan(&res.Matricola, &res.Nome, &res.Cognome)
	// --- FINE CODICE SICURO ---

	if err != nil {
		res.Success = false
		res.Error = "Matricola o password errate"
	} else {
		res.Success = true
	}

	json.NewEncoder(w).Encode(res)
}

// 2. GESTIONE PROFILO (Sicuro)
func (rt *_router) getProfilo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	matricolaStr := ps.ByName("matricola")
	matricola, _ := strconv.Atoi(matricolaStr)

	var s Studente
	var res StudentiList

	db := rt.db.GetDB()
	query := "SELECT matricola, nome, cognome, data_di_nascita, indirizzo, email FROM students WHERE matricola = ?"

	err := db.QueryRow(query, matricola).Scan(&s.Matricola, &s.Nome, &s.Cognome, &s.Data_di_nascita, &s.Indirizzo, &s.Email)

	if err != nil {
		res.Error = "Studente non trovato"
	} else {
		res.Studente = append(res.Studente, s)
	}

	json.NewEncoder(w).Encode(res)
}

// 3. GESTIONE LIBRETTO (Sicuro)
func (rt *_router) getLibretto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	matricolaStr := ps.ByName("matricola")
	matricola, _ := strconv.Atoi(matricolaStr)

	var res LibrettoList
	res.Libretto = make([]Libretto, 0)

	db := rt.db.GetDB()

	// Query corretta per il libretto (già protetta con Prepared Statement)
	query := "SELECT matricola, id_appello, insegnamento, data_registrazione, voto, CFU FROM libretto_esami WHERE matricola = ?"

	rows, err := db.Query(query, matricola)

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var l Libretto
		if err := rows.Scan(&l.Matricola, &l.Id_appello, &l.Insegnamento, &l.Data_registrazione, &l.Voto, &l.CFU); err != nil {
			fmt.Println("Errore lettura database", err)
			continue
		}
		res.Libretto = append(res.Libretto, l)
	}

	json.NewEncoder(w).Encode(res)
}

*/
