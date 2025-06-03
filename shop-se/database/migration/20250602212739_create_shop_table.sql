-- +goose Up
create table if not exists shops (
    id uuid primary key,
    name varchar not null,
    owner_id uuid not null,

    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

-- +goose Down
drop table if exists shops;