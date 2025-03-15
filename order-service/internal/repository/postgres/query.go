package postgres

const (
	createOrderQuery      = `INSERT INTO orders (user_id, amount) VALUES ($1, $2)`
	createOrderItemsQuery = `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	getOrdersQuery        = `SELECT id, amount, status, created_at FROM orders WHERE user_id = $1`
	getOrderProductsQuery = `SELECT order_items.quantity, product.name, product.price FROM order_items
JOIN products ON order_items.product_id = products.id
WHERE orders_items.order_id = $1`
	getOrderQuery = `SELECT amount, status, created_at FROM orders WHERE id = $1`
	updateOrderQuery = `UPDATE orders SET status = $1 WHERE id = $2`
)
