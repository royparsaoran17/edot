-- +goose Up
create table if not exists product_stocks
(
    id           uuid primary key,
    product_id   uuid not null,
    warehouse_id uuid not null,
    quantity     integer not null,

    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,

    constraint fk_product foreign key (product_id) references products (id),
    constraint fk_warehouse foreign key (warehouse_id) references warehouses (id)
    );

create index product_stocks_product_id_index on product_stocks (product_id);
create index product_stocks_warehouse_id_index on product_stocks (warehouse_id);

-- +goose Down
drop table if exists product_stocks;
