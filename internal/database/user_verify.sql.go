// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_verify.sql

package database

import (
	"context"
	"database/sql"
)

const getOtpRecord = `-- name: GetOtpRecord :one
SELECT verify_id, verify_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted, verify_created, verify_updated, failed_attempts, lock_until
FROM user_verify
WHERE verify_key = $1
`

func (q *Queries) GetOtpRecord(ctx context.Context, verifyKey string) (UserVerify, error) {
	row := q.db.QueryRowContext(ctx, getOtpRecord, verifyKey)
	var i UserVerify
	err := row.Scan(
		&i.VerifyID,
		&i.VerifyOtp,
		&i.VerifyKey,
		&i.VerifyKeyHash,
		&i.VerifyType,
		&i.IsVerified,
		&i.IsDeleted,
		&i.VerifyCreated,
		&i.VerifyUpdated,
		&i.FailedAttempts,
		&i.LockUntil,
	)
	return i, err
}

const insertOrUpdateOtp = `-- name: InsertOrUpdateOtp :exec
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
  lock_until = EXCLUDED.lock_until
`

type InsertOrUpdateOtpParams struct {
	VerifyKey string
	VerifyOtp string
}

func (q *Queries) InsertOrUpdateOtp(ctx context.Context, arg InsertOrUpdateOtpParams) error {
	_, err := q.db.ExecContext(ctx, insertOrUpdateOtp, arg.VerifyKey, arg.VerifyOtp)
	return err
}

const lockOtp = `-- name: LockOtp :exec
UPDATE user_verify
SET lock_until = $2,
    verify_updated = NOW()
WHERE verify_key = $1
`

type LockOtpParams struct {
	VerifyKey string
	LockUntil sql.NullTime
}

func (q *Queries) LockOtp(ctx context.Context, arg LockOtpParams) error {
	_, err := q.db.ExecContext(ctx, lockOtp, arg.VerifyKey, arg.LockUntil)
	return err
}

const updateFailedAttempts = `-- name: UpdateFailedAttempts :exec
UPDATE user_verify
SET failed_attempts = $2,
    verify_updated = NOW()
WHERE verify_key = $1
`

type UpdateFailedAttemptsParams struct {
	VerifyKey      string
	FailedAttempts int32
}

func (q *Queries) UpdateFailedAttempts(ctx context.Context, arg UpdateFailedAttemptsParams) error {
	_, err := q.db.ExecContext(ctx, updateFailedAttempts, arg.VerifyKey, arg.FailedAttempts)
	return err
}

const verifySuccessOtp = `-- name: VerifySuccessOtp :exec
UPDATE user_verify
SET is_verified = 1,
    failed_attempts = 0,
    verify_updated = NOW()
WHERE verify_key = $1
`

func (q *Queries) VerifySuccessOtp(ctx context.Context, verifyKey string) error {
	_, err := q.db.ExecContext(ctx, verifySuccessOtp, verifyKey)
	return err
}
