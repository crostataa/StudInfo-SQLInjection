FROM golang:1.22

WORKDIR /app

# Download delle dipendenze e avvio dell'applicazione
CMD sh -c "go mod tidy && go run main.go"
