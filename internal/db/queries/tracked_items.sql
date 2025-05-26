-- name: ListItems :many
SELECT id, name, query, filters AS steam_params, sent_ids, active, created_at
FROM tracked_items
WHERE active = true;

-- name: GetItemByID :one
SELECT id, name, query, filters AS steam_params, sent_ids, active, created_at
FROM tracked_items
WHERE id = $1;

-- name: CreateItem :one
INSERT INTO tracked_items (name, query, filters)
VALUES ($1, $2, $3)
RETURNING id, name, query, filters AS steam_params, sent_ids, active, created_at;

-- name: UpdateItem :one
UPDATE tracked_items
SET name = $2,
    query = $3,
    filters = $4,
    active = $5
WHERE id = $1
RETURNING id, name, query, filters AS steam_params, sent_ids, active, created_at;

-- name: DeleteItem :exec
DELETE FROM tracked_items
WHERE id = $1;
