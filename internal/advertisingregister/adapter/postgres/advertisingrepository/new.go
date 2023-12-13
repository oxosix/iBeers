package advertisingrepository

import (
	"github.com/d90ares/gloisp-web/internal/advertising-register/adapter/postgres"
	"github.com/d90ares/gloisp-web/internal/advertising-register/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.AdvertisingRepository {
	return &repository{
		db: db,
	}
}
