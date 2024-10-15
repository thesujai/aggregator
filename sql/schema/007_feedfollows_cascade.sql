-- +goose Up
ALTER TABLE feedfollows DROP CONSTRAINT IF EXISTS feedfollows_user_id_fkey;
ALTER TABLE feedfollows DROP CONSTRAINT IF EXISTS feedfollows_feed_id_fkey;

ALTER TABLE feedfollows 
    ADD CONSTRAINT feedfollows_user_id_fkey 
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE feedfollows 
    ADD CONSTRAINT feedfollows_feed_id_fkey 
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE feedfollows DROP CONSTRAINT IF EXISTS feedfollows_user_id_fkey;
ALTER TABLE feedfollows DROP CONSTRAINT IF EXISTS feedfollows_feed_id_fkey;

ALTER TABLE feedfollows 
    ADD CONSTRAINT feedfollows_user_id_fkey 
    FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE feedfollows 
    ADD CONSTRAINT feedfollows_feed_id_fkey 
    FOREIGN KEY (feed_id) REFERENCES feeds(id);

