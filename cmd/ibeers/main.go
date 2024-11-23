// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/d90ares/iBeers/api/handler"
	"github.com/d90ares/iBeers/api/middleware"
	"github.com/d90ares/iBeers/api/router"
	"github.com/d90ares/iBeers/internal/app/infra/database"
	"github.com/d90ares/iBeers/internal/app/repository"
	"github.com/d90ares/iBeers/internal/app/service"
	"github.com/d90ares/iBeers/internal/app/usecase"
	"github.com/d90ares/iBeers/pkg/art"
	"github.com/d90ares/iBeers/pkg/logs"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// content, err := os.ReadFile("ascii.txt")
	// if err != nil {
	// 	log.Fatalf("Erro ao ler o arquivo: %v", err)
	// }

	asciiArt := art.AsciiArt("OXOSIX")
	fmt.Print(asciiArt)
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

	staticDir := "./web"
	fileServer := http.FileServer(http.Dir(staticDir))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	n.UseHandler(r)

	// Start the server
	logs.Info("Serving static files from ./web")
	logs.Info("Listening on http://localhost:8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", n)
}
