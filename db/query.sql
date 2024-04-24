-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ?
);

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;

-- name: ItemGet :one
SELECT * FROM items
WHERE id = ? LIMIT 1;

-- name: NpcGet :one
SELECT * FROM npc_types
WHERE id = ? LIMIT 1;

-- name: PlayerGet :one
SELECT * FROM character_data
WHERE id = ? LIMIT 1;