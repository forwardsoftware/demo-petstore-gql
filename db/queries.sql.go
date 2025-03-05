// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const petCreate = `-- name: PetCreate :one
INSERT INTO ps_pets (name, tags)
VALUES ($1, $2)
RETURNING id, name, tags
`

func (q *Queries) PetCreate(ctx context.Context, name string, tags []string) (PsPet, error) {
	row := q.db.QueryRowContext(ctx, petCreate, name, pq.Array(tags))
	var i PsPet
	err := row.Scan(&i.ID, &i.Name, pq.Array(&i.Tags))
	return i, err
}

const petGet = `-- name: PetGet :one
SELECT id, name, tags
FROM ps_pets
WHERE id = $1
`

func (q *Queries) PetGet(ctx context.Context, id int32) (PsPet, error) {
	row := q.db.QueryRowContext(ctx, petGet, id)
	var i PsPet
	err := row.Scan(&i.ID, &i.Name, pq.Array(&i.Tags))
	return i, err
}

const petsList = `-- name: PetsList :many
SELECT id, name, tags
FROM ps_pets
ORDER BY name
LIMIT $1
`

func (q *Queries) PetsList(ctx context.Context, limit int32) ([]PsPet, error) {
	rows, err := q.db.QueryContext(ctx, petsList, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PsPet{}
	for rows.Next() {
		var i PsPet
		if err := rows.Scan(&i.ID, &i.Name, pq.Array(&i.Tags)); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
