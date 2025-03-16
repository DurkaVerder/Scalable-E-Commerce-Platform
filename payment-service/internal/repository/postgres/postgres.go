package postgres

import (
	"database/sql"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/models"
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
