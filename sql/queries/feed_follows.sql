-- name: FollowFeed :exec
INSERT INTO feedfollows (id, created_at, updated_at, user_id, feed_id)
VALUES($1, $2, $3, $4, $5);
