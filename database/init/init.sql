-- Mock Data for test
CREATE SCHEMA IF NOT EXISTS users;
CREATE TABLE IF NOT EXISTS users.users (
  "id" bigserial PRIMARY KEY,
  "username" VARCHAR (50) UNIQUE NOT NULL,
  "email" VARCHAR (50) UNIQUE NOT NULL,
  "password" VARCHAR (255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "last_login" TIMESTAMPTZ
);
INSERT INTO users.users ("username", "email", "password")
VALUES (
    'user1',
    'email1',
    -- pass1
    '$2a$10$EJt5gEzvRnNa46ZjAoE1RuQKumCbJv.ipiNsYC4g14sn4h8CzeN/e'
  ),
  (
    'user2',
    'email2',
    -- pass2
    '$2a$10$G7bXjXB.swnw5XpOWh994eYGRe9pxVjdgakzXwqVONg54P86fiE.K'
  );