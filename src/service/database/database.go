package database

import (
	"database/sql"
	"errors"
)

// AppDatabase è l'interfaccia che definisce cosa può fare il nostro database.
type AppDatabase interface {
	Ping() error
	GetDB() *sql.DB // AGGIUNTA: permette all'API di accedere direttamente al motore SQL
}

type appdbimpl struct {
	c *sql.DB
}

// New crea una nuova istanza del nostro database.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("la connessione al database è nil")
	}
	return &appdbimpl{
		c: db,
	}, nil
}

// Ping controlla che la connessione sia viva.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

// GetDB restituisce il puntatore grezzo al database MySQL.
// Lo useremo nel nostro file dell'API per lanciare la query SQL Injection!
func (db *appdbimpl) GetDB() *sql.DB {
	return db.c
}
