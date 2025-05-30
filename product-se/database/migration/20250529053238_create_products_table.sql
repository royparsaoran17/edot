-- +goose Up
create table if not exists products
(
    id          uuid primary key,
    name        varchar        not null,
    description text,
    price       numeric(12, 2) not null default 0,
    unit        varchar        not null default 'pcs',
    sku         varchar unique,
    category    varchar,
    is_active   boolean         not null default true,

    created_at  timestamp       not null,
    updated_at  timestamp       not null,
    deleted_at  timestamp
    );

-- Indexes for quick lookup
create unique index if not exists products_name_uindex on products (name);
create index if not exists products_sku_index on products (sku);
create index if not exists products_category_index on products (category);
create index if not exists products_is_active_index on products (is_active);

-- +goose Down
drop table if exists products;