-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-05-04T15:44:23.017Z

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "actions" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "title" varchar NOT NULL,
  "completed" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "subactions" (
  "id" bigserial PRIMARY KEY,
  "actions_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "completed" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "actions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "subactions" ADD FOREIGN KEY ("actions_id") REFERENCES "actions" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
