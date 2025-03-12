package postgres

const (
	getAllProductsQuery        = `SELECT id, name, price, description FROM products`
	getProductByIdQuery        = `SELECT id, name, price, description FROM products WHERE id = $1`
	getProductsByCategoryQuery = `SELECT products.id, products.name, products.price, products.description 
	FROM products
	JOIN product_categories ON products.id = product_categories.product_id
	JOIN categories ON product_categories.category_id = categories.id
	WHERE categories.name = $1`
)
