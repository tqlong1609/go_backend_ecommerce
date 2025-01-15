-- name: CreateOTP :one
INSERT INTO go_crm_user_verify (
  verify_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted, verify_created, verify_updated
) VALUES (
  $1, $2, $3, $4, 0, 0, NOW(), NOW()
)
RETURNING *;

-- name: GetInfoOTP :one
SELECT *
FROM go_crm_user_verify
WHERE verify_otp = $1 AND is_deleted = 0
LIMIT 1;

-- name: UpdateVerifyStatus :exec
UPDATE go_crm_user_verify
SET is_verified = $2, verify_updated = NOW()
WHERE verify_otp = $1 AND is_deleted = 0;

-- name: GetValidOTP :one
SELECT *
FROM go_crm_user_verify
WHERE verify_otp = $1 AND is_verified = 0 AND is_deleted = 0
LIMIT 1;
