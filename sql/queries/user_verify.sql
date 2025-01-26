-- name: InsertOrUpdateOtp :exec
INSERT INTO user_verify (
  verify_key, verify_otp, verify_key_hash, verify_type, is_verified, is_deleted, verify_created, verify_updated, failed_attempts, lock_until
) VALUES (
  $1, $2, '', 1, 0, 0, NOW(), NOW(), 0, NULL
) ON CONFLICT (verify_key) 
DO UPDATE SET 
  verify_otp = EXCLUDED.verify_otp,
  verify_key_hash = EXCLUDED.verify_key_hash,
  verify_type = EXCLUDED.verify_type,
  is_verified = EXCLUDED.is_verified,
  is_deleted = EXCLUDED.is_deleted,
  verify_updated = NOW(),
  failed_attempts = EXCLUDED.failed_attempts,
  lock_until = EXCLUDED.lock_until;

-- name: GetOtpRecord :one
SELECT *
FROM user_verify
WHERE verify_key = $1;

-- name: UpdateFailedAttempts :exec
UPDATE user_verify
SET failed_attempts = $2,
    verify_updated = NOW()
WHERE verify_key = $1;

-- name: LockOtp :exec
UPDATE user_verify
SET lock_until = $2,
    verify_updated = NOW()
WHERE verify_key = $1;

-- name: VerifySuccessOtp :exec
UPDATE user_verify
SET is_verified = 1,
    failed_attempts = 0,
    verify_updated = NOW()
WHERE verify_key = $1;

