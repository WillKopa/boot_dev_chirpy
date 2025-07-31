-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetChirpsASC :many
SELECT * FROM chirps
WHERE $1::uuid = '00000000-0000-0000-0000-000000000000'::uuid OR user_id = $1::uuid
ORDER BY created_at ASC;

-- name: GetChirpsDESC :many
SELECT * FROM chirps
WHERE $1::uuid = '00000000-0000-0000-0000-000000000000'::uuid OR user_id = $1::uuid
ORDER BY created_at DESC;

-- name: GetSingleChirp :one
SELECT * from chirps
WHERE id = $1;

-- name: DeleteChirp :exec
DELETE FROM chirps
WHERE id = $1 AND user_id = $2;