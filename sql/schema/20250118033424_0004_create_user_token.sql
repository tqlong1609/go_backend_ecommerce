-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS user_token;
CREATE TABLE IF NOT EXISTS user_token (
    user_id SERIAL PRIMARY KEY,
    refresh_token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    CONSTRAINT user_tokens_refresh_token_unique UNIQUE (refresh_token)
);
ALTER TABLE user_token
    ADD CONSTRAINT user_tokens_user_id_foreign
    FOREIGN KEY (user_id) REFERENCES user_base (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_token;
-- +goose StatementEnd
