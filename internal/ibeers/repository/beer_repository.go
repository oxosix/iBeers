package repository

import (
	"context"
	"log"
	"sync"

	"github.com/d90ares/iBeers/internal/ibeers/domain"
	"github.com/jackc/pgx/v4"
)

type BeerRepository struct {
	mu sync.Mutex
	db *pgx.Conn
}

func NewBeerRepository(db *pgx.Conn) *BeerRepository {
	return &BeerRepository{
		db: db,
	}
}

func (r *BeerRepository) GetAllBeers(ctx context.Context) ([]domain.Beer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rows, err := r.db.Query(ctx, "SELECT id, name, beer_type, beer_style FROM beers")
	if err != nil {
		log.Printf("Error querying beers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var beers []domain.Beer
	for rows.Next() {
		var beer domain.Beer
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
		if err != nil {
			log.Printf("Error scanning beer row: %v", err)
			return nil, err
		}
		beers = append(beers, beer)
	}

	return beers, nil
}
