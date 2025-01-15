-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS go_crm_user_info;
CREATE TABLE go_crm_user_info(
    user_id SERIAL PRIMARY KEY,
    user_account VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NULL,
    user_avatar VARCHAR(255) NULL,
    user_state SMALLINT NOT NULL,
    user_phone VARCHAR(255) NULL,
    user_gender SMALLINT NULL,
    user_birthday DATE NULL,
    user_email VARCHAR(255) NULL,
    user_is_auth SMALLINT NOT NULL,
    create_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);
ALTER TABLE go_crm_user_info ADD CONSTRAINT user_info_user_account_unique UNIQUE (user_account);
ALTER TABLE go_crm_user ADD CONSTRAINT user_account_foreign FOREIGN KEY (user_account) REFERENCES go_crm_user_verify (verify_key);
ALTER TABLE go_crm_user_info ADD CONSTRAINT user_info_user_id_foreign FOREIGN KEY (user_id) REFERENCES go_crm_user (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_crm_user_info;
-- +goose StatementEnd
