-- +goose Up
-- +goose StatementBegin
create table if not exists user_roles
(
    user_id uuid not null,
    role_id uuid not null,
    foreign key (user_id) references users (id),
    foreign key (role_id) references roles (id),
    unique (user_id, role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_role;
-- +goose StatementEnd
