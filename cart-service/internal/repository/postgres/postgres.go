package postgres

import (
	"cart-service/internal/models"
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) GetCart(userID int) ([]models.Product, error) {
	return nil, nil
}

func (p *Postgres) AddProductToCart(userID, productID, quantity int) error {
	return nil
}

func (p *Postgres) DeleteProductFromCart(userID, productID int) error {
	return nil
}

func (p *Postgres) UpdateProductQuantity(userID, productID, quantity int) error {
	return nil
}
