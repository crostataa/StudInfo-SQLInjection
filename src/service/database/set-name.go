package database

// SetName è un esempio di esecuzione di insert/update
func (db *appdbimpl) SetName(name string) error {
	_, err := db.c.Exec("INSERT INTO example_table (id, name) VALUES (1, ?)", name)
	return err
}
