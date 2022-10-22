CREATE SCHEMA IF NOT EXISTS users;
CREATE TABLE IF NOT EXISTS users.users (
  "id" bigserial PRIMARY KEY,
  "username" VARCHAR (50) UNIQUE NOT NULL,
  "email" VARCHAR (50) UNIQUE NOT NULL,
  "password" VARCHAR (255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "last_login" TIMESTAMPTZ
);