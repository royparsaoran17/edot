-- +goose Up
create table if not exists order_items (
    id uuid primary key,
    order_id uuid not null,
    product_id uuid not null,
    warehouse_id uuid not null,
    quantity integer not null,
    price numeric(12, 2) not null,

    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,

    constraint fk_order foreign key (order_id) references orders (id)
);

-- +goose Down
drop table if exists order_items;