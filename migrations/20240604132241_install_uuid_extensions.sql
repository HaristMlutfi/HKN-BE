-- +goose Up
-- +goose StatementBegin
create
extension if not exists "uuid-ossp";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
