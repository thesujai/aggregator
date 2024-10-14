-- +goose Up
Create Table feedfollows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    feed_id UUID NOT NULL REFERENCES feeds(id),
    UNIQUE (user_id, feed_id)
);
-- +goose Down
DROP TABLE feedfollows;
