-- name: FindUserByEmail :one
SELECT *
FROM user_base
WHERE user_account = $1;

-- name: AddUserInfo :one
INSERT INTO user_base (
  user_account, user_password, user_sait, user_login_time, user_logout_time, user_login_ip, user_created_at, user_updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, NOW(), NOW()
)
RETURNING *;

-- name: UpdateUserInfo :exec
UPDATE user_base
SET 
  user_sait = $2,
  user_login_time = $3,
  user_logout_time = $4,
  user_login_ip = $5,
  user_updated_at = NOW()
WHERE user_id = $1;

-- name: FindUserById :one
SELECT *
FROM user_base
WHERE user_id = $1;

-- name: ChangePasswordUser :exec
UPDATE user_base
SET 
  user_password = $2,
  user_updated_at = NOW()
WHERE user_id = $1;
