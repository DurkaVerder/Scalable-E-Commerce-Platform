package models

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID       int     `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderProduct struct {
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}
