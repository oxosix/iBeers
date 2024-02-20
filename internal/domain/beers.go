// domain/beer.go
package domain

type Beer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Style struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"style"`
}
