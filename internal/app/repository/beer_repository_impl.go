package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/d90ares/iBeers/internal/domain"
	"github.com/d90ares/iBeers/pkg/logs"
	"go.uber.org/zap"
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
		logs.Error("Connection error:", err)
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, `
		SELECT b.id, b.name, t.id AS type_id, t.name AS type_name, s.id AS style_id, s.name AS style_name
		FROM beer b
		JOIN beer_type t ON b.type_id = t.id
		JOIN beer_style s ON b.style_id = s.id
	`)
	if err != nil {
		logs.Error("Error executing query:", err)
		return nil, err
	}

	// defer rows.Close()

	var beers []*domain.Beer
	for rows.Next() {
		var b domain.Beer
		if err := rows.Scan(&b.ID, &b.Name, &b.Type.ID, &b.Type.Name, &b.Style.ID, &b.Style.Name); err != nil {
			return nil, fmt.Errorf("failed to mapping data: %w", err)
		}

		beers = append(beers, &b)

	}

	if beers == nil {
		logs.Info("Não há cervejas cadastradas na base de dados")
	}

	// beersJSON, err := json.Marshal(beers)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter slice para JSON: %v", err)
	}

	sugars := logs.Sugar()
	beersField := zap.Any("returnedBeers", beers)
	sugars.Console.Infow("Sucesso", beersField)
	sugars.JSON.Infow("Sucesso", beersField)

	return beers, nil
}

func (r *BeerRepository) GetByID(ctx context.Context, id int) (*domain.Beer, error) {
	// Implementar lógica para obter uma cerveja pelo seu ID

	if err := r.DB.PingContext(ctx); err != nil {
		logs.Error("erro ao conectar ao banco de dados: %w", err)
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
	var typeID int64
	errt := r.DB.QueryRowContext(ctx, "SELECT id FROM beer_type WHERE name = $1", beer.Type.Name).Scan(&typeID)
	if errt != nil {
		return nil, fmt.Errorf("error getting typeID: %w", errt)
	}

	var styleID int64
	errs := r.DB.QueryRowContext(ctx, "SELECT id FROM beer_style WHERE name = $1", beer.Style.Name).Scan(&styleID)
	if errs != nil {
		return nil, fmt.Errorf("error getting typeID: %w", errs)
	}

	_, err := r.DB.ExecContext(ctx, "INSERT INTO beers (name, type_id, style_id) VALUES ($1, $2, $3)", beer.Name, typeID, styleID)
	if err != nil {
		return nil, err
	}

	return beer, nil
}
