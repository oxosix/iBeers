package domain

type Reader interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Search(b string) (*Beer, error)
	GetBeerStyles(b string) ([]*BeerStyle, error)
	GetBeerTypes(b string) ([]*BeerType, error)
}

type Writer interface {
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

type Repository interface {
	Reader
	Writer
}

type Manager interface {
	Repository
}
