package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/d90ares/iBeers/pkg/logs"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	connStr  = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=beers sslmode=disable", host, port, user, password)
)

func NewPostgreSQLDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)

	// defer db.Close()

	if err != nil {
		log.Fatal("Error opening connections to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := RunMigrations(db); err != nil {
		logs.Error("Error on migrations: ", err)
		return nil, nil
	}

	if err := RunInitialData(db); err != nil {
		logs.Error("Error on initial data: ", err)
		return nil, nil
	}
	return db, err
}

func RunMigrations(db *sql.DB) error {

	cwd, _ := os.Getwd()
	// Leia o conteúdo do arquivo de migração
	migrationsPath := filepath.Join(cwd, "scripts", "migrations.sql")

	exists, err := verifyTablesExist(db)

	if err != nil {
		return fmt.Errorf("failed to check tables existence: %w", err)
	}
	if exists {
		// Os dados iniciais já existem, registre uma mensagem informativa
		logs.Info("Tables already exists")
		return nil
	}

	data, err := os.ReadFile(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Execute as migrações
	_, err = db.Exec(string(data))
	if err != nil {
		logs.Error("Erro ao executar migrações: ", err)
		return fmt.Errorf("failed to execute migrations: %w", err)
	}
	return nil
}

func RunInitialData(db *sql.DB) error {

	cwd, _ := os.Getwd()
	// Leia o conteúdo do arquivo de migração
	migrationsPath := filepath.Join(cwd, "scripts", "initial_data.sql")

	// Verificar se os dados iniciais já existem
	exists, err := checkInitialDataExists(db)
	if err != nil {
		return fmt.Errorf("failed to check initial data existence: %w", err)
	}
	if exists {
		// Os dados iniciais já existem, registre uma mensagem informativa
		logs.Info("Initial data already exists")
		return nil
	}

	data, err := os.ReadFile(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read initial data file: %w", err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		logs.Error("Erro ao executar dados iniciais: ", err)
		return fmt.Errorf("failed to execute initial data: %w", err)
	}
	return nil
}

func verifyTablesExist(db *sql.DB) (bool, error) {
	// Lista de tabelas a serem verificadas
	tables := []string{"beer_type", "beer_style", "beer"}
	var exists bool
	for _, table := range tables {
		// Verifica se a tabela existe no banco de dados
		query := fmt.Sprintf("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = '%s')", table)

		if err := db.QueryRow(query).Scan(&exists); err != nil {
			return false, fmt.Errorf("failed to check table existence: %w", err)
		}
	}
	return exists, nil
}

func checkInitialDataExists(db *sql.DB) (bool, error) {
	// Consulta para verificar a existência de dados nas tabelas beer_type e beer_style
	query := `SELECT EXISTS (SELECT 1 FROM beer_type) OR EXISTS (SELECT 1 FROM beer_style)`
	var exists bool
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
