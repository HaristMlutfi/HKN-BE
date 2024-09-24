-- +goose Up
-- +goose StatementBegin
with inserted_user as (
    insert into users (
                       email,
                       password,
                       name,
                       google_id,
                       is_verified
        )
        values ('superadmin@lms.lms',
                '$2a$10$JW2T1lLMPjn5HdJeDf/JnO1W1fJyFrxvWR/SFUplFYnEEp5S7Y9s2',
                'Bismillah@Lms',
                'null',
                true)
        returning id as user_id),
     inserted_role as (
         insert into roles (
                            code,
                            name,
                            description
             ) values ('superadmin',
                       'Super Admin',
                       'role_repo super admin')
             returning id as role_id)
insert
into user_roles (user_id, role_id)
select user_id, role_id
from inserted_user,
     inserted_role;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
