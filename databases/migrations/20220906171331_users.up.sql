CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL UNIQUE,
  "email" varchar(50) NOT NULL UNIQUE,
  "password" varchar(255) NOT NULL,
  "active" boolean NOT NULL,
  "verified" boolean NULL,
  "photo" text NULL,
  "role_id" int NOT NULL,
  "created_at" timestamp NULL DEFAULT (now()),
  "updated_at" timestamp NULL DEFAULT (now()),
  "deleted_at" timestamp NULL
);

ALTER TABLE "users" ADD FOREIGN KEY (role_id) REFERENCES "roles" ON DELETE RESTRICT;