-- name: PetsList :many
SELECT *
FROM ps_pets
ORDER BY name
LIMIT $1;

-- name: PetCreate :one
INSERT INTO ps_pets (name, tags)
VALUES ($1, $2)
RETURNING *;

-- name: PetGet :one
SELECT *
FROM ps_pets
WHERE id = $1;
