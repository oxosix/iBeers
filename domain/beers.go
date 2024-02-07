// domain/beer.go
package domain

type Beer struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Style string `json:"style"`
}
