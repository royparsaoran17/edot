-- +goose Up
create table if not exists payments (
    id uuid primary key,
    order_id uuid not null,
    method varchar not null, -- e.g., 'credit_card', 'bank_transfer'
    status varchar not null, -- e.g., 'pending', 'success', 'failed'
    amount numeric(12,2) not null,
    paid_at timestamp,

    created_at timestamp not null,
    updated_at timestamp not null,

    constraint fk_payment_order foreign key (order_id) references orders (id)
);

-- +goose Down
drop table if exists payments;