package postgres

const (
	createOrderQuery      = `INSERT INTO orders (user_id, amount) VALUES ($1, $2)`
	createOrderItemsQuery = `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	getOrdersQuery        = `SELECT * FROM orders WHERE user_id = $1`
	getOrderProductsQuery = `SELECT order_items.quantity, product.name, product.price FROM order_items
JOIN products ON order_items.product_id = products.id
WHERE orders_items.order_id = $1`
)
