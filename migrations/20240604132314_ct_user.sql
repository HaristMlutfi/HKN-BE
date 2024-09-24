-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id         uuid primary key default uuid_generate_v4(),
    email       varchar(50)  not null unique,
    password    varchar      not null,
    google_id  varchar(50) null unique,
    name        varchar(100) not null,
    is_verified boolean      not null default false,
    created_at  timestamp             default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
