-- +goose Up

CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VarChar(128) NOT NULL
);

-- +goose Down

DROP TABLE users;