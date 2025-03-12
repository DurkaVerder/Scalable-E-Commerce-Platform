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

	rows, err := p.db.Query(getAllProductsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
		if err != nil {
			return nil, err

		}
		products = append(products, product)
	}

	return products, nil
}

func (p *Postgres) GetProductById(id int) (models.Product, error) {
	var product models.Product
	err := p.db.QueryRow(getProductByIdQuery, id).Scan(&product.ID, &product.Name, &product.Price, &product.Description)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *Postgres) GetProductsByCategory(category string) ([]models.Product, error) {
	rows, err := p.db.Query(getProductsByCategoryQuery, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
