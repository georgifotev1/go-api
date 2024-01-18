create table users (
  "id" bigserial primary key,
  "username" varchar not null,
  "email" text UNIQUE not null,
  "password" text not null,
  "created_at" timestamptz not null default (now())
);