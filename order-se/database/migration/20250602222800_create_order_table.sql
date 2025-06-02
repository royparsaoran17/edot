-- +goose Up
create table if not exists orders (
    id uuid primary key,
    user_id uuid not null,
    status varchar not null, -- e.g., pending, paid, cancelled
    total_price numeric(12, 2) not null default 0,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
);

-- +goose Down
drop table if exists orders;