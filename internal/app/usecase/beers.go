// usecase/beer_usecase.go
package usecase

import (
	"github.com/d90ares/iBeers/internal/app/repository"
)

type UseCase interface {
	repository.Repository
}
