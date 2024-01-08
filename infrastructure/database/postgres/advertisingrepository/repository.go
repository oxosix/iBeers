package advertisingrepository

import (
	"context"
	"fmt"

	"github.com/d90ares/gloisp-web/internal/advertisingregister"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgreSQLRepository struct {
	db *gorm.DB
}

// New returns contract implementation of ProductRepository
func New(db *gorm.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		db: db,
	}
}

func (r *PostgreSQLRepository) Save(advertising *advertisingregister.Advertising) error {
	if advertising.ID == uuid.Nil {
		advertising.ID = uuid.New()
	}
	if err := r.db.Create(advertising).Error; err != nil {
		return fmt.Errorf("falha ao inserir no banco de dados: %v", err)
	}
	return nil
}

func (r *PostgreSQLRepository) GetCategories() ([]string, error) {
	var categories []string
	if err := r.db.Raw("SELECT DISTINCT category FROM advertising").Scan(&categories).Error; err != nil {
		return nil, fmt.Errorf("falha ao obter categorias do banco de dados: %v", err)
	}
	return categories, nil
}

func (r *PostgreSQLRepository) SeedCategory(ctx context.Context) error {
	categories := []string{"academias",
		"acessorios",
		"brinquedos",
		"acougue",
		"administracao",
		"advocacia",
		"alimentacao",
		"animais",
		"autopecas",
		"bar",
		"bebidas",
		"beleza",
		"borracharia",
		"comunicacao",
		"contrucao",
		"contabilidade",
		"curriculos",
		"educacao",
		"eletronicos",
		"esportes",
		"eventos",
		"farmacia",
		"funilaria",
		"gasagua",
		"grafica",
		"igreja",
		"imobiliaria",
		"industria",
		"informatica",
		"lanchonete",
		"limpeza",
		"loterica",
		"manutencao",
		"mecanina",
		"mercado",
		"metalurgica",
		"moda",
		"moradia",
		"moveis",
		"naturais",
		"odontologia",
		"ofertaempregos",
		"padaria",
		"papelaria",
		"pizzaria",
		"restaurante",
		"revestimentos",
		"saude",
		"seguranca",
		"serralheria",
		"servicos",
		"som",
		"transportes",
		"variedades",
		"veiculos",
		"vestuario",
		"vidracaria",
		"outros"}

	for _, category := range categories {
		exists, err := r.CategoryExists(ctx, category)
		if err != nil {
			return fmt.Errorf("falha ao verificar a existência da categoria: %v", err)
		}
		if !exists {
			if err := r.db.WithContext(ctx).Exec("INSERT INTO advertising (category) VALUES (?)", category).Error; err != nil {
				return fmt.Errorf("falha ao inserir categoria no banco de dados: %v", err)
			}
		}
	}
	return nil
}

func (r *PostgreSQLRepository) CategoryExists(ctx context.Context, category string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&advertisingregister.Advertising{}).Where("category = ?", category).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("falha ao verificar existência da categoria: %v", err)
	}
	return count > 0, nil
}
