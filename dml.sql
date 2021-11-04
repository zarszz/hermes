/** PRODUCT CONTEXT **/
INSERT INTO product(sku, name, display) VALUES($1, $2, $3)

SELECT * FROM product
SELECT * FROM product WHERE = your_filter

UPDATE product SET sku = $1, name = $2, display = $3 WHERE id = $4

DELETE FROM product WHERE id = $1

/** USER CONTEXT **/
INSERT INTO users (name, email, password) VALUES ($1, $2, $3)

SELECT * FROM users WHERE email = $1
SELECT * FROM users
SELECT * FROM users WHERE id = $1

UPDATE users SET name = $1, email = $2 WHERE id = $3

DELETE FROM users WHERE id = $1
