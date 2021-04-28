CREATE TABLE "diseases" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "omim_id" varchar NOT NULL,
  "link_fcbk_group" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "patients" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar,
  "disease_id" bigint,
  "age" int,
  "created_at" timestamptz DEFAULT (now())
);

ALTER TABLE "patients" ADD FOREIGN KEY ("disease_id") REFERENCES "diseases" ("id");

