package repository

import (
	"os"
	"testing"

	"github.com/user/controle-estoque/backend/internal/database"
	"github.com/user/controle-estoque/backend/internal/model"
)

func TestProductRepository(t *testing.T) {
	os.Remove("./estoque.db")
	db, _ := database.Connect()
	defer db.Close()

	repo := NewProductRepository(db)

	p := &model.Product{
		Name:  "Teclado Mecânico",
		SKU:   "TEC-001",
		Price: 250.0,
	}

	t.Run("Create", func(t *testing.T) {
		err := repo.Create(p)
		if err != nil {
			t.Fatalf("Erro ao criar produto: %v", err)
		}
		if p.ID == 0 {
			t.Fatal("ID não foi atribuído")
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		saved, err := repo.GetByID(p.ID)
		if err != nil {
			t.Fatalf("Erro ao buscar produto: %v", err)
		}
		if saved.Name != p.Name {
			t.Errorf("Esperava nome %s, recebeu %s", p.Name, saved.Name)
		}
	})

	t.Run("List", func(t *testing.T) {
		products, err := repo.List()
		if err != nil {
			t.Fatalf("Erro ao listar: %v", err)
		}
		if len(products) != 1 {
			t.Errorf("Esperava 1 produto, recebeu %d", len(products))
		}
	})
}
