package advertisingregister

type AdvertisingRepository interface {
	Save(advertising *Advertising)
	GetCategories() ([]string, error)
}
