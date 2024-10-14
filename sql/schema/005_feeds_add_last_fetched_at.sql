-- +goose Up
Alter Table feeds Add Column last_fetched_at TIMESTAMP;

-- +goose Down
Alter table feeds drop Column last_fetched_at;
