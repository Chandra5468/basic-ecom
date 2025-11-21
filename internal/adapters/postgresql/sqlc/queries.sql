-- name: ListProducts :many
SELECT
*
FROM products;

-- name: FindProductByID :one
SELECT * FROM products where id = $1;