-- +goose Up
-- +goose StatementBegin
create table if not exists items
(
    id_barang         uuid primary key default uuid_generate_v4(),
    nama_barang       varchar(50)  not null unique,
    harga    varchar      not null,
    kategori  varchar(50) null unique,
    stock        varchar(100) not null,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists items;
-- +goose StatementEnd
