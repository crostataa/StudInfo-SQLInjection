package api

// User rappresenta la struttura dei dati di un utente aziendale.
// Le "etichette" (tag) sulla destra indicano a Go come rinominare i campi quando li converte in JSON.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// SearchResult sarà il "pacchetto" completo che invieremo al frontend Vue.js.
// Conterrà sia la lista degli utenti trovati, sia eventuali messaggi di errore del database
// (il campo Error è fondamentale per far vedere a schermo l'Error-Based SQL Injection).
type SearchResult struct {
	Users []User `json:"users"`
	Error string `json:"error,omitempty"` // omitempty nasconde il campo nel JSON se non ci sono errori
}
