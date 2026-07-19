# Usiamo l'immagine ufficiale di Go (versione 1.22)
FROM golang:1.22

# Impostiamo la cartella di lavoro dentro il container
WORKDIR /app

# Usiamo un comando shell per scaricare le dipendenze (go mod tidy)
# e SUBITO DOPO avviare il server (go run main.go)
CMD sh -c "go mod tidy && go run main.go"
