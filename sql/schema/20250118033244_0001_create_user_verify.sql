-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS user_verify;
CREATE TABLE IF NOT EXISTS user_verify (
    verify_id SERIAL PRIMARY KEY,
    verify_otp VARCHAR(255) NOT NULL,
    verify_key VARCHAR(255) NOT NULL,
    verify_key_hash VARCHAR(255) NOT NULL,
    verify_type INT DEFAULT 1,
    is_verified INT,
    is_deleted INT,
    verify_created TIMESTAMP NULL,
    verify_updated TIMESTAMP NULL,
    failed_attempts INT DEFAULT 0 NOT NULL,
    lock_until TIMESTAMP NULL,
    CONSTRAINT user_verify_verify_otp_index UNIQUE (verify_otp),
    CONSTRAINT user_verify_verify_key_unique UNIQUE (verify_key)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_verify;
-- +goose StatementEnd
