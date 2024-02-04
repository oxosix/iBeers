// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/d90ares/iBeers/internal/ibeers/http/handler"
	"github.com/d90ares/iBeers/internal/ibeers/http/router"
	"github.com/d90ares/iBeers/internal/ibeers/repository"
	"github.com/d90ares/iBeers/internal/ibeers/service"
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

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	beerRepository := repository.NewBeerRepository(db)
	beerService := service.NewBeerService(beerRepository)
	beerUseCase := usecase.NewBeerUseCase(beerService)
	beerHandler := handler.NewBeerHandler(beerUseCase)

	r := mux.NewRouter()
	router.SetupRoutes(r, beerHandler)

	// Start the server
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
