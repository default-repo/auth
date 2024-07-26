-- +goose Up
create table customer
(
    id         serial primary key,
    name       text      not null,
    password   text      not null,
    email      text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp
);

-- +goose Down
drop table customer;
