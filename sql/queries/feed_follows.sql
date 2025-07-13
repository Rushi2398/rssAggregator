-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2;

-- name: GetFeedsNextToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;