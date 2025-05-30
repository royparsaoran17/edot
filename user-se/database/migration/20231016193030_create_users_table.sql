-- +goose Up
create table if not exists users
(
    id          uuid
        constraint users_pk
        primary key,
    name        varchar   not null,
    email       varchar   not null,
    phone       varchar   not null,
    password    varchar   not null,
    role_id     uuid   not null
        constraint users_role_uid_fk
        references roles (id)
        on delete cascade,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp
);

create unique index users_phone_uindex
    on users (phone);
    
create index users_email_index on users (email);
create index users_phone_index on users (phone);

-- +goose Down
drop table if exists users;
