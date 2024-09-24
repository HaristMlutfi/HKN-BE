-- +goose Up
-- +goose StatementBegin
create table if not exists roles
(
    id          uuid primary key default uuid_generate_v4(),
    code        varchar(50) not null unique,
    name        varchar(50) not null unique,
    description text,
    created_at  timestamp        default current_timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles;
-- +goose StatementEnd
