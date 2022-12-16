-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE user_login (
  id SERIAL NOT NULL,
  username VARCHAR(256),
  nama VARCHAR(256),
  "role" VARCHAR(256),
  token VARCHAR(256),
  PRIMARY KEY (id)
)

-- +migrate StatementEnd
