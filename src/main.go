package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	// Importiamo il driver MySQL. Il blank identifier "_" è necessario
	// perché Go carichi il driver dietro le quinte senza chiamarlo direttamente.

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Struttura per mappare i risultati del database
type User struct {
	ID       int
	Username string
	Email    string
}

// Struttura per passare i dati alla pagina HTML
type PageData struct {
	Query   string
	Error   string
	Results []User
}

func main() {
	var err error

	//------------CODICE VULNERABILE-------------
	// NOTA BENE: multiStatements=true è FONDAMENTALE. Permette l'esecuzione di query separate da ";" (Piggybacking).
	dsn := "appuser:apppassword@tcp(db:3306)/vulnerabile_db?multiStatements=true"

	//-----------CODICE SICURO-------------------
	//blocchiamo il piggybacking
	//dsn := "appuser:apppassword@tcp(db:3306)/vulnerabile_db"

	//-------------------------------------------

	// RETRY LOOP: Attendiamo che il container MySQL sia effettivamente pronto ad accettare connessioni
	log.Println("Tentativo di connessione al database...")
	for {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Database connesso con successo!")
				break
			}
		}
		log.Println("Database non ancora pronto. Attendo 3 secondi...")
		time.Sleep(3 * time.Second)
	}
	defer db.Close()

	// Definiamo la rotta principale dell'applicazione web
	http.HandleFunc("/", searchHandler)

	log.Println("Server avviato su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Template HTML integrato per semplicità. Genera una barra di ricerca e una tabella.
	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Laboratorio SQL Injection (Go)</title>
		<style>
			body { font-family: 'Segoe UI', Arial, sans-serif; margin: 40px; background-color: #f5f7fb; color: #333; }
			.container { max-width: 800px; margin: auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); }
			h2 { color: #2c3e50; border-bottom: 2px solid #ecf0f1; padding-bottom: 10px; }
			input[type="text"] { width: 75%; padding: 10px; border: 1px solid #bdc3c7; border-radius: 4px; font-size: 16px; }
			input[type="submit"] { width: 20%; padding: 10px; background-color: #e74c3c; color: white; border: none; border-radius: 4px; font-size: 16px; cursor: pointer; }
			input[type="submit"]:hover { background-color: #c0392b; }
			.error-box { background-color: #fadbd8; border-left: 5px solid #e74c3c; padding: 15px; margin-top: 20px; border-radius: 4px; font-family: monospace; white-space: pre-wrap; }
			table { width: 100%; border-collapse: collapse; margin-top: 25px; }
			th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
			th { background-color: #34495e; color: white; }
			tr:hover { background-color: #f2f2f2; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Ricerca Utenti Aziendali (Vulnerabile a SQLi)</h2>
			<form method="GET" action="/">
				<input type="text" name="username" placeholder="Inserisci lo username da cercare..." value="{{.Query}}">
				<input type="submit" value="Cerca">
			</form>

			{{if .Error}}
				<div class="error-box">
					<strong>Errore del Database (Utile per Error-Based SQLi):</strong><br>{{.Error}}
				</div>
			{{end}}

			{{if .Results}}
				<h3>Risultati Trovati:</h3>
				<table>
					<tr>
						<th>ID</th>
						<th>Username</th>
						<th>Email</th>
					</tr>
					{{range .Results}}
					<tr>
						<td>{{.ID}}</td>
						<td>{{.Username}}</td>
						<td>{{.Email}}</td>
					</tr>
					{{end}}
				</table>
			{{end}}
		</div>
	</body>
	</html>
	`

	// Recuperiamo il parametro "username" dall'URL (es: http://localhost:8080/?username=admin)
	queryParam := r.URL.Query().Get("username")
	data := PageData{Query: queryParam}

	//-----------CODICE VULNERABILE-------------
	if queryParam != "" {
		// VULNERABILITÀ CRITICA: L'input dell'utente viene inserito direttamente per string-concatenation
		// Non c'è alcuna validazione o parametrizzazione.
		sqlQuery := "SELECT id, username, email FROM users WHERE username = '" + queryParam + "'"
		// Eseguiamo la query.
		rows, err := db.Query(sqlQuery)

		//-----------CODICE SICURO-------------------
		//Prepared Statements (Query Parametrizzate). Con questa tecnica, inviamo al database la
		//struttura della query separata dai dati.
		//Il database tratterà i dati come semplice testo, rendendo impossibile
		//l'esecuzione di codice malevolo.

		//sqlQuery := "SELECT id, username, email FROM users WHERE username = ?"

		// Passiamo queryParam come secondo argomento.
		// Go e MySQL si assicureranno che venga trattato ESCLUSIVAMENTE come una stringa innocua.

		//rows, err := db.Query(sqlQuery, queryParam)

		//--------------------------------------------
		if err != nil {
			// Mostriamo l'errore SQL in chiaro sulla pagina web.
			// Questo permetterà all'attaccante di fare "Error-Based" SQLi o capire come correggere la sintassi dell'attacco.
			data.Error = err.Error()
		} else {
			defer rows.Close()
			for rows.Next() {
				var u User
				// Eseguiamo lo Scan delle colonne.
				// Se l'attaccante usa una UNION SELECT e altera il numero o i tipi di colonne, lo Scan potrebbe fallire internamente,
				// ma le query piggybacked distruttive verranno comunque eseguite nel DB a monte!
				if err := rows.Scan(&u.ID, &u.Username, &u.Email); err == nil {
					data.Results = append(data.Results, u)
				}
			}
		}
	}

	// Rendering della pagina
	t, _ := template.New("webpage").Parse(htmlTemplate)
	t.Execute(w, data)
}
