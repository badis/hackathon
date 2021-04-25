CREATE TABLE "diseases" (
  "id" bigserial PRIMARY KEY,
  "omim_id" varchar NOT NULL,
  "link_fcbk_group" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "patients" (
  "id" bigserial PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "disease_id" bigint NOT NULL,
  "age" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "patients" ADD FOREIGN KEY ("disease_id") REFERENCES "diseases" ("id");



