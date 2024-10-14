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

-- name: GetFollowedFeeds :many
Select * from feeds where id in (
    Select feed_id from feedfollows where feedfollows.user_id=$1
);

-- name: GetNextFeedsToFetch :many
Select url from feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at=NOW() and updated_at=NOW();
