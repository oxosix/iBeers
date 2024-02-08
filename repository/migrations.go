package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/d90ares/iBeers/config/logs"
)

func RunMigrations(db *sql.DB) error {

	cwd, _ := os.Getwd()
	// Leia o conteúdo do arquivo de migração
	migrationsPath := filepath.Join(cwd, "../..", "scripts", "migrations.sql")
	fmt.Println("Caminho do arquivo de migração:", migrationsPath)
	data, err := os.ReadFile(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Execute as migrações
	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("failed to execute migrations: %w", err)
	}
	logs.Error("Erro ao executar migrações migrações: ", err)
	return nil
}
