-- +goose Up
create table if not exists stock_reservations (
    id uuid primary key,
    order_id uuid not null,
    product_id uuid not null,
    warehouse_id uuid not null,
    quantity integer not null,
    reserved_at timestamp not null,
    expires_at timestamp not null,
    is_active boolean not null default true, -- false jika expired atau dibatalkan

    created_at timestamp not null,
    updated_at timestamp not null,

    constraint fk_order_reserve foreign key (order_id) references orders (id)
);

-- +goose Down
drop table if exists stock_reservations;