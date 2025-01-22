-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS user_base;
CREATE TABLE IF NOT EXISTS user_base (
    user_id SERIAL PRIMARY KEY,
    user_account VARCHAR(255) NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    user_sait VARCHAR(255) NOT NULL,
    user_login_time TIMESTAMP NULL,
    user_logout_time TIMESTAMP NULL,
    user_login_ip VARCHAR(255) NULL,
    user_created_at TIMESTAMP NULL,
    user_updated_at TIMESTAMP NULL,
    CONSTRAINT user_base_user_account_unique UNIQUE (user_account)
);

ALTER TABLE user_base
    ADD CONSTRAINT user_base_user_account_foreign
    FOREIGN KEY (user_account) REFERENCES user_verify (verify_key);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_base;
-- +goose StatementEnd
