-- name: CreateFeeds :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
Select * from feeds;

-- name: GetFeedByUser :many

Select * from feeds where user_id = (
    Select id from users where api_key=$1
);
