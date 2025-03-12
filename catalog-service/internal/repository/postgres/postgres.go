package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/models"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func ConnectDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func (p *Postgres) GetAllProducts() ([]models.Product, error) {
	return nil, nil
}

func (p *Postgres) GetProductById(id int) (models.Product, error) {
	return models.Product{}, nil
}

func (p *Postgres) GetProductsByCategory(category string) ([]models.Product, error) {
	return nil, nil
}
