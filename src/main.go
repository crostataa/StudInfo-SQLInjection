package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	// Driver MySQL
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	ID       int
	Username string
	Email    string
}

type PageData struct {
	Query   string
	Error   string
	Results []User
}

func main() {
	var err error

	// multiStatements=true consente query separate da ';' (piggybacking) a scopo didattico.
	// Versione sicura: omettere multiStatements per impedire stacked queries.
	dsn := "appuser:apppassword@tcp(db:3306)/vulnerabile_db?multiStatements=true"
	// dsn := "appuser:apppassword@tcp(db:3306)/vulnerabile_db" // Configurazione sicura

	// Attesa disponibilità del container MySQL
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

	http.HandleFunc("/", searchHandler)

	log.Println("Server avviato su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
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

	queryParam := r.URL.Query().Get("username")
	data := PageData{Query: queryParam}

	if queryParam != "" {
		// Vulnerabile a SQL Injection (concatenazione diretta dell'input)
		sqlQuery := "SELECT id, username, email FROM users WHERE username = '" + queryParam + "'"
		rows, err := db.Query(sqlQuery)

		// Alternativa sicura (Prepared Statement):
		// sqlQuery := "SELECT id, username, email FROM users WHERE username = ?"
		// rows, err := db.Query(sqlQuery, queryParam)

		if err != nil {
			// Espone l'errore SQL nel frontend (consente Error-Based SQLi)
			data.Error = err.Error()
		} else {
			defer rows.Close()
			for rows.Next() {
				var u User
				if err := rows.Scan(&u.ID, &u.Username, &u.Email); err == nil {
					data.Results = append(data.Results, u)
				}
			}
		}
	}

	t, _ := template.New("webpage").Parse(htmlTemplate)
	t.Execute(w, data)
}
