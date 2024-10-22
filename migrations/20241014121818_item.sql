-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items
(
    id_barang         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama_barang       VARCHAR(50) NOT NULL UNIQUE,
    harga             VARCHAR(50) NOT NULL, -- Menetapkan ukuran untuk kolom harga
    kategori          VARCHAR(50) NULL UNIQUE,
    stock             VARCHAR(100) NOT NULL -- Menghapus koma yang tidak perlu
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
