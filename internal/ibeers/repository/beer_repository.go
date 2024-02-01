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

func (r *BeerRepository) GetAllBeers(ctx context.Context) ([]*domain.Beer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rows, err := r.db.Query(ctx, "SELECT id, name, beer_type, beer_style FROM beers")
	if err != nil {
		log.Printf("Error querying beers: %v", err)
		return nil, err
	}
	defer rows.Close()

	var beers []*domain.Beer
	for rows.Next() {
		var beer domain.Beer
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)
		if err != nil {
			log.Printf("Error scanning beer row: %v", err)
			return nil, err
		}
		beers = append(beers, &beer)
	}

	return beers, nil
}

func (r *BeerRepository) Get(ID int64) (*domain.Beer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var beer domain.Beer
	err := r.db.QueryRow(context.Background(), "SELECT id, name, beer_type, beer_style FROM beers WHERE id = $1", ID).
		Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)

	switch {
	case err == pgx.ErrNoRows:
		return nil, nil // Beer not found
	case err != nil:
		log.Printf("Error querying beer by ID: %v", err)
		return nil, err
	}

	return &beer, nil
}

func (r *BeerRepository) Search(b string) (*domain.Beer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var beer domain.Beer
	err := r.db.QueryRow(context.Background(), "SELECT id, name, beer_type, beer_style FROM beers WHERE name = $1", b).
		Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)

	switch {
	case err == pgx.ErrNoRows:
		return nil, nil // Beer not found
	case err != nil:
		log.Printf("Error searching beer by name: %v", err)
		return nil, err
	}

	return &beer, nil
}

func (r *BeerRepository) Store(beer *domain.Beer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(context.Background(), "INSERT INTO beers (name, beer_type, beer_style) VALUES ($1, $2, $3)",
		beer.Name, beer.Type, beer.Style)

	if err != nil {
		log.Printf("Error storing beer: %v", err)
		return err
	}

	return nil
}

func (r *BeerRepository) GetBeerStyles() ([]*domain.BeerStyle, error) {
	var styles []*domain.BeerStyle

	rows, err := r.db.Query(context.Background(), "SELECT id, name FROM beer_styles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var style domain.BeerStyle
		if err := rows.Scan(&style.ID, &style.Name); err != nil {
			return nil, err
		}
		styles = append(styles, &style)
	}

	return styles, nil
}
