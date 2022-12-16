-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE pinjam_buku (
  id SERIAL NOT NULL,
  id_buku BIGINT,
  id_user_peminjam BIGINT,
  created_by VARCHAR(256),
  updated_by VARCHAR(256),
  updated_at VARCHAR(256),
  PRIMARY KEY (id)

)

-- +migrate StatementEnd
