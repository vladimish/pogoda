-- +goose Up
-- +goose StatementBegin
CREATE TABLE favourites
(
    id       SERIAL      NOT NULL,
    user_id  BIGINT      NOT NULL,
    name     VARCHAR(64) NOT NULL,
    lon      FLOAT       NOT NULL,
    lat      FLOAT       NOT NULL,
    selected BOOLEAN     NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE favourites;
-- +goose StatementEnd
