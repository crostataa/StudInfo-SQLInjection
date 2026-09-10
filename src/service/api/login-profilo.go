package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sicurezza/service/api/reqcontext"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ==================== CODICE VULNERABILE ====================

// 1. Gestione login (vulnerabile a tautologia)
func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	var res LoginResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = "Dati non validi"
		json.NewEncoder(w).Encode(res)
		return
	}

	db := rt.db.GetDB()

	// Query vulnerabile a SQL Injection (tautologia tramite interpolazione stringa)
	query := fmt.Sprintf("SELECT matricola, nome, cognome FROM students WHERE matricola = %d AND password = '%s'", req.Matricola, req.Password)
	err := db.QueryRow(query).Scan(&res.Matricola, &res.Nome, &res.Cognome)

	/*
		// Alternativa sicura (Prepared Statement):
		query := "SELECT matricola, nome, cognome FROM students WHERE matricola = ? AND password = ?"
		err := db.QueryRow(query, req.Matricola, req.Password).Scan(&res.Matricola, &res.Nome, &res.Cognome)
	*/

	if err != nil {
		res.Success = false
		res.Error = "Matricola o password errate"
	} else {
		res.Success = true
	}

	json.NewEncoder(w).Encode(res)
}

// 2. Gestione profilo ( vulnerabile senza Prepared Statement)
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

// 3. Gestione libretto (protetta con Prepared Statement)
func (rt *_router) getLibretto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	matricolaStr := ps.ByName("matricola")
	matricola, _ := strconv.Atoi(matricolaStr)

	var res LibrettoList
	res.Libretto = make([]Libretto, 0)

	db := rt.db.GetDB()
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
// ==================== CODICE SICURO ====================

// 1. Gestione login (messo in sicurezza con Prepared Statement)
func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	var res LoginResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res.Error = "Dati non validi"
		json.NewEncoder(w).Encode(res)
		return
	}

	db := rt.db.GetDB()

	// Prepared Statement: separa la struttura SQL dai dati dell'utente (neutralizza la tautologia)
	query := "SELECT matricola, nome, cognome FROM students WHERE matricola = ? AND password = ?"
	err := db.QueryRow(query, req.Matricola, req.Password).Scan(&res.Matricola, &res.Nome, &res.Cognome)

	if err != nil {
		res.Success = false
		res.Error = "Matricola o password errate"
	} else {
		res.Success = true
	}

	json.NewEncoder(w).Encode(res)
}

// 2. Gestione profilo (sicuro)
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

// 3. Gestione libretto (sicuro)
func (rt *_router) getLibretto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	matricolaStr := ps.ByName("matricola")
	matricola, _ := strconv.Atoi(matricolaStr)

	var res LibrettoList
	res.Libretto = make([]Libretto, 0)

	db := rt.db.GetDB()
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
