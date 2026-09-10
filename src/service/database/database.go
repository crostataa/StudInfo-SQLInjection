package database

import (
	"database/sql"
	"errors"
)

// AppDatabase definisce l'interfaccia di accesso al database.
type AppDatabase interface {
	Ping() error
	GetDB() *sql.DB // Accesso diretto alla connessione *sql.DB
}

type appdbimpl struct {
	c *sql.DB
}

// New inizializza una nuova istanza di AppDatabase.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("la connessione al database è nil")
	}
	return &appdbimpl{
		c: db,
	}, nil
}

// Ping verifica lo stato della connessione al database.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

// GetDB restituisce la connessione sottostante *sql.DB.
func (db *appdbimpl) GetDB() *sql.DB {
	return db.c
}
