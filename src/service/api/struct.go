package api

// Strutture dati per la serializzazione JSON delle entità del database e delle risposte API

type Studente struct {
	Matricola       int    `json:"matricola"`
	Nome            string `json:"nome"`
	Cognome         string `json:"cognome"`
	Data_di_nascita string `json:"data_di_nascita"`
	Indirizzo       string `json:"indirizzo"`
	Password        string `json:"password"`
	Email           string `json:"email"`
}

type StudentiList struct {
	Studente []Studente `json:"studente"`
	Error    string     `json:"error, omitempty"`
}

type Appello struct {
	ID           int    `json:"id"`
	Insegnamento string `json:"insegnamento"`
	Docente      string `json:"docente"`
	Data         string `json:"data"`
}

// AppelliList contiene l'elenco degli appelli ed eventuali errori del database (per Error-Based SQLi)
type AppelliList struct {
	Appelli []Appello `json:"appelli"`
	Error   string    `json:"error,omitempty"`
}

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

type Libretto struct {
	Matricola          int    `json:"matricola"`
	Id_appello         int    `json:"id_appello"`
	Insegnamento       string `json:"insegnamento"`
	Data_registrazione string `json:"data_registrazione"`
	Voto               int    `json:"voto"`
	CFU                int    `json:"cfu"`
}

type LibrettoList struct {
	Libretto []Libretto `json:"libretto"`
	Error    string     `json:"error, omitempty"`
}

type LoginRequest struct {
	Matricola int    `json:"matricola"`
	Password  string `json:"password"`
}

type LoginResponse struct {
	Success   bool   `json:"success"`
	Matricola int    `json:"matricola"`
	Nome      string `json:"nome"`
	Cognome   string `json:"cognome"`
	Error     string `json:"error,omitempty"`
}
