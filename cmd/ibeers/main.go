// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/d90ares/iBeers/api/handler"
	"github.com/d90ares/iBeers/api/middleware"
	"github.com/d90ares/iBeers/api/router"
	"github.com/d90ares/iBeers/internal/app/infra/database"
	"github.com/d90ares/iBeers/internal/app/repository"
	"github.com/d90ares/iBeers/internal/app/service"
	"github.com/d90ares/iBeers/internal/app/usecase"
	"github.com/d90ares/iBeers/pkg/logs"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	content, err := os.ReadFile("ascii.txt")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	asciiArt := string(content)
	fmt.Println(color.MagentaString(asciiArt) + "\n")
	logs.Info("About to Start Application")

	db, err := database.NewPostgreSQLDB()
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	beerRepository := repository.NewBeerRepository(db)
	beerService := service.NewBeerService(beerRepository)
	beerUseCase := usecase.NewBeerUseCase(beerService)
	beerHandler := handler.NewBeerHandler(beerUseCase)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(middleware.NewMiddleware())

	r := mux.NewRouter()
	router.SetupRoutes(r, beerHandler)
	router.SetupMetricsRoutes(r) // Configuração das rotas de métricas
	router.SetupHealthRoute(r, db)

	n.UseHandler(r)

	// Start the server
	logs.Info("Listening on http://localhost:8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", n)
}
