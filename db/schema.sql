CREATE TABLE "diseases" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "name" varchar NOT NULL,
  "omim_id" varchar NOT NULL UNIQUE,
  "link_fcbk_group" varchar
);

CREATE TABLE "patients" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "disease_id" bigserial NOT NULL REFERENCES diseases,
  "age" int,
  "created_at" timestamptz DEFAULT (now())
);

