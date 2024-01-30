package domain

type Beer struct {
	ID    int64      `json:"id"`
	Name  string     `json:"name"`
	Type  *BeerType  `json:"type" db:"type"`
	Style *BeerStyle `json:"style" db:"style"`
}

type BeerType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type BeerStyle struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
