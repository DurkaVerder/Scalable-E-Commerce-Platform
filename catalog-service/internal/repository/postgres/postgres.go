package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/models"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func ConnectDB(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Failed to connect to database",
				Fields: map[string]interface{}{
					"method": "ConnectDB",
					"action": "Open",
					"error":  err,
				},
			})
		panic(err)
	}

	return db
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

func (p *Postgres) Close() {
	p.db.Close()
}
