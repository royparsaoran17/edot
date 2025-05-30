-- +goose Up
create table if not exists warehouses
(
    id         uuid primary key,
    name       varchar not null,
    is_active  boolean not null default true,

    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create unique index warehouses_name_uindex on warehouses (name);
create index warehouses_is_active_index on warehouses (is_active);

-- +goose Down
drop table if exists warehouses;