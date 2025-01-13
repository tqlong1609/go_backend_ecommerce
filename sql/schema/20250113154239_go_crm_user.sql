-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS go_crm_user (
    user_id SERIAL PRIMARY KEY, -- PostgreSQL sử dụng SERIAL để tự động tăng
    user_account VARCHAR(100) NOT NULL UNIQUE,
    user_password CHAR(64) NOT NULL,
    user_salt CHAR(32) NOT NULL,
    user_login_time TIMESTAMP DEFAULT NULL,
    user_logout_time TIMESTAMP DEFAULT NULL,
    user_login_ip VARCHAR(45) DEFAULT NULL,
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS go_crm_user;
-- +goose StatementEnd
