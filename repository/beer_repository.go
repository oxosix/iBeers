// repository/beer_repository.go
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/d90ares/iBeers/config/logs"
	"github.com/d90ares/iBeers/domain"
)

type BeerRepository struct {
	DB *sql.DB
}

func NewBeerRepository(db *sql.DB) *BeerRepository {
	return &BeerRepository{
		DB: db,
	}
}

func (r *BeerRepository) GetAll(ctx context.Context) ([]*domain.Beer, error) {
	// Implementar lógica para obter todas as cervejas do banco de dados

	if err := r.DB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM beer")
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

	if beers == nil {
		logs.Info("Não há cervejas cadastradas na base de dados")
	}

	return beers, nil
}

func (r *BeerRepository) GetByID(ctx context.Context, id int) (*domain.Beer, error) {
	// Implementar lógica para obter uma cerveja pelo seu ID

	if err := r.DB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	row := r.DB.QueryRowContext(ctx, "SELECT * FROM beer WHERE id = $1", id)

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
	var typeID int
	errt := r.DB.QueryRowContext(ctx, "SELECT id FROM beer_type WHERE name = $1", beer.Type.Name).Scan(&typeID)
	if errt != nil {
		return nil, fmt.Errorf("error getting typeID: %w", errt)
	}

	var styleID int
	errs := r.DB.QueryRowContext(ctx, "SELECT id FROM beer_style WHERE name = $1", beer.Style.Name).Scan(&styleID)
	if errs != nil {
		return nil, fmt.Errorf("error getting typeID: %w", errs)
	}

	return beer, nil
}
