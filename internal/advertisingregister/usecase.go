package advertisingregister

type AdvertisingUseCase interface {
	Upload(advertisingDTO *AdvertisingDTO) (*Advertising, error)
	GetCategories() ([]string, error)
}
