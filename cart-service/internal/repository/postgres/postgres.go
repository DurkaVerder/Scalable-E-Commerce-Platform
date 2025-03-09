// Package: postgres is a repository for working with PostgreSQL database. It contains the implementation of the repository interface.
// The repository uses the database/sql package to interact with the PostgreSQL database.
package postgres

import (
	elk "cart-service/pkg/logs"
	"cart-service/internal/models"
	"database/sql"

	_ "github.com/lib/pq"
)

// Postgres is a repository for working with PostgreSQL database.
type Postgres struct {
	db *sql.DB
}

// NewPostgres creates a new Postgres repository.
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

// ConnectDB connects to the PostgreSQL database.
func ConnectDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		elk.Log.Error("Failed to connect to the database: "+err.Error(), map[string]interface{}{
			"method": "ConnectDB",
			"action": "connecting to the database",
			"dbURL":  dbURL,
		})
		return nil, err
	}
	return db, nil
}

// GetCart returns a list of products in the cart of the user with the given ID.
func (p *Postgres) GetCart(userID int) ([]models.Product, error) {
	products := make([]models.Product, 0, 1)

	rows, err := p.db.Query(getCartQuery, userID)
	if err != nil {
		elk.Log.Error("Failed to get the cart: "+err.Error(), map[string]interface{}{
			"method": "GetCart",
			"action": "getting the cart",
			"userID": userID,
		})
		return nil, err
	}

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Quantity)
		if err != nil {
			elk.Log.Error("Failed to scan the cart row: "+err.Error(), map[string]interface{}{
				"method": "GetCart",
				"action": "scanning the cart row",
			})
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// AddProductToCart adds a product with the given ID and quantity to the cart of the user with the given ID.
func (p *Postgres) AddProductToCart(userID, productID, quantity int) error {
	cartID, err := p.getCartId(userID)
	if err != nil {
		elk.Log.Error("Failed to get the cart ID: "+err.Error(), map[string]interface{}{
			"method": "AddProductToCart",
			"action": "getting the cart ID",
			"userID": userID,
		})

		return err
	}

	_, err = p.db.Exec(addProductToCartQuery, cartID, productID, quantity)
	if err != nil {
		elk.Log.Error("Failed to add the product to the cart: "+err.Error(), map[string]interface{}{
			"method":    "AddProductToCart",
			"action":    "adding the product to the cart",
			"cartID":    cartID,
			"productID": productID,
			"quantity":  quantity,
		})

		return err
	}

	return nil
}

// DeleteProductFromCart deletes a product with the given ID from the cart of the user with the given ID.
func (p *Postgres) DeleteProductFromCart(userID, productID int) error {
	cartID, err := p.getCartId(userID)
	if err != nil {
		elk.Log.Error("Failed to get the cart ID: "+err.Error(), map[string]interface{}{
			"method":    "De;eteProductFromCart",
			"action":    "getting the cart ID",
			"userID":    userID,
			"productID": productID,
		})
		return err
	}

	_, err = p.db.Exec(deleteProductFromCartQuery, cartID, productID)
	if err != nil {
		elk.Log.Error("Failed to delete the product from the cart: "+err.Error(), map[string]interface{}{
			"method":    "DeleteProductFromCart",
			"action":    "deleting the product from the cart",
			"cartID":    cartID,
			"productID": productID,
		})

		return err
	}

	return nil
}

// UpdateProductQuantity updates the quantity of a product with the given ID in the cart of the user with the given ID.
func (p *Postgres) UpdateProductQuantity(userID, productID, quantity int) error {
	cartID, err := p.getCartId(userID)
	if err != nil {
		elk.Log.Error("Failed to get the cart ID: "+err.Error(), map[string]interface{}{
			"method":    "UpdateProductQuantity",
			"action":    "getting the cart ID",
			"userID":    userID,
			"productID": productID,
			"quantity":  quantity,
		})
		return err
	}

	_, err = p.db.Exec(updateProductQuantityQuery, cartID, productID, quantity)
	if err != nil {
		elk.Log.Error("Failed to update the product quantity: "+err.Error(), map[string]interface{}{
			"method":    "UpdateProductQuantity",
			"action":    "updating the product quantity",
			"cartID":    cartID,
			"productID": productID,
			"quantity":  quantity,
		})
		return err
	}

	return nil
}

// getCartId returns the ID of the cart of the user with the given ID.
func (p *Postgres) getCartId(userID int) (int, error) {
	var cartID int
	err := p.db.QueryRow(getCartIDQuery, userID).Scan(&cartID)
	if err != nil {
		elk.Log.Error("Failed to get the cart ID: "+err.Error(), map[string]interface{}{
			"method": "getCartId",
			"action": "getting the cart ID",
			"userID": userID,
		})

		return -1, err
	}
	return cartID, nil
}
