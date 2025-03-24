package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/models"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func ConnectDB(dns string) *sql.DB {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to connect to database",
			Fields: map[string]interface{}{
				"method": "ConnectDB",
				"action": "connect_to_db",
				"error":  err.Error(),
			},
		})
		panic(err)
	}

	return db
}

func (p *Postgres) AddPayment(order models.Order, paymentIntendId string) error {
	_, err := p.db.Exec(AddPaymentQuery, order.UserID, order.ID, order.Amount, paymentIntendId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
