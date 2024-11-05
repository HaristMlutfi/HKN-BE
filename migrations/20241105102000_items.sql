-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items
(
    id   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama_barang  VARCHAR(50) NOT NULL UNIQUE,
    harga        NUMERIC(12, 2) NOT NULL, 
    kategori     VARCHAR(50),            
    stock        INTEGER NOT NULL,       
    discount     INTEGER NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
