// repository/beer_repository.go
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/d90ares/iBeers/domain"
)

type BeerRepository struct {
	db *sql.DB
}

func NewBeerRepository(db *sql.DB) *BeerRepository {
	return &BeerRepository{
		db: db,
	}
}

func (r *BeerRepository) GetAll(ctx context.Context) ([]*domain.Beer, error) {
	// Implementar lógica para obter todas as cervejas do banco de dados

	if err := r.db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM beers")
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta: %w", err)
	}
	defer rows.Close()
	var b domain.Beer
	if err := rows.Scan(&b.ID, &b.Name, &b.Style, &b.Type); err != nil {
		return nil, fmt.Errorf("erro ao mapear dados: %w", err)
	}

	fmt.Println("Repository: Obter todas as cervejas")
	return nil, nil
}

// Implementar outros métodos de repositório, como Store, Update, Remove, etc.
