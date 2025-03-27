package models

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Products  []OrderProduct
}

type Product struct {
	ID       int     `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderProduct struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Status struct {
	Status string `json:"status"`
}

type Notification struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
