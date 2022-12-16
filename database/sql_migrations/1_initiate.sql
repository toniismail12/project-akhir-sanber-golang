-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE buku (
  id SERIAL NOT NULL,
  judul_buku VARCHAR(256),
  author VARCHAR(256),
  tahun_terbit BIGINT,
  publisher VARCHAR(256),
  ISBN VARCHAR(256),
  jumlah_halaman VARCHAR(256),
  is_dipinjam VARCHAR(256),
  created_at VARCHAR(256),
  created_by VARCHAR(256),
  updated_by VARCHAR(256),
  updated_at VARCHAR(256),
  PRIMARY KEY (id)

)

-- +migrate StatementEnd
