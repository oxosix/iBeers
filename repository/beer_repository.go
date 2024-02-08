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

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM beer")
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta: %w", err)
	}

	defer rows.Close()

	var beers []*domain.Beer
	for rows.Next() {
		var b domain.Beer
		if err := rows.Scan(&b.ID, &b.Name, &b.Style, &b.Type); err != nil {
			return nil, fmt.Errorf("erro ao mapear dados: %w", err)
		}
		beers = append(beers, &b)
	}

	fmt.Println("Repository: Obter todas as cervejas")
	return beers, nil
}

func (r *BeerRepository) GetByID(ctx context.Context, id int) (*domain.Beer, error) {
	// Implementar lógica para obter uma cerveja pelo seu ID

	if err := r.db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	row := r.db.QueryRowContext(ctx, "SELECT * FROM beer WHERE id = $1", id)

	var b domain.Beer
	if err := row.Scan(&b.ID, &b.Name, &b.Style, &b.Type); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cerveja não encontrada")
		}
		return nil, fmt.Errorf("erro ao mapear dados: %w", err)
	}

	fmt.Println("Repository: Obter cerveja por ID")
	return &b, nil
}

func (r *BeerRepository) Add(ctx context.Context, beer *domain.Beer) (*domain.Beer, error) {
	// Implementar lógica para adicionar uma nova cerveja ao banco de dados
	if err := r.db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	_, err := r.db.ExecContext(ctx, "INSERT INTO beer (name, style, type) VALUES ($1, $2, $3)", beer.Name, beer.Style, beer.Type)
	if err != nil {
		return nil, fmt.Errorf("erro ao adicionar cerveja: %w", err)
	}

	fmt.Println("Repository: Adicionar nova cerveja")
	return beer, nil
}

// Implementar outros métodos de repositório, como Store, Update, Remove, etc.
