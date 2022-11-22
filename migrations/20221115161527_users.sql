-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id                BIGSERIAL     NOT NULL,
    telegram_id       BIGINT UNIQUE NOT NULL,
    first_name        VARCHAR(64)   NOT NULL,
    state             INT2          NOT NULL DEFAULT 0,
    register_date     TIMESTAMP     NOT NULL,
    last_message_date TIMESTAMP     NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
