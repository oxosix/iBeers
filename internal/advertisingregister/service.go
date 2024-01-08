package advertisingregister

import (
	"time"
)

type AdvertisingService struct {
	Repository AdvertisingRepository
	S3Service  AWSS3Service
}

func (s *AdvertisingService) Upload(advertisingDTO *AdvertisingDTO) (*Advertising, error) {
	advertising := &Advertising{
		Title:       advertisingDTO.Title,
		Description: advertisingDTO.Description,

		CreatedAt: time.Now(),
	}
}
