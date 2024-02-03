// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/d90ares/iBeers/internal/ibeers/http/handler"
	"github.com/d90ares/iBeers/internal/ibeers/repository"
	"github.com/d90ares/iBeers/internal/ibeers/usecase"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const dbURL = "host=localhost port=5432 user=postgres dbname=ibeers sslmode=disable"

func main() {
	// Abrir conexão com o banco de dados PostgreSQL
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}
	defer db.Close()

	beerRepository := repository.NewBeerRepository(db)
	beerUseCase := usecase.NewBeerUseCase(beerRepository)
	beerHandler := handler.NewBeerHandler(beerUseCase)

	router := mux.NewRouter()
	router.HandleFunc("/beers", beerHandler.GetAllBeers).Methods("GET")

	// Start the server
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
