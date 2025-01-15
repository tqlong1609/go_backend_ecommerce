-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS go_crm_user_verify;
CREATE TABLE go_crm_user_verify(
    verify_id SERIAL PRIMARY KEY,
    verify_otp VARCHAR(255) NOT NULL,
    verify_key VARCHAR(255) NOT NULL,
    verify_key_hash VARCHAR(255) NOT NULL,
    verify_type INTEGER DEFAULT 1,
    is_verified INTEGER,
    is_deleted INTEGER,
    verify_created TIMESTAMP,
    verify_updated TIMESTAMP
);
CREATE INDEX user_verify_verify_otp_index ON go_crm_user_verify(verify_otp);
CREATE UNIQUE INDEX user_verify_verify_key_unique ON go_crm_user_verify(verify_key);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_crm_user_verify;
-- +goose StatementEnd
