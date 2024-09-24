-- +goose Up
-- +goose StatementBegin
insert into roles (code, name)
values ('other', 'Other'),
       ('local-government-official', 'Local Government Official'),
       ('aplg-member', 'APLG Member'),
       ('student', 'Student');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
