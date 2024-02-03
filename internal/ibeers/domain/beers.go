// domain/beer.go
package domain

type Beer struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Style string `json:"style"`
}

type Repository interface {
	GetAll() ([]*Beer, error)
	// Adicionar outros métodos de repositório conforme necessário
}

type UseCase interface {
	GetAllBeers() ([]*Beer, error)
	// Adicionar outros métodos de caso de uso conforme necessário
}
