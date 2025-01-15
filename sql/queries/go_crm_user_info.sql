-- name: CreateUserInfo :one
INSERT INTO go_crm_user_info (
  user_account, user_name, user_avatar, user_state, user_phone, user_gender, user_birthday, user_email, user_is_auth, create_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW()
)
RETURNING *;

-- name: GetUserInfo :one
SELECT *
FROM go_crm_user_info
WHERE user_id = $1;

-- name: UpdateUserInfo :exec
UPDATE go_crm_user_info
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
