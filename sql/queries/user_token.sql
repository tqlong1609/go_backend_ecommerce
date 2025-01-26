-- name: GetRefreshToken :one
SELECT *
FROM user_token
WHERE refresh_token = $1;

-- name: DeleteRefreshToken :exec
DELETE FROM user_token
WHERE refresh_token = $1;

-- name: AddRefreshToken :exec
INSERT INTO user_token (
  user_id, refresh_token, expires_at
) VALUES (
  $1, $2, $3
)
ON CONFLICT (refresh_token) DO NOTHING;