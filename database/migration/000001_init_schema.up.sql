create table users (
  "id" bigserial primary key,
  "username" text unique not null,
  "email" text unique not null,
  "password" text not null,
  "created_at" timestamptz not null default (now()),
  "updated_at" timestamptz not null default (now())
);