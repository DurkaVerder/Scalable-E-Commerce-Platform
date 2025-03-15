package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/models"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/pkg/logs"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) CreateOrder(userID int, amount float64, products []models.Product) error {
	orderId := -1
	err := p.db.QueryRow(createOrderQuery, userID, amount).Scan(&orderId)
	if err != nil {
		elk.Log.Error("Error while creating order", map[string]interface{}{
			"method": "CreateOrder",
			"action": "queryRow",
			"error":  err.Error(),
		})
		return err
	}

	if err := p.createOrderItems(orderId, products); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) createOrderItems(orderId int, products []models.Product) error {
	for _, product := range products {
		_, err := p.db.Exec(createOrderItemsQuery, orderId, product.ID, product.Quantity)
		if err != nil {
			elk.Log.Error("Error while creating order items", map[string]interface{}{
				"method":   "createOrderItems",
				"action":   "exec",
				"order_id": orderId,
				"product":  product,
				"error":    err.Error(),
			})
			return err
		}
	}
	return nil
}

func (p *Postgres) GetOrders(userID int) ([]models.Order, error) {
	rows, err := p.db.Query(getOrdersQuery, userID)
	if err != nil {
		elk.Log.Error("Error while getting orders", map[string]interface{}{
			"method": "GetOrders",
			"action": "query",
			"error":  err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.ID, &order.Amount, &order.Status, &order.CreatedAt)
		if err != nil {
			elk.Log.Error("Error while scanning orders", map[string]interface{}{
				"method": "GetOrders",
				"action": "scan",
				"error":  err.Error(),
			})
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (p *Postgres) GetOrderProducts(orderID int) ([]models.OrderProduct, error) {
	rows, err := p.db.Query(getOrderProductsQuery, orderID)
	if err != nil {
		elk.Log.Error("Error while getting order products", map[string]interface{}{
			"method":   "GetOrderProducts",
			"action":   "query",
			"order_id": orderID,
			"error":    err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	orderProducts := make([]models.OrderProduct, 0)
	for rows.Next() {
		orderProduct := models.OrderProduct{}
		err := rows.Scan(&orderProduct.Quantity, &orderProduct.Name, &orderProduct.Price)
		if err != nil {
			elk.Log.Error("Error while scanning order products", map[string]interface{}{
				"method": "GetOrderProducts",
				"action": "scan",
				"error":  err.Error(),
			})
			return nil, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}

func (p *Postgres) UpdateOrder(orderID int, status string) error {
	_, err := p.db.Exec(updateOrderQuery, status, orderID)
	if err != nil {
		elk.Log.Error("Error while updating order", map[string]interface{}{
			"method":   "UpdateOrder",
			"action":   "exec",
			"order_id": orderID,
			"error":    err.Error(),
		})
		return err
	}
	return nil
}

func (p *Postgres) GetOrder(orderID int) (models.Order, error) {
	order := models.Order{}
	err := p.db.QueryRow(getOrderQuery, orderID).Scan(&order.Amount, &order.Status, &order.CreatedAt)
	if err != nil {
		elk.Log.Error("Error while getting order", map[string]interface{}{
			"method":   "GetOrder",
			"action":   "queryRow",
			"order_id": orderID,
			"error":    err.Error(),
		})
		return models.Order{}, err
	}
	return order, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
