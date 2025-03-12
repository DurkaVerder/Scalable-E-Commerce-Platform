package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/models"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/pkg/logs"
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
		elk.Log.Error("Error opening database connection", map[string]interface{}{
			"method":           "ConnectDB",
			"action":           "opening database connection",
			"connectionString": connectionString,
			"error":            err.Error(),
		})
		return nil, err
	}

	return db, nil
}

func (p *Postgres) GetAllProducts() ([]models.Product, error) {

	rows, err := p.db.Query(getAllProductsQuery)
	if err != nil {
		elk.Log.Error("Error getting all products", map[string]interface{}{
			"method": "GetAllProducts",
			"action": "getting all products",
			"error":  err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
		if err != nil {
			elk.Log.Error("Error scanning rows", map[string]interface{}{
				"method": "GetAllProducts",
				"action": "scanning rows",
				"error":  err.Error(),
			})
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
		elk.Log.Error("Error getting product by id", map[string]interface{}{
			"method": "GetProductById",
			"action": "getting product by id",
			"error":  err.Error(),
		})
		return models.Product{}, err
	}

	return product, nil
}

func (p *Postgres) GetProductsByCategory(category string) ([]models.Product, error) {
	rows, err := p.db.Query(getProductsByCategoryQuery, category)
	if err != nil {
		elk.Log.Error("Error getting products by category", map[string]interface{}{
			"method":   "GetProductsByCategory",
			"action":   "getting products by category",
			"category": category,
			"error":    err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
		if err != nil {
			elk.Log.Error("Error scanning rows", map[string]interface{}{
				"method": "GetProductsByCategory",
				"action": "scanning rows",
				"error":  err.Error(),
			})

			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
