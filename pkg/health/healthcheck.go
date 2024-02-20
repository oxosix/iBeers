package health

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// HealthCheckHandler é um handler HTTP para o endpoint de health check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Aqui você pode adicionar lógica para verificar o estado da aplicação
	// Por exemplo, verificar a conexão com banco de dados, integridade de serviços, etc.
	checkDatabase(db)
	// Resposta JSON indicando que a aplicação está saudável
	healthStatus := map[string]string{"status": "UP", "version": "1"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthStatus)
}

func checkDatabase(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("falha ao pingar o banco de dados: %s", err)
	}
	return nil
}
