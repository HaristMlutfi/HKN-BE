-- +goose Up
-- +goose StatementBegin
insert into roles (code, name)
values ('dealer', 'Dealer');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
