-- name: InsertUserInfo :one
INSERT INTO user_info (
  user_id, user_account, user_name, user_state, user_is_auth, create_at, updated_at
) VALUES (
  $1, $2, $3, 1, 1, NOW(), NOW()
)
RETURNING *;

-- name: UpdateUserInfo :exec
UPDATE user_info
SET 
  user_name = $2,
  user_avatar = $3,
  user_state = $4,
  user_phone = $5,
  user_gender = $6,
  user_birthday = $7,
  user_email = $8,
  user_is_auth = $9,
  updated_at = NOW()
WHERE user_id = $1;

-- name: GetUserInfoByUserId :one
SELECT *
FROM user_info
WHERE user_id = $1;
