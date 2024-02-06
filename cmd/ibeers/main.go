// main.go
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/d90ares/iBeers/config/logs"
	"github.com/d90ares/iBeers/internal/ibeers/http/handler"
	"github.com/d90ares/iBeers/internal/ibeers/http/middleware"
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

	logs.Info("About to Start Application")
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Error opening connections to database: ", err)
	}
	defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	log.Fatal("Error connecting to database: ", err)
	// }

	beerRepository := repository.NewBeerRepository(db)
	beerService := service.NewBeerService(beerRepository)
	beerUseCase := usecase.NewBeerUseCase(beerService)
	beerHandler := handler.NewBeerHandler(beerUseCase)

	// Crie uma nova instância do Negroni
	n := negroni.New()
	// Adicione o middleware de recuperação (recovery)
	n.Use(negroni.NewRecovery())

	// Adicione o middleware personalizado para tratamento de erros
	n.Use(middleware.NewMiddleware())

	r := mux.NewRouter()
	router.SetupRoutes(r, beerHandler)

	n.UseHandler(r)

	// Start the server
	logs.Info("Running in http://localhost:8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", n)
}
