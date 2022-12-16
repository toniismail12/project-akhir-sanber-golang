-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE "users" (
  id SERIAL NOT NULL,
  username VARCHAR(256),
  "password" VARCHAR(256),
  email VARCHAR(256),
  wa VARCHAR(256),
  nama VARCHAR(256),
  alamat VARCHAR(256),
  "role" VARCHAR(256),
  created_by VARCHAR(256),
  created_at VARCHAR(256),
  updated_by VARCHAR(256),
  updated_at VARCHAR(256),
  PRIMARY KEY (id)

)

-- +migrate StatementEnd
