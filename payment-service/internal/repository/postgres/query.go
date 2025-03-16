package postgres

const (
	AddPaymentQuery = `INSERT INTO payments (user_id, order_id, amount, payment_intent_id) VALUES ($1, $2, $3, $4)`
)
