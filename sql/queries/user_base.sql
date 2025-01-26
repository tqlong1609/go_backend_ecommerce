-- name: FindUserByEmail :one
SELECT *
FROM user_base
WHERE user_account = $1;

-- name: AddUserInfoBase :one
INSERT INTO user_base (
  user_account, user_password, user_salt, user_login_time, user_logout_time, user_login_ip, user_created_at, user_updated_at
) VALUES (
  $1, $2, $3, NULL, NULL, NULL, NOW(), NOW()
)
RETURNING *;

-- name: UpdateUserInfoBase :exec
UPDATE user_base
SET 
  user_salt = $2,
  user_login_time = $3,
  user_logout_time = $4,
  user_login_ip = $5,
  user_updated_at = NOW()
WHERE user_id = $1;

-- name: FindUserByIdBase :one
SELECT *
FROM user_base
WHERE user_id = $1;

-- name: ChangePasswordUserBase :exec
UPDATE user_base
SET 
  user_password = $2,
  user_updated_at = NOW()
WHERE user_id = $1;
