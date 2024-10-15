-- +goose Up
ALTER TABLE posts 
  ALTER COLUMN url SET NOT NULL,
  ADD CONSTRAINT posts_url_unique UNIQUE (url);

-- +goose Down
ALTER TABLE posts 
  ALTER COLUMN url SET NOT NULL,
  DROP CONSTRAINT posts_url_unique;
