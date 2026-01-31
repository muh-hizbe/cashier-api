package queries

var INSERT_PRODUCT = `
INSERT INTO products (name, price, stock, category_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
`

var GET_PRODUCTS = `
	SELECT * FROM products
`

var GET_PRODUCT_BY_ID = `
	SELECT * FROM products WHERE id = $1
`

var GET_PRODUCTS_WITH_CATEGORY = `
	SELECT p.*, to_jsonb(c.*) AS category FROM products p LEFT JOIN categories c ON p.category_id = c.id
`

var GET_PRODUCT_BY_ID_WITH_CATEGORY = `
	SELECT p.*, to_jsonb(c.*) AS category FROM products p LEFT JOIN categories c ON p.category_id = c.id WHERE p.id = $1
`

var UPDATE_PRODUCT = `
UPDATE products
SET
	name = COALESCE($2, name),
	price = COALESCE($3, price),
	stock = COALESCE($4, stock),
	category_id = COALESCE($5, category_id)
WHERE id = $1
RETURNING *;
`

var DELETE_PRODUCT = `
DELETE FROM products
WHERE id = $1
RETURNING *;
`
