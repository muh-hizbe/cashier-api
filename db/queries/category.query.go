package queries

var INSERT_CATEGORY = `
INSERT INTO categories (name, description)
VALUES ($1, $2)
RETURNING *;
`

var GET_CATEGORIES = `
	SELECT * FROM categories
`

var GET_CATEGORY_BY_ID = `
	SELECT * FROM categories WHERE id = $1
`

var UPDATE_CATEGORY = `
UPDATE categories
SET
	name = COALESCE($2, name),
	description = COALESCE($3, description)
WHERE id = $1
RETURNING *;
`

var DELETE_CATEGORY = `
DELETE FROM categories
WHERE id = $1
RETURNING *;
`
