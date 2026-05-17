package database

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	// Limpar banco de teste se existir
	os.Remove("./estoque.db")

	db, err := Connect()
	if err != nil {
		t.Fatalf("Erro esperado nil, recebeu %v", err)
	}
	defer db.Close()

	if db == nil {
		t.Fatal("Esperava ponteiro de banco de dados, recebeu nil")
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("Erro ao pingar banco de dados: %v", err)
	}
}
