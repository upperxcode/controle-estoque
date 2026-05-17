package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "./estoque.db")
	if err != nil {
		return nil, err
	}

	// Habilitar chaves estrangeiras no SQLite
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexão com SQLite estabelecida com sucesso")

	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		sku TEXT UNIQUE NOT NULL,
		quantity INTEGER DEFAULT 0 NOT NULL,
		price REAL NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS movements (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER NOT NULL,
		type TEXT NOT NULL CHECK(type IN ('IN', 'OUT')),
		quantity INTEGER NOT NULL CHECK(quantity > 0),
		note TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(product_id) REFERENCES products(id)
	);`
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
