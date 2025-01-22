-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS user_info;
CREATE TABLE IF NOT EXISTS user_info (
    user_id SERIAL PRIMARY KEY,
    user_account VARCHAR(255) NOT NULL,
    user_name VARCHAR(255),
    user_avatar VARCHAR(255),
    user_state SMALLINT NOT NULL,
    user_phone VARCHAR(255),
    user_gender SMALLINT,
    user_birthday DATE,
    user_email VARCHAR(255),
    user_is_auth SMALLINT NOT NULL,
    create_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    CONSTRAINT user_info_user_account_unique UNIQUE (user_account)
);

ALTER TABLE user_info
    ADD CONSTRAINT user_info_user_id_foreign
    FOREIGN KEY (user_id) REFERENCES user_base (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_info;
-- +goose StatementEnd
