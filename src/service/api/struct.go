package api

/*
Il server GO ha bisogno di strutture (struct) per capire come impacchettare i dati presi
dal database e inviarli a Vue.js (il frontend)

Le "etichette" (tag) sulla destra indicano a Go come rinominare i campi quando li converte in JSON.
*/

type Studente struct {
	matricola       int    `json:"matricola"`
	nome            string `json:"nome"`
	cognome         string `json:"cognome"`
	data_di_nascita string `json:"data_di_nascita"`
	indirizzo       string `json:"indirizzo"`
	password        string `json:"password"`
	email           string `json:"email"`
}

type StudentiList struct {
	Studente []Studente `json:"studente"`
	Error    string     `json:"error, omitempty"`
}

// -----------------------
type Appello struct {
	ID           int    `json:"id"`
	Insegnamento string `json:"insegnamento"`
	Docente      string `json:"docente"`
	Data         string `json:"data"`
}

// AppelliList sarà il "pacchetto" completo che invieremo al frontend Vue.js.
// Conterrà sia la lista degli appelli trovati, sia eventuali messaggi di errore del database
// (il campo Error è fondamentale per far vedere a schermo l'Error-Based SQL Injection).
type AppelliList struct {
	Appelli []Appello `json:"appelli"`
	Error   string    `json:"error,omitempty"` // omitempty nasconde il campo nel JSON se non ci sono errori
}

// ----------------
type Stato string

const (
	StatoConfermato Stato = "Confermato"
	StatoInAttesa         = "In attesa"
	StatoRifiutato        = "Rifiutato"
)

type Prenotazione struct {
	Id_prenotazione int    `json:"id_prenotazione"`
	Matricola       int    `json:"matricola"`
	Id_appello      int    `json:"id_appello"`
	Insegnamento    string `json:"insegnamento"`
	Data_esame      string `json:"data_esame"`
	Aula            string `json:"aula"`
	Stato           Stato  `json:"stato"`
}

type PrenotazioniList struct {
	Prenotazione []Prenotazione `json:"prenotazione"`
	Error        string         `json:"error,omitempty"`
}

//-----------------------

type Libretto struct {
	matricola          int    `json:"matricola"`
	id_appello         int    `json:"id_appello"`
	insegnamento       int    `json:"insegnamento"`
	data_registrazione string `json:"data_registrazione"`
	voto               int    `json:"voto"`
	CFU                string `json:"cfu"`
}

type LibrettoList struct {
	Libretto []Libretto `json:"libretto"`
	Error    string     `json:"error, omitempty"`
}
