package api

/*
Il server GO ha bisogno di strutture (struct) per capire come impacchettare i dati presi
dal database e inviarli a Vue.js (il frontend)

Le "etichette" (tag) sulla destra indicano a Go come rinominare i campi quando li converte in JSON.
*/
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
