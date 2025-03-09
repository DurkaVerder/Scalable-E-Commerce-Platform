package postgres

// Queries
const (
	
	getCartQuery = `SELECT products.id, products.name, products.price, products.description, cart_item.quantity 
	FROM cart_item
	JOIN products ON cart_item.product_id = products.id
	JOIN carts ON cart_item.cart_id = carts.id
	WHERE carts.user_id = $1
	ORDER BY cart_item.created_at DESC`

	addProductToCartQuery = `INSERT INTO cart_item (cart_id, product_id, quantity) VALUES ($1, $2, $3)`

	deleteProductFromCartQuery = `DELETE FROM cart_item WHERE cart_id = $1 AND product_id = $2`

	updateProductQuantityQuery = `UPDATE cart_item SET quantity = $3 WHERE cart_id = $1 AND product_id = $2`

	getCartIDQuery = `SELECT id FROM carts WHERE user_id = $1`
)
