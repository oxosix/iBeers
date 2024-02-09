// domain/beer.go
package domain

type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type BeerStyle struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
