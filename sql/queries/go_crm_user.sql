-- name: CreateNewUser :one
INSERT INTO go_crm_user (
    user_account, user_password, user_salt, user_created_at, user_updated_at
) VALUES (
    $1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
) RETURNING user_id;

-- name: GetUserAccountByUserId :one
SELECT user_account
FROM go_crm_user
WHERE user_id = $1;

-- name: UpdateUserPasswordByUserId :exec
UPDATE go_crm_user
SET user_password = $2,
    user_salt = $3,
    user_updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1;

-- name: DeleteUserByUserId :exec
DELETE FROM go_crm_user
WHERE user_id = $1;
