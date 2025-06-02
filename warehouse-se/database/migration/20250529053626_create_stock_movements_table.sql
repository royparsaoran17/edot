-- +goose Up
create table if not exists stock_movements
(
    id                uuid primary key,
    product_id        uuid not null,
    from_warehouse_id uuid,
    to_warehouse_id   uuid,
    quantity          integer not null,
    note              varchar,

    created_at timestamp not null,
    updated_at timestamp not null,

    constraint fk_from foreign key (from_warehouse_id) references warehouses (id),
    constraint fk_to foreign key (to_warehouse_id) references warehouses (id)
    );

create index stock_movements_product_id_index on stock_movements (product_id);
create index stock_movements_from_warehouse_index on stock_movements (from_warehouse_id);
create index stock_movements_to_warehouse_index on stock_movements (to_warehouse_id);

-- +goose Down
drop table if exists stock_movements;
