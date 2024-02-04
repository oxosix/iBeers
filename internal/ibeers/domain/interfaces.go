package domain

type Repository interface {
	GetAll() ([]*Beer, error)
	// Adicionar outros métodos de repositório conforme necessário
}

type UseCase interface {
	GetAllBeers() ([]*Beer, error)
	// Adicionar outros métodos de caso de uso conforme necessário
}
