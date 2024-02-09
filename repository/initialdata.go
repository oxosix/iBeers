package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/d90ares/iBeers/config/logs"
)

func RunInitialData(db *sql.DB) error {

	cwd, _ := os.Getwd()
	// Leia o conteúdo do arquivo de migração
	migrationsPath := filepath.Join(cwd, "../..", "scripts", "initial_data.sql")
	fmt.Println("Caminho do arquivo de migração:", migrationsPath)
	data, err := os.ReadFile(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read initial data file: %w", err)
	}

	// Execute as migrações
	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("failed to execute initial data: %w", err)
	}
	logs.Error("Erro ao executar dados iniciais: ", err)
	return nil
}
