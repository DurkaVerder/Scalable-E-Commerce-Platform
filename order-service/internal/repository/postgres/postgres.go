package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/models"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func ConnectDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Error connecting to database",
				Fields: map[string]interface{}{
					"method": "ConnectDB",
					"action": "connecting to database",
					"dsn":    dsn,
					"error":  err.Error(),
				},
			})
		panic(err)
	}
	return db
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) CreateOrder(userID int, amount float64, products []models.Product) error {
	orderId := -1
	err := p.db.QueryRow(createOrderQuery, userID, amount).Scan(&orderId)
	if err != nil {
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
			return err
		}
	}
	return nil
}

func (p *Postgres) GetOrders(userID int) ([]models.Order, error) {
	rows, err := p.db.Query(getOrdersQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.ID, &order.Amount, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (p *Postgres) GetOrderProducts(orderID int) ([]models.OrderProduct, error) {
	rows, err := p.db.Query(getOrderProductsQuery, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderProducts := make([]models.OrderProduct, 0)
	for rows.Next() {
		orderProduct := models.OrderProduct{}
		err := rows.Scan(&orderProduct.Quantity, &orderProduct.Name, &orderProduct.Price)
		if err != nil {
			return nil, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}

func (p *Postgres) UpdateOrder(orderID int, status string) error {
	_, err := p.db.Exec(updateOrderQuery, status, orderID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetOrder(orderID int) (models.Order, error) {
	order := models.Order{}
	err := p.db.QueryRow(getOrderQuery, orderID).Scan(&order.Amount, &order.Status, &order.CreatedAt)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
